package job

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/svc"
)

type ActiveCount struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	maxId  int64
}

func NewActiveCount(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveCount {
	return &ActiveCount{
		ctx:    ctx,
		svcCtx: svcCtx,
		maxId:  0,
	}
}

func (j *ActiveCount) Run() {
	defer func() {
		if r := recover(); r != nil {
			logc.Error(j.ctx, "ActiveCount:", fmt.Sprintf("Recovered in f, r=%v\n", r))
		}
	}()

}
