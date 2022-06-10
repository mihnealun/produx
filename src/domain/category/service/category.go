package service

import (
	"produx/domain/entity"
)

type Category interface {
	Add(app entity.Category) *entity.Category
	Update(app entity.Category) *entity.Category
	Delete(app *entity.Category) bool
	List() []*entity.Category
	Get(id string) *entity.Category
	GetByName(name string) *entity.Category
}
