package middleware

import (
	"github.com/gin-contrib/cors"

	"app/pkg/common/config"
)

var CorsConfig = cors.DefaultConfig()

func InitCorsConfig() {
    corsCfg := config.Cfg.Cors

	CorsConfig.AllowOrigins = []string{corsCfg.AllowedOrigins}
	CorsConfig.AllowMethods = []string{corsCfg.AllowedMethods}
	CorsConfig.AllowHeaders = []string{corsCfg.AllowedHeaders}
}
