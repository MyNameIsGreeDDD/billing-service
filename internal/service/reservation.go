package service

import (
	"billingService/internal/repository"
)

type ReservationService struct {
	repo repository.Reservation
}

func NewReservationService(repo repository.Reservation) *ReservationService {
	return &ReservationService{repo: repo}
}

func (r ReservationService) Reservation(userId, serviceId, orderId, value uint64) error {
	return r.repo.Reservation(userId, serviceId, orderId, value)
}

func (r ReservationService) Confirm(userId, serviceId, orderId, value uint64) error {
	return r.repo.Confirm(userId, serviceId, orderId, value)
}
