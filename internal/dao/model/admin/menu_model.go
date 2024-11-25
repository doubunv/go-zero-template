package admin

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"
	"go-api/internal/dao/dto"
	"go-api/internal/dao/model"
	"go-api/internal/dao/schema"
	"go-api/internal/svc"
	"gorm.io/gorm"
	"time"
)

type MenuModel struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewMenuModel(ctx context.Context, svcCtx *svc.ServiceContext) *MenuModel {
	return &MenuModel{
		ctx: ctx,
		db:  svcCtx.DbSelect.GetDb(ctx, model.DBAdmin),
	}
}

func (model *MenuModel) getDb() *gorm.DB {
	return model.db
}

func (model *MenuModel) FindOne(id int64) schema.Menu {
	var res schema.Menu

	dbRes := model.getDb().Model(&schema.Menu{}).Where("id = ?", id).First(&res)
	if err := dbRes.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	return res
}

func (model *MenuModel) InsertSchema(data *schema.Menu) error {
	data.CreatedAt = time.Now()
	dbRes := model.getDb().Model(&schema.Menu{}).Create(data)
	if err := dbRes.Error; err != nil {
		return err
	}

	return nil
}

func (model *MenuModel) UpdateByMap(id int64, data *schema.Menu) error {
	return model.getDb().Model(&schema.Menu{}).Where("id = ?", id).Updates(map[string]interface{}{}).Error
}

func (model *MenuModel) GetList(in *schema.AdminInfo, pageQuery *model.PageQuery) (int64, []*schema.Menu, error) {
	builder := model.getDb().Model(&schema.Menu{})
	total := int64(0)
	err := builder.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	rows := make([]*schema.Menu, 0)
	err = builder.Offset(pageQuery.Offset()).Limit(pageQuery.PageSize).Find(&rows).Error
	return total, rows, err
}

func (model *MenuModel) buildMenuTree(menus []*schema.Menu) map[int64]*dto.MenuTree {
	menuMap := make(map[int64]*dto.MenuTree)
	var tree []*schema.Menu
	if len(menus) == 0 {
		return menuMap
	}

	for i := range menus {
		menuTree := &dto.MenuTree{}
		copier.Copy(menuTree, menus[i])
		menuMap[menus[i].ID] = menuTree
	}

	for _, menu := range menus {
		if menu.MenuPid == 0 {
			tree = append(tree, menu)
		} else {
			if parent, exists := menuMap[menu.MenuPid]; exists {
				parent.Children = append(parent.Children, menu)
			}
		}
	}
	return menuMap
}

func (model *MenuModel) GetAllMenuByCache() ([]*schema.Menu, error) {
	rows := make([]*schema.Menu, 0)
	err := model.getDb().Model(&schema.Menu{}).Where("status = ? and menu_type = ?", schema.MenuStatus1, schema.MenuMenuType1).Find(&rows).Error
	return rows, err
}

func (model *MenuModel) GetMenuTreeByCache() (res map[int64]*dto.MenuTree, err error) {
	res = make(map[int64]*dto.MenuTree, 0)
	cache, err := model.GetAllMenuByCache()
	if err != nil {
		logc.Error(model.ctx, "GetMenuTreeByCache", err)
		return nil, errors.New("query menu error")
	}
	if len(cache) == 0 {
		return
	}
	return model.buildMenuTree(cache), nil
}

func (model *MenuModel) GetMenuTreeByIds(menuIds []int64) (res map[int64]*dto.MenuTree, err error) {
	rows := make([]*schema.Menu, 0)
	err = model.getDb().Model(&schema.Menu{}).Where("id in ?", menuIds).Find(rows).Error
	if err != nil {
		logc.Error(model.ctx, "GetMenuTreeByIdsCache", err)
		return nil, errors.New("query menu error")
	}

	return model.buildMenuTree(rows), nil
}
