package payments

import (
	"errors"
	"fmt"
)

type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) error {
	if p.Email == "" {
		return errors.New("email requerido para PayPal")
	}

	fmt.Printf("Pagando %.2f con PayPal (%s)\n", amount, p.Email)
	return nil
}

func (p PayPal) Name() string {
	return "paypal"
}

