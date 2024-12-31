package result

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-api/pkg/result/xcode"
	"net/http"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HttpSuccessResult(ctx context.Context, w http.ResponseWriter, resp interface{}) {
	success := Success(resp, trace.TraceIDFromContext(ctx))
	go func() {
		logSucc, _ := json.Marshal(success)
		logc.Info(ctx, "ApiResponse:", fmt.Sprintf("%s", string(logSucc)))
	}()

	httpx.WriteJsonCtx(ctx, w, http.StatusOK, success)
}

func HttpErrorResult(ctx context.Context, w http.ResponseWriter, err error) {
	var (
		xerr xcode.XCode
		code int
		msg  string
	)
	if errors.As(err, &xerr) {
		code = xerr.Code()
		msg = xerr.Error()
	} else {
		code = http.StatusInternalServerError
		msg = err.Error()
	}

	resp := Error(code, msg, trace.TraceIDFromContext(ctx))

	go func() {
		logSuc, _ := json.Marshal(resp)
		logc.Info(ctx, "ApiResponse:", string(logSuc))
	}()

	httpx.WriteJsonCtx(ctx, w, http.StatusOK, resp)
}
