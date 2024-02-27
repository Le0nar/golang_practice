package handler

import (
	"encoding/json"
	"net/http"

	customerror "github.com/Le0nar/golang_practice/http/internal/custom_error"
	"github.com/Le0nar/golang_practice/http/internal/event"
)

// TODO: move to antoher file
type resultResponse struct {
	result event.Event `json:"result"`
}
type errorResponse struct {
	error event.Event `json:"error"`
}

type service interface {
	CreateEvent(dto event.EventDto) (*event.Event, *customerror.CustomError)
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRouter() http.Handler {
	mux := http.NewServeMux()

	// POST
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)

	// GET
	mux.HandleFunc("/events_for_day", h.getEventForDay)
	mux.HandleFunc("/events_for_week", h.getEventForWeek)
	mux.HandleFunc("/events_for_month", h.getEventForMonth)

	return mux
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: mb add check for r.Method == "POST"
	var dto event.EventDto

	// TODO: add validation for requeried fields
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, customErr := h.service.CreateEvent(dto)

	// TODO: вернуть ответ 503 или 500

	// TODO: вернуть ответ 200
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventForDay(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventForWeek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventForMonth(w http.ResponseWriter, r *http.Request) {

}
