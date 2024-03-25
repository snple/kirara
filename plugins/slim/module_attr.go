package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/slim"
)

type ModuleAttr struct {
	ss *SlimSlot
}

func NewModuleAttr(ss *SlimSlot) *ModuleAttr {
	return &ModuleAttr{ss: ss}
}

func (m *ModuleAttr) Module() map[string]slim.Object {
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

func (m *ModuleAttr) view(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().View(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: attrToSlimObject(reply)}, nil
}

func (m *ModuleAttr) name(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: attrToSlimObject(reply)}, nil
}

func (m *ModuleAttr) list(args ...slim.Object) (slim.Object, error) {
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

	request := &edges.AttrListRequest{
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

	if v, ok := res["class_id"]; ok {
		class, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "class",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		request.ClassId = class
	}

	reply, err := m.ss.es.GetAttr().List(m.ss.ctx, request)
	if err != nil {
		return wrapError(err), nil
	}

	items := make([]slim.Object, 0, len(reply.GetAttr()))
	for _, item := range reply.GetAttr() {
		items = append(items, &slim.Map{Value: attrToSlimObject(item)})
	}

	return &slim.Map{Value: map[string]slim.Object{
		"total": &slim.Int{Value: int64(reply.GetCount())},
		"items": &slim.Array{Value: items},
	}}, nil
}

func (m *ModuleAttr) getValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().GetValue(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: attrValueToSlimObject(reply)}, nil
}

func (m *ModuleAttr) setValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().SetValue(m.ss.ctx,
		&pb.AttrValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleAttr) setValueByIDForce(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().SetValueForce(m.ss.ctx,
		&pb.AttrValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleAttr) getValueByName(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().GetValueByName(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: attrNameValueToSlimObject(reply)}, nil
}

func (m *ModuleAttr) setValueByName(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().SetValueByName(m.ss.ctx,
		&pb.AttrNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleAttr) setValueByNameForce(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetAttr().SetValueByNameForce(m.ss.ctx,
		&pb.AttrNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}
