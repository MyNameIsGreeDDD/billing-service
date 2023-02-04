package avito_test_case

import (
	"time"
)

type Balance struct {
	Id      uint64 `db:"id"`
	UserId  uint64 `db:"user_id"`
	Balance uint64 `db:"balance"`
}

type ValueRequest struct {
	Value  uint64 `json:"value"`
	UserId uint64 `json:"user_id"`
}

type TransferRequest struct {
	To      uint64 `json:"to"`
	From    uint64 `json:"from"`
	Value   uint64 `json:"value"`
	Comment string `json:"comment"`
}

type Transfer struct {
	Id         uint64    `db:"id" json:"id"`
	ToUserId   uint64    `db:"to_user_id" json:"to"`
	FromUserId uint64    `db:"from_user_id" json:"from"`
	Comment    string    `db:"comment" json:"comment"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	Value      uint64    `db:"value" json:"value"`
}
