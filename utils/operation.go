package utils

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

func Min(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

// BuildQueryToMap 把结构体或结构体指针转成请求query params的map，忽略结构体中为空的字段，以及把列表变成字符串格式
func BuildQueryToMap(params interface{}) map[string]interface{} {
	queryMap := make(map[string]interface{})
	types := reflect.TypeOf(params)
	values := reflect.ValueOf(params)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
		values = values.Elem()
	}
	for num := 0; num < types.NumField(); num++ {
		tag := types.Field(num).Tag.Get("json")
		title := strings.Split(tag, ",")
		value := values.Field(num)
		if value.IsZero() && strings.Contains(tag, "omitempty") {
			continue
		}
		switch value.Kind() {
		case reflect.Array, reflect.Slice, reflect.Struct:
			marshal, _ := jsoniter.Marshal(value.Interface())
			queryMap[title[0]] = string(marshal)
		default:
			queryMap[title[0]] = value.Interface()
		}
	}
	return queryMap
}
