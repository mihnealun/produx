package entity

import (
	"github.com/mindstand/gogm/v2"
)

type AttributeValue struct {
	gogm.BaseUUIDNode
	Name       string       `gogm:"name=name"`
	Image      string       `gogm:"name=image"`
	Type       string       `gogm:"name=type"`
	Value      any          `gogm:"name=value"`
	Attributes []*Attribute `gogm:"direction=INCOMING;relationship=HAS_VALUE"`
}
