package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"finapp/domains"
	"finapp/lib"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	logger  lib.Logger
	service domains.AuthService
}

// NewJWTAuthMiddleware creates new jwt auth middleware
func NewJWTAuthMiddleware(
	logger lib.Logger,
	service domains.AuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		logger:  logger,
		service: service,
	}
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality
func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Empty Authorization header",
			})
			c.Abort()
			return
		}
		t := strings.Split(authHeader, " ")
		if len(t) != 2 || t[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})
			c.Abort()
			return
		}
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := m.service.Authorize(authToken)
			if authorized {
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}
}