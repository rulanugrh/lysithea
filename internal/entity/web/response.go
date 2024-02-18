package web

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewSuccessResponse(message string, data any) error {
	return &Response{
		Code:    200,
		Message: message,
		Result:  data,
	}
}

func NewCreatedResponse(msg string, data any) error {
	return &Response{
		Code:    201,
		Message: msg,
		Result:  data,
	}
}

func NewUnauthorizedResponse(msg string) error {
	return &Response{
		Code:    401,
		Message: msg,
	}
}

func NewForbiddenResponse(msg string) error {
	return &Response{
		Code:    403,
		Message: msg,
	}
}

func NewInternalServerErrorResponse(msg string) error {
	return &Response{
		Code:    500,
		Message: msg,
	}
}

func NewStatusNotFound(msg string) error {
	return &Response{
		Code:    404,
		Message: msg,
	}
}
