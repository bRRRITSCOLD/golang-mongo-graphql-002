package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-mongo-graphql-002/internal/api/generated"
	"golang-mongo-graphql-002/internal/issue"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) Issues(ctx context.Context) ([]*issue.Issue, error) {
	foundIssues, findIssuesErr := issue.FindIssues(bson.D{})
	if findIssuesErr != nil {
		return nil, findIssuesErr
	}
	return issue.PointerIssues(foundIssues), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
