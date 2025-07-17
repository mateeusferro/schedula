package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateeusferro/schedula/internal/config"
)

func main() {
	config.LoadEnv()
	serverPort := config.EnvVariable("PORT")

	var router *gin.Engine = gin.Default()

	err := router.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Fatalf("Error while setting trusted proxies: %v", err)
	}

	router.GET("/test", serverTest)

	err1 := router.Run(":" + serverPort)

	if err1 != nil {
		log.Fatalf("Error while starting the server: %v", err1)
	}
}

func serverTest(context *gin.Context) {
	context.JSON(http.StatusOK, "OK")
}
