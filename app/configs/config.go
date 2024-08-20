package configs

import (
	"os"
	"strconv"
)

var JWT_SECRET string

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.DB_USERNAME = os.Getenv("DBUSER")
	app.DB_PASSWORD = os.Getenv("DBPASS")
	app.DB_HOSTNAME = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error convert port")
	}
	app.DB_PORT = portConv
	app.DB_NAME = os.Getenv("DBNAME")
	JWT_SECRET = os.Getenv("JWTSECRET")

	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
