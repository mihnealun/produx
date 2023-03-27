package service

import (
	"produx/domain/entity"
)

type Attribute interface {
	Add(app entity.Attribute) *entity.Attribute
	Update(app entity.Attribute) *entity.Attribute
	Delete(app entity.Attribute) bool
	List() []*entity.Attribute
	Get(id string) *entity.Attribute
	GetByName(name string) *entity.Attribute
}
