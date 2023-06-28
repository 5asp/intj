package logfx

import (
	"os"

	"golang.org/x/exp/slog"
)

// ProvideLogger to fx
func ProvideLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	return logger
}
