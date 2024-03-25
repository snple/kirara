package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/slim"
)

type ModuleClass struct {
	ss *SlimSlot
}

func NewModuleClass(ss *SlimSlot) *ModuleClass {
	return &ModuleClass{ss: ss}
}

func (m *ModuleClass) Module() map[string]slim.Object {
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
	}
}

func (m *ModuleClass) view(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetClass().View(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: classToSlimObject(reply)}, nil
}

func (m *ModuleClass) name(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetClass().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: classToSlimObject(reply)}, nil
}

func (m *ModuleClass) list(args ...slim.Object) (slim.Object, error) {
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

	request := &edges.ClassListRequest{
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

	reply, err := m.ss.es.GetClass().List(m.ss.ctx, request)
	if err != nil {
		return wrapError(err), nil
	}

	items := make([]slim.Object, 0, len(reply.GetClass()))
	for _, item := range reply.GetClass() {
		items = append(items, &slim.Map{Value: classToSlimObject(item)})
	}

	return &slim.Map{Value: map[string]slim.Object{
		"total": &slim.Int{Value: int64(reply.GetCount())},
		"items": &slim.Array{Value: items},
	}}, nil
}
