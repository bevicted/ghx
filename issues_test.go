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

func TestMapIssuesOfRepo(t *testing.T) {
	t.Parallel()

	const (
		testOwner = "testOwner"
		testRepo  = "testRepo"
	)

	type listByRepoReturnValue struct {
		Issues []*github.Issue
		Res    *github.Response
		Err    error
	}

	testCtx := context.Background()
	for _, tc := range []struct {
		name                   string
		handlerErr             error
		expectIssueCount       int
		expectErr              error
		listByRepoReturnValues []listByRepoReturnValue
	}{
		{
			name: "ok - 0 issues",
		},
		{
			name: "ok - 1 page, 1 issue",
			listByRepoReturnValues: []listByRepoReturnValue{
				{
					Issues: ghxtest.NewEmptyIssues(t, 1),
				},
			},
			expectIssueCount: 1,
		},
		{
			name: "ok - 2 pages, 8 issues",
			listByRepoReturnValues: []listByRepoReturnValue{
				{
					Issues: ghxtest.NewEmptyIssues(t, 5),
					Res:    &github.Response{NextPage: 2},
				},
				{
					Issues: ghxtest.NewEmptyIssues(t, 3),
					Res:    &github.Response{NextPage: 0},
				},
			},
			expectIssueCount: 8,
		},
		{
			name: "err - 1 page, listByRepoErr on page 1",
			listByRepoReturnValues: []listByRepoReturnValue{
				{
					Err: errors.New("listByRepoErr"),
				},
			},
			expectErr: errors.New("listByRepoErr"),
		},
		{
			name: "err - 2 pages, 8 issues, listByRepoErr on page 2",
			listByRepoReturnValues: []listByRepoReturnValue{
				{
					Issues: ghxtest.NewEmptyIssues(t, 5),
					Res:    &github.Response{NextPage: 2},
				},
				{
					Err: errors.New("listByRepoErr"),
				},
			},
			expectIssueCount: 5,
			expectErr:        errors.New("listByRepoErr"),
		},
		{
			name: "err - 2 pages, handlerErr fails before listByRepoErr",
			listByRepoReturnValues: []listByRepoReturnValue{
				{
					Issues: ghxtest.NewEmptyIssues(t, 5),
					Res:    &github.Response{NextPage: 2},
				},
				{
					Err: errors.New("listByRepoErr"),
				},
			},
			handlerErr: errors.New("handlerErr"),
			expectErr:  errors.New("handlerErr"),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opts := &github.IssueListByRepoOptions{
				Milestone: "testMilestone",
				State:     "testState",
				Assignee:  "testAssignee",
			}
			var listByRepoCallCount int
			mapByRepo := newMapByRepoF(&IssuesService{
				ListByRepo: func(actualCtx context.Context, actualOwner, actualRepo string, actualOpts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
					assert.Equal(t, testCtx, actualCtx)
					assert.Equal(t, testOwner, actualOwner)
					assert.Equal(t, testRepo, actualRepo)
					assert.Same(t, opts, actualOpts)
					require.Less(t, listByRepoCallCount, len(tc.listByRepoReturnValues), "ListByRepo called more times than it has return values mocked")
					rv := tc.listByRepoReturnValues[listByRepoCallCount]
					listByRepoCallCount++
					return rv.Issues, rv.Res, rv.Err
				},
			})

			var issueCount int
			assert.Equal(t, tc.expectErr, mapByRepo(testCtx, testOwner, testRepo, opts, func(_ *github.Issue) error {
				if tc.handlerErr == nil {
					issueCount++
				}
				return tc.handlerErr
			}))
			assert.Equal(t, tc.expectIssueCount, issueCount, "handled different amount of issues than expected")
		})
	}
}
