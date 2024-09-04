package unitofwork

import "github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"

type UnitOfWork struct {
	userRepo    repository.UserRepository
	postRepo    repository.PostRepository
	commentRepo repository.CommentRepository
	tagRepo     repository.TagRepository
}

func NewUnitOfWork(
	userRepo repository.UserRepository,
	postRepo repository.PostRepository,
	commentRepo repository.CommentRepository,
	tagRepo repository.TagRepository,
) *UnitOfWork {
	return &UnitOfWork{
		userRepo:    userRepo,
		postRepo:    postRepo,
		commentRepo: commentRepo,
		tagRepo:     tagRepo,
	}
}

func (uow *UnitOfWork) UserRepository() repository.UserRepository {
	return uow.userRepo
}

func (uow *UnitOfWork) PostRepository() repository.PostRepository {
	return uow.postRepo
}

func (uow *UnitOfWork) CommentRepository() repository.CommentRepository {
	return uow.commentRepo
}

func (uow *UnitOfWork) TagRepository() repository.TagRepository {
	return uow.tagRepo
}
