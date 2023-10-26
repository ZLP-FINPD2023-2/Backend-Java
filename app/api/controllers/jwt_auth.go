package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     domains.AuthService
	userService domains.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service domains.AuthService,
	userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// Вход

// @summary Login
// @tags auth
// @Description Вход пользователя
// @ID login
// @Accept json
// @Produce json
// @Param req body models.LoginRequest true "Данные пользователя"
// @Router /auth/login [post]
func (jwt JWTAuthController) Login(c *gin.Context) {
	// Парсинг запроса
	var q models.LoginRequest
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
		return
	}

	// Авторизация пользователя
	user, err := jwt.userService.Authorize(&q)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Отправка токена
	token := jwt.service.CreateToken(user)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Регистрация

// @summary Register
// @tags auth
// @Description Регистрация пользователя
// @ID register
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "Данные пользователя"
// @Router /auth/register [post]
func (jwt JWTAuthController) Register(c *gin.Context) {
	// Парсинг запроса
	var q models.RegisterRequest
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Хэширование пароля
	hashedPassword, err := jwt.service.HashPassword(q.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	q.Password = string(hashedPassword)

	// Регистрация пользователя
	if err := jwt.userService.Register(&q); err != nil {
		// Ошибки валидации
		var vErr validator.ValidationErrors
		if errors.As(err, &vErr) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": lib.ParseValidationErrors(vErr),
				},
			)
			return
		}

		// Ошибка уникального значения
		// TODO: Придумать обработчик получше
		if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{
					"Email": "duplicate",
				},
			})
			return
		}

		// Необработанные ошибки
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed",
		})
		return
	}

	// Отправка ответа
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}
