package ghx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchQualifiers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, `reviewed-by:"user"`, IsReviewedBy("user"))
	assert.Equal(t, `review-requested:"user"`, WithReviewRequestFromUser("user"))
	assert.Equal(t, `user-review-requested:"user"`, WithReviewRequestDirectlyFromUser("user"))
	assert.Equal(t, `team-review-requested:"team"`, WithReviewRequestFromTeam("team"))
	assert.Equal(t, `user:"user"`, InReposOwnedByUser("user"))
	assert.Equal(t, `org:"org"`, InReposOwnedByOrg("org"))
	assert.Equal(t, `repo:"owner/repo"`, InRepo("owner", "repo"))
	assert.Equal(t, `author:"author"`, IsAuthoredBy("author"))
	assert.Equal(t, `-author:"author"`, IsNotAuthoredBy("author"))
	assert.Equal(t, `assignee:"assignee"`, IsAssignedTo("assignee"))
	assert.Equal(t, `mentions:"user"`, WhereUserIsMentioned("user"))
	assert.Equal(t, `team:"team"`, WhereTeamIsMentioned("team"))
	assert.Equal(t, `commenter:"user"`, HasCommentFromUser("user"))
	assert.Equal(t, `involves:"user"`, HasUserInvolved("user"))
	assert.Equal(t, `label:"label"`, WithLabel("label"))
	assert.Equal(t, `-label:"label"`, WithoutLabel("label"))
	assert.Equal(t, `milestone:"milestone"`, WithMilestone("milestone"))
	assert.Equal(t, `project:"project"`, WithProject("project"))
	assert.Equal(t, `head:"branch"`, WithMergeSourceBranch("branch"))
	assert.Equal(t, `base:"branch"`, WithMergeTargetBranch("branch"))
	assert.Equal(t, `language:"lang"`, WrittenInLanguage("lang"))
	assert.Equal(t, `comments:"5"`, WithNumberOfComments("5"))
	assert.Equal(t, `interactions:">5"`, WithNumberOfInteractions(">5"))
	assert.Equal(t, `reactions:"5..10"`, WithNumberOfReactions("5..10"))
	assert.Equal(t, `created:"2006-01-02"`, WasCreated("2006-01-02"))
	assert.Equal(t, `updated:">2006-01-02"`, WasUpdated(">2006-01-02"))
	assert.Equal(t, `closed:"2006-01-02..2006-02-03"`, WasClosed("2006-01-02..2006-02-03"))
	assert.Equal(t, `merged:"<=2006=01-02"`, WasMerged("<=2006=01-02"))
	assert.Equal(t, `"text"`, HasText("text"))
}
