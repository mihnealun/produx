package response

import (
	"produx/domain/entity"
)

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewCategoryResponse(cat *entity.Category) CategoryResponse {
	return CategoryResponse{
		ID:   cat.UUID,
		Name: cat.Name,
		Slug: cat.Slug,
	}
}
