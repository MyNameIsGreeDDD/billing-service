package handler

import (
	"avito-test-case/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		balance := api.Group("/balance")
		{
			balance.GET("/:user_id", h.show)
			balance.GET("/:user_id/history", h.history)
			balance.POST("/enrollment", h.enrollment)
			balance.POST("/write-off", h.writeOff)
			balance.POST("/transfer", h.transfer)

		}

		reservation := api.Group("/reservation")
		{
			reservation.POST("/reserve", h.reservation)
			reservation.POST("/confirm", h.confirm)
		}
	}

	return router
}
