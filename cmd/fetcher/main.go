package main

import (
	fetcher "github.com/kzaun/intj/internal/job"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		// configfx.FxInitConfigOption("app"),
		fetcher.Module,
	).Run()
}
