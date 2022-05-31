package dto

import "produx/domain/entity"

type ProductResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}

func NewProductResponse(item *entity.Product) ProductResponse {
	return ProductResponse{
		ID:   item.UUID,
		Name: item.Name,
		Slug: item.Slug,
	}
}

func NewProductListResponse(products []*entity.Product) ProductListResponse {
	result := ProductListResponse{
		Products: []ProductResponse{},
	}

	for _, prod := range products {
		result.Products = append(result.Products, NewProductResponse(prod))
	}

	return result
}
