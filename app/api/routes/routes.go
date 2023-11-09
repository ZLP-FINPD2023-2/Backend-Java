package routes

import (
	"log"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewDocsRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewTrxRoutes),
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
	userRoutes UserRoutes,
	docsRoutes DocsRoutes,
	trxRoutes TrxRoutes,
) Routes {
	return Routes{
		authRoutes,
		userRoutes,
		docsRoutes,
		trxRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	log.Println("Setting up routes")
	for _, route := range r {
		route.Setup()
	}
}
