package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-api/internal/svc"
)

func swaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := filepath.Abs("swagger/api.json")

		_, err := os.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		http.ServeFile(w, r, s)
	}
}
