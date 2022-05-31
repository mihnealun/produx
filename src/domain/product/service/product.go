package service

import (
	"produx/domain/entity"
)

type Product interface {
	Add(app entity.Product) *entity.Product
	Update(app entity.Product) *entity.Product
	Delete(app entity.Product) bool
	List() []*entity.Product
	Get(id string) *entity.Product
}
