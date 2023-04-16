package repository

import (
	"github.com/Febrianto752/my-gram/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entity.Comment) (entity.Comment, error)
	FindById(id uint) (entity.Comment, error)
	FindAll() ([]entity.Comment, error)
	Update(comment entity.Comment, id uint) (entity.Comment, error)
	Delete(id uint) (entity.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func (r *commentRepository) Create(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Debug().Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) Delete(id uint) (entity.Comment, error) {
	var comment entity.Comment

	err := r.db.Debug().Where("id = ?", id).Delete(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) FindAll() ([]entity.Comment, error) {
	var comments []entity.Comment

	err := r.db.Debug().Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (r *commentRepository) FindById(id uint) (entity.Comment, error) {
	var comment entity.Comment
	err := r.db.Debug().First(&comment, "id = ?", id).Error
	if err != nil {
		return comment, err
	}

	return comment, err
}

func (r *commentRepository) Update(comment entity.Comment, id uint) (entity.Comment, error) {
	var photo entity.Photo
	err := r.db.Debug().First(&photo, "id = ?", comment.PhotoId).Error

	if err != nil {
		return comment, err
	}

	err = r.db.Debug().Model(&comment).Where("id = ?", id).Updates(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}
