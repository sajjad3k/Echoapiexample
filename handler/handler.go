package handler

//Handler to maintain the API Routes

import (
	"github.com/labstack/echo"
	"github.com/sajjad3k/echoapione/controller"
)

func SetRoutes() *echo.Echo {

	e := echo.New()

	api := e.Group("/shop")
	{
		api.GET("/products", controller.Getproducts)
		api.GET("/products/:code", controller.Getproduct)
		api.POST("/newproduct", controller.Createproduct)
		api.PUT("/product/:code", controller.Updateproduct)
		api.DELETE("/product/:code", controller.Deleteproduct)
	}
	return e
}
