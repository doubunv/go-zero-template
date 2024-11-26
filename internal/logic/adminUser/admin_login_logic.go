package adminUser

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/dao/model/admin"
	"go-api/internal/dao/schema"
	"go-api/internal/pkg/commonTool"
	"go-api/internal/pkg/jwtToken"
	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
	*admin.AdminLoginTokenModel
	*admin.AdminRoleModel
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		AdminInfoModel:       admin.NewAdminInfoModel(ctx, svcCtx),
		AdminLoginTokenModel: admin.NewAdminLoginTokenModel(ctx, svcCtx),
		AdminRoleModel:       admin.NewAdminRoleModel(ctx, svcCtx),
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminUserLoginReq) (resp *types.AdminUserLoginResp, err error) {
	resp = &types.AdminUserLoginResp{
		Token:            "",
		RoleMenuItemList: make([]*types.LoginRoleMenuItem, 0),
	}
	adminInfo := l.AdminInfoModel.FindByAccount(req.Account)
	if adminInfo.ID == 0 {
		return nil, errors.New("account not found")
	}
	if !commonTool.CheckPassword(adminInfo.Password, req.Password, adminInfo.PasswordSign) {
		return nil, errors.New("account password error")
	}
	if adminInfo.Status != schema.AdminInfoStatus1 {
		return nil, errors.New("account status not normal")
	}
	if adminInfo.RoleId == 0 {
		return nil, errors.New("account role is null")
	}

	// 生成token
	tokenSign := commonTool.GenerateRandomString(8)
	token, err := jwtToken.Generate2Token(adminInfo.ID, "", tokenSign, adminInfo.RoleId)
	if err != nil {
		logc.Error(l.ctx, "AdminLogin:", err)
		return nil, errors.New("token Generate error")
	}

	loginToken := &schema.AdminLoginToken{
		AdminId:   adminInfo.ID,
		TokenSign: tokenSign,
	}
	err = l.AdminLoginTokenModel.AddLoginToken(loginToken)
	if err != nil {
		logc.Error(l.ctx, "AdminLogin:", err)
		return nil, errors.New("save login token error")
	}

	//获取角色对应的菜单tree
	resp.RoleMenuItemList, err = l.AdminRoleModel.GetAdminRoleMenu(adminInfo.RoleId)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	return
}
