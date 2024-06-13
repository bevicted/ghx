package ghx

import (
	"context"
	"errors"
	"testing"

	"github.com/bevicted/ghx/ghxtest"
	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type searchIssuesReturnValue struct {
	Result *github.IssuesSearchResult
	Res    *github.Response
	Err    error
}

func TestMapSearchIssues(t *testing.T) {
	t.Parallel()

	const testQuery = "testQuery"
	testCtx := context.Background()

	for _, tc := range []struct {
		expectErr                error
		expectIssues             int
		handlerErr               error
		name                     string
		searchIssuesReturnValues []searchIssuesReturnValue
	}{
		{
			name: "ok - 0 issues",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{},
					Res:    &github.Response{},
				},
			},
		},
		{
			name: "ok - 1 page, 1 issue",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 1)},
					Res:    &github.Response{},
				},
			},
			expectIssues: 1,
		},
		{
			name:         "ok - 2 pages, 8 issues",
			expectIssues: 8,
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 5)},
					Res:    &github.Response{NextPage: 2},
				},
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 3)},
					Res:    &github.Response{},
				},
			},
		},
		{
			name: "err - 1 page, searchIssueErr on page 1",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Err: errors.New("searchIssueErr"),
				},
			},
			expectErr: errors.New("searchIssueErr"),
		},
		{
			name: "err - 2 pages, 5 issues, searchIssueErr on page 2",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 5)},
					Res:    &github.Response{NextPage: 2},
				},
				{
					Err: errors.New("searchIssueErr"),
				},
			},
			expectIssues: 5,
			expectErr:    errors.New("searchIssueErr"),
		},
		{
			name: "err - 2 pages, handlerErr fails before searchIssueErr",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 5)},
					Res:    &github.Response{NextPage: 2},
				},
				{
					Err: errors.New("searchIssueErr"),
				},
			},
			handlerErr: errors.New("handlerErr"),
			expectErr:  errors.New("handlerErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opts := &github.SearchOptions{
				Sort:  "testSort",
				Order: "testOrder",
			}
			var searchIssuesCallCount int
			mapIssues := newMapIssuesF(NewSearchService(&SearchServiceF{
				Issues: func(actualCtx context.Context, actualQuery string, actualOpts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
					assert.Equal(t, testCtx, actualCtx)
					assert.Equal(t, testQuery, actualQuery)
					assert.Same(t, opts, actualOpts)
					require.Less(t, searchIssuesCallCount, len(tc.searchIssuesReturnValues), "Issues called more times than it has return values mocked")
					rv := tc.searchIssuesReturnValues[searchIssuesCallCount]
					searchIssuesCallCount++
					return rv.Result, rv.Res, rv.Err
				},
			}))
			var issueCount int
			assert.Equal(t, tc.expectErr, mapIssues(testCtx, testQuery, opts, func(_ *github.Issue) error {
				if tc.handlerErr == nil {
					issueCount++
				}
				return tc.handlerErr
			}))
			assert.Equal(t, tc.expectIssues, issueCount, "handled different amount of issues than expected")
		})
	}
}

func TestSearchOneIssue(t *testing.T) {
	t.Parallel()

	const testQuery = "testQuery"
	testCtx := context.Background()

	for _, tc := range []struct {
		name                     string
		searchIssueErr           error
		expectErr                error
		searchIssuesReturnValues []searchIssuesReturnValue
	}{
		{
			name: "ok - no query",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{},
					Res:    &github.Response{},
				},
			},
		},
		{
			name: "ok - got issue",
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Result: &github.IssuesSearchResult{Issues: ghxtest.NewEmptyIssues(t, 1)},
					Res:    &github.Response{},
				},
			},
		},
		{
			name:      "err - searchIssueErr",
			expectErr: errors.New("searchIssueErr"),
			searchIssuesReturnValues: []searchIssuesReturnValue{
				{
					Err: errors.New("searchIssueErr"),
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var searchIssuesCallCount int
			issueF := newIssueF(NewSearchService(&SearchServiceF{
				Issues: func(actualCtx context.Context, actualQuery string, actualOpts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
					assert.Equal(t, testCtx, actualCtx)
					assert.Equal(t, testQuery, actualQuery)
					assert.Equal(t, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}}, actualOpts)
					require.Less(t, searchIssuesCallCount, len(tc.searchIssuesReturnValues), "Issues called more times than it has return values mocked")
					rv := tc.searchIssuesReturnValues[searchIssuesCallCount]
					searchIssuesCallCount++
					return rv.Result, rv.Res, rv.Err
				},
			}))

			actualIssue, err := issueF(testCtx, testQuery)
			var expectIssue *github.Issue
			r := tc.searchIssuesReturnValues[0].Result
			if r != nil && len(r.Issues) > 0 {
				expectIssue = r.Issues[0]
			}
			assert.Equal(t, expectIssue, actualIssue)
			assert.Equal(t, tc.expectErr, err)
		})
	}
}
