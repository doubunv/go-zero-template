package adminUser

import (
	"go-api/internal/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-api/internal/logic/adminUser"
	"go-api/internal/svc"
	"go-api/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAdminUserReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := adminUser.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
