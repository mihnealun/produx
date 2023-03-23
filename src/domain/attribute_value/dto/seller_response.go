package dto

import "produx/domain/entity"

type AttributeValueResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image§"`
	Address string `json:"address"`
}

type AttributeValueListResponse struct {
	AttributeValues []AttributeValueResponse `json:"attributeValues"`
}

func NewAttributeValueResponse(item *entity.AttributeValue) AttributeValueResponse {
	return AttributeValueResponse{
		ID:    item.UUID,
		Name:  item.Name,
		Image: item.Image,
	}
}

func NewAttributeValueListResponse(attributeValues []*entity.AttributeValue) AttributeValueListResponse {
	result := AttributeValueListResponse{
		AttributeValues: []AttributeValueResponse{},
	}

	for _, attributeValue := range attributeValues {
		result.AttributeValues = append(result.AttributeValues, NewAttributeValueResponse(attributeValue))
	}

	return result
}
