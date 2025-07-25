package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mateeusferro/schedula/internal/config"
	"github.com/mateeusferro/schedula/internal/database"
	"github.com/mateeusferro/schedula/internal/delivery"
)

var db *sql.DB

func main() {
	config.LoadEnv()
	serverPort := config.EnvVariable("PORT")

	var router *gin.Engine = gin.Default()

	getDb()

	err := router.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Fatalf("Error while setting trusted proxies: %v", err)
	}

	delivery.Routes(router, db)

	err = router.Run(":" + serverPort)

	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}

func getDb() {
	db = database.DbConnect()

	err := db.Ping()

	if err != nil {
		log.Fatalf("Error connecting to the db: %v", err)
	}

	fmt.Println("Connected to the db")
}
