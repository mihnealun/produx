package dto

import "produx/domain/entity"

type AttributeResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Image   string `json:"image§"`
	Address string `json:"address"`
}

type AttributeListResponse struct {
	Attributes []AttributeResponse `json:"attributes"`
}

func NewAttributeResponse(item *entity.Attribute) AttributeResponse {
	return AttributeResponse{
		ID:    item.UUID,
		Name:  item.Name,
		Slug:  item.Slug,
		Image: item.Image,
	}
}

func NewAttributeListResponse(attributes []*entity.Attribute) AttributeListResponse {
	result := AttributeListResponse{
		Attributes: []AttributeResponse{},
	}

	for _, attribute := range attributes {
		result.Attributes = append(result.Attributes, NewAttributeResponse(attribute))
	}

	return result
}
