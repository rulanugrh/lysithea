package web

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func (r *Response) Error() string {
	return r.Message
}

func Success(message string, data any) error {
	return &Response{
		Code:    200,
		Message: message,
		Result:  data,
	}
}

func Created(msg string, data any) error {
	return &Response{
		Code:    201,
		Message: msg,
		Result:  data,
	}
}

func Unauthorized(msg string) error {
	return &Response{
		Code:    401,
		Message: msg,
	}
}

func Forbidden(msg string) error {
	return &Response{
		Code:    403,
		Message: msg,
	}
}

func InternalServerError(msg string) error {
	return &Response{
		Code:    500,
		Message: msg,
	}
}

func StatusNotFound(msg string) error {
	return &Response{
		Code:    404,
		Message: msg,
	}
}

func StatusBadRequest(msg string) error {
	return &Response{
		Code:    400,
		Message: msg,
	}
}
