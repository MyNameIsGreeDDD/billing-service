package repository

import (
	avito_test_case "avito-test-case"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const BalancesTable string = "users_balances"
const TransfersTable string = "transfers"
const ValueColumn = "value"
const CreatedAtColumn = "created_at"

type BalanceRepository struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (b BalanceRepository) UserBalance(userId uint64) (uint64, error) {
	var balance avito_test_case.Balance

	query, args, err := sq.Select("*").
		From(BalancesTable).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return balance.Balance, fmt.Errorf("failed build query")
	}

	err = b.db.QueryRow(query, args...).Scan(&balance.Id, &balance.UserId, &balance.Balance)
	if err != nil {
		return balance.Balance, fmt.Errorf("not found result")
	}

	return balance.Balance, nil
}
func (b BalanceRepository) Transfer(from, to, value uint64) error {
	tx, err := b.db.Begin()
	if err != nil {
		return fmt.Errorf("cant start transfer")
	}

	query, args, err := updateUserBalance(from, value, "-")

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed write-off balance")
	}

	query, args, err = updateUserBalance(to, value, "+")

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed enrollment balance")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("cant commit changes")
	}

	return nil
}
func (b BalanceRepository) Enrollment(userId, value uint64) error {
	tx, err := b.db.Begin()
	if err != nil {
		return fmt.Errorf("cant start enrollment")
	}

	query, args, err := updateUserBalance(userId, value, "+")

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed enrollment balance")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("cant commit changes")
	}

	return nil
}
func (b BalanceRepository) WriteOff(userId, value uint64) error {
	tx, err := b.db.Begin()
	if err != nil {
		return fmt.Errorf("cant start write-off")
	}

	query, args, err := updateUserBalance(userId, value, "-")

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed write-off balance")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("cant commit changes")
	}

	return nil
}

func updateUserBalance(userId, value uint64, operator string) (string, []interface{}, error) {
	return sq.Update(BalancesTable).
		Set("balance", sq.Expr(fmt.Sprintf("%s %s %d", "balance", operator, value))).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func (b BalanceRepository) WriteTransfer(from, to, value uint64, comment string) error {
	query, args, err := sq.Insert(TransfersTable).
		Columns("from_user_id", "to_user_id", "value", "comment").
		Values(from, to, value, comment).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed query build")
	}

	_, err = b.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed write transfer")
	}

	return nil
}

func (b BalanceRepository) TransactionsHistory(userId, limit, page uint64, orderBy string) ([]avito_test_case.Transfer, error) {
	var transfers []avito_test_case.Transfer

	query, args, err := buildTransactionQuery(orderBy, userId, limit, page)
	if err != nil {
		return transfers, fmt.Errorf("failed query build")
	}
	rows, err := b.db.Query(query, args...)
	if err != nil {
		return transfers, fmt.Errorf("failed get history")
	}

	for rows.Next() {
		var transfer avito_test_case.Transfer

		err = rows.Scan(&transfer.Id, &transfer.FromUserId, &transfer.ToUserId, &transfer.Value, &transfer.Comment, &transfer.CreatedAt)
		if err != nil {
			return transfers, fmt.Errorf("failed write result")
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func buildTransactionQuery(orderBy string, userId, limit, page uint64) (string, []interface{}, error) {
	query := sq.Select("*").
		From(TransfersTable).
		Where(sq.Or{sq.Eq{"from_user_id": userId}, sq.Eq{"to_user_id": userId}})

	if orderBy == "date" {
		query = query.OrderBy(fmt.Sprintf("%s %s", CreatedAtColumn, "ASC"))
	}

	if orderBy == "sum" {
		query = query.OrderBy(fmt.Sprintf("%s %s", ValueColumn, "DESC"))
	}

	return query.Limit(limit).Offset(limit*page - limit).PlaceholderFormat(sq.Dollar).ToSql()
}
