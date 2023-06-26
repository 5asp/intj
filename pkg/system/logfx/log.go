package logfx

import (
	"go.uber.org/zap"
)

// ProvideLogger to fx
func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	return slogger
}
