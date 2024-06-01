package ghx

import (
	"context"
	"testing"

	"github.com/bevicted/ghx/ghxtest"
	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
)

func TestMapIssuesOfRepo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	for _, tc := range []struct {
		name            string
		repoIssueLister repoIssueLister
		opts            MapIssuesOfRepoOptions
		handle          IssueHandler
		expectErr       error
	}{
		{
			name: "ok",
			repoIssueLister: &ghxtest.MockRepoIssueLister{
				ReturnValues: ghxtest.NewMockRepoIssueListerReturnValues(nil, 5, 5, 5),
			},
			opts: MapIssuesOfRepoOptions{
				IssueListByRepoOptions: github.IssueListByRepoOptions{},
			},
			handle: func(i *github.Issue) error { return nil },
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.expectErr, MapIssuesOfRepo(ctx, tc.repoIssueLister, &tc.opts, tc.handle))
		})
	}
}
