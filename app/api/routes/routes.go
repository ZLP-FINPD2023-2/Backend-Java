package routes

import (
	"log"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewDocsRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	authRoutes AuthRoutes,
	docsRoutes DocsRoutes,
) Routes {
	return Routes{
		authRoutes,
		docsRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	log.Println("Setting up routes")
	for _, route := range r {
		route.Setup()
	}
}
