package service

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/unitofwork"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(repo *unitofwork.UnitOfWork) *TagService {
	return &TagService{repo: repo.TagRepository()}
}

func (s *TagService) AddTag(tag *domain.Tag) error {
	return s.repo.Create(tag)
}

func (s *TagService) FindTagByName(name string) (*domain.Tag, error) {
	return s.repo.FindByName(name)
}
