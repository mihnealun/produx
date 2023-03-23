package entity

import "github.com/mindstand/gogm/v2"

type Product struct {
	gogm.BaseUUIDNode
	Name       string       `gogm:"name=name"`
	Slug       string       `gogm:"name=slug"`
	Price      float64      `gogm:"name=price"`
	Image      string       `gogm:"name=image"`
	Category   *Category    `gogm:"direction=OUTGOING;relationship=IN_CATEGORY"`
	Seller     *Seller      `gogm:"direction=OUTGOING;relationship=SOLD_BY"`
	Attributes []*Attribute `gogm:"direction=OUTGOING;relationship=HAS_ATTRIBUTE"`
}
