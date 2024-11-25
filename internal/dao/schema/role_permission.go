package schema

import (
	"gorm.io/gorm"
)

type RolePermission struct {
	ID            int64  `gorm:"column:id"`
	RoleName      string `gorm:"column:role_name"`
	PermissionIds string `gorm:"column:permission_ids"`
	gorm.Model
}

func (u Menu) RolePermission() string {
	return "t_role_permission"
}
