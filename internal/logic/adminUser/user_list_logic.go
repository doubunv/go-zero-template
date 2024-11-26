package adminUser

import (
	"context"
	"go-api/internal/dao/model/admin"

	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
	}
}

func (l *UserListLogic) UserList(req *types.AdminUserListReq) (resp *types.AdminUserListResp, err error) {
	return
}
