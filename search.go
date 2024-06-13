package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type SearchServiceF struct {
	CodeF         func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error)
	CommitsF      func(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error)
	IssuesF       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
	LabelsF       func(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error)
	RepositoriesF func(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error)
	TopicsF       func(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error)
	UsersF        func(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error)

	IssueF     func(ctx context.Context, query string) (*github.Issue, error)
	MapIssuesF func(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error
}

type SearchService struct {
	f *SearchServiceF
}

func (s *SearchService) Code(ctx context.Context, query string, opts *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error) {
	return s.f.CodeF(ctx, query, opts)
}
func (s *SearchService) Commits(ctx context.Context, query string, opts *github.SearchOptions) (*github.CommitsSearchResult, *github.Response, error) {
	return s.f.CommitsF(ctx, query, opts)
}
func (s *SearchService) Issues(ctx context.Context, query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	return s.f.IssuesF(ctx, query, opts)
}
func (s *SearchService) Labels(ctx context.Context, repoID int64, query string, opts *github.SearchOptions) (*github.LabelsSearchResult, *github.Response, error) {
	return s.f.LabelsF(ctx, repoID, query, opts)
}
func (s *SearchService) Repositories(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error) {
	return s.f.RepositoriesF(ctx, query, opts)
}
func (s *SearchService) Topics(ctx context.Context, query string, opts *github.SearchOptions) (*github.TopicsSearchResult, *github.Response, error) {
	return s.f.TopicsF(ctx, query, opts)
}
func (s *SearchService) Users(ctx context.Context, query string, opts *github.SearchOptions) (*github.UsersSearchResult, *github.Response, error) {
	return s.f.UsersF(ctx, query, opts)
}

func (s *SearchService) Issue(ctx context.Context, query string) (*github.Issue, error) {
	return s.f.IssueF(ctx, query)
}
func (s *SearchService) MapIssues(ctx context.Context, query string, opts *github.SearchOptions, handle IssueHandler) error {
	return s.f.MapIssuesF(ctx, query, opts, handle)
}

func NewSearchService(f *SearchServiceF) *SearchService {
	return &SearchService{f: f}
}

func newSearchServicePassthrough(client *github.Client) *SearchService {
	s := &SearchService{
		f: &SearchServiceF{
			CodeF:         client.Search.Code,
			CommitsF:      client.Search.Commits,
			IssuesF:       client.Search.Issues,
			LabelsF:       client.Search.Labels,
			RepositoriesF: client.Search.Repositories,
			TopicsF:       client.Search.Topics,
			UsersF:        client.Search.Users,
		},
	}
	s.f.IssueF = newIssueF(s)
	s.f.MapIssuesF = newMapIssuesF(s)
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
