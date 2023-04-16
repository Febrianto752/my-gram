package service

import (
	"errors"

	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/helper"
	"github.com/Febrianto752/my-gram/repository"
)

type UserService interface {
	Register(payload entity.UserRequest) (entity.User, error)
	Login(payload entity.UserLogin) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Login(payload entity.UserLogin) (entity.User, error) {
	email := payload.Email
	password := payload.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, err
	}

	comparePassword := helper.ComparePass([]byte(user.Password), []byte(password))
	if !comparePassword {
		return user, errors.New("Invalid email/password")
	}

	return user, nil
}

func (s *userService) Register(payload entity.UserRequest) (entity.User, error) {

	user := entity.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
	}

	newUser, err := s.userRepository.AddUser(user)
	if err != nil {

		return newUser, err
	}

	return newUser, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}
