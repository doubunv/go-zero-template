package dto

type MenuTree struct {
	ID       int64  `gorm:"column:id"`
	MenuPid  int64  `gorm:"column:menu_pid"`
	MenuName string `gorm:"column:menu_name"`
	MenuType int    `gorm:"column:menu_type"`
	Path     string `gorm:"column:path"`
	Children []*MenuTree
}
