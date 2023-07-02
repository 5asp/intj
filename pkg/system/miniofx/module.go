package miniofx

import (
	"github.com/kzaun/intj/pkg/system/configfx"
	"go.uber.org/fx"
)

// Module provided to fx
var Module = fx.Options(
	configfx.Module,
	fx.Provide(ProvideMinio),
)
