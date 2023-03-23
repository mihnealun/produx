package entity

import (
	"github.com/mindstand/gogm/v2"
)

type Seller struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Slug     string     `gogm:"name=slug"`
	Image    string     `gogm:"name=image"`
	Address  string     `gogm:"name=address"`
	Products []*Product `gogm:"direction=INCOMING;relationship=SOLD_BY"`
}
