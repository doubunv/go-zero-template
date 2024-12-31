package appMiddleware

import (
	"bytes"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/pkg/consts"
	"go-api/pkg/headInfo"
	"go-api/pkg/result"
	"google.golang.org/grpc/metadata"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
)

type ApiHeadOption func(m *ApiHeaderMiddleware)

func CloseVerifyOption(path map[string]int) ApiHeadOption {
	return func(m *ApiHeaderMiddleware) {
		m.noVerifyPath = path
	}
}

func WithDebugOption() ApiHeadOption {
	return func(m *ApiHeaderMiddleware) {
		m.debug = true
	}
}

type ApiHeaderMiddleware struct {
	NotVerify    bool
	debug        bool
	noVerifyPath map[string]int
}

func NewApiHeaderMiddleware(arg ...ApiHeadOption) *ApiHeaderMiddleware {
	res := &ApiHeaderMiddleware{}
	for _, o := range arg {
		o(res)
	}

	return res
}

func (m *ApiHeaderMiddleware) SetNoVerify(b bool) *ApiHeaderMiddleware {
	m.NotVerify = b
	return m
}

func (m *ApiHeaderMiddleware) headInMetadata(ctx context.Context, h headInfo.Head) context.Context {
	md := metadata.Pairs(
		consts.Token, h.AuthorizationJwt,
		consts.TokenUid, h.TokenUid,
		consts.ClientIp, h.ClientIp,
		consts.ReqPath, h.ReqPath,
		consts.Version, h.Version,
		consts.Source, h.Source,
	)

	ctxNew := metadata.NewOutgoingContext(ctx, md)
	return ctxNew
}

func (m *ApiHeaderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if m.debug {
					logc.Errorf(r.Context(), "ApiHeaderMiddleware error:%v, %s", err, string(debug.Stack()))
				}
				logc.Error(r.Context(), err, string(debug.Stack()))
				result.HttpErrorResult(r.Context(), w, errors.New("Server error. "))
				return
			}
		}()

		h := headInfo.GetHead(r)
		if r.Method != http.MethodGet && !m.NotVerify && m.verifyPath(r.URL.Path) {
			if err := h.Verify(); err != nil {
				result.HttpErrorResult(r.Context(), w, err)
				return
			}
		}

		newCtx := headInfo.ContextHeadInLog(r.Context(), h)
		newCtx = m.headInMetadata(newCtx, *h)
		newReq := r.WithContext(newCtx)

		body, err := io.ReadAll(newReq.Body)
		if err != nil {
			return
		}
		logc.Info(newCtx, "ApiRequest:"+string(body))
		newReq.Body = io.NopCloser(bytes.NewBuffer(body))

		next(w, newReq)
	}
}

func (m *ApiHeaderMiddleware) verifyPath(urlPath string) bool {
	if _, ok := m.noVerifyPath[urlPath]; ok {
		return false
	}
	for path, _ := range m.noVerifyPath {
		if strings.HasPrefix(path, "/") && strings.HasSuffix(path, "*") {
			prefix := strings.TrimSuffix(path, "*")
			if strings.HasPrefix(urlPath, prefix) {
				return false
			}
		}
	}
	return true
}
