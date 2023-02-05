package service

import (
	billingService "billingService"
	"billingService/internal/repository"
)

type Balance interface {
	UserBalance(userId uint64) (uint64, error)
	Enrollment(userId, value uint64) error
	WriteOff(userId, value uint64) error
	Transfer(from, to, value uint64, comment string) error
	TransfersHistory(userId, limit, page uint64, orderBy string) ([]billingService.Transfer, error)
}

type Reservation interface {
	Reservation(userId, serviceId, orderId, value uint64) error
	Confirm(userId, serviceId, orderId, value uint64) error
}

type Service struct {
	Balance
	Reservation
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Balance:     NewBalanceService(r.Balance),
		Reservation: NewReservationService(r.Reservation),
	}
}
