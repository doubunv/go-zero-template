package middleware

import (
	"go-api/internal/config"
	"go-api/internal/pkg/appMiddleware"
	"net/http"
)

func WhiteHeaderPath() map[string]int {
	return map[string]int{
		"/user/test": 1,
	}
}

func CheckTokenHandle(c config.Config) appMiddleware.CheckRequestTokenFunc {
	return func(r *http.Request, token string) int64 {
		return 0
	}
}
