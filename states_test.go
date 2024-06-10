package ghx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		state        State
		expectString string
	}{
		{
			state:        StateOpen,
			expectString: "open",
		},
		{
			state:        StateClosed,
			expectString: "closed",
		},
		{
			state:        StateAll,
			expectString: "all",
		},
		{
			state: State(-1),
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("%v == %s", tc.state, tc.expectString), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectString, *tc.state.StringP())
		})
	}
}

func TestStateReason(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		stateReason  StateReason
		expectString string
	}{
		{
			stateReason:  StateReasonCompleted,
			expectString: "completed",
		},
		{
			stateReason:  StateReasonNotPlanned,
			expectString: "not_planned",
		},
		{
			stateReason:  StateReasonReopened,
			expectString: "reopened",
		},
		{
			stateReason: StateReason(-1),
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("%v == %s", tc.stateReason, tc.expectString), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectString, *tc.stateReason.StringP())
		})
	}
}
