package minio

import (
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

type Service struct {
}

// NewServiceContainer builds the downstream services used throughout the application.
func NewService(logger *slog.Logger) *Service {
	return &Service{}
}

var Module = fx.Options(
	fx.Provide(NewService),
)
