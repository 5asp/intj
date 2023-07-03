package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kzaun/intj/internal/minio/service"
	"github.com/kzaun/intj/pkg/lib"
	reutrns "github.com/kzaun/intj/pkg/lib/returns"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UploadHandler struct {
	log *zap.Logger
	svc *service.Service
}

type UploadParams struct {
	fx.In
	Logger  *zap.Logger `optional:"true"`
	Service *service.Service
}

func ProvideUploadHandler(p UploadParams) *UploadHandler {
	return &UploadHandler{
		log: p.Logger,
		svc: p.Service,
	}
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte(reutrns.OK.WithData("successful").ToString()))
	case "POST":
		bucketName := r.FormValue("bucket")
		file, head, err := r.FormFile("image")
		if err != nil {
			fmt.Printf("Failed to get file data %s\n", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create newFile data %s\n", err.Error())
			return
		}

		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			msg := fmt.Sprintf("failed to save into newFile=%s failed,err=%s", head.Filename, err.Error())
			w.Write([]byte(reutrns.Err.WithData(msg).ToString()))
		}

		contentType, _ := lib.GetFileContentType(newFile)
		if contentType == "" {
			h.log.Info(fmt.Sprintf("contentType: %s", "not be null."))
			w.Write([]byte(reutrns.Err.WithData("content type not be null.").ToString()))
			return
		}
		fileOssURL, err := h.svc.UploadToOss(bucketName, newFile.Name(), contentType)
		if err != nil {
			msg := fmt.Sprintf("upload filename=%s failed. err=%s", head.Filename, err.Error())
			w.Write([]byte(reutrns.Err.WithData(msg).ToString()))
		}
		w.Write([]byte(reutrns.OK.WithData(fileOssURL).ToString()))
	}
	return
}
func (*UploadHandler) Pattern() string {
	return "/upload"
}
