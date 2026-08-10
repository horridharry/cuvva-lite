package quote

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/horridharry/cuvva-lite/backend/internal/vehicle"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	q, err := h.service.Create(
		r.Context(),
		request,
	)

	if errors.Is(err, InvalidQuoteRequest) {
		http.Error(
			w,
			"Invalid quote request",
			http.StatusBadRequest,
		)
		return
	}

	if errors.Is(err, vehicle.VehicleNotFound) {
		http.Error(
			w,
			"Vehicle not found",
			http.StatusNotFound,
		)
		return
	}

	if err != nil {
		log.Printf("Failed to create quote: %v", err)

		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(q); err != nil {
		return
	}
}
