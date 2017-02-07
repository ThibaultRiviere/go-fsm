// Package fsm is a library tool for finite state nachine
package fsm

import (
	"errors"
)

// Fsm is a finitate state machine
// Fsm define the possible transitions, actions, the states possible and the current state
type Fsm struct {
	// current state of the fsm
	current string
	// states possible for the fsm
	states []string
	// transitions possible for the fsm
	transitions map[string]transition
	// action possible for the fsm
	actions map[string]action
}

// isValideState is used to know if a state is part of the fsm possible states
func isValideState(states []string, state string) bool {
	for _, val := range states {
		if val == state {
			return true
		}
	}
	return false
}

// New will initialize all the properties of the fsm.
// Check if the current state is part of the possible states
func New(states []string, current string) (*Fsm, error) {
	if isValideState(states, current) == false {
		return nil, errors.New("Default state is not part of the states")
	}
	transitions := make(map[string]transition)
	actions := make(map[string]action)
	return &Fsm{current, states, transitions, actions}, nil
}

// GetState return the current state of the fsm
func (Fsm *Fsm) GetState() string {
	return Fsm.current
}
