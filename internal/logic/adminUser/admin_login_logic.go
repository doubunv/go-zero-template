package adminUser

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/dao/model/mysql"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"go-api/internal/types"
	"go-api/pkg/commonTool"
	"go-api/pkg/jwtToken"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*mysql.AdminInfoModel
	*mysql.AdminLoginTokenModel
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		AdminInfoModel:       mysql.NewAdminInfoModel(ctx, svcCtx),
		AdminLoginTokenModel: mysql.NewAdminLoginTokenModel(ctx, svcCtx),
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminUserLoginReq) (resp *types.AdminUserLoginResp, err error) {
	resp = &types.AdminUserLoginResp{
		Token: "",
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

	resp.Token = token
	return
}

func (l *AdminLoginLogic) CheckToken(token string) int64 {
	parseToken, err := jwtToken.ParseRefreshToken(token)
	if err != nil {
		return 0
	}
	adminInfo := l.AdminInfoModel.FindOne(parseToken.Uid)
	if adminInfo.ID == 0 ||
		adminInfo.Status != schema.AdminInfoStatus1 {
		return 0
	}

	loginTokenInfo := l.AdminLoginTokenModel.FindOneByAdminId(adminInfo.ID)
	if loginTokenInfo.ID == 0 || parseToken.Sign != loginTokenInfo.TokenSign {
		return 0
	}

	return parseToken.Uid
}
