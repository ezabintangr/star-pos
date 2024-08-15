package main

import (
	"star-pos/app/configs"
	"star-pos/app/databases"
	"star-pos/app/migrations"
	"star-pos/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMySql := databases.InitMySql(cfg)

	e := echo.New()

	migrations.InitMigration()
	routes.InitRouter(e, dbMySql)

	e.Logger.Fatal(e.Start(":8080"))
}
