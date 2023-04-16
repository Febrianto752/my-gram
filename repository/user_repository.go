package repository

import (
	"github.com/Febrianto752/my-gram/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	IsUsernameExists(username string) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) IsUsernameExists(username string) error {
	var user entity.User
	err := r.db.Debug().Select("id", "username").Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Debug().First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) AddUser(user entity.User) (entity.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {

		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
