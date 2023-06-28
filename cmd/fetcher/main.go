package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/kzaun/intj/internal/minio"
	"github.com/kzaun/intj/pkg/system/logfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// ...
		fx.Provide(logfx.ProvideLogger),
		// fx.Provide(minio.NewRouter),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	l *slog.Logger,
	// r *chi.Mux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				l.InfoCtx(ctx, "ok")
				ser := minio.NewService(l)
				r := minio.NewRouter(l, ser)
				l.InfoCtx(ctx, fmt.Sprintf("routes successfully initialized, now listening on port %d", 9000))
				go http.ListenAndServe(fmt.Sprintf(":%d", 9000), r)
				return nil
			},
			OnStop: func(context.Context) error {

				return nil
			},
		},
	)
}
