package svc

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Service defines the extension handlers.
type Service interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	GetMe(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
}

// NewService returns a service implementation for Service.
func NewService(opts ...Option) Service {
	options := newOptions(opts...)

	m := chi.NewMux()
	m.Use(options.Middleware...)

	svc := Graph{
		config: options.Config,
		mux:    m,
		logger: &options.Logger,
	}

	m.Route("/v1.0", func(r chi.Router) {
		r.Get("/me", svc.GetMe)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", svc.GetUsers)
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(svc.UserCtx)
				r.Get("/", svc.GetUser)
			})
		})
	})

	return svc
}
