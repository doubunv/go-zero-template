package middleware

import (
	"context"
	"go-api/internal/logic/adminUser"
	"go-api/internal/pkg/appMiddleware"
	"go-api/internal/svc"
	"net/http"
)

func WhiteHeaderPath() map[string]int {
	return map[string]int{
		"/adminUser/userLogin": 1,
	}
}

func CheckTokenHandle(serverCtx *svc.ServiceContext) appMiddleware.CheckRequestTokenFunc {
	return func(r *http.Request, token string) int64 {
		return adminUser.NewAdminLoginLogic(context.Background(), serverCtx).CheckToken(token)
	}
}
