package model

import "github.com/spf13/cast"

const (
	DB      = "chgame"
	DBAdmin = "chgame_admin"
)

type PageQuery struct {
	Page     int
	PageSize int
}

func NewPageQuery(page, pageSize int) *PageQuery {
	return &PageQuery{
		Page:     page,
		PageSize: pageSize,
	}
}

func (s *PageQuery) Offset() int {
	return cast.ToInt((s.Page - 1) * s.PageSize)
}
