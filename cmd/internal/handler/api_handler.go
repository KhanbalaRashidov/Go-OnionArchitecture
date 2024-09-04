package handler

import (
	"encoding/json"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/domain"
	"github.com/KhanbalaRashidov/Go-OnionArcitechture.git/cmd/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type APIHandler struct {
	userService    *service.UserService
	postService    *service.PostService
	commentService *service.CommentService
	tagService     *service.TagService
}

func NewAPIHandler(userService *service.UserService, postService *service.PostService, commentService *service.CommentService, tagService *service.TagService) *APIHandler {
	return &APIHandler{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		tagService:     tagService,
	}
}

func (h *APIHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.userService.Register(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.postService.Create(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.ID = uint(postID)
	if err := h.postService.Update(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *APIHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	if err := h.postService.Delete(uint(postID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *APIHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	if err := h.postService.LikePost(uint(postID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *APIHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.commentService.AddComment(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) AddTag(w http.ResponseWriter, r *http.Request) {
	var tag domain.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.tagService.AddTag(&tag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
