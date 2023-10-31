package controllers

import (
	"finapp/lib"
	"finapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	logger      lib.Logger
	userService services.UserService
}

// NewUserController creates a new user controller
func NewUserController(logger lib.Logger) UserController {
	return UserController{
		logger: logger,
	}
}

// getUser returns the fields of the authorized user (with the exception of the fields from gorm and Password)
func (ctrl UserController) GetUser(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing userID",
		})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userID",
		})
		return
	}

	user, err := ctrl.userService.GetUser(userIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
