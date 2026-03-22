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
	if s.Client == nil {
		return fmt.Errorf("stripe client no inicializado")
	}

	fmt.Printf("Procesando pago con Stripe: %.2f\n", amount)

	if err := s.Client.Charge(amount); err != nil {
		return fmt.Errorf("error al procesar pago con Stripe: %w", err)
	}

	return nil
}

func (s Stripe) Name() string {
	return "stripe"
}