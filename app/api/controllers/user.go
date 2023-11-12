package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
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

	user, err := uc.service.Get(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	response := models.GetResponse{
		Email:      user.Email,
		First_name: user.FirstName,
		Last_name:  user.LastName,
		Patronymic: user.Patronymic,
		Gender:     user.Gender,
		Birthday:   user.Birthday.Format(models.DateFormat),
	}

	c.JSON(http.StatusOK, response)
}

// Обновление

//	@Security		ApiKeyAuth
//	@summary		Update user
//	@tags			user
//	@Description	Обновление пользователя
//	@ID				update
//	@Accept			json
//	@Produce		json
//	@Router			/user [patch]

func (uc UserController) Update(c *gin.Context) {
	// Парсинг запроса
	var updateRequest models.UpdateRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update request"})
		return
	}

	// Получение идентификатора пользователя из контекста
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Получение текущего пользователя из службы
	user, err := uc.service.Get(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Обновление полей пользователя на основе запроса на обновление
	user.FirstName = updateRequest.FirstName
	user.LastName = updateRequest.LastName
	user.Patronymic = updateRequest.Patronymic
	user.Gender = updateRequest.Gender
	// Преобразование строки с датой рождения в формат time.Time
	updatedBirthday, err := time.Parse(models.DateFormat, updateRequest.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birthday format"})
		return
	}
	user.Birthday = updatedBirthday

	// Вызов метода Update для сохранения изменений в базе данных
	if err := uc.service.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Получение данных о транзакции из запроса
	var transactionRequest models.TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction request"})
		return
	}

	// Создание финансовой транзакции
	err = uc.service.CreateTransaction(userID.(uint), transactionRequest.Amount, transactionRequest.Currency, transactionRequest.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Ответ успешного обновления
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
