package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalmiddlrwares []Middleware
}

func Newmanager() *Manager {
	return &Manager{
		globalmiddlrwares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalmiddlrwares = append(mngr.globalmiddlrwares, middlewares...)
	return mngr
}

func (mngr *Manager) With(Handler http.Handler, middlewares ...Middleware) http.Handler {

	h := Handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	for _, globalmiddleware := range mngr.globalmiddlrwares {
		h = globalmiddleware(h)
	}
	return h

}

func (mngr *Manager) Wrapmux(Handler http.Handler, middlewares ...Middleware) http.Handler {

	h := Handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h

}
