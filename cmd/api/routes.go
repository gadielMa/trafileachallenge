package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

type AppHandler struct {
	DivisibleHandler     http.HandlerFunc
	NumberValueHandler   http.HandlerFunc
	CreateValueHandler   http.HandlerFunc
	GetCollectionHandler http.HandlerFunc
}

func (a *AppHandler) InitializeRoutes(r *chi.Mux) {
	r.Route("/trafilea", func(r chi.Router) {
		r.Get("/divisible", a.DivisibleHandler)

		r.Get("/number", a.NumberValueHandler)
		r.Post("/collection", a.CreateValueHandler)
		r.Get("/collection", a.GetCollectionHandler)
	})
}
