package middleware

import (
	"app/pkg/common/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "Empty Authorization header"},
			)
			c.Abort()
			return
		}

		splittedToken := strings.Split(tokenHeader, " ")
		if len(splittedToken) != 2 || splittedToken[0] != "Bearer" {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid token format"},
			)
			c.Abort()
			return
		}

		tokenString := splittedToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"},
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
