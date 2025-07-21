package config

import (
	"log"
	"strconv"
)

func ReturnDbConfig() (string, int, string, string, string) {
	host := EnvVariable("DB_HOST")
	port, err := strconv.Atoi(EnvVariable("DB_PORT"))
	user := EnvVariable("DB_USER")
	password := EnvVariable("DB_PWD")
	name := EnvVariable("DB_NAME")

	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	return host, port, user, password, name
}
