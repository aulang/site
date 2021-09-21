package model

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

func FailWithCodeAndError(code int, err error) Response {
	return Response{
		Code: code,
		Msg:  err.Error(),
	}
}

func FailWithError(err error) Response {
	return Response{
		Code: 500,
		Msg:  err.Error(),
	}
}
