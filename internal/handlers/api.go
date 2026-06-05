package handlers

import (
	"github.com/Nika-commits/Go_api/internal/middleware"
	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(mw.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
