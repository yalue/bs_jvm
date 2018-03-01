package bs_jvm

// This file contains functions and types related to JVM thread stacks.

import (
	"math"
)

// An interface for a function call stack. A thread can keep this separate from
// its data stack.
type CallStack interface {
	// Used to push a return method and instruction index onto the call stack.
	Push(method *Method, returnIndex uint) error
	// Used to pop a return method and instruction index from the stack.
	Pop() (*Method, uint, error)
}

// Implements the CallStack interface.
type basicCallStack struct {
	methods       []*Method
	returnIndices []uint
}

// Returns a new CallStack instance, that can hold up to the given number of
// return addresses.
func NewCallStack(capacity uint32) CallStack {
	return &basicCallStack{
		methods:       make([]*Method, 0, capacity),
		returnIndices: make([]uint, 0, capacity),
	}
}

func (s *basicCallStack) Push(method *Method, returnIndex uint) error {
	if len(s.methods) >= cap(s.methods) {
		return StackOverflowError
	}
	s.methods = append(s.methods, method)
	s.returnIndices = append(s.returnIndices, returnIndex)
	return nil
}

func (s *basicCallStack) Pop() (*Method, uint, error) {
	if len(s.methods) == 0 {
		return nil, 0, StackEmptyError
	}
	method := s.methods[len(s.methods)-1]
	returnIndex := s.returnIndices[len(s.returnIndices)-1]
	s.methods = s.methods[0 : len(s.methods)-1]
	s.returnIndices = s.returnIndices[0 : len(s.returnIndices)-1]
	return method, returnIndex, nil
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
}

// Implements the stack interface.
type basicDataStack struct {
	data []int32
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
