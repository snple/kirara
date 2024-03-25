package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/slim"
)

type ModuleControl struct {
	ss *SlimSlot
}

func NewModuleControl(ss *SlimSlot) *ModuleControl {
	return &ModuleControl{ss: ss}
}

func (m *ModuleControl) Module() map[string]slim.Object {
	return map[string]slim.Object{
		"get_tag_value": &slim.UserFunction{
			Name:  "get_tag_value",
			Value: m.getTagValue,
		},
		"set_tag_value": &slim.UserFunction{
			Name:  "set_tag_value",
			Value: m.setTagValue,
		},
	}
}

func (m *ModuleControl) getTagValue(args ...slim.Object) (slim.Object, error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	id, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	reply, err := m.ss.es.GetControl().GetTagValue(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagValueToSlimObject(reply)}, nil
}

func (m *ModuleControl) setTagValue(args ...slim.Object) (slim.Object, error) {
	if len(args) != 2 {
		return nil, slim.ErrWrongNumArguments
	}

	id, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	value, ok := slim.ToString(args[1])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	reply, err := m.ss.es.GetControl().SetTagValue(m.ss.ctx, &pb.TagValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagValueToSlimObject(reply)}, nil
}
