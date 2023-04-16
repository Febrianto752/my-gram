package service

import (
	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/repository"
)

type SocialMediaService interface {
	Create(payload entity.SocialMediaRequest, userId uint) (entity.SocialMedia, error)
	GetById(id uint) (entity.SocialMedia, error)
	GetAll() ([]entity.SocialMedia, error)
	Update(payload entity.SocialMediaRequest, id uint, userId uint) (entity.SocialMedia, error)
	Delete(socialMedia entity.SocialMedia)
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

// Create implements SocialMediaService
func (s *socialMediaService) Create(payload entity.SocialMediaRequest, userId uint) (entity.SocialMedia, error) {
	socialMedia := entity.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         userId,
	}

	newSocialMedia, err := s.socialMediaRepository.Create(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}

	return newSocialMedia, nil
}

// Delete implements SocialMediaService
func (s *socialMediaService) Delete(socialMedia entity.SocialMedia) {
	s.socialMediaRepository.Delete(socialMedia.ID)
}

// GetAll implements SocialMediaService
func (s *socialMediaService) GetAll() ([]entity.SocialMedia, error) {
	socialMedias, err := s.socialMediaRepository.FindAll()
	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

// GetById implements SocialMediaService
func (s *socialMediaService) GetById(id uint) (entity.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(id)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

// Update implements SocialMediaService
func (s *socialMediaService) Update(payload entity.SocialMediaRequest, id uint, userId uint) (entity.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	// var socialMedia entity.SocialMedia

	socialMedia.Name = payload.Name
	socialMedia.SocialMediaUrl = payload.SocialMediaUrl
	socialMedia.UserId = userId

	updatedSocialMedia, err := s.socialMediaRepository.Update(socialMedia, id)
	if err != nil {
		return updatedSocialMedia, err
	}

	return updatedSocialMedia, nil
}

func NewSocialMediaService(socialMedia repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository: socialMedia}
}
