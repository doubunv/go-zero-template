package middleware

import (
	"context"
	"go-api/internal/logic/adminUser"
	"go-api/internal/svc"
	"go-api/pkg/appMiddleware"
	"net/http"
)

func WhiteHeaderPath() map[string]int {
	return map[string]int{
		"/swagger":             1,
		"/adminUser/userLogin": 1,
	}
}

func CheckTokenHandle(serverCtx *svc.ServiceContext) appMiddleware.CheckRequestTokenFunc {
	return func(r *http.Request, token string) int64 {
		return adminUser.NewAdminLoginLogic(context.Background(), serverCtx).CheckToken(token)
	}
}
