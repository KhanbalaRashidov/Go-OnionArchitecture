package repository

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
