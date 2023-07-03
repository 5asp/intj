package service

import (
	"context"
	"path"
	"strconv"
	"time"

	"github.com/kzaun/intj/pkg/system/miniofx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Service struct {
	OssClient *miniofx.MinioClient
	Log       *zap.Logger
}
type Params struct {
	fx.In
	MinioSDK *miniofx.MinioClient
	Logger   *zap.Logger
}

type Result struct {
	fx.Out
	svc *Service
}

// ProvideServiceContainer builds the downstream services used throughout the application.
func ProvideService(p Params) *Service {
	// p.Logg.
	return &Service{
		OssClient: p.MinioSDK,
		Log:       p.Logger,
	}
}

func (s *Service) UploadToOss(bucketName, tmpPath, contentType string) (imageURL string, err error) {
	ctx := context.TODO()
	s.OssClient.CreateBucket(ctx, bucketName)
	filesuffix := path.Ext(tmpPath)
	objectName := strconv.Itoa(int(time.Now().Unix())) + filesuffix
	// Upload the zip file with FPutObject
	imageURL, err = s.OssClient.FPutObject(ctx, bucketName, objectName, tmpPath, contentType)
	if err != nil {
		s.Log.Error(err.Error())
		return "", err
	}
	return imageURL, nil
}
