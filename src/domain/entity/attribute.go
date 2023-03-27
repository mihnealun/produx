package entity

import (
	"github.com/mindstand/gogm/v2"
)

type Attribute struct {
	gogm.BaseUUIDNode
	Name         string            `gogm:"name=name"`
	Slug         string            `gogm:"name=slug"`
	Image        string            `gogm:"name=image"`
	DefaultValue string            `gogm:"name=default_value"`
	Values       []*AttributeValue `gogm:"direction=OUTGOING;relationship=HAS_VALUE"`
	Groups       []*AttributeGroup `gogm:"direction=INCOMING;relationship=HAS_ATTRIBUTE"`
}
