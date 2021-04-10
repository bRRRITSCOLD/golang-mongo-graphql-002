package issue

func PointerIssue(iss Issue) *Issue {
	i := iss
	return &i
}

func PointerIssues(issues []Issue) []*Issue {
	var pointerIssues []*Issue
	for _, issue := range issues {
		pointerIssues = append(pointerIssues, PointerIssue(issue))
	}
	return pointerIssues
}

func MapToIssue(issue Issue) Issue {
	return Issue{
		ID:          issue.ID,
		IssueID:     issue.IssueID,
		CreatedAt:   issue.CreatedAt,
		UpdatedAt:   issue.UpdatedAt,
		Title:       issue.Title,
		Code:        issue.Code,
		Description: issue.Description,
		Completed:   issue.Completed,
	}
}
