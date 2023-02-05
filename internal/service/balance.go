package service

import (
	billingService "billingService"
	"billingService/internal/repository"
	"fmt"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (u BalanceService) UserBalance(userId uint64) (uint64, error) {
	return u.repo.UserBalance(userId)
}

func (u BalanceService) Transfer(from, to, value uint64, comment string) error {
	err := u.repo.Transfer(from, to, value)
	if err != nil {
		return err
	}
	return u.repo.WriteTransfer(from, to, value, comment)
}

func (u BalanceService) Enrollment(userId, value uint64) error {
	return u.repo.Enrollment(userId, value)
}

func (u BalanceService) WriteOff(userId, value uint64) error {
	return u.repo.WriteOff(userId, value)
}

func (u BalanceService) TransfersHistory(userId, limit, page uint64, orderBy string) ([]billingService.Transfer, error) {
	transfers, err := u.repo.TransactionsHistory(userId, limit, page, orderBy)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	if len(transfers) == 0 {
		return nil, fmt.Errorf("%s", "no result set")
	}

	return transfers, nil
}
