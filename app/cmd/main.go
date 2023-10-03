package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testmessage",
		})
	})

	r.Run(config.Cfg.Port)
}
