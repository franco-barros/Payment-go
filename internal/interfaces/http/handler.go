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

// Helper para respuestas de error en JSON
func writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

// Helper para respuestas exitosas
func writeJSON(w http.ResponseWriter, data any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}

// Handler principal
func (h *PaymentHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Método permitido
	if r.Method != http.MethodPost {
		writeError(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if req.Method == "" {
		writeError(w, "el método es requerido", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		writeError(w, "el monto debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	// Obtener método desde factory
	method, err := h.factory.Get(req.Method)
	if err != nil {
		writeError(w, "método de pago no soportado", http.StatusBadRequest)
		return
	}

	// Procesar pago
	if err := h.service.Process(method, req.Amount); err != nil {
		// Acá podrías mapear errores de dominio más adelante
		writeError(w, "error al procesar el pago", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	writeJSON(w, map[string]string{
		"message": "pago procesado correctamente",
	}, http.StatusOK)
}