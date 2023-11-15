package controllers

import (
	"finapp/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"finapp/domains"
	"finapp/lib"
)

type TrxController struct {
	logger      lib.Logger
	service     domains.TrxService
	authService domains.AuthService
}

func NewTrxController(
	logger lib.Logger,
	service domains.TrxService,
	authService domains.AuthService,
) TrxController {
	return TrxController{
		logger:      logger,
		service:     service,
		authService: authService,
	}
}

// Получение

// @Security		ApiKeyAuth
// @summary		Get trx
// @tags			trx
// @Description	Получение транзакции
// @ID				get
// @Accept			json
// @Produce		json
// @Router			/trx [get]
func (tc TrxController) Get(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid name",
		})
		return
	}

	trx, err := tc.service.Get(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get transaction",
		})
		return
	}
	trxResponse := models.TrxResponse{
		Name:   trx.Name,
		Date:   trx.Date,
		Amount: trx.Amount,
	}
	c.JSON(http.StatusOK, trxResponse)
}

// Создание

// @Security		ApiKeyAuth
// @summary		Create trx
// @tags			trx
// @Description	Создание транзакции
// @ID				post
// @Accept			json
// @Produce		json
// @Router			/trx [post]
func (tc TrxController) Post(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}
