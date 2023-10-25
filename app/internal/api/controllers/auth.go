package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"app/internal/config"
	"app/internal/db"
	"app/internal/models"
)

// Генерация токена
func generateToken(ID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = ID
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(config.Cfg.SecretKey))
}

// Вход

// Структура запроса
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @summary Login
// @tags auth
// @Description Вход пользователя
// @ID login
// @Accept json
// @Produce json
// @Param req body loginRequest true "Данные пользователя"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	// Парсинг запроса
	var q loginRequest
	if err := c.BindJSON(&q); err != nil {
		c.JSON(
			http.StatusBadRequest,
			models.ErrorResponse{
				Error: "Invalid request body",
			},
		)
		return
	}

	// Поиск пользователя
	var user models.User
	if err := db.DB.Where("email = ?", q.Email).First(&user).Error; err != nil {
		c.JSON(
			http.StatusUnauthorized,
			models.ErrorResponse{
				Error: "Invalid email or password",
			},
		)
		return
	}

	// Сравнение хэша и пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(q.Password)); err != nil {
		c.JSON(
			http.StatusUnauthorized,
			models.ErrorResponse{
				Error: "Invalid email or password",
			},
		)
		return
	}

	// Генерация JWT токена
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.ErrorResponse{
				Error: "Failed to generate token",
			},
		)
		return
	}

	// Отправка токена
	c.JSON(
		http.StatusOK,
		gin.H{"token": token},
	)
}

// Регистрация

// Структура запроса
type registerRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        uint8  `json:"age"`
	Gender     bool   `json:"gender"`
}

// @summary Register
// @tags auth
// @Description Регистрация пользователя
// @ID register
// @Accept json
// @Produce json
// @Param user body registerRequest true "Данные пользователя"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	// Парсинг запроса
	var q registerRequest
	if err := c.BindJSON(&q); err != nil {
		c.JSON(
			http.StatusBadRequest,
			models.ErrorResponse{
				Error: "Invalid request body",
			},
		)
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(q.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.ErrorResponse{
				Error: "Failed to hash password",
			},
		)
		return
	}
	q.Password = string(hashedPassword)

	// Сохранение пользователя в БД
	user := models.User{
		Email:      q.Email,
		Password:   q.Password,
		FirstName:  q.FirstName,
		LastName:   q.LastName,
		Patronymic: q.Patronymic,
		Age:        q.Age,
		Gender:     q.Gender,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err,
				},
			)
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			models.ErrorResponse{
				Error: "Validation failed",
			},
		)
		return
	}

	// Отправка ответа
	c.JSON(
		http.StatusOK,
		models.SuccessResponse{
			Message: "User registered successfully",
		},
	)
}
