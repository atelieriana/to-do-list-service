package http_response

type BadRequestResponse struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
}

type ValidationErrorResponse struct {
	HttpCode int         `json:"code"`
	Message  interface{} `json:"message"`
}
