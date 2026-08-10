package vehicle

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (handler *Handler) GetByRegistration(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
) {
	registration := httpRequest.PathValue("registration")

	vehicle, repositoryError := handler.repository.FindByRegistration(
		httpRequest.Context(),
		registration,
	)

	if errors.Is(repositoryError, VehicleNotFound) {
		http.Error(responseWriter, "Vehicle not found", http.StatusNotFound)
		return
	}

	if repositoryError != nil {
		http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	if encoderError := json.NewEncoder(responseWriter).Encode(vehicle); encoderError != nil {
		http.Error(responseWriter, "Failed to encode the response", http.StatusInternalServerError)
	}
}
