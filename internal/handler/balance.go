package handler

import (
	avito_test_case "avito-test-case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) transfer(c *gin.Context) {
	var transfer avito_test_case.TransferRequest

	err := c.ShouldBindJSON(&transfer)
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "failed validation")
		return
	}

	err = h.services.Balance.Transfer(transfer.From, transfer.To, transfer.Value, transfer.Comment)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (h *Handler) enrollment(c *gin.Context) {
	var valueRequest avito_test_case.ValueRequest

	err := c.ShouldBindJSON(&valueRequest)
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "failed validation")
		return
	}

	err = h.services.Enrollment(valueRequest.UserId, valueRequest.Value)
	if err != nil {
		newErrorResponse(c, http.StatusUnsupportedMediaType, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
func (h *Handler) writeOff(c *gin.Context) {
	var valueRequest avito_test_case.ValueRequest

	err := c.ShouldBindJSON(&valueRequest)
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "failed validation")
		return
	}

	err = h.services.WriteOff(valueRequest.UserId, valueRequest.Value)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (h *Handler) show(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Params.ByName("user_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "failed validation")
		return
	}

	balance, err := h.services.Balance.UserBalance(userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"balance": balance,
	})
}

func (h *Handler) history(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Params.ByName("user_id"), 10, 64)
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "15"), 10, 64)
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	orderBy := c.DefaultQuery("orderBy", "date")
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "failed validation")
		return
	}

	transactions, err := h.services.Balance.TransfersHistory(userId, limit, page, orderBy)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) proceeds(c *gin.Context) {
	proceeds, err := h.services.Balance.GetProceeds(c.Query("date"))
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}
	path, err := h.services.Balance.WriteProceedsToCSV(proceeds, c.Query("date"))

	c.JSON(http.StatusOK, map[string]string{
		"success": path,
	})
}
