package ghx

import "github.com/google/go-github/v62/github"

const MaxPerPage = 100

type IssueHandler func(*github.Issue) error
type PullRequestHandler func(*github.PullRequest) error

func PTR[T comparable](v T) *T {
	return &v
}
