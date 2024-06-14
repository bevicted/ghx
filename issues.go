package ghx

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type IssuesServiceF struct {
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
	RemoveAssignees        func(ctx context.Context, owner string, repo string, number int, assignees []string) (*github.Issue, *github.Response, error)
	RemoveLabelForIssue    func(ctx context.Context, owner string, repo string, number int, label string) (*github.Response, error)
	RemoveLabelsForIssue   func(ctx context.Context, owner string, repo string, number int) (*github.Response, error)
	RemoveMilestone        func(ctx context.Context, owner string, repo string, issueNumber int) (*github.Issue, *github.Response, error)
	ReplaceLabelsForIssue  func(ctx context.Context, owner string, repo string, number int, labels []string) ([]*github.Label, *github.Response, error)
	Unlock                 func(ctx context.Context, owner string, repo string, number int) (*github.Response, error)

	MapByRepo func(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions, handle IssueHandler) error
}

type IssuesService struct {
	f *IssuesServiceF
}

func (i *IssuesService) AddAssignees(ctx context.Context, owner string, repo string, number int, assignees []string) (*github.Issue, *github.Response, error) {
	return i.f.AddAssignees(ctx, owner, repo, number, assignees)
}
func (i *IssuesService) AddLabelsToIssue(ctx context.Context, owner string, repo string, number int, labels []string) ([]*github.Label, *github.Response, error) {
	return i.f.AddLabelsToIssue(ctx, owner, repo, number, labels)
}
func (i *IssuesService) Create(ctx context.Context, owner string, repo string, issue *github.IssueRequest) (*github.Issue, *github.Response, error) {
	return i.f.Create(ctx, owner, repo, issue)
}
func (i *IssuesService) CreateComment(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	return i.f.CreateComment(ctx, owner, repo, number, comment)
}
func (i *IssuesService) CreateLabel(ctx context.Context, owner string, repo string, label *github.Label) (*github.Label, *github.Response, error) {
	return i.f.CreateLabel(ctx, owner, repo, label)
}
func (i *IssuesService) CreateMilestone(ctx context.Context, owner string, repo string, milestone *github.Milestone) (*github.Milestone, *github.Response, error) {
	return i.f.CreateMilestone(ctx, owner, repo, milestone)
}
func (i *IssuesService) DeleteComment(ctx context.Context, owner string, repo string, commentID int64) (*github.Response, error) {
	return i.f.DeleteComment(ctx, owner, repo, commentID)
}
func (i *IssuesService) DeleteLabel(ctx context.Context, owner string, repo string, name string) (*github.Response, error) {
	return i.f.DeleteLabel(ctx, owner, repo, name)
}
func (i *IssuesService) DeleteMilestone(ctx context.Context, owner string, repo string, number int) (*github.Response, error) {
	return i.f.DeleteMilestone(ctx, owner, repo, number)
}
func (i *IssuesService) Edit(ctx context.Context, owner string, repo string, number int, issue *github.IssueRequest) (*github.Issue, *github.Response, error) {
	return i.f.Edit(ctx, owner, repo, number, issue)
}
func (i *IssuesService) EditComment(ctx context.Context, owner string, repo string, commentID int64, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	return i.f.EditComment(ctx, owner, repo, commentID, comment)
}
func (i *IssuesService) EditLabel(ctx context.Context, owner string, repo string, name string, label *github.Label) (*github.Label, *github.Response, error) {
	return i.f.EditLabel(ctx, owner, repo, name, label)
}
func (i *IssuesService) EditMilestone(ctx context.Context, owner string, repo string, number int, milestone *github.Milestone) (*github.Milestone, *github.Response, error) {
	return i.f.EditMilestone(ctx, owner, repo, number, milestone)
}
func (i *IssuesService) Get(ctx context.Context, owner string, repo string, number int) (*github.Issue, *github.Response, error) {
	return i.f.Get(ctx, owner, repo, number)
}
func (i *IssuesService) GetComment(ctx context.Context, owner string, repo string, commentID int64) (*github.IssueComment, *github.Response, error) {
	return i.f.GetComment(ctx, owner, repo, commentID)
}
func (i *IssuesService) GetEvent(ctx context.Context, owner string, repo string, id int64) (*github.IssueEvent, *github.Response, error) {
	return i.f.GetEvent(ctx, owner, repo, id)
}
func (i *IssuesService) GetLabel(ctx context.Context, owner string, repo string, name string) (*github.Label, *github.Response, error) {
	return i.f.GetLabel(ctx, owner, repo, name)
}
func (i *IssuesService) GetMilestone(ctx context.Context, owner string, repo string, number int) (*github.Milestone, *github.Response, error) {
	return i.f.GetMilestone(ctx, owner, repo, number)
}
func (i *IssuesService) IsAssignee(ctx context.Context, owner string, repo string, user string) (bool, *github.Response, error) {
	return i.f.IsAssignee(ctx, owner, repo, user)
}
func (i *IssuesService) List(ctx context.Context, all bool, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error) {
	return i.f.List(ctx, all, opts)
}
func (i *IssuesService) ListAssignees(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.User, *github.Response, error) {
	return i.f.ListAssignees(ctx, owner, repo, opts)
}
func (i *IssuesService) ListByOrg(ctx context.Context, org string, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error) {
	return i.f.ListByOrg(ctx, org, opts)
}
func (i *IssuesService) ListByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
	return i.f.ListByRepo(ctx, owner, repo, opts)
}
func (i *IssuesService) ListComments(ctx context.Context, owner string, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	return i.f.ListComments(ctx, owner, repo, number, opts)
}
func (i *IssuesService) ListIssueEvents(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.IssueEvent, *github.Response, error) {
	return i.f.ListIssueEvents(ctx, owner, repo, number, opts)
}
func (i *IssuesService) ListIssueTimeline(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Timeline, *github.Response, error) {
	return i.f.ListIssueTimeline(ctx, owner, repo, number, opts)
}
func (i *IssuesService) ListLabels(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Label, *github.Response, error) {
	return i.f.ListLabels(ctx, owner, repo, opts)
}
func (i *IssuesService) ListLabelsByIssue(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Label, *github.Response, error) {
	return i.f.ListLabelsByIssue(ctx, owner, repo, number, opts)
}
func (i *IssuesService) ListLabelsForMilestone(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Label, *github.Response, error) {
	return i.f.ListLabelsForMilestone(ctx, owner, repo, number, opts)
}
func (i *IssuesService) ListMilestones(ctx context.Context, owner string, repo string, opts *github.MilestoneListOptions) ([]*github.Milestone, *github.Response, error) {
	return i.f.ListMilestones(ctx, owner, repo, opts)
}
func (i *IssuesService) ListRepositoryEvents(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.IssueEvent, *github.Response, error) {
	return i.f.ListRepositoryEvents(ctx, owner, repo, opts)
}
func (i *IssuesService) Lock(ctx context.Context, owner string, repo string, number int, opts *github.LockIssueOptions) (*github.Response, error) {
	return i.f.Lock(ctx, owner, repo, number, opts)
}
func (i *IssuesService) RemoveAssignees(ctx context.Context, owner string, repo string, number int, assignees []string) (*github.Issue, *github.Response, error) {
	return i.f.RemoveAssignees(ctx, owner, repo, number, assignees)
}
func (i *IssuesService) RemoveLabelForIssue(ctx context.Context, owner string, repo string, number int, label string) (*github.Response, error) {
	return i.f.RemoveLabelForIssue(ctx, owner, repo, number, label)
}
func (i *IssuesService) RemoveLabelsForIssue(ctx context.Context, owner string, repo string, number int) (*github.Response, error) {
	return i.f.RemoveLabelsForIssue(ctx, owner, repo, number)
}
func (i *IssuesService) RemoveMilestone(ctx context.Context, owner string, repo string, issueNumber int) (*github.Issue, *github.Response, error) {
	return i.f.RemoveMilestone(ctx, owner, repo, issueNumber)
}
func (i *IssuesService) ReplaceLabelsForIssue(ctx context.Context, owner string, repo string, number int, labels []string) ([]*github.Label, *github.Response, error) {
	return i.f.ReplaceLabelsForIssue(ctx, owner, repo, number, labels)
}
func (i *IssuesService) Unlock(ctx context.Context, owner string, repo string, number int) (*github.Response, error) {
	return i.f.Unlock(ctx, owner, repo, number)
}

func (i *IssuesService) MapByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions, handle IssueHandler) error {
	return i.f.MapByRepo(ctx, owner, repo, opts, handle)
}

func NewIssuesService(f *IssuesServiceF) *IssuesService {
	return &IssuesService{f: f}
}

func newIssuesServicePassthrough(client *github.Client) *IssuesService {
	i := NewIssuesService(&IssuesServiceF{
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
		RemoveAssignees:        client.Issues.RemoveAssignees,
		RemoveLabelForIssue:    client.Issues.RemoveLabelForIssue,
		RemoveLabelsForIssue:   client.Issues.RemoveLabelsForIssue,
		RemoveMilestone:        client.Issues.RemoveMilestone,
		ReplaceLabelsForIssue:  client.Issues.ReplaceLabelsForIssue,
		Unlock:                 client.Issues.Unlock,
	})
	i.f.MapByRepo = newMapByRepoF(i)
	return i
}

type issueLister interface {
	ListByRepo(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error)
}

func newMapByRepoF(issuesService issueLister) func(ctx context.Context, owner string, repo string, opts *github.IssueListByRepoOptions, handle IssueHandler) error {
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
