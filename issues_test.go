package ghx

import (
	"context"
	"errors"
	"testing"

	"github.com/bevicted/ghx/ghxtest"
	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
)

func TestMapIssuesOfRepo(t *testing.T) {
	t.Parallel()

	const (
		testOwner = "testOwner"
		testRepo  = "testRepo"
	)

	testCtx := context.Background()
	for _, tc := range []struct {
		name          string
		lastPageErr   error
		issuesOnPages []int
		handlerErr    error
		expectIssues  int
		expectErr     error
	}{
		{
			name: "ok - 0 issues",
		},
		{
			name:          "ok - 1 page, 1 issue",
			issuesOnPages: []int{1},
			expectIssues:  1,
		},
		{
			name:          "ok - 3 pages, 15 issues",
			issuesOnPages: []int{5, 5, 5},
			expectIssues:  15,
		},
		{
			name:        "err - 1 page, repoIssueListerErr on page 1",
			lastPageErr: errors.New("repoIssueListerErr"),
			expectErr:   errors.New("repoIssueListerErr"),
		},
		{
			name:          "err - 3 pages, 10 issues, repoIssueListerErr on page 3",
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 5, 0},
			expectIssues:  10,
			expectErr:     errors.New("repoIssueListerErr"),
		},
		{
			name:          "err - 2 pages, handlerErr fails before repoIssueListerErr",
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 0},
			handlerErr:    errors.New("handlerErr"),
			expectErr:     errors.New("handlerErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var issueCount int
			opts := &MapIssuesOfRepoOptions{
				Owner: testOwner,
				Repo:  testRepo,
				IssueListByRepoOptions: github.IssueListByRepoOptions{
					Milestone: "testMilestone",
					State:     "testState",
					Assignee:  "testAssignee",
				},
			}
			mockRepoIssueListerTestFunc := func(t *testing.T, ctx context.Context, owner, repo string, issueListByRepoOptions *github.IssueListByRepoOptions) {
				t.Helper()
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testOwner, owner)
				assert.Equal(t, testRepo, repo)
				assert.Equal(t, opts.IssueListByRepoOptions, *issueListByRepoOptions)
			}
			mockRepoIssueLister := &ghxtest.MockRepoIssueLister{
				T:            t,
				TestFunc:     mockRepoIssueListerTestFunc,
				ReturnValues: ghxtest.NewMockRepoIssueListerReturnValues(tc.lastPageErr, tc.issuesOnPages...),
			}

			assert.Equal(t, tc.expectErr, MapIssuesOfRepo(testCtx, mockRepoIssueLister, opts, func(_ *github.Issue) error {
				if tc.handlerErr == nil {
					issueCount++
				}
				return tc.handlerErr
			}))
			assert.Equal(t, tc.expectIssues, issueCount, "handled different amount of issues than expected")
		})
	}
}

func TestMapSearchIssues(t *testing.T) {
	t.Parallel()

	testCtx := context.Background()

	for _, tc := range []struct {
		name             string
		searchQualifiers SearchQualifiers
		issuesOnPages    []int
		lastPageErr      error
		handlerErr       error
		expectIssues     int
		expectErr        error
	}{
		{
			name: "ok - 0 issues",
		},
		{
			name: "ok - 1 page, 1 issue",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
				HasText("some very specific text"),
			},
			issuesOnPages: []int{1},
			expectIssues:  1,
		},
		{
			name: "ok - 3 pages, 15 issues",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			issuesOnPages: []int{5, 5, 5},
			expectIssues:  15,
		},
		{
			name: "err - 1 page, repoIssueListerErr on page 1",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr: errors.New("repoIssueListerErr"),
			expectErr:   errors.New("repoIssueListerErr"),
		},
		{
			name: "err - 3 pages, 10 issues, repoIssueListerErr on page 3",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 5, 0},
			expectIssues:  10,
			expectErr:     errors.New("repoIssueListerErr"),
		},
		{
			name: "err - 2 pages, handlerErr fails before repoIssueListerErr",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 0},
			handlerErr:    errors.New("handlerErr"),
			expectErr:     errors.New("handlerErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var issueCount int
			opts := &MapSearchIssuesOptions{
				SearchQualifiers: tc.searchQualifiers,
				SearchOptions: github.SearchOptions{
					Sort:  "testSort",
					Order: "testOrder",
				},
			}
			mockIssueSearcherTestFunc := func(t *testing.T, ctx context.Context, query string, searchOptions *github.SearchOptions) {
				t.Helper()
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, tc.searchQualifiers.String(), query)
				assert.Equal(t, opts.SearchOptions, *searchOptions)
			}
			issueSearcher := &ghxtest.MockIssueSearcher{
				T:            t,
				TestFunc:     mockIssueSearcherTestFunc,
				ReturnValues: ghxtest.NewMockIssueSearcherReturnValues(tc.lastPageErr, tc.issuesOnPages...),
			}
			assert.Equal(t, tc.expectErr, MapSearchIssues(testCtx, issueSearcher, opts, func(_ *github.Issue) error {
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

	testCtx := context.Background()

	for _, tc := range []struct {
		name             string
		searchQualifiers SearchQualifiers
		issuesOnPages    []int
		lastPageErr      error
		handlerErr       error
		expectIssues     int
		expectErr        error
		mockIssue        *github.Issue
	}{
		{
			name: "ok - 0 issues",
		},
		{
			name: "ok - 1 page, 1 issue",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
				HasText("some very specific text"),
			},
			issuesOnPages: []int{1},
			expectIssues:  1,
		},
		{
			name: "ok - 3 pages, 15 issues",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			issuesOnPages: []int{5, 5, 5},
			expectIssues:  15,
		},
		{
			name: "err - 1 page, repoIssueListerErr on page 1",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr: errors.New("repoIssueListerErr"),
			expectErr:   errors.New("repoIssueListerErr"),
		},
		{
			name: "err - 3 pages, 10 issues, repoIssueListerErr on page 3",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 5, 0},
			expectIssues:  10,
			expectErr:     errors.New("repoIssueListerErr"),
		},
		{
			name: "err - 2 pages, handlerErr fails before repoIssueListerErr",
			searchQualifiers: SearchQualifiers{
				IsIssue,
				IsOpen,
				RestrictTextSearchToBody,
			},
			lastPageErr:   errors.New("repoIssueListerErr"),
			issuesOnPages: []int{5, 0},
			handlerErr:    errors.New("handlerErr"),
			expectErr:     errors.New("handlerErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockIssueSearcherTestFunc := func(t *testing.T, ctx context.Context, query string, searchOptions *github.SearchOptions) {
				t.Helper()
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, tc.searchQualifiers.String(), query)
				assert.Equal(t, github.SearchOptions{ListOptions: github.ListOptions{}}, *searchOptions)
			}
			issueSearcher := &ghxtest.MockIssueSearcher{
				T:        t,
				TestFunc: mockIssueSearcherTestFunc,
				ReturnValues: []*ghxtest.MockIssueSearcherReturnValues{
					{
						IssueSearchResult: &github.IssuesSearchResult{
							Issues: []*github.Issue{tc.mockIssue},
						},
					},
				},
			}

			issue, err := SearchOneIssue(testCtx, issueSearcher, tc.searchQualifiers)
			assert.Equal(t, tc.mockIssue, issue)
			assert.Equal(t, tc.expectErr, err)
		})
	}
}
