package http_common

type HttpResponse[T any] struct {
	Success bool  `json:"Success"`
	Data    *T    `json:"data"`
	Error   Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewErrorResponse(error Error) HttpResponse[any] {
	return HttpResponse[any]{
		Success: false,
		Error:   error,
	}
}

func NewSuccessResponse[T any](data *T) HttpResponse[T] {
	return HttpResponse[T]{
		Success: true,
		Data:    data,
	}
}
