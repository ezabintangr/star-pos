package routes

import (
	categoryHandler "star-pos/features/categories/handler"
	userHandler "star-pos/features/user/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/hello", userHandler.Hello)
	e.POST("/login", userHandler.Login)

	e.POST("/users", userHandler.CreateAccount)
	e.GET("/users", userHandler.GetAllProfile)
	e.GET("/users/:id", userHandler.GetProfile)
	e.PATCH("/users/:id", userHandler.UpdateProfile)
	e.DELETE("/users/:id", userHandler.DeleteAccount)

	e.POST("/category", categoryHandler.CreateCategories)
	e.GET("/category", categoryHandler.GetAllCategories)
	e.GET("/category/:id", categoryHandler.GetCurrentCategory)
	e.PATCH("/category/:id", categoryHandler.UpdateCategory)
	e.DELETE("/category/:id", categoryHandler.DeleteCategory)
}
