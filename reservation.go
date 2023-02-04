package avito_test_case

type ReservationRequest struct {
	Id        uint64 `db:"id" json:"id"`
	UserId    uint64 `db:"user_id" json:"user_id"`
	OrderId   uint64 `db:"order_id" json:"order_id"`
	ServiceId uint64 `db:"service_id" json:"service_id"`
	Value     uint64 `db:"value" json:"value"`
}
