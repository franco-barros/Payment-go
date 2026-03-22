package domain

type PaymentMethod interface {
	Pay(amount float64) error
}