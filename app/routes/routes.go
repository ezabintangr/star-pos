package routes

import (
	userHandler "star-pos/features/user/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/hello", userHandler.Hello)
	e.POST("/users", userHandler.CreateAccount)
	e.POST("/login", userHandler.Login)
	e.GET("/users", userHandler.GetAllProfile)
	e.GET("/users/:id", userHandler.GetProfile)
	e.PATCH("/users/:id", userHandler.UpdateProfile)
	e.DELETE("/users/:id", userHandler.DeleteAccount)
}
