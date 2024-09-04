package repository

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	GetByPostID(postID uint) ([]domain.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *domain.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetByPostID(postID uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
