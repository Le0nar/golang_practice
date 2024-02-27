package repository

import (
	"fmt"

	"github.com/Le0nar/golang_practice/http/internal/event"
)

type Repository struct {
	EventList []event.Event
}

func NewRepository() *Repository {
	return &Repository{EventList: make([]event.Event, 0)}
}

func (r *Repository) CreateEvent(e event.Event) error {
	r.EventList = append(r.EventList, e)

	fmt.Printf("r.EventList: %v\n", r.EventList)

	return nil
}
