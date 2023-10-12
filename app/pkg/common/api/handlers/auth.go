package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"app/pkg/common/config"
	"app/pkg/common/db"
	"app/pkg/common/models"
)

// Генерация токена
func GenerateToken(ID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(config.Cfg.SecretKey))
}

// Вход

// Структура запроса
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @summary Login
// @tags auth
// @Description Обработка запроса на вход пользователя
// @ID login
// @Accept json
// @Produce json
// @Param req body LoginRequest true "Данные пользователя (email и пароль)"
// @Router /api/v1/auth/login [post]
// TODO - Сделать для ошибок/успеха свои json
func Login(c *gin.Context) {
	// Парсинг запроса
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid request body"},
		)
		return
	}

	// Поиск пользователя
	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Invalid email or password"},
		)
		return
	}

	// Сравнение хэша и пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Invalid email or password"},
		)
		return
	}

	// Генерация JWT токена
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to generate token"},
		)
		return
	}

	// Отправка токена
	c.JSON(
		http.StatusOK,
		gin.H{"token": token},
	)
}

// @summary Register
// @tags auth
// @Description Обработка запроса на регистрацию пользователя
// @ID register
// @Accept json
// @Produce json
// @Param user body models.User true "Данные о пользователе"
// @Router /api/v1/auth/register [post]
// TODO - Сделать для ошибок/успеха свои структуры
func Register(c *gin.Context) {
	// Парсинг запроса
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid request body",
			},
		)
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to hash password"},
		)
		return
	}
	user.Password = string(hashedPassword)

	// Сохранение пользователя в БД
	// TODO: улучшить обработку ошибок
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Failed to safe user",
			},
		)
		return
	}

	// Отправка ответа
	c.JSON(
		http.StatusOK,
		gin.H{"message": "User registered successfully"},
	)
}

// Выход

// @summary Logout
// @tags auth
// @Description Обработка запроса выход пользователя из системы
// @ID logout
// @Produce json
// @Router /api/v1/auth/logout [post]
// TODO - Сделать для ошибок/успеха свои структуры
func Logout(c *gin.Context) {
	// TODO: Реализовать выход
	c.JSON(
		http.StatusInternalServerError,
		gin.H{"error": "Not realized"},
	)
}
