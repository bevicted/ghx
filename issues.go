package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type MapIssuesOfRepoOptions struct {
	// The account owner of the repository. The name is not case sensitive.
	Owner string

	// The name of the repository without the .git extension. The name is not case sensitive.
	Repo string

	*github.IssueListByRepoOptions
}

type IssueHandler func(*github.Issue) error

func MapIssuesOfRepo(ctx context.Context, client *github.Client, opts *MapIssuesOfRepoOptions, handle IssueHandler) error {
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

type MapSearchIssuesOptions struct {
	SearchQualifiers SearchQualifiers

	*github.SearchOptions
}

func MapSearchIssues(ctx context.Context, client *github.Client, opts *MapSearchIssuesOptions, handle IssueHandler) error {
	for {
		issuesSearchResult, resp, err := client.Search.Issues(ctx, opts.SearchQualifiers.Join(), opts.SearchOptions)
		if err != nil {
			return err
		}
		for _, issue := range issuesSearchResult.Issues {
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

func SearchOneIssue(ctx context.Context, client *github.Client, searchQualifiers SearchQualifiers) (*github.Issue, error) {
	issuesSearchResult, _, err := client.Search.Issues(ctx, searchQualifiers.Join(), &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}})
	if err != nil || len(issuesSearchResult.Issues) < 1 {
		return nil, err
	}

	return issuesSearchResult.Issues[0], nil
}
