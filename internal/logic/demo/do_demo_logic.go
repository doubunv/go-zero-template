package demo

import (
	"context"
	"go-api/internal/dao/model"

	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*model.DemoModel
}

func NewDoDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoDemoLogic {
	return &DoDemoLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		DemoModel: model.NewDemoModel(ctx, svcCtx.DbSelect),
	}
}

func (l *DoDemoLogic) DoDemo(req *types.DemoReq) (resp *types.DemoResp, err error) {
	info := l.DemoModel.FindOne(1)
	return &types.DemoResp{
		Name: info.Name,
	}, nil
}
