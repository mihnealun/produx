package response

import "produx/domain/entity"

type CommentResponse struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	UserName string `json:"username"`
	Status   string `json:"status"`
	Target   struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"target"`
	User struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"user"`
	AppName string `json:"app"`
}

func NewCommentResponse(comment *entity.Comment) CommentResponse {
	return CommentResponse{
		Body:     comment.Body,
		UserName: comment.User.Name,
		Status:   comment.Status,
		Target: struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		}{
			ID:   comment.Target.UUID,
			Type: comment.Target.Type,
		},
		User: struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Name   string `json:"name"`
			Status string `json:"status"`
		}{
			ID:     comment.User.UUID,
			Type:   comment.User.Type,
			Name:   comment.User.Name,
			Status: comment.User.Status,
		},
		AppName: comment.App.Name,
	}
}
