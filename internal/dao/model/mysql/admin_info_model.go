package mysql

import (
	"context"
	"errors"
	"go-api/internal/dao/model"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"gorm.io/gorm"
	"time"
)

type AdminInfoModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewAdminInfoModel(ctx context.Context, svcCtx *svc.ServiceContext) *AdminInfoModel {
	return &AdminInfoModel{
		ctx:    ctx,
		svcCtx: svcCtx,
		db:     svcCtx.DbSelect.GetDb(ctx, model.DB),
	}
}

func (model *AdminInfoModel) getDb() *gorm.DB {
	return model.db
}

func (model *AdminInfoModel) FindOne(id int64) schema.AdminInfo {
	var res schema.AdminInfo
	if id <= 0 {
		return res
	}

	dbRes := model.getDb().Model(&schema.AdminInfo{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *AdminInfoModel) FindByAccount(account string) schema.AdminInfo {
	var res schema.AdminInfo

	dbRes := model.getDb().Model(&schema.AdminInfo{}).Where("account = ?", account).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *AdminInfoModel) InsertSchema(data *schema.AdminInfo) error {
	data.CreatedAt = time.Now().Unix()
	dbRes := model.getDb().Model(&schema.AdminInfo{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *AdminInfoModel) UpdateByMap(id int64, data *schema.AdminInfo) error {
	if id <= 0 {
		return errors.New("id error")
	}
	updateData := map[string]interface{}{
		"account": data.Account,
		"name":    data.Name,
		"status":  data.Status,
		"role_id": data.RoleId,
	}
	if data.Password != "" {
		updateData["password"] = data.Password
		updateData["password_sign"] = data.PasswordSign
	}
	return model.getDb().Model(&schema.AdminInfo{}).Where("id = ?", id).Updates(updateData).Error
}

func (model *AdminInfoModel) GetList(in *schema.AdminInfo, pageQuery *model.PageQuery) (int64, []*schema.AdminInfo, error) {
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

func (model *AdminInfoModel) DeleteById(id int64) error {
	if id <= 0 {
		return errors.New("id error")
	}
	return model.getDb().Model(&schema.AdminInfo{}).Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": time.Now()}).Error
}

func (model *AdminInfoModel) FindByIds(ids []int64) map[int64]*schema.AdminInfo {
	rows := make([]*schema.AdminInfo, 0)
	if len(ids) == 0 {
		return make(map[int64]*schema.AdminInfo)
	}
	model.getDb().Model(&schema.AdminInfo{}).Where("id in ?", ids).Find(&rows)
	res := make(map[int64]*schema.AdminInfo)
	for _, v := range rows {
		res[v.ID] = v
	}
	return res
}
