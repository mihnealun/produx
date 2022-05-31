package service

import (
	"produx/domain/entity"
)

type User interface {
	Add(user entity.User) *entity.User
	Update(user entity.User) *entity.User
	Delete(user entity.User) bool
	List() []*entity.User
}
