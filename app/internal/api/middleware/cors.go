package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"app/internal/config"
)

func CorsMiddleware() gin.HandlerFunc {
	cfg := config.Cfg
	CorsConfig := cors.DefaultConfig()

	CorsConfig.AllowOrigins = []string{cfg.CORSAllowedOrigins}
	CorsConfig.AllowMethods = []string{cfg.CORSAllowedMethods}
	CorsConfig.AllowHeaders = []string{cfg.CORSAllowedHeaders}

	return cors.New(CorsConfig)
}
