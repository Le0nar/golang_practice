package response

import "github.com/Le0nar/golang_practice/http/internal/event"

type ResultResponse struct {
	Result event.Event `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
