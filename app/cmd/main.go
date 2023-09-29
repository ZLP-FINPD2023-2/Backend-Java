package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/pkg/common/config"
	"app/pkg/common/db"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	_, err = db.Init(cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testmessage",
		})
	})

	r.Run(cfg.Port)
}
