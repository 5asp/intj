package main

import (
	"github.com/kzaun/intj/internal/fetcher"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		// configfx.FxInitConfigOption("app"),
		fetcher.Module,
	).Run()
}
