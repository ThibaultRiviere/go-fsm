package fsm

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	positions = []string{"locked", "unlocked"}
)

func _testAddTransition(name string, from string, to string, errExpected error) {
	Convey(name, func() {
		turnstile, err := New(positions, "locked")
		So(err, ShouldEqual, nil)
		err = turnstile.AddTransition("new", from, to, func() {})
		if err == nil {
			So(err, ShouldEqual, errExpected)
		} else {
			So(err.Error(), ShouldEqual, errExpected.Error())
		}
	})
}

func _testHandleTransition(name string, current string, trans string, errExpected error) {
	Convey(name, func() {
		turnstile, err := New(positions, current)
		So(err, ShouldEqual, nil)
		err = turnstile.AddTransition("unlock", "locked", "unlocked", func() {})
		So(err, ShouldEqual, nil)
		state, err2 := turnstile.HandleTransition(trans)
		if err2 == nil {
			So(err2, ShouldEqual, errExpected)
			So(state, ShouldNotEqual, current)
		} else {
			So(err2.Error(), ShouldEqual, errExpected.Error())
			So(current, ShouldEqual, turnstile.GetState())
		}
	})
}

func TestAddTransition(t *testing.T) {
	Convey("Testing AddTransition to fsm", t, func() {
		_testAddTransition("With unexisting from", "unexist", "locked", ErrUnknowState)
		_testAddTransition("With unexisting to", "locked", "unexist", ErrUnknowState)
		_testAddTransition("With good values", "locked", "unlocked", nil)
	})
}

func TestHandleTransition(t *testing.T) {
	Convey("Testing HandleTransition", t, func() {
		_testHandleTransition("With unexisting event", "locked", "unexist", ErrUnknowTransition)
		_testHandleTransition("with bad state", "unlocked", "unlock", ErrBadState)
		_testHandleTransition("with good state", "locked", "unlock", nil)
	})
}
