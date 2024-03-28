package entity

import "github.com/mindstand/gogm/v2"

type User struct {
	gogm.BaseUUIDNode
	Name             string     `gogm:"name=name"`
	Type             string     `gogm:"name=type"`
	Status           string     `gogm:"name=status"`
	Comments         []*Comment `gogm:"direction=outgoing;relationship=CREATED"`
	DislikedComments []*Comment `gogm:"direction=outgoing;relationship=DISLIKED"`
	LikedComments    []*Comment `gogm:"direction=outgoing;relationship=LIKED"`
	ReportedComments []*Comment `gogm:"direction=outgoing;relationship=REPORTED"`
	DislikedTarget   []*Comment `gogm:"direction=outgoing;relationship=DISLIKED"`
	LikedTarget      []*Comment `gogm:"direction=outgoing;relationship=LIKED"`
	ReportedTarget   []*Comment `gogm:"direction=outgoing;relationship=REPORTED"`
}
