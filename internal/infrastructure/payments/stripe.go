package payments

import (
	"fmt"

	"payment-system/internal/domain"
)

type StripeClient interface {
	Charge(amount float64) error
}

type Stripe struct {
	Client StripeClient
}

var _ domain.PaymentMethod = (*Stripe)(nil)

func (s Stripe) Pay(amount float64) error {
	// 🔴 Validación básica
	if amount <= 0 {
		return fmt.Errorf("stripe: monto inválido")
	}

	// 🔴 Validación de dependencia
	if s.Client == nil {
		return fmt.Errorf("stripe: client no inicializado")
	}

	fmt.Printf("Procesando pago con Stripe: %.2f\n", amount)

	// 🔴 Error externo envuelto
	if err := s.Client.Charge(amount); err != nil {
		return fmt.Errorf("stripe: error al procesar pago: %w", err)
	}

	return nil
}

func (s Stripe) Name() string {
	return "stripe"
}