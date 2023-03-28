package service

import (
	"produx/domain/entity"
)

type AttributeGroup interface {
	Add(app entity.AttributeGroup) *entity.AttributeGroup
	Update(app entity.AttributeGroup) *entity.AttributeGroup
	Delete(app entity.AttributeGroup) bool
	List() []*entity.AttributeGroup
	Get(id string) *entity.AttributeGroup
	GetByName(name string) *entity.AttributeGroup
}
