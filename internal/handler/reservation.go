package handler

import (
	billingService "billingService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reservation(c *gin.Context) {
	var reservation billingService.ReservationRequest

	err := c.ShouldBindJSON(&reservation)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "failed validation")
		return
	}

	err = h.services.Reservation.Reservation(reservation.UserId, reservation.ServiceId, reservation.OrderId, reservation.Value)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (h *Handler) confirm(c *gin.Context) {
	var reservation billingService.ReservationRequest

	err := c.ShouldBindJSON(&reservation)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "failed validation")
		return
	}

	err = h.services.Reservation.Confirm(reservation.UserId, reservation.ServiceId, reservation.OrderId, reservation.Value)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
