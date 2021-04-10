package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-mongo-graphql-002/internal/issue"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateIssue(ctx context.Context, input issue.NewIssueInput) (*issue.Issue, error) {
	createdIssues, createIssuesErr := issue.CreateIssues([]issue.Issue{
		{
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
	_, deleteIssuesErr := issue.DeleteIssues(issue.Issue{
		IssueID: issueID,
	})
	if deleteIssuesErr != nil {
		return false, deleteIssuesErr
	}
	return true, nil
}

func (r *mutationResolver) UpdateIssue(ctx context.Context, issueID string, input issue.UpdateIssueInput) (bool, error) {
	_, updateIssuesErr := issue.UpdateIssues(
		issue.Issue{
			IssueID: issueID,
		},
		issue.Issue{
			Title:       input.Title,
			Code:        input.Code,
			Description: input.Description,
			Completed:   input.Completed,
		},
	)
	if updateIssuesErr != nil {
		return false, updateIssuesErr
	}
	return true, nil
}

func (r *queryResolver) Issues(ctx context.Context) ([]*issue.Issue, error) {
	foundIssues, findIssuesErr := issue.FindIssues(bson.D{})
	if findIssuesErr != nil {
		return nil, findIssuesErr
	}
	return issue.PointerIssues(foundIssues), nil
}

func (r *queryResolver) Issue(ctx context.Context, issueID string) (*issue.Issue, error) {
	foundIssues, findIssuesErr := issue.FindIssues(issue.Issue{
		IssueID: issueID,
	})
	if findIssuesErr != nil {
		return nil, findIssuesErr
	}
	return issue.PointerIssue(foundIssues[0]), nil
}
