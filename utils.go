package dgraphland

import (
	"context"
	"reflect"
	"strings"
)

const (
	TagDgland = "dgland"
	TagJson   = "json"
)

func ctx() context.Context {
	return context.Background()
}

func Schema(model interface{}) string {
	t := reflect.TypeOf(model)

	if t.NumField() == 0 {
		return ""
	}

	var strs []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValueDglang := field.Tag.Get(TagDgland)
		tagValueJson := field.Tag.Get(TagJson)
		if tagValueDglang != "" && tagValueJson != "" {
			strs = append(strs, field.Name+" "+tagValueDglang+" .")
		}
	}
	return strings.Join(strs, "\n")
}
