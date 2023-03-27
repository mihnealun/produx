package response

import (
	"produx/domain/entity"
)

type AttributeListResponse struct {
	Attributes []AttributeResponse `json:"attributes"`
}

func NewAttributeListResponse(attributes []*entity.Attribute) AttributeListResponse {
	result := AttributeListResponse{
		Attributes: []AttributeResponse{},
	}

	for _, att := range attributes {
		result.Attributes = append(result.Attributes, NewAttributeResponse(att))
	}

	return result
}
