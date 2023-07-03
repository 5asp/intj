package minio

import (
	"github.com/kzaun/intj/internal/minio/handler"
	"github.com/kzaun/intj/internal/minio/service"
	"github.com/kzaun/intj/pkg/lib"
	"go.uber.org/fx"
)

var Module = fx.Module("minio",
	fx.Provide(
		service.ProvideService,
		// lib.AsRoute(handler.ProvideEchoHandler),
		lib.AsRoute(handler.ProvideUploadHandler),
	),
)
