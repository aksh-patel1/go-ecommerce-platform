package main

import (
	"os"

	"github.com/aksh-patel1/go-ecommerce-platform/controllers"
	"github.com/aksh-patel1/go-ecommerce-platform/database"
	"github.com/aksh-patel1/go-ecommerce-platform/middleware"
	"github.com/aksh-patel1/go-ecommerce-platform/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}