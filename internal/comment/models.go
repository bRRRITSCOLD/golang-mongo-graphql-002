package comment

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Comment struct {
	ID        string    `json:"_id" bson:"_id,omitempty"`
	CommentID string    `json:"commentId" bson:"commentId,omitempty"`
	IssueID   string    `json:"issueId" bson:"issueId,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	Body      string    `json:"body" bson:"body,omitempty"`
}
