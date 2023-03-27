package response

import (
	"produx/domain/entity"
)

type AttributeResponse struct {
	ID     string                   `json:"id"`
	Name   string                   `json:"name"`
	Slug   string                   `json:"slug"`
	Groups []*entity.AttributeGroup `json:"groups"`
}

func NewAttributeResponse(attribute *entity.Attribute) AttributeResponse {
	return AttributeResponse{
		ID:     attribute.UUID,
		Name:   attribute.Name,
		Slug:   attribute.Slug,
		Groups: attribute.Groups,
	}
}
