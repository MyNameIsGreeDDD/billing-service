package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const (
	ReservationTable = "reservations"
	PurchasesTable   = "purchases"
)

type ReservationRepository struct {
	db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r ReservationRepository) Reservation(userId, serviceId, orderId, value uint64) error {
	query, args, err := sq.Insert(ReservationTable).
		Columns("user_id", "order_id", "service_id", "value").
		Values(userId, orderId, serviceId, value).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("fail building query")
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return fmt.Errorf("fail exec query")
	}

	return nil
}

func (r ReservationRepository) Confirm(userId, serviceId, orderId, value uint64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed start transactions")
	}

	query, args, err := sq.Delete(ReservationTable).
		Where(sq.And{
			sq.Eq{"order_id": orderId},
			sq.Eq{"user_id": userId},
			sq.Eq{"service_id": serviceId},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("fail building query")
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed delete reservation")
	}

	query, args, err = sq.Insert(PurchasesTable).
		Columns("user_id", "order_id", "service_id", "value").
		Values(userId, orderId, serviceId, value).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("fail building query")
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		fmt.Printf("%s", err.Error())
		tx.Rollback()
		return fmt.Errorf("failed create operation")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("commit failed")
	}

	return nil
}
