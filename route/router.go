package route

import (
	"github.com/Febrianto752/my-gram/handler"
	"github.com/Febrianto752/my-gram/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler handler.UserHandler, photoHandler handler.PhotoHandler, socialMediaHandler handler.SocialMediaHandler, commentHandler handler.CommentHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.PostUserRegisterHandler)
	router.POST("/signin", userHandler.PostUserLoginHandler)

	photo := router.Group("/photo")
	{
		photo.Use(middleware.Authentication())
		photo.POST("", photoHandler.PostPhotoHandler)
		photo.GET("", photoHandler.GetPhotosHandler)
		photo.GET("/:id", photoHandler.GetPhotoHandler)
		photo.PUT("/:id", middleware.PhotoAuthorization(), photoHandler.PutPhotoHandler)
		photo.DELETE("/:id", middleware.PhotoAuthorization(), photoHandler.DeletePhotoHandler)
	}

	socialMedia := router.Group("/socialmedia")
	{
		socialMedia.Use(middleware.Authentication())
		socialMedia.POST("", socialMediaHandler.PostSocialMediaHandler)
		socialMedia.GET("", socialMediaHandler.GetSocialMediasHandler)
		socialMedia.GET("/:id", socialMediaHandler.GetSocialMediaHandler)
		socialMedia.PUT("/:id", middleware.SocialMediaAuthorization(), socialMediaHandler.PutSocialMediaHandler)
		socialMedia.DELETE("/:id", middleware.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMediaHandler)
	}

	comment := router.Group("/comment")
	{
		comment.Use(middleware.Authentication())
		comment.POST("", commentHandler.PostCommentHandler)
		comment.GET("", commentHandler.GetCommentsHandler)
		comment.GET("/:id", commentHandler.GetCommentHandler)
		comment.PUT("/:id", middleware.CommentAuthorization(), commentHandler.PutCommentHandler)
		comment.DELETE("/:id", middleware.CommentAuthorization(), commentHandler.DeleteCommentHandler)
	}

	return router
}
