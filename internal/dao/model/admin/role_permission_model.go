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

type RoleModel struct {
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	db                *gorm.DB
	RolePermissionKey string
}

func NewRoleModel(ctx context.Context, svcCtx *svc.ServiceContext) *RoleModel {
	return &RoleModel{
		ctx:               ctx,
		svcCtx:            svcCtx,
		db:                svcCtx.DbSelect.GetDb(ctx, model.DBAdmin),
		RolePermissionKey: "rolePermission:",
	}
}

func (model *RoleModel) getDb() *gorm.DB {
	return model.db
}

func (model *RoleModel) FindOne(id int64) schema.RolePermission {
	var res schema.RolePermission

	dbRes := model.getDb().Model(&schema.RolePermission{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *RoleModel) InsertSchema(data *schema.RolePermission) error {
	data.CreatedAt = time.Now()
	dbRes := model.getDb().Model(&schema.RolePermission{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *RoleModel) UpdateByMap(id int64, data *schema.RolePermission) error {
	return model.getDb().Model(&schema.RolePermission{}).Where("id = ?", id).Updates(map[string]interface{}{}).Error
}

func (model *RoleModel) GetList(in *schema.RolePermission, pageQuery *model.PageQuery) (int64, []*schema.RolePermission, error) {
	builder := model.getDb().Model(&schema.RolePermission{})
	total := int64(0)
	err := builder.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	rows := make([]*schema.RolePermission, 0)
	err = builder.Offset(pageQuery.Offset()).Limit(pageQuery.PageSize).Find(&rows).Error
	return total, rows, err
}

func (model *RoleModel) GetAllRoleByCache() ([]*schema.RolePermission, error) {
	redisClient := model.svcCtx.RedisClient
	//cacheData, _ := model.svcCtx.RedisClient.HGetAll(model.ctx, "RolePermission").Result()
	//

	rows := make([]*schema.RolePermission, 0)
	err := model.getDb().Model(&schema.RolePermission{}).Find(&rows).Error
	if len(rows) > 0 {
		for _, v := range rows {
			redisClient.HSet(model.ctx, model.RolePermissionKey, v)
		}
	}
	return rows, err
}

func (model *RoleModel) GetAllRolePermissionByCache() (map[int64]*schema.RolePermission, error) {
	cache, err := model.GetAllRoleByCache()
	if err != nil {
		return nil, err
	}
	roleMap := make(map[int64]*schema.RolePermission)
	for _, v := range cache {
		roleMap[v.ID] = v
	}
	return roleMap, nil
}

func (model *RoleModel) GetRolePermissionById(roleId int64) (res []int64, err error) {
	res = make([]int64, 0)
	roleInfo := model.FindOne(roleId)
	if roleInfo.ID == 0 {
		return res, nil
	}
	json.Unmarshal([]byte(roleInfo.PermissionIds), &res)
	return
}
