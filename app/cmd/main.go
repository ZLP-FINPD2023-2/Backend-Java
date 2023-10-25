package main

import (
	"log"
	"reflect"
	"strings"

	"app/docs"
	"app/internal/api/routes"
	"app/internal/config"
	"app/internal/db"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	docs.SwaggerInfo.Host = config.Cfg.Host

	r.Run(config.Cfg.Port)
}
