package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"finapp/domains"
	"finapp/lib"
)

// UserController struct
type UserController struct {
	logger      lib.Logger
	service     domains.UserService
	authService domains.AuthService
}

// NewUserController creates new controller
func NewUserController(
	logger lib.Logger,
	service domains.UserService,
	userService domains.AuthService,
) UserController {
	return UserController{
		logger:      logger,
		service:     service,
		authService: userService,
	}
}

// Удаление

// @Security		ApiKeyAuth
// @summary		Delete user
// @tags			user
// @Description	Удаление пользователя
// @ID				delete
// @Accept			json
// @Produce		json
// @Router			/user [delete]
func (uc UserController) Delete(c *gin.Context) {
	// Парсинг запроса
	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	// Удаление пользователя
	if err := uc.service.Delete(userId.(uint)); err != nil {
		// Необработанные ошибки
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	// Отправка ответа
	c.JSON(http.StatusNoContent, gin.H{
		"message": "User deleted successfully",
	})
}

// Получение

// @Security		ApiKeyAuth
// @summary		Get user
// @tags			user
// @Description	Получение пользователя
// @ID				get
// @Accept			json
// @Produce		json
// @Router			/user [get]
func (uc UserController) Get(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	user, err := uc.service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	response := gin.H{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"patronymic": user.Patronymic,
		"gender":     user.Gender,
		"birthday":   user.Birthday,
	}

	c.JSON(http.StatusOK, response)
}

// Обновление

// @Security		ApiKeyAuth
// @summary		Update user
// @tags			user
// @Description	Обновление пользователя
// @ID				update
// @Accept			json
// @Produce		json
// @Router			/user [patch]
func (uc UserController) Update(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
