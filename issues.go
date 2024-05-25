package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type MapIssuesOfRepoOpts struct {
	// The account owner of the repository. The name is not case sensitive.
	Owner string

	// The name of the repository without the .git extension. The name is not case sensitive.
	Repo string

	*github.IssueListByRepoOptions
}

type MapIssueOfRepoHandler func(*github.Issue) error

func MapIssuesOfRepo(ctx context.Context, client *github.Client, opts *MapIssuesOfRepoOpts, handle MapIssueOfRepoHandler) error {
	for {
		issues, resp, err := client.Issues.ListByRepo(ctx, opts.Owner, opts.Repo, opts.IssueListByRepoOptions)
		if err != nil {
			return err
		}
		for _, issue := range issues {
			if err := handle(issue); err != nil {
				return err
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil
}
