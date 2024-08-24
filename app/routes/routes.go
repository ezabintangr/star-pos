package routes

import (
	"star-pos/app/middlewares"
	userHandler "star-pos/features/user/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/hello", userHandler.Hello)
	e.POST("/users", userHandler.CreateAccount)
	e.POST("/login", userHandler.Login)
	e.GET("/users", userHandler.GetAllProfile)
	e.GET("/users/:id", userHandler.GetProfile, middlewares.JWTMiddleware())
	e.PATCH("/users/:id", userHandler.UpdateProfile, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandler.DeleteAccount, middlewares.JWTMiddleware())
}
