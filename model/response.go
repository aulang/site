package model

import "log"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success() Response {
	return Response{
		Code: 0,
		Data: "success",
	}
}

func SuccessWithData(data interface{}) Response {
	return Response{
		Code: 0,
		Data: data,
	}
}

func Fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

func FailWithError(err error) Response {
	log.Printf("接口调用失败，%v", err)
	return Response{
		Code: -1,
		Msg:  err.Error(),
	}
}
