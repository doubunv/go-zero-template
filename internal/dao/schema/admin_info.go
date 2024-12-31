package schema

type AdminInfo struct {
	ID           int64  `gorm:"column:id"`
	Account      string `gorm:"column:account"`
	Name         string `gorm:"column:name"`
	Password     string `gorm:"column:password"`
	PasswordSign string `gorm:"column:password_sign"`
	RoleId       int64  `gorm:"column:role_id"`
	Status       int    `gorm:"column:status"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`
	DeletedAt    int64  `gorm:"column:deleted_at"`
}

func (u AdminInfo) TableName() string {
	return "t_admin_info"
}

const (
	AdminInfoStatus1 = 1 //正常
	AdminInfoStatus2 = 2 //异常
)
