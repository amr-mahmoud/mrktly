package server

import (
	"github.com/amr-mahmoud/mrktly/internal/interface/controllers/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer(userController *user.UserController) *chi.Mux {
	r := chi.NewRouter()

	// Essential middleware
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,


	)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userController.Routes())
	})




	return r
}