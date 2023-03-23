package entity

import (
	"github.com/mindstand/gogm/v2"
)

type Category struct {
	gogm.BaseUUIDNode
	Name          string      `gogm:"name=name"`
	Slug          string      `gogm:"name=slug"`
	Parent        *Category   `gogm:"direction=OUTGOING;relationship=CHILD_OF"`
	Subcategories []*Category `gogm:"direction=INCOMING;relationship=CHILD_OF"`
	Products      []*Product  `gogm:"direction=INCOMING;relationship=IN_CATEGORY"`
}
