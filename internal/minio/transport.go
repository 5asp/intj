package minio

import (
	"github.com/go-chi/chi"
	"golang.org/x/exp/slog"
)

func MakeHTTPHandler(s *Service, logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	return router
}
