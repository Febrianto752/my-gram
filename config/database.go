package config

import (
	"fmt"
	"log"
	"os"

	entities "github.com/Febrianto752/my-gram/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database = ", err)
	} else {
		log.Println("Successfully connected to database")
	}

	db.Debug().AutoMigrate(&entities.User{}, &entities.Photo{}, &entities.SocialMedia{}, &entities.Comment{})

	return db
}
