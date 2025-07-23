package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mateeusferro/schedula/internal/config"
)

func DbConnect() *sql.DB {
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
