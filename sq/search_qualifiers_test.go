package sq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchQualifiers(t *testing.T) {
	t.Parallel()

	eq := func(s string, q SearchQualifier) {
		assert.Equal(t, s, string(q))
	}

	eq(`reviewed-by:"user"`, IsReviewedBy("user"))
	eq(`review-requested:"user"`, WithReviewRequestFromUser("user"))
	eq(`user-review-requested:"user"`, WithReviewRequestDirectlyFromUser("user"))
	eq(`team-review-requested:"team"`, WithReviewRequestFromTeam("team"))
	eq(`user:"user"`, InReposOwnedByUser("user"))
	eq(`org:"org"`, InReposOwnedByOrg("org"))
	eq(`repo:"owner/repo"`, InRepo("owner", "repo"))
	eq(`author:"author"`, IsAuthoredBy("author"))
	eq(`-author:"author"`, IsNotAuthoredBy("author"))
	eq(`assignee:"assignee"`, IsAssignedTo("assignee"))
	eq(`mentions:"user"`, WhereUserIsMentioned("user"))
	eq(`team:"team"`, WhereTeamIsMentioned("team"))
	eq(`commenter:"user"`, HasCommentFromUser("user"))
	eq(`involves:"user"`, HasUserInvolved("user"))
	eq(`label:"label"`, WithLabel("label"))
	eq(`-label:"label"`, WithoutLabel("label"))
	eq(`milestone:"milestone"`, WithMilestone("milestone"))
	eq(`project:"project"`, WithProject("project"))
	eq(`head:"branch"`, WithMergeSourceBranch("branch"))
	eq(`base:"branch"`, WithMergeTargetBranch("branch"))
	eq(`language:"lang"`, WrittenInLanguage("lang"))
	eq(`comments:"5"`, WithNumberOfComments("5"))
	eq(`interactions:">5"`, WithNumberOfInteractions(">5"))
	eq(`reactions:"5..10"`, WithNumberOfReactions("5..10"))
	eq(`created:"2006-01-02"`, WasCreated("2006-01-02"))
	eq(`updated:">2006-01-02"`, WasUpdated(">2006-01-02"))
	eq(`closed:"2006-01-02..2006-02-03"`, WasClosed("2006-01-02..2006-02-03"))
	eq(`merged:"<=2006=01-02"`, WasMerged("<=2006=01-02"))
	eq(`"text"`, HasText("text"))
}
