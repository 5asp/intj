package service

import (
	"fmt"

	"github.com/kzaun/intj/pkg/system/miniofx"
	"go.uber.org/fx"
)

type Service struct {
	MinioSDK *miniofx.MinioClient
	// Minio *minio.Client
}
type Params struct {
	fx.In
	MinioSDK *miniofx.MinioClient
	// Logg     *logfx.Client
	// Oss    *minio.Client
	// Logger *zap.Logger `optional:"true"`
}

type Result struct {
	fx.Out
	svc *Service
	// Client *Client
}

// ProvideServiceContainer builds the downstream services used throughout the application.
func ProvideService(p Params) *Service {
	// p.Logg.
	return &Service{
		MinioSDK: p.MinioSDK,
	}
}

func (s *Service) Upload() {
	// miniofx.ProvideMinio()
	fmt.Println("welcome")
}
