package routes

import (
	"github.com/aksh-patel1/go-ecommerce-platform/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingroutes *gin.Engine) {
	incomingroutes.POST("/users/signup", controllers.SignUp())
	incomingroutes.POST("/users/login", controllers.Login())
	incomingroutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingroutes.GET("/users/productview", controllers.SearchProduct())
	incomingroutes.GET("/users/search", controllers.SearchProductByQuery())
}