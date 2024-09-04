package main

import (
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/config"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/handler"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/repository"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/service"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/unitofwork"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := mux.NewRouter()

	userRepo := repository.NewGormUserRepository(config.DB)
	postRepo := repository.NewGormPostRepository(config.DB)
	commentRepo := repository.NewGormCommentRepository(config.DB)
	tagRepo := repository.NewGormTagRepository(config.DB)

	uow := unitofwork.NewUnitOfWork(userRepo, postRepo, commentRepo, tagRepo)

	userService := service.NewUserService(uow)
	postService := service.NewPostService(uow)
	commentService := service.NewCommentService(uow)
	tagService := service.NewTagService(uow)

	apiHandler := handler.NewAPIHandler(userService, postService, commentService, tagService)

	r.HandleFunc("/register", apiHandler.Register).Methods("POST")
	r.HandleFunc("/posts", apiHandler.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", apiHandler.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", apiHandler.DeletePost).Methods("DELETE")
	r.HandleFunc("/posts/{id}/like", apiHandler.LikePost).Methods("POST")
	r.HandleFunc("/comments", apiHandler.AddComment).Methods("POST")
	r.HandleFunc("/tags", apiHandler.AddTag).Methods("POST")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
