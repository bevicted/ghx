package ghxtest

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type MockRepoIssueLister struct {
	cursor       int
	ReturnValues []*MockRepoIssueListerReturnValues
}

type MockRepoIssueListerReturnValues struct {
	Issues   []*github.Issue
	Response *github.Response
	Err      error
}

func (m *MockRepoIssueLister) ListByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
	if m.cursor > len(m.ReturnValues) {
		panic("MockRepoIssueLister.ListByRepo called more times than it has ReturnValues mocked")
	}
	rv := m.ReturnValues[m.cursor]
	m.cursor++
	return rv.Issues, rv.Response, rv.Err
}

type MockIssueSearcher struct {
	cursor       int
	ReturnValues []*MockIssueSearcherReturnValues
}

type MockIssueSearcherReturnValues struct {
	IssueSearchResult *github.IssuesSearchResult
	Response          *github.Response
	Err               error
}

func (m *MockIssueSearcher) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	if m.cursor > len(m.ReturnValues) {
		panic("MockIssueSearcher.Issues called more times than it has ReturnValues mocked")
	}
	rv := m.ReturnValues[m.cursor]
	m.cursor++
	return rv.IssueSearchResult, rv.Response, rv.Err
}

func NewEmptyIssues(num int) []*github.Issue {
	issues := make([]*github.Issue, num)
	for i := 1; i <= num; i++ {
		issues[i-1] = &github.Issue{Number: &i}
	}
	return issues
}
