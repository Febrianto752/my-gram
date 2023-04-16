package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/Febrianto752/my-gram/entity"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	FindById(id uint) (entity.SocialMedia, error)
	FindAll() ([]entity.SocialMedia, error)
	Update(socialMedia entity.SocialMedia, id uint) (entity.SocialMedia, error)
	Delete(id uint)
}

type socialMediaRepository struct {
	db *gorm.DB
}

// Create implements SocialMediaRepository
func (r *socialMediaRepository) Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Debug().Create(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

// Delete implements SocialMediaRepository
func (r *socialMediaRepository) Delete(id uint) {
	var socialMedia entity.SocialMedia

	err := r.db.Debug().Where("id = ?", id).Delete(&socialMedia).Error
	if err != nil {
		log.Fatalln("error deleting data", err)
		return
	}
}

// FindAll implements SocialMediaRepository
func (r *socialMediaRepository) FindAll() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia

	err := r.db.Debug().Preload("User").Find(&socialMedias).Error
	if err != nil {
		// log.Fatal("error getting all social medias data: ", err)
		return socialMedias, err
	}
	return socialMedias, nil
}

// FindById implements SocialMediaRepository
func (r *socialMediaRepository) FindById(id uint) (entity.SocialMedia, error) {
	var socialMedia entity.SocialMedia
	err := r.db.Debug().Preload("User").First(&socialMedia, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("social media not found")
		}
		log.Fatal("error getting data :", err)
	}

	return socialMedia, err
}

// Update implements SocialMediaRepository
func (r *socialMediaRepository) Update(socialMedia entity.SocialMedia, id uint) (entity.SocialMedia, error) {

	err := r.db.Debug().Model(&socialMedia).Where("id = ?", id).Updates(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}

	fmt.Println("repo", socialMedia)

	return socialMedia, nil
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}
