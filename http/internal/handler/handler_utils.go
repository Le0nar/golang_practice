package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Le0nar/golang_practice/http/internal/event"
)

// TODO: mb rewrite with generics for use for several dto structs
// func [T any] serializeDto (someStr T) ([]byte, err)  {
func serializeCreateEventDto(r *http.Request) (*event.CreateEventDto, error) {
	var dto event.CreateEventDto

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dto)
	if err != nil {
		return nil, err
	}

	// Сомнительный момент. Обычно возвращают 404 когда не совпадает метод (тут мы возвращаем 400)
	if r.Method != "POST" {
		return nil, errors.New("wrong method")
	}

	return &dto, nil
}
