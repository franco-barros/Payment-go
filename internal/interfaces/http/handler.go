package http

import (
	"encoding/json"
	"net/http"

	"payment-system/internal/application"
)

type PaymentHandler struct {
	service *application.PaymentService
	factory *application.PaymentFactory
}

// Constructor (inyección de dependencias)
func NewPaymentHandler(
	service *application.PaymentService,
	factory *application.PaymentFactory,
) *PaymentHandler {
	return &PaymentHandler{
		service: service,
		factory: factory,
	}
}

// Request DTO
type PaymentRequest struct {
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
}

// Handler principal
func (h *PaymentHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// 🔥 Ahora usamos la factory
	method, err := h.factory.Get(req.Method)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Process(method, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Pago procesado correctamente"))
}