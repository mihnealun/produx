package response

import "produx/domain/entity"

type SellerResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Image   string `json:"image§"`
	Address string `json:"address"`
}

type SellerListResponse struct {
	Sellers []SellerResponse `json:"sellers"`
}

func NewSellerResponse(item *entity.Seller) SellerResponse {
	return SellerResponse{
		ID:      item.UUID,
		Name:    item.Name,
		Slug:    item.Slug,
		Image:   item.Image,
		Address: item.Address,
	}
}

func NewSellerListResponse(sellers []*entity.Seller) SellerListResponse {
	result := SellerListResponse{
		Sellers: []SellerResponse{},
	}

	for _, seller := range sellers {
		result.Sellers = append(result.Sellers, NewSellerResponse(seller))
	}

	return result
}
