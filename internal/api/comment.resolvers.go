package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-mongo-graphql-002/internal/comment"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input comment.NewCommentInput) (*comment.Comment, error) {
	createdComments, createCommentssErr := comment.CreateComments([]comment.Comment{
		{
			IssueID: input.IssueID,
			Body:    input.Body,
		},
	})
	if createCommentssErr != nil {
		return nil, createCommentssErr
	}
	return comment.PointerComment(createdComments[0]), nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, commentID string) (bool, error) {
	_, deleteCommentsErr := comment.DeleteComments(comment.Comment{
		CommentID: commentID,
	})
	if deleteCommentsErr != nil {
		return false, deleteCommentsErr
	}
	return true, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, commentID string, input comment.UpdateCommentInput) (bool, error) {
	_, updateCommentsErr := comment.UpdateComments(
		comment.Comment{
			CommentID: commentID,
		},
		comment.Comment{
			IssueID: input.IssueID,
			Body:    input.Body,
		},
	)
	if updateCommentsErr != nil {
		return false, updateCommentsErr
	}
	return true, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*comment.Comment, error) {
	foundComments, findCommentsErr := comment.FindComments(bson.D{})
	if findCommentsErr != nil {
		return nil, findCommentsErr
	}
	return comment.PointerComments(foundComments), nil
}

func (r *queryResolver) Comment(ctx context.Context, commentID string) (*comment.Comment, error) {
	foundComments, findCommentsErr := comment.FindComments(comment.Comment{
		CommentID: commentID,
	})
	if findCommentsErr != nil {
		return nil, findCommentsErr
	}
	return comment.PointerComment(foundComments[0]), nil
}
