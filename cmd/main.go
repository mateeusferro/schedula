package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mateeusferro/schedula/internal/config"
)

var db *sql.DB

func main() {
	config.LoadEnv()
	serverPort := config.EnvVariable("PORT")

	db = dbConnect()
	defer db.Close()

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatalf("Error connecting to the db: %v", pingErr)
	}

	fmt.Println("Connected to the db")

	var router *gin.Engine = gin.Default()

	err := router.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Fatalf("Error while setting trusted proxies: %v", err)
	}

	err1 := router.Run(":" + serverPort)

	if err1 != nil {
		log.Fatalf("Error while starting the server: %v", err1)
	}
}

func dbConnect() *sql.DB {
	host, port, user, password, name := config.ReturnDbConfig()

	pgConnStr := fmt.Sprintf(`
	host=%s port=%d user=%s password=%s
	dbname=%s sslmode=disable
	`, host, port, user, password, name)

	conn, err := sql.Open("postgres", pgConnStr)

	if err != nil {
		log.Fatalf("Error opening db connection: %v", err)
	}

	return conn
}
