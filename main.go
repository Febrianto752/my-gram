package main

import (
	"log"

	"github.com/Febrianto752/my-gram/config"
	"github.com/Febrianto752/my-gram/handler"
	"github.com/Febrianto752/my-gram/repository"
	"github.com/Febrianto752/my-gram/route"
	"github.com/Febrianto752/my-gram/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("gagal mengambil .env %v", err)

	}

	db := config.InitializeDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository, photoRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	router := route.NewRouter(userHandler, photoHandler, socialMediaHandler, commentHandler)

	router.Run(":8080")
}
