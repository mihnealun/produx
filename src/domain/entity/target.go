package entity

import (
	"github.com/mindstand/gogm/v2"
)

type Target struct {
	gogm.BaseUUIDNode
	Name      string     `gogm:"name=name"`
	Type      string     `gogm:"name=type"`
	Url       string     `gogm:"name=url"`
	Comments  []*Comment `gogm:"direction=incoming;relationship=TARGETS"`
	Likers    []*User    `gogm:"direction=incoming;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=incoming;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=incoming;relationship=REPORTED"`
}
