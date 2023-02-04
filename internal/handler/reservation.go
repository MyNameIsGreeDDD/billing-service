package handler

import (
	avito_test_case "avito-test-case"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reservation(c *gin.Context) {
	var reservation avito_test_case.ReservationRequest

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
	var reservation avito_test_case.ReservationRequest

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
