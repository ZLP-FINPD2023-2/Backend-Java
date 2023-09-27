package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    port := ":8080" // В последствии из конфига
	// dbUrl := *из конфига*

	r := gin.Default()
	// db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testmessage",
		})
	})

	r.Run(port)
}
