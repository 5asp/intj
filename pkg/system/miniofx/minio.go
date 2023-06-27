package miniofx

import (
	"go.uber.org/fx"
)

// ProvideConfig to fx
func ProvideMinio() fx.Option {
	return *&fx.NopLogger
}
