package fsm

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	degres = []string{"-100", "0", "100"}
)

func _testAddAction(name string, from string, errExpected error) {
	Convey(name, func() {
		water, err := New(degres, "0")
		So(err, ShouldEqual, nil)
		err = water.AddAction(name, from, func() {})
		if err == nil {
			So(err, ShouldEqual, errExpected)

		} else {
			So(err.Error(), ShouldEqual, errExpected.Error())
		}
	})
}

func TestAddAction(t *testing.T) {
	Convey("Testing addAction to fsm", t, func() {
		var err error

		err = errors.New("Default state is not part of the states")
		_testAddAction("with unexisting state", "42", err)

		err = nil
		_testAddAction("with existing state", "100", err)
	})
}

func _testHandleAction(name string, current string, action string, errExpected error) {
	Convey(name, func() {
		var err error
		water, err := New(degres, current)
		So(err, ShouldEqual, nil)
		water.AddAction("freeze", "-100", func() {})
		_, err = water.HandleAction(action)
		if err == nil {
			So(err, ShouldEqual, errExpected)
		} else {
			So(err.Error(), ShouldEqual, errExpected.Error())
		}
	})
}

func TestHandleAction(t *testing.T) {
	Convey("Testing HandleAction to fsm", t, func() {
		var err error

		err = errors.New("Event doesn't exist")
		_testHandleAction("with unexisting action", "0", "unexist", err)

		err = errors.New("Bad state refuse event")
		_testHandleAction("with action and bad state", "0", "freeze", err)

		err = nil
		_testHandleAction("with action and good state", "-100", "freeze", err)
	})
}
