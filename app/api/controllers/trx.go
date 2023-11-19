package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
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

//	@Security		ApiKeyAuth
//	@summary		Get trx
//	@tags			trx
//	@Description	Получение транзакции
//	@ID				get_trx
//	@Accept			json
//	@Produce		json
//	@Param			amount_min	query	number	false	"Минимальное значение суммы транзакции"
//	@Param			amount_max	query	number	false	"Максимальное значение суммы транзакции"
//	@Param			date_from	query	string	false	"Дата начала периода в формате 18-10-2004"
//	@Param			date_to		query	string	false	"Дата окончания периода в формате 18-10-2004"
//	@Success		200			{array}	models.TrxResponse
//	@Router			/trx [get]
func (tc TrxController) Get(c *gin.Context) {
	trxs, err := tc.service.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get transaction",
		})
		return
	}

	var trxResponses []models.TrxResponse
	for _, trx := range trxs {
		trxResponses = append(trxResponses, models.TrxResponse{
			Name:   trx.Name,
			Date:   trx.Date,
			Amount: trx.Amount,
		})
	}

	c.JSON(http.StatusOK, trxResponses)
}

// Создание

//	@Security		ApiKeyAuth
//	@summary		Create trx
//	@tags			trx
//	@Description	Создание транзакции
//	@ID				post
//	@Accept			json
//	@Produce		json
//	@Router			/trx [post]
func (tc TrxController) Post(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}
