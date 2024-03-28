package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	nethttp "net/http"
	_ "net/http/pprof"

	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/snple/kirara"
	"github.com/snple/kirara/bin/config"
	"github.com/snple/kirara/bin/log"
	"github.com/snple/kirara/db"
	"github.com/snple/kirara/edge"
	"github.com/snple/kirara/http"
	"github.com/snple/kirara/http/api"
	"github.com/snple/kirara/http/web"
	"github.com/snple/kirara/plugins/emu"
	"github.com/snple/kirara/plugins/slim"
	"github.com/snple/kirara/slot"
	"github.com/snple/kirara/util"
	"github.com/snple/kirara/util/compress/zstd"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "version", "-V":
			fmt.Printf("kirara edge version: %v\n", kirara.Version)
			return
		}
	}

	rand.Seed(time.Now().Unix())

	config.Parse()

	log.Init(config.Config.Debug)

	log.Logger.Info("main: Started")
	defer log.Logger.Info("main: Completed")

	bundb, err := db.ConnectSqlite(config.Config.DB.File, config.Config.DB.Debug)
	if err != nil {
		log.Logger.Sugar().Fatalf("connecting to db: %v", err)
	}

	defer bundb.Close()

	if err = edge.CreateSchema(bundb); err != nil {
		log.Logger.Sugar().Fatalf("create schema: %v", err)
	}

	if err = edge.Seed(bundb); err != nil {
		log.Logger.Sugar().Fatalf("seed: %v", err)
	}

	command := flag.Arg(0)
	switch command {
	case "seed":
		log.Logger.Sugar().Infof("seed: Completed")
		return
	case "pull", "push":
		if err := cli(command, bundb); err != nil {
			log.Logger.Sugar().Errorf("error: shutting down: %s", err)
		}

		return
	}

	edgeOpts := make([]edge.EdgeOption, 0)

	{
		edgeOpts = append(edgeOpts, edge.WithDeviceID(config.Config.DeviceID, config.Config.Secret))
		edgeOpts = append(edgeOpts, edge.WithLinkTTL(time.Second*time.Duration(config.Config.Status.LinkTTL)))

		badgerOptions := func() badger.Options {
			if config.Config.BadgerDB.InMemory {
				return badger.DefaultOptions("").WithInMemory(true)
			}

			return badger.DefaultOptions(config.Config.BadgerDB.Path)
		}()

		edgeOpts = append(edgeOpts, edge.WithBadger(badgerOptions))
	}

	es, err := edge.Edge(bundb, edgeOpts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewEdgeService: %v", err)
	}

	es.Start()
	defer es.Stop()

	zstd.Register()

	if config.Config.EdgeService.Enable {
		edgeGrpcOpts := []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		}

		if config.Config.EdgeService.TLS {
			tlsConfig, err := util.LoadServerCert(config.Config.EdgeService.CA, config.Config.EdgeService.Cert, config.Config.EdgeService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			edgeGrpcOpts = append(edgeGrpcOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
		}

		lis, err := net.Listen("tcp", config.Config.EdgeService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(edgeGrpcOpts...)

		es.Register(s)

		go func() {
			log.Logger.Sugar().Infof("edge grpc start: %v, tls: %v", config.Config.EdgeService.Addr, config.Config.EdgeService.TLS)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	if config.Config.SlotService.Enable {
		slotOpts := make([]slot.SlotOption, 0)

		ns, err := slot.Slot(es, slotOpts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewSlotService: %v", err)
		}

		ns.Start()
		defer ns.Stop()

		slotGrpcOpts := []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		}

		if config.Config.SlotService.TLS {
			tlsConfig, err := util.LoadServerCert(config.Config.SlotService.CA, config.Config.SlotService.Cert, config.Config.SlotService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			slotGrpcOpts = append(slotGrpcOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
		}

		lis, err := net.Listen("tcp", config.Config.SlotService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(slotGrpcOpts...)

		ns.RegisterGrpc(s)

		go func() {
			log.Logger.Sugar().Infof("slot grpc start: %v", config.Config.SlotService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	if !config.Config.Gin.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.Config.WebService.Enable {
		opts := make([]http.HttpServerOption, 0)

		opts = append(opts, http.WithAppName("web"))
		opts = append(opts, http.WithAddr(config.Config.WebService.Addr))
		opts = append(opts, http.WithDebug(config.Config.WebService.Debug))

		if config.Config.WebService.TLS {
			if config.Config.WebService.CA != "" {
				pool := x509.NewCertPool()

				ca, err := os.ReadFile(config.Config.WebService.CA)
				if err != nil {
					log.Logger.Sugar().Fatal(err)
				}

				if ok := pool.AppendCertsFromPEM(ca); !ok {
					log.Logger.Sugar().Fatal(err)
				}

				tlsConfig := &tls.Config{
					ClientAuth: tls.RequireAndVerifyClientCert,
					ClientCAs:  pool,
				}

				opts = append(opts, http.WithTLSConfig(tlsConfig))
			}

			opts = append(opts, http.WithTLS(config.Config.WebService.Cert, config.Config.WebService.Key))
		}

		hs, err := http.NewHttpServer(es.Context(), opts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewHttpServer: %v", err)
		}

		{
			ws, err := web.NewWebService(es)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewWebService: %v", err)
			}

			ws.Register(hs.Engine())

			go ws.Start()
			defer ws.Stop()

			as, err := api.NewApiService(es)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewApiService: %v", err)
			}

			apiGroup := hs.Engine().Group("/api", ws.GetAuth().MiddlewareFunc())
			as.Register(apiGroup)

			go as.Start()
			defer as.Stop()
		}

		go hs.Start()
		defer hs.Stop()
	}

	if config.Config.ApiService.Enable {
		opts := make([]http.HttpServerOption, 0)

		opts = append(opts, http.WithAppName("api"))
		opts = append(opts, http.WithAddr(config.Config.ApiService.Addr))
		opts = append(opts, http.WithDebug(config.Config.ApiService.Debug))

		if config.Config.ApiService.TLS {
			if config.Config.ApiService.CA != "" {
				pool := x509.NewCertPool()

				ca, err := os.ReadFile(config.Config.ApiService.CA)
				if err != nil {
					log.Logger.Sugar().Fatal(err)
				}

				if ok := pool.AppendCertsFromPEM(ca); !ok {
					log.Logger.Sugar().Fatal(err)
				}

				tlsConfig := &tls.Config{
					ClientAuth: tls.RequireAndVerifyClientCert,
					ClientCAs:  pool,
				}

				opts = append(opts, http.WithTLSConfig(tlsConfig))
			}

			opts = append(opts, http.WithTLS(config.Config.ApiService.Cert, config.Config.ApiService.Key))
		}

		hs, err := http.NewHttpServer(es.Context(), opts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewHttpServer: %v", err)
		}

		{
			as, err := api.NewApiService(es)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewApiService: %v", err)
			}

			as.Register(hs.Engine())

			go as.Start()
			defer as.Stop()
		}

		go hs.Start()
		defer hs.Stop()
	}

	for _, static := range config.Config.Statics {
		if !static.Enable {
			continue
		}

		engine := gin.New()
		engine.Use(gin.Recovery())
		engine.Static("/", static.Path)
		engine.NoRoute(func(ctx *gin.Context) {
			ctx.File(static.Path + "/index.html")
		})

		log.Logger.Sugar().Infof("static server start： %v", static.Addr)

		if static.TLS {
			go engine.RunTLS(static.Addr, static.Cert, static.Key)
		} else {
			go engine.Run(static.Addr)
		}
	}

	// if config.Config.GoS7.Enable {
	// 	plugin, err := gos7.GoS7(es,
	// 		gos7.WithTickerInterval(time.Second*time.Duration(config.Config.GoS7.Interval)),
	// 		gos7.WithReadDataInterval(time.Second*time.Duration(config.Config.GoS7.ReadInterval)))
	// 	if err != nil {
	// 		log.Logger.Sugar().Fatalf("GoS7: %v", err)
	// 	}

	// 	go plugin.Start()
	// 	defer plugin.Stop()
	// }

	if config.Config.Slim.Enable {
		slim, err := slim.Slim(es,
			slim.WithTickerInterval(time.Second*time.Duration(config.Config.Slim.Interval)),
			slim.WithCacheTTL(time.Second*time.Duration(config.Config.Slim.CacheTTL)),
			slim.WithBBolt(config.Config.Slim.BBolt),
		)
		if err != nil {
			log.Logger.Sugar().Fatalf("Slim: %v", err)
		}

		go slim.Start()
		defer slim.Stop()
	}

	if config.EnableEmu {
		plugin, err := emu.Emu(es, emu.WithTickerInterval(time.Second*10))
		if err != nil {
			log.Logger.Sugar().Fatalf("Emu: %v", err)
		}

		go plugin.Start()
		defer plugin.Stop()
	}

	go func() {
		nethttp.ListenAndServe(":6060", nil)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}

func cli(command string, bundb *bun.DB) error {
	log.Logger.Sugar().Infof("cli %v: Started", command)
	defer log.Logger.Sugar().Infof("cli %v : Completed", command)

	edgeOpts := make([]edge.EdgeOption, 0)
	edgeOpts = append(edgeOpts, edge.WithDeviceID(config.Config.DeviceID, config.Config.Secret))

	es, err := edge.Edge(bundb, edgeOpts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewEdgeService: %v", err)
	}

	_ = es

	return nil
}
