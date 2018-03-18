package bs_jvm

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewDataStack(3)
	e := s.Push(-1337)
	if e != nil {
		t.Logf("Failed pushing to empty stack: %s\n", e)
		t.FailNow()
	}
	value32, e := s.Pop()
	if e != nil {
		t.Logf("Failed popping from empty stack: %s\n", e)
		t.FailNow()
	}
	if value32 != -1337 {
		t.Logf("Popped %d from stack, wanted -1337.\n", value32)
		t.Fail()
	}
	_, e = s.Pop()
	if e == nil {
		t.Logf("Didn't get an error for popping empty stack.\n")
		t.FailNow()
	}
	if e != StackEmptyError {
		t.Logf("Didn't get stack empty error, but %s.\n", e)
		t.Fail()
	}
	e = s.PushFloat(1.7e25)
	if e != nil {
		t.Logf("Failed pushing float: %s\n", e)
		t.FailNow()
	}
	v := uint64(0xf000133780001337)
	e = s.PushLong(Long(v))
	if e != nil {
		t.Logf("Failed pushing long: %s\n", e)
		t.FailNow()
	}
	e = s.Push(1)
	if e == nil {
		t.Logf("Didn't fail pushing to full stack.\n")
		t.FailNow()
	}
	if e != StackOverflowError {
		t.Logf("Didn't get stack overflow error, but %s\n", e)
		t.Fail()
	}
	value64, e := s.PopLong()
	if e != nil {
		t.Logf("Failed popping long int: %s\n", e)
		t.FailNow()
	}
	if uint64(value64) != v {
		t.Logf("Didn't pop correct long int: %d\n", value64)
		t.Fail()
	}
	valueFloat, e := s.PopFloat()
	if valueFloat != 1.7e25 {
		t.Logf("Didn't pop correct float: %f\n", valueFloat)
		t.Fail()
	}
	e = s.PushDouble(0.13371337e-25)
	if e != nil {
		t.Logf("Failed pushing double: %s\n", e)
		t.FailNow()
	}
	valueDouble, e := s.PopDouble()
	if e != nil {
		t.Logf("Failed popping double: %s\n", e)
		t.FailNow()
	}
	if valueDouble != 0.13371337e-25 {
		t.Logf("Didn't pop correct double: %f\n", valueDouble)
		t.Fail()
	}
}
