package handler

import (
	"net/http"

	"github.com/cgressang/go-api-skeleton/usecase/interactor"
	"github.com/cgressang/go-api-skeleton/usecase/log"
	"github.com/go-chi/chi"
)

type HomeHandler struct {
	Logger     log.StdLogger
	Interactor interactor.Home
}

func NewHomeHandler(logger log.StdLogger, hi interactor.Home) *HomeHandler {
	return &HomeHandler{
		Logger:     logger,
		Interactor: hi,
	}
}

func (c *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(c.Interactor.GetMessage()))
}

func (c *HomeHandler) SetUpRoutes(r *chi.Mux) {
	r.Get("/", c.Home)
}
