package fsm

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
	if isValideState(fsm.states, from) == false {
		return ErrUnknowState
	}
	action := action{from, handler}
	fsm.actions[name] = action
	return nil
}

// HandleAction will check if the current stat is the allowed state.
// If it is it will call the handler
func (fsm *Fsm) HandleAction(name string) (string, error) {
	action, exist := fsm.actions[name]
	if exist == false {
		return fsm.current, ErrUnknowAction
	}
	if action.from == fsm.current {
		action.handler()
		return fsm.current, nil
	}
	return fsm.current, ErrBadState
}
