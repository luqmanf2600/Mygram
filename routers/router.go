package routers

import (
	"Mygram/controllers"
	"Mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitApplication() *gin.Engine {
	r := gin.Default()
	// user/register
	// user/login
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.UserRegistration)
		userGroup.POST("/login", controllers.UserLogin)
		userGroup.PUT("/login", controllers.UserLogin)
		userGroup.DELETE("/login", controllers.UserLogin)
	}

	productGroup := r.Group("/product")
	{
		productGroup.Use(middlewares.Authentication())
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.POST("/:productId", middlewares.Authorization(), controllers.UpdateProduct)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("/:id", controllers.GetCommentByID)
	}

	return r
}
