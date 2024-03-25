package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/slim"
)

type ModuleConst struct {
	ss *SlimSlot
}

func NewModuleConst(ss *SlimSlot) *ModuleConst {
	return &ModuleConst{ss: ss}
}

func (m *ModuleConst) Module() map[string]slim.Object {
	return map[string]slim.Object{
		"view": &slim.UserFunction{
			Name:  "view",
			Value: m.view,
		},
		"name": &slim.UserFunction{
			Name:  "name",
			Value: m.name,
		},
		"list": &slim.UserFunction{
			Name:  "list",
			Value: m.list,
		},
		"get_value_by_id": &slim.UserFunction{
			Name:  "get_value_by_id",
			Value: m.getValueByID,
		},
		"set_value_by_id": &slim.UserFunction{
			Name:  "set_value_by_id",
			Value: m.setValueByID,
		},
		"set_value_by_id_force": &slim.UserFunction{
			Name:  "set_value_by_id_force",
			Value: m.setValueByIDForce,
		},
		"get_value": &slim.UserFunction{
			Name:  "get_value",
			Value: m.getValueByName,
		},
		"set_value": &slim.UserFunction{
			Name:  "set_value",
			Value: m.setValueByName,
		},
		"set_value_force": &slim.UserFunction{
			Name:  "set_value_force",
			Value: m.setValueByNameForce,
		},
	}
}

func (m *ModuleConst) view(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetConst().View(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: constToSlimObject(reply)}, nil
}

func (m *ModuleConst) name(args ...slim.Object) (slim.Object, error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	name, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	reply, err := m.ss.es.GetConst().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: constToSlimObject(reply)}, nil
}

func (m *ModuleConst) list(args ...slim.Object) (slim.Object, error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	res := make(map[string]slim.Object)

	switch o := args[0].(type) {
	case *slim.Map:
		res = o.Value
	case *slim.ImmutableMap:
		res = o.Value
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	page := pb.Page{
		Limit:  10,
		Offset: 0,
		Search: "",
	}

	if v, ok := res["limit"]; ok {
		limit, ok := slim.ToInt(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "limit",
				Expected: "int(compatible)",
				Found:    v.TypeName(),
			}
		}

		page.Limit = uint32(limit)
	}

	if v, ok := res["offset"]; ok {
		offset, ok := slim.ToInt(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "offset",
				Expected: "int(compatible)",
				Found:    v.TypeName(),
			}
		}

		page.Offset = uint32(offset)
	}

	if v, ok := res["search"]; ok {
		search, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "search",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		page.Search = search
	}

	request := &edges.ConstListRequest{
		Page: &page,
		Tags: "",
		Type: "",
	}

	if v, ok := res["tags"]; ok {
		tags, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "tags",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		request.Tags = tags
	}

	if v, ok := res["type"]; ok {
		t, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "type",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		request.Type = t
	}

	reply, err := m.ss.es.GetConst().List(m.ss.ctx, request)
	if err != nil {
		return wrapError(err), nil
	}

	items := make([]slim.Object, 0, len(reply.GetConst()))
	for _, item := range reply.GetConst() {
		items = append(items, &slim.Map{Value: constToSlimObject(item)})
	}

	return &slim.Map{Value: map[string]slim.Object{
		"total": &slim.Int{Value: int64(reply.GetCount())},
		"items": &slim.Array{Value: items},
	}}, nil
}

func (m *ModuleConst) getValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetConst().GetValue(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: constValueToSlimObject(reply)}, nil
}

func (m *ModuleConst) setValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetConst().SetValue(m.ss.ctx,
		&pb.ConstValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleConst) setValueByIDForce(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetConst().SetValueForce(m.ss.ctx,
		&pb.ConstValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleConst) getValueByName(args ...slim.Object) (slim.Object, error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	name, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	reply, err := m.ss.es.GetConst().GetValueByName(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: constNameValueToSlimObject(reply)}, nil
}

func (m *ModuleConst) setValueByName(args ...slim.Object) (slim.Object, error) {
	if len(args) != 2 {
		return nil, slim.ErrWrongNumArguments
	}

	name, ok := slim.ToString(args[0])
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

	reply, err := m.ss.es.GetConst().SetValueByName(m.ss.ctx,
		&pb.ConstNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleConst) setValueByNameForce(args ...slim.Object) (slim.Object, error) {
	if len(args) != 2 {
		return nil, slim.ErrWrongNumArguments
	}

	name, ok := slim.ToString(args[0])
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

	reply, err := m.ss.es.GetConst().SetValueByNameForce(m.ss.ctx,
		&pb.ConstNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}
