package entity

import (
	"github.com/mindstand/gogm/v2"
)

type User struct {
	gogm.BaseUUIDNode
	Name             string     `gogm:"name=name"`
	Type             string     `gogm:"name=type"`
	Status           string     `gogm:"name=status"`
	Comments         []*Comment `gogm:"direction=OUTGOING;relationship=CREATED"`
	DislikedComments []*Comment `gogm:"direction=OUTGOING;relationship=DISLIKED"`
	LikedComments    []*Comment `gogm:"direction=OUTGOING;relationship=LIKED"`
	ReportedComments []*Comment `gogm:"direction=OUTGOING;relationship=REPORTED"`
	DislikedTarget   []*Comment `gogm:"direction=OUTGOING;relationship=DISLIKED"`
	LikedTarget      []*Comment `gogm:"direction=OUTGOING;relationship=LIKED"`
	ReportedTarget   []*Comment `gogm:"direction=OUTGOING;relationship=REPORTED"`
}

type Target struct {
	gogm.BaseUUIDNode
	Name      string     `gogm:"name=name"`
	Type      string     `gogm:"name=type"`
	Url       string     `gogm:"name=url"`
	Comments  []*Comment `gogm:"direction=INCOMING;relationship=TARGETS"`
	Likers    []*User    `gogm:"direction=INCOMING;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=INCOMING;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=INCOMING;relationship=REPORTED"`
}

type App struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Slug     string     `gogm:"name=slug"`
	Comments []*Comment `gogm:"direction=INCOMING;relationship=POSTED_ON"`
}

type Comment struct {
	gogm.BaseUUIDNode
	Body      string     `gogm:"name=body"`
	Type      string     `gogm:"name=type"`
	Status    string     `gogm:"name=status"`
	CreatedAt int        `gogm:"name=created_at"`
	User      *User      `gogm:"direction=INCOMING;relationship=CREATED"`
	Target    *Target    `gogm:"direction=OUTGOING;relationship=TARGETS"`
	App       *App       `gogm:"direction=OUTGOING;relationship=POSTED_ON"`
	Replies   []*Comment `gogm:"direction=INCOMING;relationship=REPLIES_TO"`
	Likers    []*User    `gogm:"direction=INCOMING;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=INCOMING;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=INCOMING;relationship=REPORTED"`
	Parent    *Comment   `gogm:"direction=OUTGOING;relationship=REPLIES_TO"`
}
