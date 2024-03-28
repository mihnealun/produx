package entity

import "github.com/mindstand/gogm/v2"

type Comment struct {
	gogm.BaseUUIDNode
	Body      string     `gogm:"name=body"`
	Type      string     `gogm:"name=type"`
	Status    string     `gogm:"name=status"`
	CreatedAt int        `gogm:"name=created_at"`
	User      *User      `gogm:"direction=incoming;relationship=CREATED"`
	Target    *Target    `gogm:"direction=outgoing;relationship=TARGETS"`
	App       *App       `gogm:"direction=outgoing;relationship=POSTED_ON"`
	Replies   []*Comment `gogm:"direction=incoming;relationship=REPLIES_TO"`
	Likers    []*User    `gogm:"direction=incoming;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=incoming;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=incoming;relationship=REPORTED"`
	Parent    *Comment   `gogm:"direction=outgoing;relationship=REPLIES_TO"`
}
