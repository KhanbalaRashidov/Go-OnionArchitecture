package service

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/unitofwork"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo *unitofwork.UnitOfWork) *CommentService {
	return &CommentService{repo: repo.CommentRepository()}
}

func (s *CommentService) AddComment(comment *domain.Comment) error {
	return s.repo.Create(comment)
}

func (s *CommentService) GetCommentsByPostID(postID uint) ([]domain.Comment, error) {
	return s.repo.GetByPostID(postID)
}
