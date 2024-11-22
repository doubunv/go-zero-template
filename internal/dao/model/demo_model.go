package model

import (
	"context"
	"errors"
	"go-api/internal/dao/schema"
	"go-api/internal/pkg/dbM"
	"gorm.io/gorm"
)

type DemoModel struct {
	ctx context.Context
	db  *gorm.DB
}

func NewDemoModel(ctx context.Context, db *dbM.SelectDb) *DemoModel {
	return &DemoModel{
		ctx: ctx,
		db:  db.GetDb(ctx, DBAdmin),
	}
}

func (model *DemoModel) getDb() *gorm.DB {
	return model.db
}

func (model *DemoModel) FindOne(id int64) schema.User {
	var res schema.User

	dbRes := model.getDb().Model(&schema.User{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *DemoModel) InsertSchema(data *schema.User) error {
	dbRes := model.getDb().Model(&schema.User{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}
