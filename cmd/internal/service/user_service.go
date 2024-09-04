package service

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/unitofwork"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo *unitofwork.UnitOfWork) *UserService {
	return &UserService{repo: repo.UserRepository()}
}

func (s *UserService) Register(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil || user.Password != password {
		return nil, err
	}
	return user, nil
}
