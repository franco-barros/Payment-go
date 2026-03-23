package application

// AccountRepository define el contrato para acceder a cuentas
type AccountRepository interface {
	GetBalance(userID string) (float64, error)
	Debit(userID string, amount float64) error
}