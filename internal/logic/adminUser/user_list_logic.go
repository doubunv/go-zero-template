package adminUser

import (
	"context"
	"github.com/jinzhu/copier"
	"go-api/internal/dao/model"
	"go-api/internal/dao/model/admin"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"go-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*admin.AdminInfoModel
	*admin.RoleModel
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		AdminInfoModel: admin.NewAdminInfoModel(ctx, svcCtx),
		RoleModel:      admin.NewRoleModel(ctx, svcCtx),
	}
}

func (l *UserListLogic) UserList(req *types.AdminUserListReq) (resp *types.AdminUserListResp, err error) {
	resp = &types.AdminUserListResp{
		List:     make([]types.AdminUserListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}
	qData := &schema.AdminInfo{
		Account:      "",
		Name:         "",
		Password:     "",
		PasswordSign: "",
		RoleId:       0,
		Status:       0,
	}
	total, list, err := l.AdminInfoModel.GetList(qData, model.NewPageQuery(int(req.Page), int(req.PageSize)))
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return
	}

	roleList, err := l.RoleModel.GetAllRoleByCache()
	if err != nil {
		return nil, err
	}

	roleListMap := make(map[int64]string)
	for _, v := range roleList {
		roleListMap[v.ID] = v.RoleName
	}

	resp.Total = total
	for _, v := range list {
		item := types.AdminUserListItem{}
		copier.Copy(item, v)

		if name, ok := roleListMap[v.RoleId]; ok {
			item.RoleName = name
		}
		resp.List = append(resp.List, item)
	}

	return
}
