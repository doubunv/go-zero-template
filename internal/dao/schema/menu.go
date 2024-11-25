package schema

import (
	"gorm.io/gorm"
)

type Menu struct {
	ID       int64  `gorm:"column:id"`
	MenuPid  int64  `gorm:"column:menu_pid"`
	MenuName string `gorm:"column:menu_name"`
	MenuType int64  `gorm:"column:menu_type"`
	Path     string `gorm:"column:path"`
	Status   int64  `gorm:"column:status"`
	Sort     int64  `gorm:"column:sort"`
	gorm.Model
}

func (u Menu) TableName() string {
	return "t_menu"
}

const (
	MenuStatus1 = 1 //启用
	MenuStatus2 = 2 //禁用

	MenuMenuType1 = 1 //页面
	MenuMenuType2 = 2 //按钮
)
