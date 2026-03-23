package payments

import (
	"errors"
	"fmt"
)

type CreditCard struct {
	Holder string
	Number string
}
 
func (c CreditCard) Pay(amount float64) error {
	if amount > 1000 {
		return errors.New("monto excede el límite de la tarjeta")
	}

	fmt.Printf("Pagando %.2f con tarjeta %s\n", amount, c.Number)
	return nil
}

func (c CreditCard) Name() string {
	return "credit_card"
}