package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigStruct struct {
	Debug       bool        `toml:"debug"`
	DeviceID    string      `toml:"device_id"`
	Secret      string      `toml:"secret"`
	DB          DB          `toml:"db"`
	BadgerDB    BadgerDB    `toml:"badger"`
	NodeClient  GRPCClient  `toml:"node"`
	EdgeService GRPCService `toml:"edge"`
	SlotService GRPCService `toml:"slot"`
	Sync        Sync        `toml:"sync"`
	Status      Status      `toml:"status"`
	Gin         Gin         `toml:"gin"`
	WebService  HttpService `toml:"web"`
	ApiService  HttpService `toml:"api"`
	Statics     []Static    `toml:"static"`
	Upload      Upload      `toml:"upload"`
	Slim        Slim        `toml:"slim"`
	Source      Source      `toml:"source"`
}

type DB struct {
	Debug bool   `toml:"debug"`
	File  string `toml:"file"`
}

type BadgerDB struct {
	Path     string `toml:"path"`
	InMemory bool   `toml:"in_memory"`
}

type GRPCClient struct {
	Enable             bool   `toml:"enable"`
	Addr               string `toml:"addr"`
	TLS                bool   `toml:"tls"`
	CA                 string `toml:"ca"`
	Cert               string `toml:"cert"`
	Key                string `toml:"key"`
	ServerName         string `toml:"server_name"`
	InsecureSkipVerify bool   `toml:"insecure_skip_verify"`
}

type GRPCService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	TLS    bool   `toml:"tls"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Sync struct {
	TokenRefresh int  `toml:"token_refresh"`
	Link         int  `toml:"link"`
	Interval     int  `toml:"interval"`
	Realtime     bool `toml:"realtime"`
}

type Status struct {
	LinkTTL int `toml:"link_ttl"`
}

type Gin struct {
	Debug bool `toml:"debug"`
}

type HttpService struct {
	Enable bool   `toml:"enable"`
	Debug  bool   `toml:"debug"`
	Addr   string `toml:"addr"`
	TLS    bool   `toml:"tls"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Static struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	Path   string `toml:"path"`
	TLS    bool   `toml:"tls"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Upload struct {
	Enable   bool `toml:"enable"`
	Interval int  `toml:"interval"`
	Batch    int  `toml:"batch"`
}

type Slim struct {
	Enable   bool   `toml:"enable"`
	Interval int    `toml:"interval"`
	BBolt    string `toml:"bbolt"`
	CacheTTL int    `toml:"cache_ttl"`
}

type Source struct {
	Enable       bool `toml:"enable"`
	Interval     int  `toml:"interval"`
	ReadInterval int  `toml:"read_interval"`
	GoS7         bool `toml:"gos7"`
	Modbus       bool `toml:"modbus"`
}

func DefaultConfig() ConfigStruct {
	return ConfigStruct{
		Debug: false,
		DB: DB{
			File: "store.db",
		},
		BadgerDB: BadgerDB{
			Path:     "badger",
			InMemory: true,
		},
		NodeClient: GRPCClient{
			Addr: "127.0.0.1:6007",
			TLS:  true,
			CA:   "certs/ca.crt",
			Cert: "certs/client.crt",
			Key:  "certs/client.key",
		},
		EdgeService: GRPCService{
			Addr: "127.0.0.1:6010",
			TLS:  true,
			CA:   "certs/ca.crt",
			Cert: "certs/server.crt",
			Key:  "certs/server.key",
		},
		SlotService: GRPCService{
			Addr: "127.0.0.1:6011",
			TLS:  true,
			CA:   "certs/ca.crt",
			Cert: "certs/server.crt",
			Key:  "certs/server.key",
		},
		Sync: Sync{
			TokenRefresh: 60 * 30,
			Link:         60,
			Interval:     60,
			Realtime:     false,
		},
		Status: Status{
			LinkTTL: 3 * 60,
		},
		WebService: HttpService{
			Addr: ":8010",
		},
		ApiService: HttpService{
			Addr: ":8012",
		},
		Upload: Upload{
			Enable:   false,
			Interval: 60,
			Batch:    1000,
		},
		Source: Source{
			Interval:     30,
			ReadInterval: 30,
		},
		Slim: Slim{
			Interval: 60,
			BBolt:    "slim.db",
			CacheTTL: 60 * 60 * 24 * 7,
		},
	}
}

var Config = DefaultConfig()

func (c *ConfigStruct) Validate() error {
	if len(c.DeviceID) == 0 {
		return errors.New("DeviceID must be specified")
	}

	if len(c.Secret) == 0 {
		return errors.New("Secret must be specified")
	}

	if c.EdgeService.Enable {
		if len(c.EdgeService.Addr) == 0 {
			return errors.New("EdgeService.Addr must be specified")
		}

		if c.EdgeService.TLS {
			if len(c.EdgeService.CA) == 0 {
				return errors.New("EdgeService.CA must be specified")
			}

			if len(c.EdgeService.Cert) == 0 {
				return errors.New("EdgeService.Cert must be specified")
			}

			if len(c.EdgeService.Key) == 0 {
				return errors.New("EdgeService.Key must be specified")
			}
		}
	}

	return nil
}

var EnableEmu = false

func Parse() {
	var err error

	configFile := flag.String("c", "config.toml", "config file")
	flag.BoolVar(&EnableEmu, "emu", false, "-emu")

	flag.Parse()

	if _, err = toml.DecodeFile(*configFile, &Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = Config.Validate(); err != nil {
		fmt.Println("config:", err)
		os.Exit(1)
	}
}
