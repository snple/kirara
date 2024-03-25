package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/slim"
)

type ModuleTag struct {
	ss *SlimSlot
}

func NewModuleTag(ss *SlimSlot) *ModuleTag {
	return &ModuleTag{ss: ss}
}

func (m *ModuleTag) Module() map[string]slim.Object {
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

func (m *ModuleTag) view(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().View(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagToSlimObject(reply)}, nil
}

func (m *ModuleTag) name(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagToSlimObject(reply)}, nil
}

func (m *ModuleTag) list(args ...slim.Object) (slim.Object, error) {
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

	request := &edges.TagListRequest{
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

	if v, ok := res["source_id"]; ok {
		source, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "source",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		request.SourceId = source
	}

	reply, err := m.ss.es.GetTag().List(m.ss.ctx, request)
	if err != nil {
		return wrapError(err), nil
	}

	items := make([]slim.Object, 0, len(reply.GetTag()))
	for _, item := range reply.GetTag() {
		items = append(items, &slim.Map{Value: tagToSlimObject(item)})
	}

	return &slim.Map{Value: map[string]slim.Object{
		"total": &slim.Int{Value: int64(reply.GetCount())},
		"items": &slim.Array{Value: items},
	}}, nil
}

func (m *ModuleTag) getValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().GetValue(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagValueToSlimObject(reply)}, nil
}

func (m *ModuleTag) setValueByID(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().SetValue(m.ss.ctx,
		&pb.TagValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleTag) setValueByIDForce(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().SetValueForce(m.ss.ctx,
		&pb.TagValue{Id: id, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleTag) getValueByName(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().GetValueByName(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: tagNameValueToSlimObject(reply)}, nil
}

func (m *ModuleTag) setValueByName(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().SetValueByName(m.ss.ctx,
		&pb.TagNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}

func (m *ModuleTag) setValueByNameForce(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetTag().SetValueByNameForce(m.ss.ctx,
		&pb.TagNameValue{Name: name, Value: value})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}
