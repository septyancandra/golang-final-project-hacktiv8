package routers

import (
	//"Mygram/controllers"
	//"Mygram/middlewares"

	"Mygram/controllers"
	"Mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.DeleteUser)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/add", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMediaBySocialMediaId)
		socialMediaRouter.PUT("/update/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/delete/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/add", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.GET("/:photoId", controllers.GetPhotoByPhotoId)
		photoRouter.PUT("/update/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/add", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.GET("/:commentId", controllers.GetCommentByCommentId)
		commentRouter.PUT("/update/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/delete//:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
