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
	eq(`review-requested:"user"`, HasUserReviewRequested("user"))
	eq(`user-review-requested:"user"`, HasUserReviewDirectlyRequested("user"))
	eq(`team-review-requested:"team"`, HasTeamReviewRequested("team"))
	eq(`user:"user"`, InReposOwnedByUser("user"))
	eq(`org:"org"`, InReposOwnedByOrg("org"))
	eq(`repo:"owner/repo"`, InRepo("owner", "repo"))
	eq(`author:"author"`, IsAuthoredBy("author"))
	eq(`-author:"author"`, IsNotAuthoredBy("author"))
	eq(`assignee:"assignee"`, IsAssignedTo("assignee"))
	eq(`mentions:"user"`, HasUserMention("user"))
	eq(`team:"team"`, HasTeamMention("team"))
	eq(`commenter:"user"`, HasCommentFromUser("user"))
	eq(`involves:"user"`, HasUserInvolved("user"))
	eq(`label:"label"`, HasLabel("label"))
	eq(`label:"label1,label2"`, HasLabels("label1", "label2"))
	eq(`-label:"label"`, LacksLabel("label"))
	eq(`-label:"label1,label2"`, LacksLabels("label1", "label2"))
	eq(`milestone:"milestone"`, InMilestone("milestone"))
	eq(`project:"project"`, InProject("project"))
	eq(`head:"branch"`, HasMergeSourceBranch("branch"))
	eq(`base:"branch"`, HasMergeTargetBranch("branch"))
	eq(`language:"lang"`, IsWrittenInLanguage("lang"))
	eq(`comments:"5"`, HasNumberOfComments("5"))
	eq(`interactions:">5"`, HasNumberOfInteractions(">5"))
	eq(`reactions:"5..10"`, HasNumberOfReactions("5..10"))
	eq(`created:"2006-01-02"`, WasCreated("2006-01-02"))
	eq(`updated:">2006-01-02"`, WasUpdated(">2006-01-02"))
	eq(`closed:"2006-01-02..2006-02-03"`, WasClosed("2006-01-02..2006-02-03"))
	eq(`merged:"<=2006=01-02"`, WasMerged("<=2006=01-02"))
	eq(`"text"`, HasText("text"))
}
