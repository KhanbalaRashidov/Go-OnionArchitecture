package service

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/unitofwork"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo *unitofwork.UnitOfWork) *PostService {
	return &PostService{repo: repo.PostRepository()}
}

func (s *PostService) Create(post *domain.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) Update(post *domain.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *PostService) GetAll() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) LikePost(id uint) error {
	return s.repo.IncrementLikes(id)
}

func (s *PostService) FindByID(id uint) (*domain.Post, error) {
	return s.repo.FindByID(id)
}
