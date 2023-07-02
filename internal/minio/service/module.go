package service

import (
	"github.com/kzaun/intj/pkg/system/miniofx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	miniofx.Module,
	fx.Provide(ProvideService),
)
