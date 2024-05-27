package ghx

type (
	State       int
	StateReason int
)

const (
	StateOpen State = iota
	StateClosed
	StateAll
)

const (
	StateReasonCompleted StateReason = iota
	StateReasonNotPlanned
	StateReasonReopened
)

func (s State) String() string {
	switch s {
	case StateOpen:
		return "open"
	case StateClosed:
		return "closed"
	case StateAll:
		return "all"
	default:
		return ""
	}
}

func (sr StateReason) String() string {
	switch sr {
	case StateReasonCompleted:
		return "completed"
	case StateReasonNotPlanned:
		return "not_planned"
	case StateReasonReopened:
		return "reopened"
	default:
		return ""
	}
}

func (s State) StringP() *string {
	v := s.String()
	return &v
}

func (sr StateReason) StringP() *string {
	v := sr.String()
	return &v
}
