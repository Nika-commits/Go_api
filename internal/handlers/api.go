package handlers

import (
	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
	// "github.com/Nika-commits/Go_api/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(mw.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins" GetCoinBalance)
	})
}


