package response

type ErrorResponse struct {
	Type   string `json:"type"`
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}
