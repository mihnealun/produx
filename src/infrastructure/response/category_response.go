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
	var subs []CategoryResponse

	for _, c := range cat.Subcategories {
		subs = append(subs, NewCategoryResponse(c))
	}

	return CategoryResponse{
		ID:            cat.UUID,
		Name:          cat.Name,
		Slug:          cat.Slug,
		Subcategories: subs,
		ProductCount:  len(cat.Products),
	}
}
