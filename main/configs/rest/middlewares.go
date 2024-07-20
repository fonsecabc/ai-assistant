package configs

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func PrepareMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentEncoding("json"))
}
