package fsm

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	states = []string{"a", "b", "c"}
)

func _testFsm(name string, states []string, current string, errExpected error) {
	Convey(name, func() {
		_, err := NewFsm(states, current)
		if err != nil {
			So(err.Error(), ShouldEqual, errExpected.Error())
		} else {
			So(err, ShouldEqual, errExpected)
		}
	})
}

func TestFsm(t *testing.T) {
	Convey("Testing initialization of the fsm", t, func() {
		var err error

		err = errors.New("Default state is not part of the states")
		_testFsm("with unexisting default", states, "d", err)

		err = nil
		_testFsm("with valid state", states, states[1], err)
	})
}

func _testGetState(name string, states []string, current string) {
	Convey(name, func() {
		fsm, err := NewFsm(states, current)
		So(err, ShouldEqual, nil)
		So(fsm.GetState(), ShouldEqual, current)
	})
}

func TestGetState(t *testing.T) {
	Convey("Testing GetState", t, func() {
		_testGetState("current as first states", states, "a")
		_testGetState("current as middle states", states, "b")
		_testGetState("current as end states", states, "c")
	})
}
