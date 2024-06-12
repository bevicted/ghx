package ghxtest

import (
	"testing"

	"github.com/google/go-github/v62/github"
)

func NewEmptyIssues(t *testing.T, n int) []*github.Issue {
	t.Helper()
	issues := make([]*github.Issue, n)
	for i := range n {
		i := i
		issues[i] = &github.Issue{Number: &i}
	}
	return issues
}
