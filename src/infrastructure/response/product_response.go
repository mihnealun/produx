package response

import (
	"produx/domain/entity"
)

type ProductResponse struct {
	ID       string                   `json:"id"`
	Name     string                   `json:"name"`
	Slug     string                   `json:"slug"`
	Seller   SellerResponse           `json:"seller"`
	Category *ProductCategoryResponse `json:"category"`
}

type ProductListResponse struct {
	Count    int               `json:"count"`
	Products []ProductResponse `json:"products"`
}

func NewProductResponse(item *entity.Product) ProductResponse {
	return ProductResponse{
		ID:       item.UUID,
		Name:     item.Name,
		Slug:     item.Slug,
		Category: NewProductCategoryResponse(item.Category),
		Seller:   NewSellerResponse(item.Seller),
	}
}

func NewProductListResponse(products []*entity.Product) ProductListResponse {
	result := ProductListResponse{
		Products: []ProductResponse{},
	}

	for _, prod := range products {
		result.Products = append(result.Products, NewProductResponse(prod))
	}

	result.Count = len(result.Products)

	return result
}
