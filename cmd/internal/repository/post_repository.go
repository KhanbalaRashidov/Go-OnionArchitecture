package repository

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *domain.Post) error
	Update(post *domain.Post) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	IncrementLikes(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Update(post *domain.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Post{}, id).Error
}

func (r *postRepository) FindByID(id uint) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) IncrementLikes(id uint) error {
	return r.db.Model(&domain.Post{}).Where("id = ?", id).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}
