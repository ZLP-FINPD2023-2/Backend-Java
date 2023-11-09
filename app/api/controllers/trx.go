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
	trxService domains.AuthService,
) TrxController {
	return TrxController{
		logger:      logger,
		service:     service,
		authService: trxService,
	}
}

func (tc TrxController) Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "GET responce",
	})
}

func (tc TrxController) Post(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "POST responce",
	})
}
