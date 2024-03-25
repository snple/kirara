package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/slim"
)

type ModuleSource struct {
	ss *SlimSlot
}

func NewModuleSource(ss *SlimSlot) *ModuleSource {
	return &ModuleSource{ss: ss}
}

func (m *ModuleSource) Module() map[string]slim.Object {
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
		"link": &slim.UserFunction{
			Name:  "link",
			Value: m.link,
		},
	}
}

func (m *ModuleSource) view(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetSource().View(m.ss.ctx, &pb.Id{Id: id})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: sourceToSlimObject(reply)}, nil
}

func (m *ModuleSource) name(args ...slim.Object) (slim.Object, error) {
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

	reply, err := m.ss.es.GetSource().Name(m.ss.ctx, &pb.Name{Name: name})
	if err != nil {
		return wrapError(err), nil
	}

	return &slim.Map{Value: sourceToSlimObject(reply)}, nil
}

func (m *ModuleSource) list(args ...slim.Object) (slim.Object, error) {
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

	request := &edges.SourceListRequest{
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

	if v, ok := res["source"]; ok {
		source, ok := slim.ToString(v)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     "source",
				Expected: "string(compatible)",
				Found:    v.TypeName(),
			}
		}

		request.Source = source
	}

	reply, err := m.ss.es.GetSource().List(m.ss.ctx, request)
	if err != nil {
		return wrapError(err), nil
	}

	items := make([]slim.Object, 0, len(reply.GetSource()))
	for _, item := range reply.GetSource() {
		items = append(items, &slim.Map{Value: sourceToSlimObject(item)})
	}

	return &slim.Map{Value: map[string]slim.Object{
		"total": &slim.Int{Value: int64(reply.GetCount())},
		"items": &slim.Array{Value: items},
	}}, nil
}

func (m *ModuleSource) link(args ...slim.Object) (slim.Object, error) {
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

	status, ok := slim.ToInt(args[1])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	reply, err := m.ss.es.GetSource().Link(m.ss.ctx,
		&edges.SourceLinkRequest{Id: id, Status: int32(status)})
	if err != nil {
		return wrapError(err), nil
	}

	if reply.GetBool() {
		return slim.TrueValue, nil
	}

	return slim.FalseValue, nil
}
