package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type SearchServiceF struct {
	Code         func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error)
	Commits      func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error)
	Issues       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
	Labels       func(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error)
	Repositories func(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error)
	Topics       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error)
	Users        func(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error)

	Issue     func(ctx context.Context, query string) (*github.Issue, error)
	MapIssues func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error
}

type SearchService struct {
	f *SearchServiceF
}

func (s *SearchService) Code(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error) {
	return s.f.Code(ctx, query, opts)
}
func (s *SearchService) Commits(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error) {
	return s.f.Commits(ctx, query, opts)
}
func (s *SearchService) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	return s.f.Issues(ctx, query, opts)
}
func (s *SearchService) Labels(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error) {
	return s.f.Labels(ctx, repoID, query, opts)
}
func (s *SearchService) Repositories(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error) {
	return s.f.Repositories(ctx, query, opts)
}
func (s *SearchService) Topics(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error) {
	return s.f.Topics(ctx, query, opts)
}
func (s *SearchService) Users(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error) {
	return s.f.Users(ctx, query, opts)
}

func (s *SearchService) Issue(ctx context.Context, query string) (*github.Issue, error) {
	return s.f.Issue(ctx, query)
}
func (s *SearchService) MapIssues(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
	return s.f.MapIssues(ctx, query, opts, handle)
}

func NewSearchService(f *SearchServiceF) *SearchService {
	return &SearchService{f: f}
}

func newSearchServicePassthrough(client *github.Client) *SearchService {
	s := &SearchService{f: &SearchServiceF{
		Code:         client.Search.Code,
		Commits:      client.Search.Commits,
		Issues:       client.Search.Issues,
		Labels:       client.Search.Labels,
		Repositories: client.Search.Repositories,
		Topics:       client.Search.Topics,
		Users:        client.Search.Users,
	}}
	s.f.Issue = newIssueF(s)
	s.f.MapIssues = newMapIssuesF(s)
	return s
}

type issueSearcher interface {
	Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
}

func newIssueF(searchService issueSearcher) func(ctx context.Context, query string) (*github.Issue, error) {
	return func(ctx context.Context, query string) (*github.Issue, error) {
		issuesSearchResult, _, err := searchService.Issues(ctx, query, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}})
		if err != nil || len(issuesSearchResult.Issues) < 1 {
			return nil, err
		}

		return issuesSearchResult.Issues[0], nil
	}
}

func newMapIssuesF(searchService issueSearcher) func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
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
