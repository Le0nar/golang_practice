package service

import (
	"errors"
	"net/http"
	"time"

	customerror "github.com/Le0nar/golang_practice/http/internal/custom_error"
	"github.com/Le0nar/golang_practice/http/internal/event"
)

type repository interface {
	CreateEvent(e event.Event) error
}

type Serivce struct {
	repository repository
}

func NewService(r repository) *Serivce {
	return &Serivce{repository: r}
}

func (s *Serivce) CreateEvent(dto event.EventDto) (*event.Event, *customerror.CustomError) {

	eventId := time.Now().String()
	e := event.Event{UserId: dto.UserId, Date: dto.Date, Content: dto.Content, Id: eventId}

	charLen := len([]rune(e.Content))

	if charLen > 1000 {
		return nil, &customerror.CustomError{StatusCode: http.StatusServiceUnavailable, Err: errors.New("Value has more than 1000 characters")}
	}

	err := s.repository.CreateEvent(e)
	if err != nil {
		return nil, &customerror.CustomError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return &e, nil
}
