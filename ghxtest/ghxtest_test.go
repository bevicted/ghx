package ghxtest

import (
	"errors"
	"testing"

	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
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
			name: "oops, all empty",
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
