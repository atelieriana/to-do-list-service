package http_response

type SingleResponse struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
}

type ArrayResponse[T any] struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
	Data     []T    `json:"data"`
}

type EmptyResponse struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
}
