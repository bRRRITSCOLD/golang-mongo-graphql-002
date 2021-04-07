package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-mongo-graphql-002/internal/api/generated"
	"golang-mongo-graphql-002/internal/issue"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateIssue(ctx context.Context, input issue.NewIssue) (*issue.Issue, error) {
	createdIssues, createIssuesErr := issue.CreateIssues([]issue.Issue{
		{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       input.Title,
			Code:        input.Code,
			Description: input.Description,
			Completed:   input.Completed,
		},
	})
	if createIssuesErr != nil {
		return nil, createIssuesErr
	}
	return issue.PointerIssue(createdIssues[0]), nil
}

func (r *mutationResolver) DeleteIssue(ctx context.Context, issueID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Issues(ctx context.Context) ([]*issue.Issue, error) {
	foundIssues, findIssuesErr := issue.FindIssues(bson.D{})
	if findIssuesErr != nil {
		return nil, findIssuesErr
	}
	return issue.PointerIssues(foundIssues), nil
}

func (r *queryResolver) Issue(ctx context.Context, issueID string) (*issue.Issue, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *issueResolver) IssueID(ctx context.Context, obj *issue.Issue) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

type issueResolver struct{ *Resolver }

func (r *queryResolver) DeleteIssue(ctx context.Context, issueID string) (*issue.Issue, error) {
	panic(fmt.Errorf("not implemented"))
}
