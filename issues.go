package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type mapByRepoF func(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions, handle IssueHandler) error

type IssuesService struct {
	AddAssignees           func(ctx context.Context, owner string, repo string, number int, assignees []string) (*github.Issue, *github.Response, error)
	AddLabelsToIssue       func(ctx context.Context, owner string, repo string, number int, labels []string) ([]*github.Label, *github.Response, error)
	Create                 func(ctx context.Context, owner string, repo string, issue *github.IssueRequest) (*github.Issue, *github.Response, error)
	CreateComment          func(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
	CreateLabel            func(ctx context.Context, owner string, repo string, label *github.Label) (*github.Label, *github.Response, error)
	CreateMilestone        func(ctx context.Context, owner string, repo string, milestone *github.Milestone) (*github.Milestone, *github.Response, error)
	DeleteComment          func(ctx context.Context, owner string, repo string, commentID int64) (*github.Response, error)
	DeleteLabel            func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	DeleteMilestone        func(ctx context.Context, owner string, repo string, number int) (*github.Response, error)
	Edit                   func(ctx context.Context, owner string, repo string, number int, issue *github.IssueRequest) (*github.Issue, *github.Response, error)
	EditComment            func(ctx context.Context, owner string, repo string, commentID int64, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
	EditLabel              func(ctx context.Context, owner string, repo string, name string, label *github.Label) (*github.Label, *github.Response, error)
	EditMilestone          func(ctx context.Context, owner string, repo string, number int, milestone *github.Milestone) (*github.Milestone, *github.Response, error)
	Get                    func(ctx context.Context, owner string, repo string, number int) (*github.Issue, *github.Response, error)
	GetComment             func(ctx context.Context, owner string, repo string, commentID int64) (*github.IssueComment, *github.Response, error)
	GetEvent               func(ctx context.Context, owner string, repo string, id int64) (*github.IssueEvent, *github.Response, error)
	GetLabel               func(ctx context.Context, owner string, repo string, name string) (*github.Label, *github.Response, error)
	GetMilestone           func(ctx context.Context, owner string, repo string, number int) (*github.Milestone, *github.Response, error)
	IsAssignee             func(ctx context.Context, owner string, repo string, user string) (bool, *github.Response, error)
	List                   func(ctx context.Context, all bool, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error)
	ListAssignees          func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListByOrg              func(ctx context.Context, org string, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error)
	ListByRepo             func(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error)
	ListComments           func(ctx context.Context, owner string, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)
	ListIssueEvents        func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.IssueEvent, *github.Response, error)
	ListIssueTimeline      func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Timeline, *github.Response, error)
	ListLabels             func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Label, *github.Response, error)
	ListLabelsByIssue      func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Label, *github.Response, error)
	ListLabelsForMilestone func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Label, *github.Response, error)
	ListMilestones         func(ctx context.Context, owner string, repo string, opts *github.MilestoneListOptions) ([]*github.Milestone, *github.Response, error)
	ListRepositoryEvents   func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.IssueEvent, *github.Response, error)
	Lock                   func(ctx context.Context, owner string, repo string, number int, opts *github.LockIssueOptions) (*github.Response, error)
	MapByRepo              mapByRepoF
	RemoveAssignees        func(ctx context.Context, owner string, repo string, number int, assignees []string) (*github.Issue, *github.Response, error)
	RemoveLabelForIssue    func(ctx context.Context, owner string, repo string, number int, label string) (*github.Response, error)
	RemoveLabelsForIssue   func(ctx context.Context, owner string, repo string, number int) (*github.Response, error)
	RemoveMilestone        func(ctx context.Context, owner string, repo string, issueNumber int) (*github.Issue, *github.Response, error)
	ReplaceLabelsForIssue  func(ctx context.Context, owner string, repo string, number int, labels []string) ([]*github.Label, *github.Response, error)
	Unlock                 func(ctx context.Context, owner string, repo string, number int) (*github.Response, error)
}

func newIssuesService(client *github.Client) *IssuesService {
	return &IssuesService{
		AddAssignees:           client.Issues.AddAssignees,
		AddLabelsToIssue:       client.Issues.AddLabelsToIssue,
		Create:                 client.Issues.Create,
		CreateComment:          client.Issues.CreateComment,
		CreateLabel:            client.Issues.CreateLabel,
		CreateMilestone:        client.Issues.CreateMilestone,
		DeleteComment:          client.Issues.DeleteComment,
		DeleteLabel:            client.Issues.DeleteLabel,
		DeleteMilestone:        client.Issues.DeleteMilestone,
		Edit:                   client.Issues.Edit,
		EditComment:            client.Issues.EditComment,
		EditLabel:              client.Issues.EditLabel,
		EditMilestone:          client.Issues.EditMilestone,
		Get:                    client.Issues.Get,
		GetComment:             client.Issues.GetComment,
		GetEvent:               client.Issues.GetEvent,
		GetLabel:               client.Issues.GetLabel,
		GetMilestone:           client.Issues.GetMilestone,
		IsAssignee:             client.Issues.IsAssignee,
		List:                   client.Issues.List,
		ListAssignees:          client.Issues.ListAssignees,
		ListByOrg:              client.Issues.ListByOrg,
		ListByRepo:             client.Issues.ListByRepo,
		ListComments:           client.Issues.ListComments,
		ListIssueEvents:        client.Issues.ListIssueEvents,
		ListIssueTimeline:      client.Issues.ListIssueTimeline,
		ListLabels:             client.Issues.ListLabels,
		ListLabelsByIssue:      client.Issues.ListLabelsByIssue,
		ListLabelsForMilestone: client.Issues.ListLabelsForMilestone,
		ListMilestones:         client.Issues.ListMilestones,
		ListRepositoryEvents:   client.Issues.ListRepositoryEvents,
		Lock:                   client.Issues.Lock,
		MapByRepo:              newMapByRepoF(&IssuesService{ListByRepo: client.Issues.ListByRepo}),
		RemoveAssignees:        client.Issues.RemoveAssignees,
		RemoveLabelForIssue:    client.Issues.RemoveLabelForIssue,
		RemoveLabelsForIssue:   client.Issues.RemoveLabelsForIssue,
		RemoveMilestone:        client.Issues.RemoveMilestone,
		ReplaceLabelsForIssue:  client.Issues.ReplaceLabelsForIssue,
		Unlock:                 client.Issues.Unlock,
	}
}

func newMapByRepoF(issuesService *IssuesService) mapByRepoF {
	return func(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions, handle IssueHandler) error {
		for {
			issues, resp, err := issuesService.ListByRepo(ctx, owner, repo, opts)
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
}
