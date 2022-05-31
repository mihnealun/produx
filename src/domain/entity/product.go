package entity

import (
	"github.com/mindstand/gogm/v2"
)

type Product struct {
	gogm.BaseUUIDNode
	Name     string    `gogm:"name=name"`
	Slug     string    `gogm:"name=slug"`
	Price    float64   `gogm:"name=price"`
	Image    string    `gogm:"name=image"`
	Category *Category `gogm:"direction=outgoing;relationship=IN_CATEGORY"`
	Seller   *Seller   `gogm:"direction=outgoing;relationship=SOLD_BY"`
}
