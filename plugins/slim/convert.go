package slim

import (
	"github.com/snple/kirara/pb"
	"github.com/snple/slim"
)

func wrapError(err error) slim.Object {
	if err == nil {
		return slim.TrueValue
	}
	return &slim.Error{Value: &slim.String{Value: err.Error()}}
}

func fnToSlimObject(src *pb.Fn) map[string]slim.Object {
	return map[string]slim.Object{
		"id":     &slim.String{Value: src.GetId()},
		"name":   &slim.String{Value: src.GetName()},
		"desc":   &slim.String{Value: src.GetDesc()},
		"tags":   &slim.String{Value: src.GetTags()},
		"type":   &slim.String{Value: src.GetType()},
		"config": &slim.String{Value: src.GetConfig()},
		"link":   &slim.Int{Value: int64(src.GetLink())},
		"status": &slim.Int{Value: int64(src.GetStatus())},
		"debug":  &slim.Int{Value: int64(src.GetDebug())},
	}
}

func logicToSlimObject(src *pb.Logic) map[string]slim.Object {
	return map[string]slim.Object{
		"id":     &slim.String{Value: src.GetId()},
		"name":   &slim.String{Value: src.GetName()},
		"desc":   &slim.String{Value: src.GetDesc()},
		"tags":   &slim.String{Value: src.GetTags()},
		"type":   &slim.String{Value: src.GetType()},
		"status": &slim.Int{Value: int64(src.GetStatus())},
	}
}

func deviceToSlimObject(src *pb.Device) map[string]slim.Object {
	return map[string]slim.Object{
		"id":       &slim.String{Value: src.GetId()},
		"name":     &slim.String{Value: src.GetName()},
		"desc":     &slim.String{Value: src.GetDesc()},
		"tags":     &slim.String{Value: src.GetTags()},
		"type":     &slim.String{Value: src.GetType()},
		"location": &slim.String{Value: src.GetLocation()},
		"config":   &slim.String{Value: src.GetConfig()},
		"link":     &slim.Int{Value: int64(src.GetLink())},
		"status":   &slim.Int{Value: int64(src.GetStatus())},
	}
}

func slotToSlimObject(src *pb.Slot) map[string]slim.Object {
	return map[string]slim.Object{
		"id":       &slim.String{Value: src.GetId()},
		"name":     &slim.String{Value: src.GetName()},
		"desc":     &slim.String{Value: src.GetDesc()},
		"tags":     &slim.String{Value: src.GetTags()},
		"type":     &slim.String{Value: src.GetType()},
		"location": &slim.String{Value: src.GetLocation()},
		"config":   &slim.String{Value: src.GetConfig()},
		"link":     &slim.Int{Value: int64(src.GetLink())},
		"status":   &slim.Int{Value: int64(src.GetStatus())},
	}
}

func sourceToSlimObject(src *pb.Source) map[string]slim.Object {
	return map[string]slim.Object{
		"id":     &slim.String{Value: src.GetId()},
		"name":   &slim.String{Value: src.GetName()},
		"desc":   &slim.String{Value: src.GetDesc()},
		"tags":   &slim.String{Value: src.GetTags()},
		"type":   &slim.String{Value: src.GetType()},
		"source": &slim.String{Value: src.GetSource()},
		"params": &slim.String{Value: src.GetParams()},
		"config": &slim.String{Value: src.GetConfig()},
		"link":   &slim.Int{Value: int64(src.GetLink())},
		"status": &slim.Int{Value: int64(src.GetStatus())},
		"save":   &slim.Int{Value: int64(src.GetSave())},
	}
}

func tagToSlimObject(src *pb.Tag) map[string]slim.Object {
	return map[string]slim.Object{
		"id":        &slim.String{Value: src.GetId()},
		"name":      &slim.String{Value: src.GetName()},
		"desc":      &slim.String{Value: src.GetDesc()},
		"tags":      &slim.String{Value: src.GetTags()},
		"type":      &slim.String{Value: src.GetType()},
		"data_type": &slim.String{Value: src.GetDataType()},
		"address":   &slim.String{Value: src.GetAddress()},
		"value":     &slim.String{Value: src.GetValue()},
		"h_value":   &slim.String{Value: src.GetHValue()},
		"l_value":   &slim.String{Value: src.GetLValue()},
		"config":    &slim.String{Value: src.GetConfig()},
		"status":    &slim.Int{Value: int64(src.GetStatus())},
		"access":    &slim.Int{Value: int64(src.GetAccess())},
		"save":      &slim.Int{Value: int64(src.GetSave())},
	}
}

func constToSlimObject(src *pb.Const) map[string]slim.Object {
	return map[string]slim.Object{
		"id":        &slim.String{Value: src.GetId()},
		"name":      &slim.String{Value: src.GetName()},
		"desc":      &slim.String{Value: src.GetDesc()},
		"tags":      &slim.String{Value: src.GetTags()},
		"type":      &slim.String{Value: src.GetType()},
		"data_type": &slim.String{Value: src.GetDataType()},
		"value":     &slim.String{Value: src.GetValue()},
		"h_value":   &slim.String{Value: src.GetHValue()},
		"l_value":   &slim.String{Value: src.GetLValue()},
		"config":    &slim.String{Value: src.GetConfig()},
		"status":    &slim.Int{Value: int64(src.GetStatus())},
		"access":    &slim.Int{Value: int64(src.GetAccess())},
	}
}

func classToSlimObject(src *pb.Class) map[string]slim.Object {
	return map[string]slim.Object{
		"id":     &slim.String{Value: src.GetId()},
		"name":   &slim.String{Value: src.GetName()},
		"desc":   &slim.String{Value: src.GetDesc()},
		"tags":   &slim.String{Value: src.GetTags()},
		"type":   &slim.String{Value: src.GetType()},
		"config": &slim.String{Value: src.GetConfig()},
		"status": &slim.Int{Value: int64(src.GetStatus())},
		"save":   &slim.Int{Value: int64(src.GetSave())},
	}
}

func attrToSlimObject(src *pb.Attr) map[string]slim.Object {
	return map[string]slim.Object{
		"id":        &slim.String{Value: src.GetId()},
		"name":      &slim.String{Value: src.GetName()},
		"desc":      &slim.String{Value: src.GetDesc()},
		"tags":      &slim.String{Value: src.GetTags()},
		"type":      &slim.String{Value: src.GetType()},
		"data_type": &slim.String{Value: src.GetDataType()},
		"value":     &slim.String{Value: src.GetValue()},
		"h_value":   &slim.String{Value: src.GetHValue()},
		"l_value":   &slim.String{Value: src.GetLValue()},
		"config":    &slim.String{Value: src.GetConfig()},
		"status":    &slim.Int{Value: int64(src.GetStatus())},
		"access":    &slim.Int{Value: int64(src.GetAccess())},
		"save":      &slim.Int{Value: int64(src.GetSave())},
	}
}

func optionToSlimObject(src *pb.Option) map[string]slim.Object {
	return map[string]slim.Object{
		"id":     &slim.String{Value: src.GetId()},
		"name":   &slim.String{Value: src.GetName()},
		"desc":   &slim.String{Value: src.GetDesc()},
		"tags":   &slim.String{Value: src.GetTags()},
		"type":   &slim.String{Value: src.GetType()},
		"status": &slim.Int{Value: int64(src.GetStatus())},
	}
}

func tagValueToSlimObject(src *pb.TagValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}

func tagNameValueToSlimObject(src *pb.TagNameValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"name":    &slim.String{Value: src.GetName()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}

func constValueToSlimObject(src *pb.ConstValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}

func constNameValueToSlimObject(src *pb.ConstNameValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"name":    &slim.String{Value: src.GetName()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}

func attrValueToSlimObject(src *pb.AttrValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}

func attrNameValueToSlimObject(src *pb.AttrNameValue) map[string]slim.Object {
	return map[string]slim.Object{
		"id":      &slim.String{Value: src.GetId()},
		"name":    &slim.String{Value: src.GetName()},
		"value":   &slim.String{Value: src.GetValue()},
		"updated": &slim.Int{Value: src.GetUpdated() / 1000000},
	}
}
