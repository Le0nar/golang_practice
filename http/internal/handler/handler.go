package handler

import (
	"encoding/json"
	"net/http"

	customerror "github.com/Le0nar/golang_practice/http/internal/custom_error"
	"github.com/Le0nar/golang_practice/http/internal/event"
	"github.com/Le0nar/golang_practice/http/internal/logging"
	"github.com/Le0nar/golang_practice/http/internal/response"
)

// TODO: move to antoher file

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

	// TODO: add midleware logger for handler function
	// POST
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)

	// GET
	mux.HandleFunc("/events_for_day", h.getEventForDay)
	mux.HandleFunc("/events_for_week", h.getEventForWeek)
	mux.HandleFunc("/events_for_month", h.getEventForMonth)

	// adding middleware
	router := logging.Logging(mux)

	return router
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: mb add check for r.Method == "POST"
	var dto event.EventDto

	// TODO: add validation for requeried fields
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dto)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errorData, _ := json.Marshal(response.ErrorResponse{Error: err.Error()})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorData)
		return
	}

	e, customErr := h.service.CreateEvent(dto)

	if customErr != nil {
		errorData, _ := json.Marshal(response.ErrorResponse{Error: customErr.Err.Error()})

		w.WriteHeader(customErr.StatusCode)
		w.Write(errorData)
		return
	}

	successData, _ := json.Marshal(*e)

	w.Write(successData)
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
