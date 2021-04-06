package issue

func PointerIssue(intrface Issue) *Issue {
	i := intrface
	return &i
}

func PointerIssues(issues []Issue) []*Issue {
	var pointerIssues []*Issue
	for _, issue := range issues {
		pointerIssues = append(pointerIssues, PointerIssue(issue))
	}
	return pointerIssues
}
