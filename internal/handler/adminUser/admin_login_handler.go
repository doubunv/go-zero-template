package adminUser

import (
	"go-api/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-api/internal/logic/adminUser"
	"go-api/internal/svc"
	"go-api/internal/types"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminUserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := adminUser.NewAdminLoginLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogin(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
