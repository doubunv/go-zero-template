package adminUser

import (
	"context"

	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneUserLogic {
	return &GetOneUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneUserLogic) GetOneUser(req *types.GetOneAdminUserReq) (resp *types.GetOneAdminUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
