package main

import (
	"github.com/kzaun/intj/internal/oss"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		// configfx.FxInitConfigOption("app"),
		oss.Module,
	).Run()
}
