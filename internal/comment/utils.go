package comment

func PointerComment(cmt Comment) *Comment {
	i := cmt
	return &i
}

func PointerComments(comments []Comment) []*Comment {
	var pointerComments []*Comment
	for _, comment := range comments {
		pointerComments = append(pointerComments, PointerComment(comment))
	}
	return pointerComments
}

func MapToComment(comment Comment) Comment {
	return Comment{
		ID:        comment.ID,
		CommentID: comment.CommentID,
		IssueID:   comment.IssueID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Body:      comment.Body,
	}
}
