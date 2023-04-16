package service

import (
	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/repository"
)

type PhotoService interface {
	Create(payload entity.PhotoRequest, userId uint) (entity.Photo, error)
	GetById(id uint) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(payload entity.PhotoRequest, id uint, userId uint) (entity.Photo, error)
	Delete(photo entity.Photo)
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func (s *photoService) Create(payload entity.PhotoRequest, userId uint) (entity.Photo, error) {
	photo := entity.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserId:   userId,
	}

	newPhoto, err := s.photoRepository.Create(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

func (s *photoService) Delete(photo entity.Photo) {
	s.photoRepository.Delete(photo.ID)
}

func (s *photoService) GetAll() ([]entity.Photo, error) {
	photos, err := s.photoRepository.FindAll()
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (s *photoService) GetById(id uint) (entity.Photo, error) {
	photo, err := s.photoRepository.FindById(id)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *photoService) Update(payload entity.PhotoRequest, id uint, userId uint) (entity.Photo, error) {
	photo, err := s.photoRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	photo.Title = payload.Title
	photo.Caption = payload.Caption
	photo.PhotoUrl = payload.PhotoUrl
	photo.UserId = userId

	updatedPhoto, err := s.photoRepository.Update(photo, id)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil
}

func NewPhotoService(photo repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository: photo}
}
