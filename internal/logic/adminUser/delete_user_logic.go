package adminUser

import (
	"context"
	"errors"
	"go-api/internal/dao/model/admin"

	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteAdminUserReq) (resp *types.DeleteAdminUserResp, err error) {
	if req.UserId == 0 && req.Account != "" {
		adminInfo := l.AdminInfoModel.FindByAccount(req.Account)
		if adminInfo.ID == 0 {
			return nil, errors.New("account not found")
		}
		req.UserId = adminInfo.ID
	}

	err = l.AdminInfoModel.DeleteById(req.UserId)
	if err != nil {
		return nil, err
	}
	return
}
