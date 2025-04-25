package main

import (
	"star-pos/app/configs"
	"star-pos/app/databases"
	"star-pos/app/migrations"
	"star-pos/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	databases.InitMySql(cfg)
	migrations.InitMigration()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	routes.InitRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
