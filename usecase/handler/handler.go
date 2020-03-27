package handler

import "github.com/go-chi/chi"

type Base interface {
	SetUpRoutes(*chi.Mux)
}
