package lib

import (
	"context"
	"net"
	"net/http"

	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, slog *zap.SugaredLogger, conf *koanf.Koanf) *http.Server {
	port := conf.String("Port")
	srv := &http.Server{Addr: port, Handler: mux}
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
