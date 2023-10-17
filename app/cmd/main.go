package main

import (
	"log"

	"app/pkg/common/api/routes"
	"app/pkg/common/config"
	"app/pkg/common/db"
)

// @title Finapp API
// @version 1.0

// @host zlp-lfp.ru
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
