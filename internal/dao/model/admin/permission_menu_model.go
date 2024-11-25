package admin

import (
	"context"
	"encoding/json"
	"errors"
	"go-api/internal/dao/model"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"gorm.io/gorm"
	"time"
)

type PermissionMenuModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewPermissionMenuModel(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionMenuModel {
	return &PermissionMenuModel{
		ctx: ctx,
		db:  svcCtx.DbSelect.GetDb(ctx, model.DBAdmin),
	}
}

func (model *PermissionMenuModel) getDb() *gorm.DB {
	return model.db
}

func (model *PermissionMenuModel) FindOne(id int64) schema.PermissionMenu {
	var res schema.PermissionMenu

	dbRes := model.getDb().Model(&schema.PermissionMenu{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *PermissionMenuModel) InsertSchema(data *schema.PermissionMenu) error {
	data.CreatedAt = time.Now()
	dbRes := model.getDb().Model(&schema.PermissionMenu{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *PermissionMenuModel) UpdateByMap(id int64, data *schema.PermissionMenu) error {
	return model.getDb().Model(&schema.PermissionMenu{}).Where("id = ?", id).Updates(map[string]interface{}{}).Error
}

func (model *PermissionMenuModel) GetList(in *schema.PermissionMenu, pageQuery *model.PageQuery) (int64, []*schema.PermissionMenu, error) {
	builder := model.getDb().Model(&schema.PermissionMenu{})
	total := int64(0)
	err := builder.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	rows := make([]*schema.PermissionMenu, 0)
	err = builder.Offset(pageQuery.Offset()).Limit(pageQuery.PageSize).Find(&rows).Error
	return total, rows, err
}

func (model *PermissionMenuModel) GetMenuIds(permissionIds []int64) (res []int64, err error) {
	res = make([]int64, 0)
	rows := make([]*schema.PermissionMenu, 0)
	err = model.getDb().Model(&schema.PermissionMenu{}).Where("id in ?", permissionIds).Find(&rows).Error
	if err != nil {
		return
	}

	for _, v := range rows {
		ids := make([]int64, 0)
		json.Unmarshal([]byte(v.MenuIds), &ids)
		res = append(res, ids...)
	}

	return res, nil
}
