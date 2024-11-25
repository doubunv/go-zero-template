package schema

import (
	"gorm.io/gorm"
)

type PermissionMenu struct {
	ID             int64  `gorm:"column:id"`
	PermissionName int64  `gorm:"column:permission_name"`
	MenuIds        string `gorm:"column:menu_ids"`
	gorm.Model
}

func (u Menu) PermissionMenu() string {
	return "t_permission_menu"
}
