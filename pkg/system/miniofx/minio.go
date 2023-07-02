package miniofx

import (
	"github.com/knadh/koanf/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MinioClient struct {
	*minio.Client
}
type Params struct {
	fx.In
	Config *koanf.Koanf
	Logger *zap.Logger `optional:"true"`
}

// ProvideMinio builds the downstream services used throughout the application.
func ProvideMinio(p Params) *MinioClient {
	client, err := minio.New(p.Config.String("OSS.Endpoint"), &minio.Options{
		Creds:  credentials.NewStaticV4(p.Config.String("OSS.AccessKey"), p.Config.String("OSS.SecretKey"), ""),
		Secure: p.Config.Bool("OSS.UseSSL"),
	})
	if err != nil {
		panic(err)
	}
	p.Logger.Debug("minio client connect successful.")
	return &MinioClient{
		client,
	}
}
