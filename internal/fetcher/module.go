package fetcher

import "go.uber.org/fx"

// Module exported for go-fx depdency injection
var Module = fx.Options(
	// service.Module,
	// repository.Module,
	// handler.Module,
	fx.Invoke(registerHooks),
)
