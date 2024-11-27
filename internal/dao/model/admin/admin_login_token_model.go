package admin

import (
	"context"
	"errors"
	"go-api/internal/dao/model"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"gorm.io/gorm"
	"time"
)

type AdminLoginTokenModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewAdminLoginTokenModel(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginTokenModel {
	return &AdminLoginTokenModel{
		ctx:    ctx,
		svcCtx: svcCtx,
		db:     svcCtx.DbSelect.GetDb(ctx, model.DBAdmin),
	}
}

func (model *AdminLoginTokenModel) getDb() *gorm.DB {
	return model.db
}

func (model *AdminLoginTokenModel) AddLoginToken(data *schema.AdminLoginToken) error {
	var res schema.AdminLoginToken
	err := model.getDb().Model(&schema.AdminLoginToken{}).Where("admin_id = ?", data.AdminId).First(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if res.ID == 0 {
		err = model.InsertSchema(data)
		if err != nil {
			return err
		}
	} else {
		err = model.UpdateByIdMap(res.ID, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (model *AdminLoginTokenModel) InsertSchema(data *schema.AdminLoginToken) error {
	data.CreatedAt = time.Now()
	dbRes := model.getDb().Model(&schema.AdminLoginToken{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *AdminLoginTokenModel) UpdateByIdMap(id int64, data *schema.AdminLoginToken) error {
	return model.getDb().Model(&schema.AdminLoginToken{}).Where("id = ?", id).Updates(map[string]interface{}{
		"token_sign": data.TokenSign,
		"updated_at": time.Now(),
	}).Error
}

func (model *AdminLoginTokenModel) GetList(in *schema.AdminInfo, pageQuery *model.PageQuery) (int64, []*schema.AdminLoginToken, error) {
	builder := model.getDb().Model(&schema.AdminLoginToken{})
	total := int64(0)
	err := builder.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	rows := make([]*schema.AdminLoginToken, 0)
	err = builder.Offset(pageQuery.Offset()).Limit(pageQuery.PageSize).Find(&rows).Error
	return total, rows, err
}

func (model *AdminLoginTokenModel) FindOneByAdminId(adminId int64) schema.AdminLoginToken {
	var res schema.AdminLoginToken
	if adminId <= 0 {
		return res
	}
	dbRes := model.getDb().Model(&schema.AdminLoginToken{}).Where("admin_id = ?", adminId).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}
