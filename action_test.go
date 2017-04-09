package fsm

import (
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

func _testHandleConcurrentsAction(water *Fsm, action string, ret chan error) {
	_, err := water.HandleAction(action)
	ret <- err
}

func _testMutex(name string, concurrents int) {
	Convey(name, func() {
		isBoil := false
		water, err := New(degres, "0")
		So(err, ShouldEqual, nil)
		water.AddTransition("boil", "0", "100", func() { isBoil = true })
		water.AddAction("evaporate", "100", func() {
			if isBoil == false {
				panic("bad state something wrong here")
			}
		})
		rets := make(chan error, concurrents)
		for i := 0; i < concurrents; i++ {
			go _testHandleConcurrentsAction(water, "evaporate", rets)
		}
		go water.HandleTransition("boil")
		for i := 0; i < concurrents; i++ {
			err = <-rets
			if err != nil {
				So(err, ShouldEqual, ErrBadState)
			}
		}
	})
}

func TestAddAction(t *testing.T) {
	Convey("Testing addAction to fsm", t, func() {
		_testAddAction("with unexisting state", "42", ErrUnknowState)
		_testAddAction("with existing state", "100", nil)
	})
}

func TestHandleAction(t *testing.T) {
	Convey("Testing HandleAction to fsm", t, func() {
		_testHandleAction("with unexisting action", "0", "unexist", ErrUnknowAction)
		_testHandleAction("with action and bad state", "0", "freeze", ErrBadState)
		_testHandleAction("with action and good state", "-100", "freeze", nil)
	})
}

func TestMutexAction(t *testing.T) {
	Convey("Tesing Muxtex to fsm for action", t, func() {
		_testMutex("with 2 concurrences", 2)
		_testMutex("with 5 concurrences", 5)
		_testMutex("with 10 concurrences", 10)
		_testMutex("with 15 concurrences", 15)
		_testMutex("with 25 concurrences", 25)
	})
}
