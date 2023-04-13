package router

import (
	"go-jwt/controllers"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

 func StartApp() *gin.Engine{
	r := gin.Default()

	r.POST("/register", controllers.UserRegister)
	r.POST("/login", controllers.UserLogin)

	productRouter := r.Group("/products")
	{
		userRouter := productRouter.Group("/user")
		{
			userRouter.GET("/", controllers.GetAllProduct)
			userRouter.GET("/:productId", controllers.GetProductById)
			userRouter.Use(middlewares.Authentication())
			userRouter.POST("/", controllers.CreateProduct)
	
		}
	
		adminRouter := productRouter.Group("/admin")
		{
			adminRouter.Use(middlewares.Authentication())
			adminRouter.GET("/", controllers.GetAllProduct)
			adminRouter.GET("/:productId", controllers.GetProductById)
			adminRouter.POST("/", controllers.CreateProduct)
			adminRouter.PUT("/:productId",middlewares.ProductAuthorization(), controllers.UpdateProduct)
			adminRouter.DELETE("/:productId", middlewares.ProductAuthorization(),controllers.DeleteProduct)
		}
	}

	

	return r
 }