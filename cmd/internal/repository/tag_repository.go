package repository

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"gorm.io/gorm"
)

type TagRepository interface {
	Create(tag *domain.Tag) error
	FindByName(name string) (*domain.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewGormTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(tag *domain.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) FindByName(name string) (*domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
