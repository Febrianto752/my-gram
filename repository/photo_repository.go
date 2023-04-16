package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/Febrianto752/my-gram/entity"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo entity.Photo) (entity.Photo, error)
	FindById(id uint) (entity.Photo, error)
	FindAll() ([]entity.Photo, error)
	Update(photo entity.Photo, id uint) (entity.Photo, error)
	Delete(id uint)
	IsPhotoExist(id uint) error
}

type photoRepository struct {
	db *gorm.DB
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(id uint) error {
	var photo entity.Photo
	err := r.db.Debug().First(&photo, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

// Create implements PhotoRepository
func (r *photoRepository) Create(photo entity.Photo) (entity.Photo, error) {
	err := r.db.Debug().Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Delete implements PhotoRepository
func (r *photoRepository) Delete(id uint) {
	var photo entity.Photo

	err := r.db.Debug().Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		log.Fatalln("error deleting data", err)
		return
	}
}

// FindAll implements PhotoRepository
func (r *photoRepository) FindAll() ([]entity.Photo, error) {
	var photos []entity.Photo

	err := r.db.Debug().Find(&photos).Error
	if err != nil {
		log.Fatal("error getting all social medias data: ", err)
	}
	return photos, nil
}

// FindById implements PhotoRepository
func (r *photoRepository) FindById(id uint) (entity.Photo, error) {
	var photo entity.Photo
	err := r.db.Debug().First(&photo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("social media not found")
		}
		log.Fatal("error getting data :", err)
	}

	return photo, err
}

// Update implements PhotoRepository
func (r *photoRepository) Update(photo entity.Photo, id uint) (entity.Photo, error) {

	err := r.db.Debug().Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		return photo, err
	}

	fmt.Println("repo", photo)

	return photo, nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}
