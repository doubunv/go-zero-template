package headInfo

import (
	"context"
	"encoding/json"
	"go-api/internal/pkg/consts"
	"go.opentelemetry.io/otel/trace"
	"net"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	HeaderToken = consts.HeaderToken
)

type Head struct {
	AuthorizationJwt string `json:"authorization_jwt"` // 用户token
	Version          string `json:"version"`           // APP版本
	Source           string `json:"source"`            // 来源渠道	* Android * Ios * Pc
	ClientIp         string `json:"client_ip"`         // 客户端IP
	Trace            string `json:"trace"`             // 链路路由
	TokenUid         string `json:"token_uid"`         // 用户ID
}

func GetHead(r *http.Request) *Head {
	header := r.Header
	return &Head{
		AuthorizationJwt: strings.Trim(header.Get(consts.HeaderToken), " "),
		Version:          strings.Trim(header.Get("Version"), " "),
		Source:           strings.Trim(header.Get("Source"), " "),
		ClientIp:         getClientIP(r),
		TokenUid:         strings.Trim(header.Get("TokenUid"), " "),
		Trace:            trace.SpanContextFromContext(r.Context()).TraceID().String(),
	}
}

func (h *Head) Verify() error {
	return nil
}

func (h *Head) String() string {
	data, _ := json.Marshal(h)
	return string(data)
}

func ContextHeadInLog(ctx context.Context, h *Head) context.Context {
	ctxNew := logx.ContextWithFields(ctx,
		logx.Field(consts.HeaderToken, h.AuthorizationJwt),
		logx.Field("Version", h.Version),
		logx.Field("Source", h.Source),
		logx.Field("ClientIp", h.ClientIp),
		logx.Field("Trance", h.Trace),
		logx.Field("TokenUid", h.TokenUid),
	)
	return ctxNew
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("x_forwarded_realip")
	if ip == "" {
		ip = r.Header.Get("X-Real-Ip")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func GetFullHead(r *http.Request) map[string][]string {
	headers := make(map[string][]string)

	for k, v := range r.Header {
		headers[k] = v
	}

	return headers
}
