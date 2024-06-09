package ghxtest

import (
	"context"
	"testing"

	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/require"
)

type MockRepoIssueLister struct {
	cursor       int
	T            *testing.T
	TestFunc     func(t *testing.T, ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions)
	ReturnValues []*MockRepoIssueListerReturnValues
}

type MockRepoIssueListerReturnValues struct {
	Issues   []*github.Issue
	Response *github.Response
	Err      error
}

func NewMockRepoIssueListerReturnValues(lastPageErr error, issuesOnPages ...int) []*MockRepoIssueListerReturnValues {
	if len(issuesOnPages) == 0 && lastPageErr != nil {
		issuesOnPages = append(issuesOnPages, 0)
	}
	rv := make([]*MockRepoIssueListerReturnValues, len(issuesOnPages))
	for idx, issueOnPage := range issuesOnPages {
		var (
			issues   []*github.Issue
			res      *github.Response
			err      error
			nextPage = idx + 2
		)
		if len(issuesOnPages) <= idx+1 {
			err = lastPageErr
			nextPage = 0
		}
		if err == nil {
			issues = NewEmptyIssues(issueOnPage)
			res = &github.Response{NextPage: nextPage}
		}
		rv[idx] = &MockRepoIssueListerReturnValues{
			Issues:   issues,
			Response: res,
			Err:      err,
		}
	}
	return rv
}

func (m *MockRepoIssueLister) ListByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
	m.T.Helper()
	require.Less(m.T, m.cursor, len(m.ReturnValues), "MockRepoIssueLister.ListByRepo called more times than it has ReturnValues mocked")
	if m.TestFunc != nil {
		m.TestFunc(m.T, ctx, owner, repo, opts)
	}
	rv := m.ReturnValues[m.cursor]
	m.cursor++
	return rv.Issues, rv.Response, rv.Err
}

type MockIssueSearcher struct {
	cursor       int
	T            *testing.T
	TestFunc     func(t *testing.T, ctx context.Context, query string, opts *github.SearchOptions)
	ReturnValues []*MockIssueSearcherReturnValues
}

type MockIssueSearcherReturnValues struct {
	IssueSearchResult *github.IssuesSearchResult
	Response          *github.Response
	Err               error
}

func NewMockIssueSearcherReturnValues(lastPageErr error, issuesOnPages ...int) []*MockIssueSearcherReturnValues {
	issueListerRV := NewMockRepoIssueListerReturnValues(lastPageErr, issuesOnPages...)
	rv := make([]*MockIssueSearcherReturnValues, len(issueListerRV))
	for idx, baseRV := range issueListerRV {
		var issuesSearchResult *github.IssuesSearchResult
		if baseRV.Issues != nil {
			issuesSearchResult = &github.IssuesSearchResult{Issues: baseRV.Issues}
		}
		rv[idx] = &MockIssueSearcherReturnValues{
			IssueSearchResult: issuesSearchResult,
			Response:          baseRV.Response,
			Err:               baseRV.Err,
		}
	}
	return rv
}

func (m *MockIssueSearcher) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	m.T.Helper()
	require.Less(m.T, m.cursor, len(m.ReturnValues), "MockIssueSearcher.Issues called more times than it has ReturnValues mocked")
	if m.TestFunc != nil {
		m.TestFunc(m.T, ctx, query, opts)
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
