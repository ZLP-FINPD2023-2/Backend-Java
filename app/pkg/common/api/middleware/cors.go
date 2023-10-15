package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"app/pkg/common/config"
)

func CorsMiddleware() gin.HandlerFunc {
	corsCfg := config.Cfg.Cors
	CorsConfig := cors.DefaultConfig()

	CorsConfig.AllowOrigins = []string{corsCfg.AllowedOrigins}
	CorsConfig.AllowMethods = []string{corsCfg.AllowedMethods}
	CorsConfig.AllowHeaders = []string{corsCfg.AllowedHeaders}

	return cors.New(CorsConfig)
}
