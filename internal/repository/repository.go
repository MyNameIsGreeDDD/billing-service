package repository

import (
	billingService "billingService"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	UserBalance(userId uint64) (uint64, error)
	Enrollment(userId, value uint64) error
	WriteOff(userId, value uint64) error
	Transfer(from, to, value uint64) error
	WriteTransfer(from, to, value uint64, comment string) error
	TransactionsHistory(userId, limit, page uint64, orderBy string) ([]billingService.Transfer, error)
}
type Reservation interface {
	Reservation(userId, serviceId, orderId, value uint64) error
	Confirm(userId, serviceId, orderId, value uint64) error
}

type Repository struct {
	Balance
	Reservation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance:     NewBalanceRepository(db),
		Reservation: NewReservationRepository(db),
	}
}
