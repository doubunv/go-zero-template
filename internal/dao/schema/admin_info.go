package schema

import "gorm.io/gorm"

type AdminInfo struct {
	ID           int64  `gorm:"column:id"`
	Account      string `gorm:"column:account"`
	Name         string `gorm:"column:name"`
	Password     string `gorm:"column:password"`
	PasswordSign string `gorm:"column:password_sign"`
	RoleId       int64  `gorm:"column:role_id"`
	Status       int64  `gorm:"column:status"`
	gorm.Model
}

func (u AdminInfo) TableName() string {
	return "t_admin_info"
}

const (
	AdminInfoStatus1 = 1 //正常
	AdminInfoStatus2 = 2 //异常
)
