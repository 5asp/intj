package miniofx

import (
	"context"
	"log"

	"github.com/knadh/koanf/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MinioClient struct {
	Oss      *minio.Client
	Location string
	Endpoint string
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
		Oss:      client,
		Location: p.Config.String("OSS.Location"),
		Endpoint: p.Config.String("OSS.Endpoint"),
	}
}

// type Oss struct {
// }

func (m *MinioClient) CreateBucket(ctx context.Context, bucketName string) {
	err := m.Oss.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: m.Location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.Oss.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func (m *MinioClient) FPutObject(ctx context.Context, bucketName, objectName, tmpPath, contentType string) (imageURL string, err error) {
	_, err = m.Oss.FPutObject(ctx, bucketName, objectName, tmpPath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	imageURL = "https://" + m.Endpoint + "/" + bucketName + "/" + objectName
	return imageURL, nil
}
