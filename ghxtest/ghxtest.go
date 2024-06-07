package ghxtest

import (
	"context"
	"testing"

	"github.com/google/go-github/v62/github"
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
	var rv []*MockRepoIssueListerReturnValues
	if len(issuesOnPages) == 0 && lastPageErr != nil {
		issuesOnPages = append(issuesOnPages, 0)
	}
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
		rv = append(rv, &MockRepoIssueListerReturnValues{
			Issues:   issues,
			Response: res,
			Err:      err,
		})
	}
	return rv
}

func (m *MockRepoIssueLister) ListByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
	if len(m.ReturnValues) == 0 {
		return []*github.Issue{}, &github.Response{}, nil
	}
	if len(m.ReturnValues) <= m.cursor {
		panic("MockRepoIssueLister.ListByRepo called more times than it has ReturnValues mocked")
	}
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
	var rv []*MockIssueSearcherReturnValues
	for _, baseRV := range NewMockRepoIssueListerReturnValues(lastPageErr, issuesOnPages...) {
		rv = append(rv, &MockIssueSearcherReturnValues{
			IssueSearchResult: &github.IssuesSearchResult{Issues: baseRV.Issues},
			Response:          baseRV.Response,
			Err:               baseRV.Err,
		})
	}
	return rv
}

func (m *MockIssueSearcher) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	if len(m.ReturnValues) == 0 {
		return &github.IssuesSearchResult{}, &github.Response{}, nil
	}
	if len(m.ReturnValues) <= m.cursor {
		panic("MockIssueSearcher.Issues called more times than it has ReturnValues mocked")
	}
	if m.TestFunc != nil {
		m.TestFunc(m.T, ctx, query, opts)
	}
	rv := m.ReturnValues[m.cursor]
	m.cursor++
	return rv.IssueSearchResult, rv.Response, rv.Err
}

func NewEmptyIssues(num int) []*github.Issue {
	var issues []*github.Issue
	for i := 1; i <= num; i++ {
		issues = append(issues, &github.Issue{Number: &i})
	}
	return issues
}
