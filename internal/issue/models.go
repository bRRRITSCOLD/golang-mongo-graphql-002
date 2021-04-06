package issue

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Issue struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Code        string    `json:"code" bson:"code,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Completed   bool      `json:"completed" bson:"completed,omitempty"`
}

type NewIssue struct {
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Code        string    `json:"code" bson:"code,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Completed   bool      `json:"completed" bson:"completed,omitempty"`
}
