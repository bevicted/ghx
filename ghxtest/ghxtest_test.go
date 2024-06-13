package ghxtest

import (
	"testing"

	"github.com/bevicted/ghx"
	"github.com/google/go-github/v62/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIssueFactory(t *testing.T) {
	t.Parallel()

	const length = 5

	baseIssue := github.Issue{
		Number: ghx.PTR(123),
		State:  ghx.StateOpen.StringP(),
		User: &github.User{
			Name: ghx.PTR("username"),
		},
	}
	f := &IssueFactory{
		Issue: baseIssue,
	}

	i1 := f.NewIssue(t)
	assert.Equal(t, baseIssue, *i1)
	i2 := f.NewIssue(t)
	assert.Equal(t, baseIssue, *i2)
	assert.NotSame(t, i1, i2)
	i1.User.Name = ghx.PTR("some other username")
	assert.NotEqual(t, *i1.User.Name, *i2.User.Name)

	i3 := f.NewIssueWithNumber(t, 321)
	assert.Equal(t, 321, *i3.Number)
	i3.Number = ghx.PTR(*baseIssue.Number)
	assert.Equal(t, baseIssue, *i3)

	issues := f.NewIssues(t, length)
	assert.Len(t, issues, length)
	for idx, i := range issues {
		assert.Equal(t, *baseIssue.Number+idx, *i.Number)
	}
}

func TestNewEmptyIssues(t *testing.T) {
	t.Parallel()

	const length = 5

	issues := NewEmptyIssues(t, length)
	require.Len(t, issues, length)
	for idx, issue := range issues {
		assert.Equal(t, idx, *issue.Number)
	}
}
