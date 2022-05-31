package response

import (
	"produx/domain/entity"
)

type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

func NewCategoryListResponse(categories []*entity.Category) CategoryListResponse {
	result := CategoryListResponse{
		Categories: []CategoryResponse{},
	}

	for _, cat := range categories {
		result.Categories = append(result.Categories, NewCategoryResponse(cat))
	}

	return result
}
