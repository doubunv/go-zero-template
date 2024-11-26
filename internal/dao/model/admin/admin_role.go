package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/dao/model"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"go-api/internal/types"
	"gorm.io/gorm"
)

type AdminRoleModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
	*MenuModel
	*PermissionMenuModel
	*RoleModel
}

func NewAdminRoleModel(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRoleModel {
	return &AdminRoleModel{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		db:                  svcCtx.DbSelect.GetDb(ctx, model.DBAdmin),
		MenuModel:           NewMenuModel(ctx, svcCtx),
		PermissionMenuModel: NewPermissionMenuModel(ctx, svcCtx),
		RoleModel:           NewRoleModel(ctx, svcCtx),
	}
}

func (model *AdminRoleModel) getAdminRoleMenuByCache(roleId int64) (res []*types.LoginRoleMenuItem, err error) {
	return
}

func (model *AdminRoleModel) setAdminRoleMenuByCache(roleId int64, data []*types.LoginRoleMenuItem) {
	return
}

func (model *AdminRoleModel) DeleteAdminRoleMenuByCache(roleId int64) {
	return
}

func (model *AdminRoleModel) GetAdminRoleMenu(roleId int64) (res []*types.LoginRoleMenuItem, err error) {
	res, err = model.getAdminRoleMenuByCache(roleId)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		return
	}

	permissionIds, err := model.RoleModel.GetRolePermissionById(roleId)
	if err != nil {
		return nil, err
	}
	if len(permissionIds) == 0 {
		logc.Info(model.ctx, "GetAdminRoleMenu", "role not found permission id")
		return
	}
	menuIds, err := model.PermissionMenuModel.GetMenuIds(permissionIds)
	if err != nil {
		return nil, err
	}
	if len(menuIds) == 0 {
		logc.Info(model.ctx, "GetAdminRoleMenu", "role not found menu id")
		return
	}

	menuInfo, err := model.MenuModel.GetMenuByIds(menuIds)
	if err != nil {
		return nil, err
	}
	if len(menuInfo) == 0 {
		logc.Info(model.ctx, "GetAdminRoleMenu", "role not found menu info list")
		return
	}

	menuTree := model.buildMenuTree(menuInfo)
	model.setAdminRoleMenuByCache(roleId, menuTree)

	return menuTree, nil
}

func (model *AdminRoleModel) buildMenuTree(menus []*schema.Menu) []*types.LoginRoleMenuItem {
	MenuTree := make(map[int64]*types.LoginRoleMenuItem)
	MenuTreeTemp := make([]*types.LoginRoleMenuItem, 0)
	var tree []*types.LoginRoleMenuItem
	if len(menus) == 0 {
		return tree
	}

	for _, v := range menus {
		menuTree := &types.LoginRoleMenuItem{}
		copier.Copy(menuTree, v)
		MenuTree[v.ID] = menuTree
		MenuTreeTemp = append(MenuTreeTemp, menuTree)
	}

	for _, menu := range MenuTreeTemp {
		if menu.MenuPid == 0 {
			tree = append(tree, menu)
		} else {
			if parent, exists := MenuTree[menu.MenuPid]; exists {
				parent.Children = append(parent.Children, menu)
			}
		}
	}
	return tree
}
