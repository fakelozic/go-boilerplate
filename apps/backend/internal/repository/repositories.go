package repository

import "github.com/fakelozic/go-boilerplate/internal/server"

type Repositories struct {}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}