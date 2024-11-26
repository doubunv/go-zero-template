package adminUser

import (
	"context"
	"github.com/pkg/errors"
	"go-api/internal/dao/model/admin"
	"go-api/internal/dao/schema"
	"go-api/internal/pkg/commonTool"
	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateAdminUserReq) (resp *types.CreateAdminUserResp, err error) {
	resp = &types.CreateAdminUserResp{}
	accountInfo := l.AdminInfoModel.FindByAccount(req.Account)
	if accountInfo.ID != 0 {
		return nil, errors.New("Account already exists")
	}
	sign := commonTool.GenerateRandomString(8)
	admin := &schema.AdminInfo{
		Account:      req.Account,
		Name:         req.Name,
		Password:     commonTool.BuildPassword(req.Password, sign),
		PasswordSign: sign,
		RoleId:       req.RoleId,
	}
	err = l.AdminInfoModel.InsertSchema(admin)
	if err != nil {
		return nil, err
	}
	resp.Id = admin.ID
	return
}
