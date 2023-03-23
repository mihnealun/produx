package service

import (
	"produx/domain/entity"
)

type AttributeValue interface {
	Add(app entity.AttributeValue) *entity.AttributeValue
	Update(app entity.AttributeValue) *entity.AttributeValue
	Delete(app entity.AttributeValue) bool
	List() []*entity.AttributeValue
	Get(id string) *entity.AttributeValue
}
