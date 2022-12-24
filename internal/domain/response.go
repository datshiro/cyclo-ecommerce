package domain

type GetResponse struct {
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Error   ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Message interface{} `json:"message"`
}
