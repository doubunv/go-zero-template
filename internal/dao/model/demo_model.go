package model

import (
	"context"
	"errors"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"gorm.io/gorm"
	"time"
)

type DemoInfoModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewDemoInfoModel(ctx context.Context, svcCtx *svc.ServiceContext) *DemoInfoModel {
	return &DemoInfoModel{
		ctx: ctx,
		db:  svcCtx.DbSelect.GetDb(ctx, DBAdmin),
	}
}

func (model *DemoInfoModel) getDb() *gorm.DB {
	return model.db
}

func (model *DemoInfoModel) FindOne(id int64) schema.AdminInfo {
	var res schema.AdminInfo

	dbRes := model.getDb().Model(&schema.AdminInfo{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *DemoInfoModel) InsertSchema(data *schema.AdminInfo) error {
	data.CreatedAt = time.Now()
	dbRes := model.getDb().Model(&schema.AdminInfo{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *DemoInfoModel) UpdateByMap(id int64, data *schema.AdminInfo) error {
	return model.getDb().Model(&schema.AdminInfo{}).Where("id = ?", id).Updates(map[string]interface{}{}).Error
}

func (model *DemoInfoModel) GetList(in *schema.AdminInfo, pageQuery *PageQuery) (int64, []*schema.AdminInfo, error) {
	builder := model.getDb().Model(&schema.AdminInfo{})
	total := int64(0)
	err := builder.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	rows := make([]*schema.AdminInfo, 0)
	err = builder.Offset(pageQuery.Offset()).Limit(pageQuery.PageSize).Find(&rows).Error
	return total, rows, err
}
