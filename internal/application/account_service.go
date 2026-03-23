package application

import "payment-system/internal/domain"

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Debit(userID string, amount float64) error {
	balance, err := s.repo.GetBalance(userID)
	if err != nil {
		return err
	}

	if balance < amount {
		return domain.ErrInsufficientFunds
	}

	return s.repo.Debit(userID, amount)
}