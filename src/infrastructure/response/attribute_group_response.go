package response

import (
	"produx/domain/entity"
)

type AttributeGroupResponse struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	Slug       string              `json:"slug"`
	Attributes []*entity.Attribute `json:"groups"`
}

func NewAttributeGroupResponse(attribute *entity.AttributeGroup) AttributeGroupResponse {
	return AttributeGroupResponse{
		ID:         attribute.UUID,
		Name:       attribute.Name,
		Slug:       attribute.Slug,
		Attributes: attribute.Attributes,
	}
}

type AttributeGroupListResponse struct {
	Groups []AttributeGroupResponse `json:"groups"`
}

func NewAttributeGroupListResponse(attributes []*entity.AttributeGroup) AttributeGroupListResponse {
	result := AttributeGroupListResponse{
		Groups: []AttributeGroupResponse{},
	}

	for _, att := range attributes {
		result.Groups = append(result.Groups, NewAttributeGroupResponse(att))
	}

	return result
}
