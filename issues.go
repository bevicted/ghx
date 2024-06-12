package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type IssueHandler func(*github.Issue) error

type MapIssuesOfRepoOptions struct {
	// The account owner of the repository. The name is not case sensitive.
	Owner string

	// The name of the repository without the .git extension. The name is not case sensitive.
	Repo string

	github.IssueListByRepoOptions
}

func (is *IssuesService) MapIssuesOfRepo(ctx context.Context, opts *MapIssuesOfRepoOptions, handle IssueHandler) error {
	for {
		issues, resp, err := is.ListByRepo(ctx, opts.Owner, opts.Repo, &opts.IssueListByRepoOptions)
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

	github.SearchOptions
}

func MapSearchIssues(ctx context.Context, search IssueSearcher, opts *MapSearchIssuesOptions, handle IssueHandler) error {
	for {
		issuesSearchResult, resp, err := search.Issues(ctx, opts.SearchQualifiers.String(), &opts.SearchOptions)
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

func SearchOneIssue(ctx context.Context, search IssueSearcher, searchQualifiers SearchQualifiers) (*github.Issue, error) {
	issuesSearchResult, _, err := search.Issues(ctx, searchQualifiers.String(), &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}})
	if err != nil || len(issuesSearchResult.Issues) < 1 {
		return nil, err
	}

	return issuesSearchResult.Issues[0], nil
}
