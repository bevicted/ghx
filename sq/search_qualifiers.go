// Based on https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests

package sq

import (
	"fmt"
	"strings"
)

type SearchQualifier string

func (s SearchQualifier) String() string {
	return string(s)
}

type SearchQualifiers []SearchQualifier

func (s SearchQualifiers) String() string {
	var query strings.Builder
	for idx, qualifier := range s {
		if idx != 0 {
			if _, err := query.WriteString(" "); err != nil {
				panic(err)
			}
		}
		if _, err := query.WriteString(qualifier.String()); err != nil {
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

	IsClosedBecauseCompleted  SearchQualifier = "reason:completed"
	IsClosedBecauseNotPlanned SearchQualifier = "reason:\"not planned\""

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

	HasStatusPending SearchQualifier = "status:pending"
	HasStatusSuccess SearchQualifier = "status:success"
	HasStatusFailing SearchQualifier = "status:failure"

	IsDraft    SearchQualifier = "draft:true"
	IsNotDraft SearchQualifier = "draft:false"

	HasNoReviews     SearchQualifier = "review:none"
	RequiresReview   SearchQualifier = "review:required"
	IsApproved       SearchQualifier = "review:approved"
	HasChangeRequest SearchQualifier = "review:changes_requested"

	IsArchived    SearchQualifier = "archived:true"
	IsNotArchived SearchQualifier = "archived:false"

	HasNoLabels     SearchQualifier = "no:label"
	HasNoMilestones SearchQualifier = "no:milestone"
	HasNoAssignees  SearchQualifier = "no:assignee"
	HasNoProjects   SearchQualifier = "no:project"

	IsAuthoredBySelf SearchQualifier = "author:@me"
)

func newSearchQualifier(key string, value string) SearchQualifier {
	return SearchQualifier(fmt.Sprintf("%s:%q", key, value))
}

func IsReviewedBy(user string) SearchQualifier {
	return newSearchQualifier("reviewed-by", user)
}

func HasUserReviewRequested(user string) SearchQualifier {
	return newSearchQualifier("review-requested", user)
}

func HasUserReviewDirectlyRequested(user string) SearchQualifier {
	return newSearchQualifier("user-review-requested", user)
}

func HasTeamReviewRequested(team string) SearchQualifier {
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

func HasUserMention(user string) SearchQualifier {
	return newSearchQualifier("mentions", user)
}

func HasTeamMention(team string) SearchQualifier {
	return newSearchQualifier("team", team)
}

func HasCommentFromUser(user string) SearchQualifier {
	return newSearchQualifier("commenter", user)
}

func HasUserInvolved(user string) SearchQualifier {
	return newSearchQualifier("involves", user)
}

func HasLabel(label string) SearchQualifier {
	return newSearchQualifier("label", label)
}

func HasLabels(labels ...string) SearchQualifier {
	return HasLabel(strings.Join(labels, ","))
}

func LacksLabel(label string) SearchQualifier {
	return newSearchQualifier("-label", label)
}

func LacksLabels(labels ...string) SearchQualifier {
	return LacksLabel(strings.Join(labels, ","))
}

func InMilestone(milestone string) SearchQualifier {
	return newSearchQualifier("milestone", milestone)
}

func InProject(project string) SearchQualifier {
	return newSearchQualifier("project", project)
}

func HasMergeSourceBranch(branch string) SearchQualifier {
	return newSearchQualifier("head", branch)
}

func HasMergeTargetBranch(branch string) SearchQualifier {
	return newSearchQualifier("base", branch)
}

func IsWrittenInLanguage(language string) SearchQualifier {
	return newSearchQualifier("language", language)
}

// e.g.: 0, >100, 500..1000
func HasNumberOfComments(quantity string) SearchQualifier {
	return newSearchQualifier("comments", quantity)
}

// e.g.: 0, >100, 500..1000
func HasNumberOfInteractions(quantity string) SearchQualifier {
	return newSearchQualifier("interactions", quantity)
}

// e.g.: 0, >100, 500..1000
func HasNumberOfReactions(quantity string) SearchQualifier {
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

func WithText(text string) SearchQualifier {
	return SearchQualifier(fmt.Sprintf("%q", text))
}
