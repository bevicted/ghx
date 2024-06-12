package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type (
	mapIssuesF func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error
	issueF     func(ctx context.Context, query string) (*github.Issue, error)
)

type SearchService struct {
	Code         func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error)
	Commits      func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error)
	Issue        issueF
	Issues       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
	Labels       func(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error)
	MapIssues    mapIssuesF
	Repositories func(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error)
	Topics       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error)
	Users        func(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error)
}

func newSearchService(client *github.Client) *SearchService {
	issueSearcher := &SearchService{Issues: client.Search.Issues}
	return &SearchService{
		Code:         client.Search.Code,
		Commits:      client.Search.Commits,
		Issue:        newIssueF(issueSearcher),
		Issues:       client.Search.Issues,
		Labels:       client.Search.Labels,
		MapIssues:    newMapIssuesF(issueSearcher),
		Repositories: client.Search.Repositories,
		Topics:       client.Search.Topics,
		Users:        client.Search.Users,
	}
}

func newMapIssuesF(searchService *SearchService) mapIssuesF {
	return func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
		for {
			issuesSearchResult, resp, err := searchService.Issues(ctx, query, opts)
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
}

func newIssueF(searchService *SearchService) issueF {
	return func(ctx context.Context, query string) (*github.Issue, error) {
		issuesSearchResult, _, err := searchService.Issues(ctx, query, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}})
		if err != nil || len(issuesSearchResult.Issues) < 1 {
			return nil, err
		}

		return issuesSearchResult.Issues[0], nil
	}
}
