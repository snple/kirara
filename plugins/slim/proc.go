package slim

import (
	"context"
	"sync"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/pb"
	"github.com/snple/slim"
)

type Proc struct {
	ss    *SlimSlot
	fn    *pb.Fn
	scope *slim.Scope
	// modules *DynamicModules

	exit  bool
	debug int32
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newProc(ss *SlimSlot, fn *pb.Fn) (*Proc, error) {
	ctx, cancel := context.WithCancel(ss.ctx)

	p := &Proc{
		ss:     ss,
		fn:     fn,
		debug:  fn.GetDebug(),
		ctx:    ctx,
		cancel: cancel,
	}

	vars := slim.NewVars()
	vars.SetAny("proc", fnToSlimObject(fn))
	vars.SetAny("info", p.userFunctionInfo)
	vars.SetAny("debug", p.userFunctionDebug)
	vars.SetAny("exit", p.userFunctionExit)

	modules := NewDynamicModules(ss, fn.GetId())

	p.scope = slim.NewScope(modules, vars)

	return p, nil
}

func (p *Proc) start() {
	defer func() {
		if re := recover(); re != nil {
			p.ss.logger().Sugar().Errorf("proc: %v, run error: %v", p.fn.GetName(), re)
		}
	}()

	p.closeWG.Add(1)
	defer p.closeWG.Done()

	defer p.cancel()

	go p.waitExit()

	compiled, err := p.scope.Complie("main", []byte(p.fn.GetMain()))
	if err != nil {
		p.ss.logger().Sugar().Errorf("slim complie main: %v", err)
		return
	}

	err = p.scope.Run(compiled)
	if err != nil {
		p.ss.logger().Sugar().Errorf("slim run main: %v", err)
		return
	}
}

func (p *Proc) stop() {
	p.lock.Lock()
	p.exit = true
	p.lock.Unlock()

	p.cancel()
	p.closeWG.Wait()
}

func (p *Proc) waitExit() {
	<-p.ctx.Done()

	p.lock.Lock()
	defer p.lock.Unlock()

	p.exit = true
}

// func (p *Proc) reset() {
// 	ctx, cancel := context.WithCancel(p.ss.ctx)

// 	p.exit = false
// 	p.ctx = ctx
// 	p.cancel = cancel
// }

func (p *Proc) setDebug(value int32) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.debug = value
}

func (p *Proc) getDebug() int32 {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.debug
}

func (p *Proc) userFunctionInfo(args ...slim.Object) (slim.Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, slim.ErrWrongNumArguments
	}

	format, ok := args[0].(*slim.String)
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// fmt.Println(format)
		p.ss.logger().Sugar().Infof("%v: %v", p.fn.Name, format)
		return nil, nil
	}

	s, err := slim.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	// fmt.Println(s)
	p.ss.logger().Sugar().Infof("%v: %v", p.fn.Name, s)
	return nil, nil
}

func (p *Proc) userFunctionDebug(args ...slim.Object) (slim.Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, slim.ErrWrongNumArguments
	}

	if p.getDebug() != consts.ON {
		return nil, nil
	}

	format, ok := args[0].(*slim.String)
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// fmt.Println(format)
		p.ss.logger().Sugar().Debugf("%v: %v", p.fn.Name, format)
		return nil, nil
	}

	s, err := slim.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	// fmt.Println(s)
	p.ss.logger().Sugar().Debugf("%v: %v", p.fn.Name, s)
	return nil, nil
}

func (p *Proc) userFunctionExit(args ...slim.Object) (slim.Object, error) {
	if len(args) != 0 {
		return nil, slim.ErrWrongNumArguments
	}

	p.lock.RLock()
	defer p.lock.RUnlock()

	if p.exit {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}
