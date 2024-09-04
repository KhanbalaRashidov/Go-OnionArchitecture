# Go Onion Architecture  Blog API

A professional blog API built with Go, utilizing PostgreSQL for data storage, and implementing best practices such as Repository and Unit of Work patterns, along with Onion Architecture principles. This API supports user registration, login, post creation, updating, deletion, liking, commenting, tagging, and JWT authentication.


````
/Go-OnionArcitechture
├── /cmd
│   └── main.go
├── /internal
│   ├── /domain
│   │   ├── post.go
│   │   ├── user.go
│   │   ├── comment.go
│   │   └── tag.go
│   ├── /repository
│   │   ├── post_repository.go
│   │   ├── user_repository.go
│   │   ├── comment_repository.go
│   │   └── tag_repository.go
│   ├── /service
│   │   ├── post_service.go
│   │   ├── user_service.go
│   │   ├── comment_service.go
│   │   └── tag_service.go
│   ├── /handler
│   │   └── api_handler.go
│   ├── /middleware
│   │   └── auth_middleware.go
│   ├── /unitofwork
│   │   └── unit_of_work.go
│   ├── /config
│   │   └── config.go
│   └── /storage
│       └── file_storage.go
└── go.mod

````

## Features

- User registration and login
- Create, update, delete posts
- Like posts
- Comment on posts
- Tagging posts
- JWT authentication
- PostgreSQL as the database

## Tech Stack

- Go
- GORM (ORM for Go)
- PostgreSQL
- Gorilla Mux (for routing)
- JWT for authentication
- Docker
- Kubernetes

## Prerequisites

- Go installed on your machine.
- PostgreSQL server running.
- Docker installed.
- Kubernetes cluster (local or remote).

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/KhanbalaRashidov/Go-OnionArcitechture.git
cd Go-OnionArcitechture
