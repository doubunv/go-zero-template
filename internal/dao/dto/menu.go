package dto

import "go-api/internal/dao/schema"

type MenuTree struct {
	schema.Menu
	Children []*schema.Menu
}
