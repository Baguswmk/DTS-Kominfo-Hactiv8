package routers

import (
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/database"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/middleware"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/controller"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	db := database.ConnectDB()
	router := gin.Default()
	user := controller.NewUserController(db)
	social := controller.NewSocialController(db)
	photo := controller.NewPhotoController(db)
	comment := controller.NewCommentController(db)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", user.UserLogin)
		userGroup.POST("/register", user.CreateUser)
		userGroup.PUT("/", middleware.Authorization(), user.UpdateUser)
		userGroup.DELETE("/", middleware.Authorization(), user.DeleteUser)
	}

	socialGroup := router.Group("/socials")
	{
		socialGroup.GET("/", middleware.Authorization(), social.FindAllSocial)
		socialGroup.POST("/", middleware.Authorization(), social.CreateSocial)
		socialGroup.PUT("/:socialMediaId", middleware.Authorization(), middleware.SocialAuthorization(), social.UpdateSocial)
		socialGroup.DELETE("/:socialMediaId", middleware.Authorization(), middleware.SocialAuthorization(),social.DeleteSocial)
	}

	photoGroup := router.Group("/photos")
	{
		photoGroup.GET("/", middleware.Authorization(), photo.FindAllPhoto)
		photoGroup.POST("/", middleware.Authorization(), photo.CreatePhoto)
		photoGroup.PUT("/:photoId", middleware.Authorization(), middleware.PhotoAuthorization(), photo.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middleware.Authorization(),middleware.PhotoAuthorization(), photo.DeletePhoto)
	}

	commentGroup := router.Group("/comments")
	{
		commentGroup.GET("/", middleware.Authorization(), comment.FindAllComment)
		commentGroup.POST("/", middleware.Authorization(), comment.CreateComment)
		commentGroup.PUT("/:commentId", middleware.Authorization(),middleware.CommentAuthorization(), comment.UpdateComment)
		commentGroup.DELETE("/:commentId", middleware.Authorization(),middleware.CommentAuthorization(), comment.DeleteComment)
	}

	return router
}
