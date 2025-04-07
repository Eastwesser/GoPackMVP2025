// internal/controller/geocoder.go
package controller

import (
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GeocoderController struct {
	svc service.GeocoderService
}

func NewGeocoderController(svc service.GeocoderService) *GeocoderController {
	return &GeocoderController{svc: svc}
}

// Geocode godoc
// @Summary Геокодирование адреса
// @Description Определение координат по адресу (только для России)
// @Tags geocoder
// @Accept  json
// @Produce  json
// @Param input body GeocodeRequest true "Адрес для геокодирования"
// @Security ApiKeyAuth
// @Success 200 {object} entity.Address
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/geocode [post]
func (c *GeocoderController) Geocode(ctx *gin.Context) {
	var req GeocodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetUint("userID")

	address, err := c.svc.Geocode(ctx.Request.Context(), req.Address, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

// GetHistory godoc
// @Summary История запросов
// @Description Получение истории геокодирований пользователя
// @Tags geocoder
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entity.Address
// @Failure 500 {object} ErrorResponse
// @Router /api/history [get]
func (c *GeocoderController) GetHistory(ctx *gin.Context) {
	userID := ctx.GetUint("userID")

	history, err := c.svc.GetHistory(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, history)
}

// Другие обработчики...
