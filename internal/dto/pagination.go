package dto

import "wallet-api/internal/model"

type Pagination struct {
	Limit int    `json:"limit" valid:"-"`
	Page  int    `json:"page" valid:"-"`
	Sort  string `json:"sort" valid:"-"`
} // @name Pagination

func (dto *Pagination) ConvertToVO() *model.Pagination {
	result := &model.Pagination{}

	result.Limit = dto.Limit
	result.Page = dto.Page
	result.Sort = dto.Sort

	return result
}
