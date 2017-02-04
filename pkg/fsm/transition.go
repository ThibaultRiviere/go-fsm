package fsm

import (
	"errors"
)

// transition is to go from a state to an other
// It will check the current state to ensure that the transition is
// allowed execute the handler and then change the current state to
// a new state
type transition struct {
	// from represent the state needed for allowing the transition
	from string
	// to represente the state of the fsm after the transition
	to string
	// handler to call if the transition is allowed
	handler func()
}

// AddTransition will add to the fsm a new transition.
func (fsm *Fsm) AddTransition(name string, from string, to string, handler func()) error {
	if IsValideState(fsm.states, from) == false {
		return errors.New("State from is not part of the states")
	}

	if IsValideState(fsm.states, to) == false {
		return errors.New("State to is not part of the states")
	}

	transition := transition{from, to, handler}
	fsm.transitions[name] = transition
	return nil
}

// HandleTransition will check if the transition is possible.
// Check if the current state is the state require, call the handler and
// change the current state to the new state.
func (fsm *Fsm) HandleTransition(name string) (error, string) {
	transition, exist := fsm.transitions[name]
	if exist == false {
		return errors.New("Event doesn't exist"), fsm.current
	}
	if transition.from == fsm.current {
		transition.handler()
		fsm.current = transition.to
		return nil, fsm.current
	}
	return errors.New("Bad state refuse event"), fsm.current
}
