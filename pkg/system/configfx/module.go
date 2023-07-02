package configfx

import "go.uber.org/fx"

// Module provided to fx
var Module = fx.Option(
	fx.Provide(ProvideConfig),
)
