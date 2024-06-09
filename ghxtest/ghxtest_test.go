package ghxtest

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMockRepoIssueListerReturnValues(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name          string
		lastPageErr   error
		issuesOnPages []int
		expectValues  []*MockRepoIssueListerReturnValues
	}{
		{
			name:          "5 issues",
			issuesOnPages: []int{5},
			expectValues: []*MockRepoIssueListerReturnValues{
				{
					Issues:   NewEmptyIssues(5),
					Response: &github.Response{NextPage: 0},
				},
			},
		},
		{
			name:          "5 issues, 5 issues, 2 issues",
			issuesOnPages: []int{5, 5, 2},
			expectValues: []*MockRepoIssueListerReturnValues{
				{
					Issues:   NewEmptyIssues(5),
					Response: &github.Response{NextPage: 2},
				},
				{
					Issues:   NewEmptyIssues(5),
					Response: &github.Response{NextPage: 3},
				},
				{
					Issues:   NewEmptyIssues(2),
					Response: &github.Response{NextPage: 0},
				},
			},
		},
		{
			name:          "5 issues, 5 issues, 0 issues + err",
			issuesOnPages: []int{5, 5, 999},
			lastPageErr:   errors.New("3rd page err"),
			expectValues: []*MockRepoIssueListerReturnValues{
				{
					Issues:   NewEmptyIssues(5),
					Response: &github.Response{NextPage: 2},
				},
				{
					Issues:   NewEmptyIssues(5),
					Response: &github.Response{NextPage: 3},
				},
				{
					Err: errors.New("3rd page err"),
				},
			},
		},
		{
			name:         "oops, all empty",
			expectValues: []*MockRepoIssueListerReturnValues{},
		},
		{
			name:        "error only",
			lastPageErr: errors.New("error only"),
			expectValues: []*MockRepoIssueListerReturnValues{
				{
					Err: errors.New("error only"),
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.EqualValues(t, tc.expectValues, NewMockRepoIssueListerReturnValues(tc.lastPageErr, tc.issuesOnPages...))
		})
	}
}

func TestMockRepoIssueListerListByRepo(t *testing.T) {
	t.Parallel()

	var (
		testCtx   = context.Background()
		testOwner = "testOwner"
		testRepo  = "testRepo"
	)
	for _, tc := range []struct {
		name                string
		callCount           int
		expectIssues        []int
		expectErrOnLastCall error
	}{
		{
			name:         "single call, 5 issues, no err",
			callCount:    1,
			expectIssues: []int{5},
		},
		{
			name:         "multiple calls, 5-3-2 issues, no err",
			callCount:    3,
			expectIssues: []int{5, 3, 2},
		},
		{
			name:                "single call, err",
			callCount:           1,
			expectIssues:        []int{0},
			expectErrOnLastCall: errors.New("lastPageErr"),
		},
		{
			name:                "multiple calls, 5-3-err",
			callCount:           3,
			expectIssues:        []int{5, 3, 0},
			expectErrOnLastCall: errors.New("lastPageErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Len(t, tc.expectIssues, tc.callCount)

			opts := &github.IssueListByRepoOptions{
				Milestone:   "testMilestone",
				State:       "testState",
				Assignee:    "testAssignee",
				ListOptions: github.ListOptions{},
			}
			mril := &MockRepoIssueLister{
				T: t,
				TestFunc: func(testFuncT *testing.T, ctx context.Context, owner, repo string, testFuncOpts *github.IssueListByRepoOptions) { //nolint:thelper // have to differentiate T
					t.Helper()
					assert.Same(t, t, testFuncT)
					assert.Equal(t, testCtx, ctx)
					assert.Equal(t, testOwner, owner)
					assert.Equal(t, testRepo, repo)
					assert.Same(t, opts, testFuncOpts)
				},
				ReturnValues: NewMockRepoIssueListerReturnValues(tc.expectErrOnLastCall, tc.expectIssues...),
			}
			for i := 1; i <= tc.callCount; i++ {
				issues, res, err := mril.ListByRepo(testCtx, testOwner, testRepo, opts)

				assert.Len(t, issues, tc.expectIssues[i-1])

				isLastCall := i == tc.callCount

				var expectNextPage int
				if !isLastCall {
					expectNextPage = i + 1
				}

				if tc.expectErrOnLastCall != nil && isLastCall {
					assert.Nil(t, res)
				} else {
					assert.Equal(t, expectNextPage, res.NextPage)
					opts.Page = res.NextPage
				}

				var expectErr error
				if isLastCall {
					expectErr = tc.expectErrOnLastCall
				}
				assert.Equal(t, expectErr, err)
			}
		})
	}
}

func TestNewMockIssueSearcherReturnValues(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name          string
		lastPageErr   error
		issuesOnPages []int
		expectValues  []*MockIssueSearcherReturnValues
	}{
		{
			name:          "5 issues",
			issuesOnPages: []int{5},
			expectValues: []*MockIssueSearcherReturnValues{
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(5)},
					Response:          &github.Response{NextPage: 0},
				},
			},
		},
		{
			name:          "5 issues, 5 issues, 2 issues",
			issuesOnPages: []int{5, 5, 2},
			expectValues: []*MockIssueSearcherReturnValues{
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(5)},
					Response:          &github.Response{NextPage: 2},
				},
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(5)},
					Response:          &github.Response{NextPage: 3},
				},
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(2)},
					Response:          &github.Response{NextPage: 0},
				},
			},
		},
		{
			name:          "5 issues, 5 issues, 0 issues + err",
			issuesOnPages: []int{5, 5, 999},
			lastPageErr:   errors.New("3rd page err"),
			expectValues: []*MockIssueSearcherReturnValues{
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(5)},
					Response:          &github.Response{NextPage: 2},
				},
				{
					IssueSearchResult: &github.IssuesSearchResult{Issues: NewEmptyIssues(5)},
					Response:          &github.Response{NextPage: 3},
				},
				{
					Err: errors.New("3rd page err"),
				},
			},
		},
		{
			name:         "oops, all empty",
			expectValues: []*MockIssueSearcherReturnValues{},
		},
		{
			name:        "error only",
			lastPageErr: errors.New("error only"),
			expectValues: []*MockIssueSearcherReturnValues{
				{
					Err: errors.New("error only"),
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.EqualValues(t, tc.expectValues, NewMockIssueSearcherReturnValues(tc.lastPageErr, tc.issuesOnPages...))
		})
	}
}

func TestNewEmptyIssues(t *testing.T) {
	t.Parallel()

	numP := func(n int) *int {
		return &n
	}

	for _, tc := range []struct {
		name         string
		num          int
		expectIssues []*github.Issue
	}{
		{
			name: "single issue",
			num:  1,
			expectIssues: []*github.Issue{
				{Number: numP(1)},
			},
		},
		{
			name: "multiple issues",
			num:  3,
			expectIssues: []*github.Issue{
				{Number: numP(1)},
				{Number: numP(2)},
				{Number: numP(3)},
			},
		},
		{
			name:         "oops, all zeroes",
			expectIssues: []*github.Issue{},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectIssues, NewEmptyIssues(tc.num))
		})
	}
}
