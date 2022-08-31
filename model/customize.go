package model

import jsoniter "github.com/json-iterator/go"

type ListOfInt64 []int64

func (l ListOfInt64) Marshal() interface{} {
	if len(l) == 0 {
		return l
	}
	marshal, _ := jsoniter.Marshal(l)
	return string(marshal)
}
