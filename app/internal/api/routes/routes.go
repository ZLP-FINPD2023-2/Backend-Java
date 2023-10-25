package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "app/docs"
	"app/internal/api/handlers"
	"app/internal/api/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/schemes/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			auth := v1.Group("/auth")
			{
				auth.POST("/login", handlers.Login)
				auth.POST("/register", handlers.Register)

				auth.Use(middleware.AuthMiddleware())
				auth.POST("/logout", handlers.Logout)
			}
		}
	}

	return r
}
