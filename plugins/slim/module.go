package slim

import (
	"errors"
	"fmt"
	"sync"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/pb"
	"github.com/snple/slim"
	"github.com/snple/slim/stdlib"
)

type DynamicModules struct {
	mods     slim.ModuleGetter
	fallback func(name string) slim.Importable
}

func (dm *DynamicModules) Get(name string) slim.Importable {
	if mod := dm.mods.Get(name); mod != nil {
		return mod
	}
	return dm.fallback(name)
}

type SourceModule struct {
	ss   *SlimSlot
	fnID string
	Name string
}

// Import returns a module source code.
func (m *SourceModule) Import(name string) (interface{}, error) {
	reply, err := m.ss.es.GetLogic().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return nil, err
	}

	if reply.GetExec() != EXEC {
		return nil, errors.New("exec is not slim")
	}

	if reply.GetStatus() != consts.ON {
		return nil, fmt.Errorf("logic %v is not enable", name)
	}

	m.ss.moduleDpds.Add(reply.GetId(), m.fnID)

	return []byte(reply.GetMain()), nil
}

func NewDynamicModules(ss *SlimSlot, fnID string) *DynamicModules {
	moduleMap := slim.NewModuleMap()

	moduleMap.AddMap(stdlib.GetModuleMap(stdlib.AllModuleNames()...))

	moduleMap.AddBuiltinModule("store", ss.store.Module())
	moduleMap.AddSourceModule("cache", []byte(`export func() { store := import("store"); return store.cache }`))

	moduleMap.AddBuiltinModule("device", ss.moduleDevice.Module())
	moduleMap.AddBuiltinModule("source", ss.moduleSource.Module())
	moduleMap.AddBuiltinModule("tag", ss.moduleTag.Module())
	moduleMap.AddBuiltinModule("const", ss.moduleConst.Module())
	moduleMap.AddBuiltinModule("class", ss.moduleClass.Module())
	moduleMap.AddBuiltinModule("attr", ss.moduleAttr.Module())
	moduleMap.AddBuiltinModule("control", ss.moduleControl.Module())

	modules := &DynamicModules{
		mods: moduleMap,
		fallback: func(name string) slim.Importable {
			return &SourceModule{ss: ss, fnID: fnID, Name: name}
		},
	}

	return modules
}

type ModuleDpds struct {
	lock sync.Mutex
	maps map[string]map[string]struct{}
}

func NewModuleDpds() *ModuleDpds {
	return &ModuleDpds{
		maps: make(map[string]map[string]struct{}),
	}
}

func (m *ModuleDpds) Add(logicId, procId string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if ids, ok := m.maps[logicId]; ok {
		ids[procId] = struct{}{}
	} else {
		m.maps[logicId] = map[string]struct{}{procId: {}}
	}
}

func (m *ModuleDpds) Take(logicId string) map[string]struct{} {
	m.lock.Lock()
	defer m.lock.Unlock()

	if ids, ok := m.maps[logicId]; ok {
		delete(m.maps, logicId)
		return ids
	}

	return map[string]struct{}{}
}
