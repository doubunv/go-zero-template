package adminUser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-api/internal/logic/adminUser"
	"go-api/internal/svc"
	"go-api/internal/types"
)

func GetOneUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOneAdminUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminUser.NewGetOneUserLogic(r.Context(), svcCtx)
		resp, err := l.GetOneUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
