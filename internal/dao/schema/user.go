package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   int64  `gorm:"primarykey"`
	Name string `json:"name"`
}

func (u User) TableName() string {
	return "userx"
}
