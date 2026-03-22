package payments

import "fmt"

type Crypto struct {
	WalletAddress string
}

func (c Crypto) Pay(amount float64) error {
	fmt.Printf("Pagando %.2f con Crypto wallet %s\n", amount, c.WalletAddress)
	return nil
}

func (c Crypto) Name() string {
	return "crypto"
}