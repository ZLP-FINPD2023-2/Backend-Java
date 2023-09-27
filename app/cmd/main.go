package main

import (
	"app/pkg/common/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/pkg/common/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	port := cfg.Port

	r := gin.Default()
	_, err = db.Init(cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testmessage",
		})
	})

	r.Run(port)
}
