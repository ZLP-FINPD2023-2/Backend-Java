package routes

import (
	"finapp/api/controllers"
	"finapp/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1")
	{
		user := root.Group("/user")
		{
			user.GET("/", s.userController.GetUser)
		}
	}
}

// NewUserRoutes creates new user routes
func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		logger:         logger,
		handler:        handler,
		userController: userController,
	}
}
