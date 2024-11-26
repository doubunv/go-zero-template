package adminUser

import (
	"context"
	"go-api/internal/dao/model/admin"
	"go-api/internal/dao/schema"
	"go-api/internal/pkg/commonTool"
	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateAdminUserReq) (resp *types.UpdateAdminUserResp, err error) {
	upData := &schema.AdminInfo{
		Account:      req.Account,
		Name:         req.Name,
		Password:     "",
		PasswordSign: "",
		RoleId:       req.RoleId,
		Status:       req.Status,
	}

	if req.Password != "" {
		sign := commonTool.GenerateRandomString(8)
		upData.Password = commonTool.BuildPassword(req.Password, sign)
		upData.PasswordSign = sign
	}

	err = l.AdminInfoModel.UpdateByMap(req.Id, upData)
	if err != nil {
		return nil, err
	}
	return
}
