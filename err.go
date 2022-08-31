package SimqianchuanSDK

import (
	"errors"
	"fmt"
)

type Unite struct {
	QCError
	Data interface{} `json:"data"`
}

// QCError 错误结构体
type QCError struct {
	Code      int64  `json:"code"`                 // 错误码
	Message   string `json:"message"`              // 错误码描述
	RequestId string `json:"request_id,omitempty"` // 错误码描述r
}

func (e *QCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func (e *QCError) NewError() error {
	return errors.New(fmt.Sprintf("%d: %s", e.Code, e.Message))
}
