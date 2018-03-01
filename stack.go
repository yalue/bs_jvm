package bs_jvm

// This file contains functions and types related to JVM thread stacks.

import (
	"math"
)

// Holds everything necessary to restore a previous method's frame and return
// address.
type MethodFrame struct {
	Method         *Method
	ReturnIndex    uint
	ReferenceFrame int
	DataFrame      int
}

// An interface for a function call stack. A thread can keep this separate from
// its data stack.
type CallStack interface {
	// Used to push a return method and instruction index onto the call stack.
	Push(f MethodFrame) error
	// Used to pop a return method and instruction index from the stack.
	Pop() (MethodFrame, error)
}

// Implements the CallStack interface.
type basicCallStack struct {
	frames []MethodFrame
}

// Returns a new CallStack instance, that can hold up to the given number of
// return locations.
func NewCallStack(capacity uint32) CallStack {
	return &basicCallStack{
		frames: make([]MethodFrame, 0, capacity),
	}
}

func (s *basicCallStack) Push(f MethodFrame) error {
	if len(s.frames) >= cap(s.frames) {
		return StackOverflowError
	}
	s.frames = append(s.frames, f)
	return nil
}

func (s *basicCallStack) Pop() (MethodFrame, error) {
	if len(s.frames) == 0 {
		return MethodFrame{}, StackEmptyError
	}
	toReturn := s.frames[len(s.frames)-1]
	s.frames = s.frames[0 : len(s.frames)-1]
	return toReturn, nil
}

// An interface for a thread's stack of references. This can be separate from
// the data stack just for the sake of type checking.
type ReferenceStack interface {
	Push(r Reference) error
	Pop() (Reference, error)
	// Returns the current stack top indicator, which can be restored later.
	GetFrame() int
	// Sets the top of the stack, used to restore a method frame. This can
	// only be used to reduce the current stack contents, otherwise returns an
	// error.
	SetFrame(n int) error
}

// Implements the ReferenceStack interface.
type basicReferenceStack struct {
	references []Reference
}

func (s *basicReferenceStack) Push(r Reference) error {
	if len(s.references) >= cap(s.references) {
		return StackOverflowError
	}
	s.references = append(s.references, r)
	return nil
}

func (s *basicReferenceStack) Pop() (Reference, error) {
	if len(s.references) == 0 {
		return nil, StackEmptyError
	}
	toReturn := s.references[len(s.references)-1]
	s.references = s.references[0 : len(s.references)-1]
	return toReturn, nil
}

func (s *basicReferenceStack) GetFrame() int {
	return len(s.references)
}

func (s *basicReferenceStack) SetFrame(n int) error {
	if (n < 0) || (n > len(s.references)) {
		return BadFrameError(n)
	}
	s.references = s.references[0:n]
	return nil
}

// An interface for a thread's data stack. Returns an error if a stack overflow
// occurs, or if a stack is empty.
type DataStack interface {
	Push(v int32) error
	Pop() (int32, error)
	PushLong(v int64) error
	PopLong() (int64, error)
	PushFloat(v float32) error
	PopFloat() (float32, error)
	PushDouble(v float64) error
	PopDouble() (float64, error)
	// Returns the current stack top indicator, which can be restored later.
	GetFrame() int
	// Sets the top of the stack, used to restore a method frame. This can
	// only be used to reduce the current stack contents, otherwise returns an
	// error.
	SetFrame(n int) error
}

// Implements the stack interface.
type basicDataStack struct {
	data []int32
}

func (s *basicDataStack) GetFrame() int {
	return len(s.data)
}

func (s *basicDataStack) SetFrame(n int) error {
	if (n <= 0) || (n > len(s.data)) {
		return BadFrameError(n)
	}
	s.data = s.data[0:n]
	return nil
}

func (s *basicDataStack) Push(v int32) error {
	if len(s.data) >= cap(s.data) {
		return StackOverflowError
	}
	s.data = append(s.data, v)
	return nil
}

func (s *basicDataStack) Pop() (int32, error) {
	if len(s.data) < 1 {
		return 0, StackEmptyError
	}
	toReturn := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return toReturn, nil
}

func (s *basicDataStack) PushLong(v int64) error {
	if (len(s.data) + 1) >= cap(s.data) {
		return StackOverflowError
	}
	lowBits := int32(v)
	highBits := int32(v >> 32)
	s.data = append(s.data, lowBits, highBits)
	return nil
}

func (s *basicDataStack) PopLong() (int64, error) {
	if len(s.data) < 2 {
		return 0, StackEmptyError
	}
	highBits := s.data[len(s.data)-1]
	lowBits := s.data[len(s.data)-2]
	s.data = s.data[0 : len(s.data)-2]
	// Cast low bits to an unsigned value to avoid sign extension.
	return (int64(highBits) << 32) | int64(uint32(lowBits)), nil
}

func (s *basicDataStack) PushFloat(v float32) error {
	return s.Push(int32(math.Float32bits(v)))
}

func (s *basicDataStack) PopFloat() (float32, error) {
	bits, e := s.Pop()
	if e != nil {
		return 0, e
	}
	return math.Float32frombits(uint32(bits)), nil
}

func (s *basicDataStack) PushDouble(v float64) error {
	return s.PushLong(int64(math.Float64bits(v)))
}

func (s *basicDataStack) PopDouble() (float64, error) {
	bits, e := s.PopLong()
	if e != nil {
		return 0, e
	}
	return math.Float64frombits(uint64(bits)), nil
}

// Takes a capacity, in a number of 32-bit integers, and returns a new empty
// stack.
func NewDataStack(capacity uint32) DataStack {
	return &basicDataStack{
		data: make([]int32, 0, capacity),
	}
}
