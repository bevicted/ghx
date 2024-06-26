package ghxtest

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/require"
)

type IssueFactory struct {
	github.Issue
}

// NewIssue deepcopies IssueFactory.Issue by marshaling and unmarshaling it
// NOTE: performance heavy, only use for detailed testing
func (f *IssueFactory) NewIssue(t *testing.T) *github.Issue {
	t.Helper()
	var i github.Issue
	remarshal(t, f.Issue, &i)
	return &i
}

func (f *IssueFactory) NewIssueWithOverwrite(t *testing.T, overwrite github.Issue) *github.Issue {
	t.Helper()
	var i github.Issue

	remarshal(t, f.Issue, &i)
	remarshal(t, overwrite, &i)

	return &i
}

func (f *IssueFactory) NewIssueWithNumber(t *testing.T, number int) *github.Issue {
	t.Helper()
	i := f.NewIssue(t)
	i.Number = &number
	return i
}

func (f *IssueFactory) NewIssues(t *testing.T, n int) []*github.Issue {
	t.Helper()
	var baseNumber int
	if f.Number != nil {
		baseNumber = *f.Number
	}
	issues := make([]*github.Issue, n)
	for i := range n {
		issues[i] = f.NewIssueWithNumber(t, baseNumber+i)
	}
	return issues
}

func NewEmptyIssues(t *testing.T, n int) []*github.Issue {
	t.Helper()
	issues := make([]*github.Issue, n)
	for i := range n {
		i := i
		issues[i] = &github.Issue{Number: &i}
	}
	return issues
}

func remarshal[T any](t *testing.T, marshalTarget T, unmarshalTarget *T) {
	t.Helper()

	b, err := json.Marshal(marshalTarget)
	require.NoError(t, err, "error marshalling in remarshal")
	require.NoError(t, json.Unmarshal(b, unmarshalTarget), "error unmarshalling in remarshal")
}
