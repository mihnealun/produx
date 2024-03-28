package entity

import "github.com/mindstand/gogm/v2"

type App struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Slug     string     `gogm:"name=slug"`
	Comments []*Comment `gogm:"direction=incoming;relationship=POSTED_ON"`
}
