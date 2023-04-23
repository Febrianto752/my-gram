package route

import (
	_ "github.com/Febrianto752/my-gram/docs"
	"github.com/Febrianto752/my-gram/handler"
	"github.com/Febrianto752/my-gram/middleware"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Mygram
// @version 1.0
// @description Final Project Digitalent x Hactive8
// @termsOfService http://swagger.io/terms/
// @contact.name febrianto
// @contact.email febrianto.bekasi@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @schemes http
// @BasePath /
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description	How to input in swagger : 'Bearer <insert_your_token_here>'
func NewRouter(userHandler handler.UserHandler, photoHandler handler.PhotoHandler, socialMediaHandler handler.SocialMediaHandler, commentHandler handler.CommentHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.PostUserRegisterHandler)
	router.POST("/signin", userHandler.PostUserLoginHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
