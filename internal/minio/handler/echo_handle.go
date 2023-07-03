package handler

import (
	"net/http"

	"github.com/kzaun/intj/internal/minio/service"
	reutrns "github.com/kzaun/intj/pkg/lib/returns"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger
	svc *service.Service
}

type Params struct {
	fx.In
	Logger  *zap.Logger `optional:"true"`
	Service *service.Service
}

func ProvideEchoHandler(p Params) *EchoHandler {
	return &EchoHandler{
		log: p.Logger,
		svc: p.Service,
	}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte(reutrns.OK.WithData("successful").ToString()))
	case "POST":

	}
	h.svc.Upload()
	// if _, err := io.Copy(w, r.Body); err != nil {
	// 	h.log.Warn("Failed to handle request", zap.Error(err))
	// }
}
func (*EchoHandler) Pattern() string {
	return "/echo"
}
