package routes

import (
	cartHandler "star-pos/features/cart/handler"
	categoryHandler "star-pos/features/categories/handler"
	discountHandler "star-pos/features/discount/handler"
	outletHandler "star-pos/features/outlet/handler"
	productHandler "star-pos/features/product/handler"
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

	e.POST("/product", productHandler.CreateProduct)
	e.GET("/product", productHandler.GetAllProducts)
	e.GET("/product/:id", productHandler.GetProduct)
	e.PATCH("/product/:id", productHandler.UpdateProduct)
	e.DELETE("/product/:id", productHandler.DeleteProduct)

	e.POST("/discounts", discountHandler.CreateDiscounts)
	e.GET("/discounts", discountHandler.GetAllDiscounts)
	e.GET("/discounts/:id", discountHandler.GetCurrentDiscount)
	e.PATCH("/discounts/:id", discountHandler.UpdateDiscount)
	e.DELETE("/discounts/:id", discountHandler.DeleteDiscount)

	e.POST("/outlets", outletHandler.CreateOutlet)
	e.GET("/outlets", outletHandler.GetAllOutlets)
	e.GET("/outlets/:id", outletHandler.GetOutlet)
	e.PATCH("/outlets/:id", outletHandler.UpdateOutlet)
	e.DELETE("/outlets/:id", outletHandler.DeleteOutlet)

	e.POST("/cart/:id/product", cartHandler.AddProductToCartHandler)
}
