package routes

import (
	userHandler "star-pos/features/user/handler"
	userData "star-pos/features/user/repository"
	userService "star-pos/features/user/service"
	encrypts "star-pos/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	userData := userData.New(db)
	hashData := encrypts.NewHashService()
	userService := userService.New(userData, hashData)
	userHandler := userHandler.New(userService)
	e.GET("/hello", userHandler.Hello)
	e.POST("/users", userHandler.CreateAccount)
}
