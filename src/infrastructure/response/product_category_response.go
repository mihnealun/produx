package response

import "produx/domain/entity"

type ProductCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProductCategoryResponse(item *entity.Category) *ProductCategoryResponse {
	return &ProductCategoryResponse{
		ID:   item.UUID,
		Name: item.Name,
	}
}
