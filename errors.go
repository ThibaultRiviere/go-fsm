package fsm

import "errors"

var (
	// ErrUnknowState is return when a state used is not part of the possible states for the fsm
	//
	// ## Example:
	//
	// fsm, err := fsm.New([]string{"a", "b"}, "b")
	// fsm.addAction("oups", "c")
	//
	// Here the state c is not part of the possible states. fsm.AddAction will return a ErrUnknowState
	ErrUnknowState = errors.New("state doesn't exist")

	// ErrUnknowAction is return when the fsm is trying to handle an action who has not be define
	//
	// ## example
	//
	// fsm, err := fsm.New([]string{"a", "b"}, "b")
	// fsm.HandleAction("Action is not define")
	//
	// Here the action have never be added. the function fsm.HandleAction will return an ErrUnknowAction
	ErrUnknowAction = errors.New("action unknow never be define")

	// ErrUnknowTransition is return when the fsm is trying to handle a transition who has not be define
	//
	// ## exanple
	//
	// fsm, err := fsm.New([]string{"a", "b"}, "b")
	// fsm.HandleTransition("Transition is not define")
	//
	// Here the transition have never be added. the function fsm.HandleTransition will return an ErrUnknowAction
	ErrUnknowTransition = errors.New("transition unknow never be define")

	// ErrBadState is return when the current state doesn't allow to handler to do his thing
	//
	// ## Example
	//
	// fsm, err := fsm.New([]string{"a", "b"}, "b")
	// fsm.AddAction("test", "a", func() {
	// 		fmt.Println("this is an action")
	// })
	// fsm.HandleAction("test")
	//
	// Here the action test is not allowed since the current state is b and not a.
	// The function will return an ErrBadState
	ErrBadState = errors.New("The current state doesn't allow this")
)
