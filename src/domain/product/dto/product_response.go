package dto

import (
	"produx/domain/entity"
	"produx/domain/seller/dto"
)

type ProductResponse struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Slug       string                 `json:"slug"`
	Sellers    dto.SellerListResponse `json:"sellers"`
	Attributes dto.SellerListResponse `json:"attributes"`
	Category   string                 `json:"category"`
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
		Category: item.Category.Name,
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
