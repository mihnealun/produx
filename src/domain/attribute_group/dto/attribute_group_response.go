package dto

import "produx/domain/entity"

type AttributeGroupResponse struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	Slug       string              `json:"slug"`
	Image      string              `json:"image§"`
	Attributes []*entity.Attribute `json:"attributes"`
}

type AttributeGroupListResponse struct {
	Groups []AttributeGroupResponse `json:"groups"`
}

func NewAttributeGroupResponse(item *entity.Attribute) AttributeGroupResponse {
	return AttributeGroupResponse{
		ID:    item.UUID,
		Name:  item.Name,
		Slug:  item.Slug,
		Image: item.Image,
	}
}

func NewAttributeGroupListResponse(attributes []*entity.Attribute) AttributeGroupListResponse {
	result := AttributeGroupListResponse{
		Groups: []AttributeGroupResponse{},
	}

	for _, attribute := range attributes {
		result.Groups = append(result.Groups, NewAttributeGroupResponse(attribute))
	}

	return result
}
