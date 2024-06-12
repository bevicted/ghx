package sq

import (
	"fmt"
	"strings"
)

type SearchQualifier string
type SearchQualifiers []SearchQualifier

func (s SearchQualifiers) String() string {
	var query strings.Builder
	for idx, qualifier := range s {
		if idx != 0 {
			if _, err := query.WriteString(" "); err != nil {
				panic(err)
			}
		}
		if _, err := query.WriteString(string(qualifier)); err != nil {
			panic(err)
		}
	}
	return query.String()
}

const (
	IsIssue SearchQualifier = "type:issue"
	IsPR    SearchQualifier = "type:pr"

	RestrictTextSearchToTitle    SearchQualifier = "in:title"
	RestrictTextSearchToBody     SearchQualifier = "in:body"
	RestrictTextSearchToComments SearchQualifier = "in:comments"

	IsOpen   SearchQualifier = "state:open"
	IsClosed SearchQualifier = "state:closed"

	ClosedBecauseCompleted  SearchQualifier = "reason:completed"
	ClosedBecauseNotPlanned SearchQualifier = "reason:\"not planned\""

	IsQueued   SearchQualifier = "is:queued"
	IsPublic   SearchQualifier = "is:public"
	IsPrivate  SearchQualifier = "is:private"
	IsMerged   SearchQualifier = "is:merged"
	IsUnmerged SearchQualifier = "is:unmerged"
	IsLocked   SearchQualifier = "is:locked"
	IsUnlocked SearchQualifier = "is:unlocked"

	HasLinkedPR      SearchQualifier = "linked:pr"
	HasLinkedIssue   SearchQualifier = "linked:issue"
	HasNoLinkedPR    SearchQualifier = "-linked:pr"
	HasNoLinkedIssue SearchQualifier = "-linked:issue"

	WithStatusPending SearchQualifier = "status:pending"
	WithStatusSuccess SearchQualifier = "status:success"
	WithStatusFailure SearchQualifier = "status:failure"

	IsDraft    SearchQualifier = "draft:true"
	IsNotDraft SearchQualifier = "draft:false"

	HasNoReview      SearchQualifier = "review:none"
	RequiresReview   SearchQualifier = "review:required"
	IsApproved       SearchQualifier = "review:approved"
	HasChangeRequest SearchQualifier = "review:changes_requested"

	IsArchived    SearchQualifier = "archived:true"
	IsNotArchived SearchQualifier = "archived:false"

	WithoutAnyLabels     SearchQualifier = "no:label"
	WithoutAnyMilestones SearchQualifier = "no:milestone"
	WithoutAnyAssignees  SearchQualifier = "no:assignee"
	WithoutAnyProjects   SearchQualifier = "no:project"

	IsAuthoredBySelf SearchQualifier = "author:@me"
)

func newSearchQualifier(key string, value string) SearchQualifier {
	return SearchQualifier(fmt.Sprintf("%s:%q", key, value))
}

func IsReviewedBy(user string) SearchQualifier {
	return newSearchQualifier("reviewed-by", user)
}

func WithReviewRequestFromUser(user string) SearchQualifier {
	return newSearchQualifier("review-requested", user)
}

func WithReviewRequestDirectlyFromUser(user string) SearchQualifier {
	return newSearchQualifier("user-review-requested", user)
}

func WithReviewRequestFromTeam(team string) SearchQualifier {
	return newSearchQualifier("team-review-requested", team)
}

func InReposOwnedByUser(user string) SearchQualifier {
	return newSearchQualifier("user", user)
}

func InReposOwnedByOrg(org string) SearchQualifier {
	return newSearchQualifier("org", org)
}

func InRepo(owner string, repo string) SearchQualifier {
	return newSearchQualifier("repo", fmt.Sprintf("%s/%s", owner, repo))
}

func IsAuthoredBy(user string) SearchQualifier {
	return newSearchQualifier("author", user)
}

func IsNotAuthoredBy(user string) SearchQualifier {
	return newSearchQualifier("-author", user)
}

func IsAssignedTo(assignee string) SearchQualifier {
	return newSearchQualifier("assignee", assignee)
}

func WhereUserIsMentioned(user string) SearchQualifier {
	return newSearchQualifier("mentions", user)
}

func WhereTeamIsMentioned(team string) SearchQualifier {
	return newSearchQualifier("team", team)
}

func HasCommentFromUser(user string) SearchQualifier {
	return newSearchQualifier("commenter", user)
}

func HasUserInvolved(user string) SearchQualifier {
	return newSearchQualifier("involves", user)
}

func WithLabel(label string) SearchQualifier {
	return newSearchQualifier("label", label)
}

func WithoutLabel(label string) SearchQualifier {
	return newSearchQualifier("-label", label)
}

func WithMilestone(milestone string) SearchQualifier {
	return newSearchQualifier("milestone", milestone)
}

func WithProject(project string) SearchQualifier {
	return newSearchQualifier("project", project)
}

func WithMergeSourceBranch(branch string) SearchQualifier {
	return newSearchQualifier("head", branch)
}

func WithMergeTargetBranch(branch string) SearchQualifier {
	return newSearchQualifier("base", branch)
}

func WrittenInLanguage(language string) SearchQualifier {
	return newSearchQualifier("language", language)
}

// e.g.: 0, >100, 500..1000
func WithNumberOfComments(quantity string) SearchQualifier {
	return newSearchQualifier("comments", quantity)
}

// e.g.: 0, >100, 500..1000
func WithNumberOfInteractions(quantity string) SearchQualifier {
	return newSearchQualifier("interactions", quantity)
}

// e.g.: 0, >100, 500..1000
func WithNumberOfReactions(quantity string) SearchQualifier {
	return newSearchQualifier("reactions", quantity)
}

// YYYY-MM-DD, >=YYYY-MM-DD, YYYY-MM-DD..YYYY-MM-DD, *..YYYY-MM-DD
func WasCreated(iso8601Date string) SearchQualifier {
	return newSearchQualifier("created", iso8601Date)
}

// YYYY-MM-DD, >=YYYY-MM-DD, YYYY-MM-DD..YYYY-MM-DD, *..YYYY-MM-DD
func WasUpdated(iso8601Date string) SearchQualifier {
	return newSearchQualifier("updated", iso8601Date)
}

// YYYY-MM-DD, >=YYYY-MM-DD, YYYY-MM-DD..YYYY-MM-DD, *..YYYY-MM-DD
func WasClosed(iso8601Date string) SearchQualifier {
	return newSearchQualifier("closed", iso8601Date)
}

// YYYY-MM-DD, >=YYYY-MM-DD, YYYY-MM-DD..YYYY-MM-DD, *..YYYY-MM-DD
func WasMerged(iso8601Date string) SearchQualifier {
	return newSearchQualifier("merged", iso8601Date)
}

func HasText(text string) SearchQualifier {
	return SearchQualifier(fmt.Sprintf("%q", text))
}
