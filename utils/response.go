package utils

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
