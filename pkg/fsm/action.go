package fsm

import (
	"errors"
)

// action define an action on the state machine
// It will check the current state of the fsm and see if
// the action is allowed.
type action struct {
	// from represent the state needed for allowing the action
	from string
	// handler to call if the action is allowed
	handler func()
}

// AddAction will add to the fsm a new action.
func (fsm *Fsm) AddAction(name string, from string, handler func()) error {
	if IsValideState(fsm.states, from) == false {
		return errors.New("Default state is not part of the states")
	}
	action := action{from, handler}
	fsm.actions[name] = action
	return nil
}

// HandleAction will check if the current stat is the allowed state.
// If it is it will call the handler
func (fsm *Fsm) HandleAction(name string) (error, string) {
	action, exist := fsm.actions[name]
	if exist == false {
		return errors.New("Event doesn't exist"), fsm.current
	}
	if action.from == fsm.current {
		action.handler()
		return nil, fsm.current
	}
	return errors.New("Bad state refuse event"), fsm.current
}
