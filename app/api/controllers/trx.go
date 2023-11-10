package controllers

import (
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

//	@Security		ApiKeyAuth
//	@summary		Get trx
//	@tags			trx
//	@Description	Получение транзакции
//	@ID				get
//	@Accept			json
//	@Produce		json
//	@Router			/trx [get]
func (tc TrxController) Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
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
