package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/slim"
)

type ModuleDevice struct {
	ss *SlimSlot
}

func NewModuleDevice(ss *SlimSlot) *ModuleDevice {
	return &ModuleDevice{ss: ss}
}

func (m *ModuleDevice) Module() map[string]slim.Object {
	return map[string]slim.Object{
		"view": &slim.UserFunction{
			Name:  "view",
			Value: m.view,
		},
	}
}

func (m *ModuleDevice) view(args ...slim.Object) (slim.Object, error) {
	if len(args) != 0 {
		return nil, slim.ErrWrongNumArguments
	}

	reply, err := m.ss.es.GetDevice().View(m.ss.ctx, &pb.MyEmpty{})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: deviceToSlimObject(reply)}, nil
}
