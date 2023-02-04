package service

import (
	avito_test_case "avito-test-case"
	"avito-test-case/internal/repository"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
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

func (u BalanceService) TransfersHistory(userId, limit, page uint64, orderBy string) ([]avito_test_case.Transfer, error) {
	return u.repo.TransactionsHistory(userId, limit, page, orderBy)
}

func (u BalanceService) GetProceeds(date string) ([]avito_test_case.Proceeds, error) {
	firstDate, err := time.Parse("2006-01", date)
	if err != nil {
		return nil, fmt.Errorf("%s", "cant parse date")
	}
	lastDate := firstDate.AddDate(0, 1, 0)

	return u.repo.GetProceeds(firstDate, lastDate)
}

func (u BalanceService) WriteProceedsToCSV(proceeds []avito_test_case.Proceeds, slug string) (string, error) {
	columns := []string{"order_id", "sum"}
	fileName := fmt.Sprintf("%s%s", slug, ".scv")

	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("%s", "failed to create file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write(columns); err != nil {
		return "", fmt.Errorf("%s", "failed write to file")
	}

	for _, proceed := range proceeds {
		orderId := strconv.FormatUint(proceed.OrderId, 10)
		sum := strconv.FormatUint(proceed.Sum, 10)

		if err := w.Write([]string{orderId, sum}); err != nil {
			return "", fmt.Errorf("%s", "failed write to file")
		}
	}

	return fileName, nil
}
