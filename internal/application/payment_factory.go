package application

import (
	"errors"

	"payment-system/internal/domain"
)

type PaymentFactory struct {
	methods map[string]domain.PaymentMethod
}

func NewPaymentFactory(methods map[string]domain.PaymentMethod) *PaymentFactory {
	return &PaymentFactory{methods: methods}
}

func (f *PaymentFactory) Get(method string) (domain.PaymentMethod, error) {
	m, ok := f.methods[method]
	if !ok {
		return nil, errors.New("método no soportado")
	}
	return m, nil
}