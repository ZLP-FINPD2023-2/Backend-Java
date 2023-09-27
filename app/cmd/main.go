package main

import (
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

    port := cfg.Port // В последствии из конфига
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
