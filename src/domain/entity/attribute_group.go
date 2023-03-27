package entity

import (
	"github.com/mindstand/gogm/v2"
)

type AttributeGroup struct {
	gogm.BaseUUIDNode
	Name       string       `gogm:"name=name"`
	Slug       string       `gogm:"name=slug"`
	Products   []*Product   `gogm:"direction=INCOMING;relationship=HAS_ATTRIBUTE_GROUP"`
	Attributes []*Attribute `gogm:"direction=OUTGOING;relationship=HAS_ATTRIBUTE"`
}
