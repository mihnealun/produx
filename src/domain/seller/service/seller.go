package service

import (
	"produx/domain/entity"
)

type Seller interface {
	Add(app entity.Seller) *entity.Seller
	Update(app entity.Seller) *entity.Seller
	Delete(app entity.Seller) bool
	List() []*entity.Seller
	Get(id string) *entity.Seller
}
