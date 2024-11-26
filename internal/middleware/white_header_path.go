package middleware

import (
	"go-api/internal/config"
	"go-api/internal/pkg/appMiddleware"
	"net/http"
)

func WhiteHeaderPath() map[string]int {
	return map[string]int{
		"/adminUser/userLogin": 1,
	}
}

func CheckTokenHandle(c config.Config) appMiddleware.CheckRequestTokenFunc {
	return func(r *http.Request, token string) int64 {
		return 1
	}
}
