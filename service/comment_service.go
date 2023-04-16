package service

import (
	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/repository"
)

type CommentService interface {
	Create(payload entity.CommentRequest, userId uint) (entity.Comment, error)
	GetById(id uint) (entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	Update(payload entity.CommentRequest, id uint, userId uint) (entity.Comment, error)
	Delete(comment entity.Comment) (entity.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

// Create implements CommentService
func (s *commentService) Create(payload entity.CommentRequest, userId uint) (entity.Comment, error) {
	err := s.photoRepository.IsPhotoExist(payload.PhotoId)

	comment := entity.Comment{
		Message: payload.Message,
		PhotoId: payload.PhotoId,
		UserId:  userId,
	}

	if err != nil {
		return comment, err
	}

	newComment, err := s.commentRepository.Create(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

// Delete implements CommentService
func (s *commentService) Delete(comment entity.Comment) (entity.Comment, error) {
	comment, err := s.commentRepository.Delete(comment.ID)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

// GetAll implements CommentService
func (s *commentService) GetAll() ([]entity.Comment, error) {
	comments, err := s.commentRepository.FindAll()
	if err != nil {
		return comments, err
	}

	return comments, nil
}

// GetById implements CommentService
func (s *commentService) GetById(id uint) (entity.Comment, error) {
	comment, err := s.commentRepository.FindById(id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

// Update implements CommentService
func (s *commentService) Update(payload entity.CommentRequest, id uint, userId uint) (entity.Comment, error) {
	comment, err := s.commentRepository.FindById(id)
	if err != nil {
		return comment, err
	}

	comment.Message = payload.Message
	comment.PhotoId = payload.PhotoId
	comment.UserId = userId

	updatedComment, err := s.commentRepository.Update(comment, id)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func NewCommentService(comment repository.CommentRepository, photoRepository repository.PhotoRepository) CommentService {
	return &commentService{
		commentRepository: comment,
		photoRepository:   photoRepository,
	}
}
