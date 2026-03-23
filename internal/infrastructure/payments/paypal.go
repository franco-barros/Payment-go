package payments

import (
	"errors"
	"fmt"

	"payment-system/internal/domain"
)

type PayPal struct {
	Email string
}

var _ domain.PaymentMethod = (*PayPal)(nil)

func (p PayPal) Pay(amount float64) error {
	// 🔴 Validaciones
	if p.Email == "" {
		return errors.New("paypal: email requerido")
	}

	if amount <= 0 {
		return errors.New("paypal: monto inválido")
	}

	fmt.Printf("Pagando %.2f con PayPal (%s)\n", amount, p.Email)
	return nil
}

func (p PayPal) Name() string {
	return "paypal"
}