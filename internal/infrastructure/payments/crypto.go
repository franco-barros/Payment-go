package payments

import (
	"errors"
	"fmt"

	"payment-system/internal/domain"
)

type Crypto struct {
	WalletAddress string
}

var _ domain.PaymentMethod = (*Crypto)(nil)

func (c Crypto) Pay(amount float64) error {
	// 🔴 Validación de wallet
	if c.WalletAddress == "" {
		return errors.New("crypto: wallet address requerida")
	}

	// 🔴 Validación de monto
	if amount <= 0 {
		return errors.New("crypto: monto inválido")
	}

	fmt.Printf("Pagando %.2f con Crypto wallet %s\n", amount, c.WalletAddress)
	return nil
}

func (c Crypto) Name() string {
	return "crypto"
}