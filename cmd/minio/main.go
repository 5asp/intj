package main

import (
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
			lib.NewHTTPServer,
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
