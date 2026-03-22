package application

import "payment-system/internal/domain"

type PaymentService struct{}

func (s *PaymentService) Process(p domain.PaymentMethod, amount float64) error {
	return p.Pay(amount)
}

