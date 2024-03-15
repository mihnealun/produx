package response

import "produx/domain/entity"

type CommentListResponse struct {
	Comments []CommentResponse `json:"comments"`
}

func NewCommentListResponse(comments []*entity.Comment) CommentListResponse {
	result := CommentListResponse{
		Comments: []CommentResponse{},
	}

	for _, c := range comments {
		result.Comments = append(result.Comments, NewCommentResponse(c))
	}

	return result
}
