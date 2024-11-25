package admin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/dao/dto"
	"go-api/internal/dao/model"
	"go-api/internal/svc"
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

func (model *AdminRoleModel) getAdminRoleMenuByCache(roleId int64) (res map[int64]*dto.MenuTree, err error) {
	return
}

func (model *AdminRoleModel) setAdminRoleMenuByCache(roleId int64, data map[int64]*dto.MenuTree) {
	return
}

func (model *AdminRoleModel) DeleteAdminRoleMenuByCache(roleId int64) {
	return
}

func (model *AdminRoleModel) GetAdminRoleMenu(roleId int64) (res map[int64]*dto.MenuTree, err error) {
	res, err = model.getAdminRoleMenuByCache(roleId)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		return
	}

	res = make(map[int64]*dto.MenuTree)
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

	menuInfo, err := model.MenuModel.GetMenuTreeByIds(menuIds)
	if err != nil {
		return nil, err
	}
	if len(menuInfo) == 0 {
		logc.Info(model.ctx, "GetAdminRoleMenu", "role not found menu info list")
		return
	}

	model.setAdminRoleMenuByCache(roleId, menuInfo)

	return menuInfo, nil
}
