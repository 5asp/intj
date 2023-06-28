package minio

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

func NewRouter(logger *slog.Logger, s *Service) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}

var RModule = fx.Options(
	fx.Provide(NewRouter),
)
