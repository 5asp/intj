package oss

import (
	"context"

	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *slog.Logger,
) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.InfoCtx(ctx, "start")
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
