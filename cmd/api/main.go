package main

import (
	"log"
	nethttp "net/http"

	app "payment-system/internal/application"
	"payment-system/internal/domain"
	infra "payment-system/internal/infrastructure/payments"
	handler "payment-system/internal/interfaces/http"
)

// Mock de StripeClient
type MockStripeClient struct{}

func (m MockStripeClient) Charge(amount float64) error {
	return nil
}

func main() {
	// 🔹 Infraestructura
	stripe := infra.Stripe{
		Client: MockStripeClient{},
	}

	paypal := infra.PayPal{
		Email: "test@paypal.com",
	}

	crypto := infra.Crypto{
		WalletAddress: "wallet123",
	}

	creditCard := infra.CreditCard{
		Holder: "Franco",
		Number: "1234-5678-9999",
	}

	// 🔹 Registro de métodos
	methods := map[string]domain.PaymentMethod{
		"stripe":      stripe,
		"paypal":      paypal,
		"crypto":      crypto,
		"credit_card": creditCard,
	}

	// 🔹 Factory
	factory := app.NewPaymentFactory(methods)

	// 🔹 Service
	service := &app.PaymentService{}

	// 🔹 Handler
	paymentHandler := handler.NewPaymentHandler(service, factory)

	// 🔹 Router
	mux := nethttp.NewServeMux()
	mux.HandleFunc("/pay", paymentHandler.Handle)

	log.Println("Servidor corriendo en :8080")
	if err := nethttp.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}