package logfx

import (
	"os"

	"golang.org/x/exp/slog"
)

// ProvideLogger to fx
func ProvideLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	return logger
}
