package middleware

import "admin-service/config"

type middleware struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *middleware {
	return &middleware{
		cnf: cnf,
	}
}
