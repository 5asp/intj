package configfx

import "go.uber.org/fx"

// ProvideConfig to fx
func ProvideConfig() {
	// logger, _ := zap.NewProduction()
	// slogger := logger.Sugar()

	return
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
