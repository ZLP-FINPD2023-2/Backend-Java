package main

import (
	"log"

	"app/pkg/common/api/routes"
	"app/pkg/common/config"
	"app/pkg/common/db"
)

// @title Todo App API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := db.Init(); err != nil {
		log.Fatalln(err)
	}

	r := routes.InitRouter()

	r.Run(config.Cfg.Port)
}
