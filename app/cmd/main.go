package main

import (
	"log"

    "app/pkg/common/api/middleware"
	"app/pkg/common/api/routes"
	"app/pkg/common/config"
	"app/pkg/common/db"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := db.Init(); err != nil {
		log.Fatalln(err)
	}

    middleware.InitCorsConfig()

	r := routes.InitRouter()

	r.Run(config.Cfg.Port)
}
