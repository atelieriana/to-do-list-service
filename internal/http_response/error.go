package http_response

type BadRequestResponse struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
}
