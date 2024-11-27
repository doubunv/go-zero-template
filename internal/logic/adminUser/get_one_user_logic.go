package adminUser

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-api/internal/dao/model/admin"
	"go-api/internal/svc"
	"go-api/internal/types"
	"go-api/pkg/headInfo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
	*admin.RoleModel
}

func NewGetOneUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneUserLogic {
	return &GetOneUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
		RoleModel:      admin.NewRoleModel(ctx, svcCtx),
	}
}

func (l *GetOneUserLogic) GetOneUser(req *types.GetOneAdminUserReq) (resp *types.GetOneAdminUserResp, err error) {
	resp = &types.GetOneAdminUserResp{}
	if req.UserId == 0 && req.Account != "" {
		adminInfo := l.AdminInfoModel.FindByAccount(req.Account)
		if adminInfo.ID == 0 {
			return nil, errors.New("account not found")
		}
		roleInfo := l.RoleModel.FindOne(adminInfo.RoleId)
		copier.Copy(resp, adminInfo)
		resp.RoleName = roleInfo.RoleName
		return
	}
	adminId := req.UserId
	if adminId == 0 {
		adminId = headInfo.GetTokenUid(l.ctx)
	}
	adminInfo := l.AdminInfoModel.FindOne(adminId)
	if adminInfo.ID == 0 {
		return nil, errors.New("account not found")
	}
	roleInfo := l.RoleModel.FindOne(adminInfo.RoleId)
	copier.Copy(resp, adminInfo)
	resp.RoleName = roleInfo.RoleName
	return
}
