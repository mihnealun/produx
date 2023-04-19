package response

import (
	"produx/domain/entity"
)

type CategoryResponse struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Slug          string             `json:"slug"`
	Subcategories []CategoryResponse `json:"categories"`
	ProductCount  int                `json:"product_count"`
}

func NewCategoryResponse(cat *entity.Category) CategoryResponse {
	subcategories := make([]CategoryResponse, 0)

	for _, c := range cat.Subcategories {
		subcategories = append(subcategories, NewCategoryResponse(c))
	}

	return CategoryResponse{
		ID:            cat.UUID,
		Name:          cat.Name,
		Slug:          cat.Slug,
		Subcategories: subcategories,
		ProductCount:  len(cat.Products),
	}
}
