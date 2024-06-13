package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type SearchService struct {
	codeF         func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error)
	commitsF      func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error)
	issuesF       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
	labelsF       func(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error)
	repositoriesF func(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error)
	topicsF       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error)
	usersF        func(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error)

	issueF     func(ctx context.Context, query string) (*github.Issue, error)
	mapIssuesF func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error
}

func (s *SearchService) Code(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error) {
	return s.codeF(ctx, query, opts)
}
func (s *SearchService) Commits(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error) {
	return s.commitsF(ctx, query, opts)
}
func (s *SearchService) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	return s.issuesF(ctx, query, opts)
}
func (s *SearchService) Labels(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error) {
	return s.labelsF(ctx, repoID, query, opts)
}
func (s *SearchService) Repositories(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error) {
	return s.repositoriesF(ctx, query, opts)
}
func (s *SearchService) Topics(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error) {
	return s.topicsF(ctx, query, opts)
}
func (s *SearchService) Users(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error) {
	return s.usersF(ctx, query, opts)
}

func (s *SearchService) Issue(ctx context.Context, query string) (*github.Issue, error) {
	return s.issueF(ctx, query)
}
func (s *SearchService) MapIssues(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
	return s.mapIssuesF(ctx, query, opts, handle)
}

func newSearchService(client *github.Client) *SearchService {
	s := &SearchService{
		codeF:         client.Search.Code,
		commitsF:      client.Search.Commits,
		issuesF:       client.Search.Issues,
		labelsF:       client.Search.Labels,
		repositoriesF: client.Search.Repositories,
		topicsF:       client.Search.Topics,
		usersF:        client.Search.Users,
	}
	s.issueF = newIssueF(s)
	s.mapIssuesF = newMapIssuesF(s)
	return s
}

func newIssueF(searchService *SearchService) func(ctx context.Context, query string) (*github.Issue, error) {
	return func(ctx context.Context, query string) (*github.Issue, error) {
		issuesSearchResult, _, err := searchService.Issues(ctx, query, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}})
		if err != nil || len(issuesSearchResult.Issues) < 1 {
			return nil, err
		}

		return issuesSearchResult.Issues[0], nil
	}
}

func newMapIssuesF(searchService *SearchService) func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
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
