package user

import (
	"admin-service/config"
	"admin-service/repo"
)

type Handler struct {
	cnf      *config.Config
	userrepo repo.Userrepo
}

func NewHandler(cnf *config.Config, userRepo repo.Userrepo) *Handler {
	return &Handler{
		cnf:      cnf,
		userrepo: userRepo,
	}
}
