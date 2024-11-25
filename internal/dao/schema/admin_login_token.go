package schema

import (
	"gorm.io/gorm"
)

type AdminLoginToken struct {
	ID        int64  `gorm:"column:id"`
	AdminId   int64  `gorm:"column:admin_id"`
	TokenSign string `gorm:"column:token_sign"`
	gorm.Model
}

func (u AdminLoginToken) TableName() string {
	return "t_admin_login_token"
}
