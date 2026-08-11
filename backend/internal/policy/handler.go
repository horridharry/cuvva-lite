package policy

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/horridharry/cuvva-lite/backend/internal/quote"
)

type PurchaseRequest struct {
	PaymentMethod string `json:"paymentMethod"`
}

type Handler struct {
	quoteRepository *quote.Repository
	service         *Service
}

func NewHandler(
	quoteRepository *quote.Repository,
	service *Service,
) *Handler {
	return &Handler{
		quoteRepository: quoteRepository,
		service:         service,
	}
}

func (h *Handler) Purchase(
	w http.ResponseWriter,
	r *http.Request,
) {
	quoteID, err := strconv.ParseInt(
		r.PathValue("id"),
		10,
		64,
	)

	if err != nil {
		http.Error(
			w,
			"invalid quote id",
			http.StatusBadRequest,
		)
		return
	}

	log.Printf("Purchasing quote id: %d", quoteID)

	var request PurchaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	q, err := h.quoteRepository.FindByID(
		r.Context(),
		quoteID,
	)

	if errors.Is(err, quote.ErrQuoteExpired) {
		http.Error(
			w,
			"quote not found",
			http.StatusNotFound,
		)
		return
	}

	if err != nil {
		log.Printf("failed to find quote: %v", err)

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	p, err := h.service.CreateFromQuote(
		r.Context(),
		q,
		request.PaymentMethod,
	)

	if errors.Is(err, ErrQuoteExpired) {
		http.Error(
			w,
			"quote expired",
			http.StatusConflict,
		)
		return
	}

	if errors.Is(err, ErrAlreadyPurchased) {
		http.Error(
			w,
			"quote already purchased",
			http.StatusConflict,
		)
		return
	}

	if errors.Is(err, ErrPaymentDeclined) {
		http.Error(
			w,
			"payment declined",
			http.StatusPaymentRequired,
		)
		return
	}

	if err != nil {
		log.Printf("failed to create policy: %v", err)

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("failed to encode policy: %v", err)
	}
}
