package main

import (
	"log"

	"app/docs"
	"app/pkg/common/api/routes"
	"app/pkg/common/config"
	"app/pkg/common/db"
)

// @title Finapp API
// @version 1.0

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

	docs.SwaggerInfo.Host = config.Cfg.Swagger.Host

	r.Run(config.Cfg.Port)
}
