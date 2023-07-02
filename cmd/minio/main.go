package main

import (
	"context"
	"net"
	"net/http"

	"github.com/kzaun/intj/internal/minio"
	"github.com/kzaun/intj/pkg/lib"
	"github.com/kzaun/intj/pkg/system/configfx"
	"github.com/kzaun/intj/pkg/system/logfx"
	"github.com/kzaun/intj/pkg/system/miniofx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		minio.Module,
		fx.Provide(
			NewHTTPServer,
			configfx.ProvideConfig,
			miniofx.ProvideMinio,
			logfx.ProvideLogger,
			fx.Annotate(
				lib.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			zap.NewDevelopment,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, slog *zap.SugaredLogger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			slog.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := syncLog(slog)
			if err != nil {
				return err
			}
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
func syncLog(slog *zap.SugaredLogger) error {
	return slog.Sync()
}
