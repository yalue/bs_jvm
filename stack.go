package bs_jvm

// This file contains functions and types related to JVM thread stacks.

import (
	"math"
)

const (
	DefaultDataStackCapacity      = 4096 * 4
	DefaultReferenceStackCapacity = 4096
	DefaultCallStackCapacity      = 1024
)

// Holds the state of a thread's stack, to be used when calling or returning
// from methods.
type StackSizes struct {
	DataStackSize      int
	ReferenceStackSize int
	IsRefStackSize     int
}

// Holds everything necessary to restore a previous method's frame and return
// address.
type ReturnInfo struct {
	Method         *Method
	ReturnIndex    uint
	StackState     StackSizes
	LocalVariables []Object
}

// An interface for a function call stack. A thread can keep this separate from
// its data stack.
type CallStack interface {
	// Used to push a return method and instruction index onto the call stack.
	PushFrame(f ReturnInfo) error
	// Used to pop a return method and instruction index from the stack.
	PopFrame() (ReturnInfo, error)
}

// Implements the CallStack interface.
type basicCallStack struct {
	frames []ReturnInfo
}

// Returns a new CallStack instance, that can hold up to the given number of
// return locations.
func NewCallStack(capacity uint32) CallStack {
	return &basicCallStack{
		frames: make([]ReturnInfo, 0, capacity),
	}
}

func (s *basicCallStack) PushFrame(f ReturnInfo) error {
	if len(s.frames) >= cap(s.frames) {
		return StackOverflowError
	}
	s.frames = append(s.frames, f)
	return nil
}

func (s *basicCallStack) PopFrame() (ReturnInfo, error) {
	if len(s.frames) == 0 {
		return ReturnInfo{}, StackEmptyError
	}
	toReturn := s.frames[len(s.frames)-1]
	s.frames = s.frames[0 : len(s.frames)-1]
	return toReturn, nil
}

// An interface for a thread's stack of references. This can be separate from
// the data stack just for the sake of type checking.
type ReferenceStack interface {
	PushRef(r Object) error
	PopRef() (Object, error)
	// Returns the current stack top indicator, which can be restored later.
	// Returns the current stack size, which can be restored later.
	GetSize() int
	// Sets the size of the stack, used for discarding multiple values at once.
	SetSize(n int) error
}

// Implements the ReferenceStack interface.
type basicReferenceStack struct {
	references []Object
}

func (s *basicReferenceStack) PushRef(r Object) error {
	if len(s.references) >= cap(s.references) {
		return StackOverflowError
	}
	s.references = append(s.references, r)
	return nil
}

func (s *basicReferenceStack) PopRef() (Object, error) {
	if len(s.references) == 0 {
		return nil, StackEmptyError
	}
	toReturn := s.references[len(s.references)-1]
	s.references = s.references[0 : len(s.references)-1]
	return toReturn, nil
}

func (s *basicReferenceStack) GetSize() int {
	return len(s.references)
}

func (s *basicReferenceStack) SetSize(n int) error {
	if (n < 0) || (n > len(s.references)) {
		return BadStackSizeError(n)
	}
	s.references = s.references[0:n]
	return nil
}

func NewReferenceStack(capacity uint32) ReferenceStack {
	return &basicReferenceStack{
		references: make([]Object, 0, capacity),
	}
}

// An interface for a thread's data stack. Returns an error if a stack overflow
// occurs, or if a stack is empty.
type DataStack interface {
	Push(v Int) error
	Pop() (Int, error)
	PushLong(v Long) error
	PopLong() (Long, error)
	PushFloat(v Float) error
	PopFloat() (Float, error)
	PushDouble(v Double) error
	PopDouble() (Double, error)
	// Returns the current stack top indicator, which can be restored later.
	GetSize() int
	// Sets the top of the stack, used to restore a method frame. This can
	// only be used to reduce the current stack contents, otherwise returns an
	// error.
	SetSize(n int) error
}

// Implements the stack interface.
type basicDataStack struct {
	data []int32
}

func (s *basicDataStack) GetSize() int {
	return len(s.data)
}

func (s *basicDataStack) SetSize(n int) error {
	if (n <= 0) || (n > len(s.data)) {
		return BadStackSizeError(n)
	}
	s.data = s.data[0:n]
	return nil
}

func (s *basicDataStack) Push(v Int) error {
	if len(s.data) >= cap(s.data) {
		return StackOverflowError
	}
	s.data = append(s.data, int32(v))
	return nil
}

func (s *basicDataStack) Pop() (Int, error) {
	if len(s.data) < 1 {
		return 0, StackEmptyError
	}
	toReturn := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return Int(toReturn), nil
}

func (s *basicDataStack) PushLong(v Long) error {
	if (len(s.data) + 1) >= cap(s.data) {
		return StackOverflowError
	}
	lowBits := int32(v)
	highBits := int32(v >> 32)
	s.data = append(s.data, lowBits, highBits)
	return nil
}

func (s *basicDataStack) PopLong() (Long, error) {
	if len(s.data) < 2 {
		return 0, StackEmptyError
	}
	highBits := s.data[len(s.data)-1]
	lowBits := s.data[len(s.data)-2]
	s.data = s.data[0 : len(s.data)-2]
	// Cast low bits to an unsigned value to avoid sign extension.
	return (Long(highBits) << 32) | Long(uint32(lowBits)), nil
}

func (s *basicDataStack) PushFloat(v Float) error {
	return s.Push(Int(math.Float32bits(float32(v))))
}

func (s *basicDataStack) PopFloat() (Float, error) {
	bits, e := s.Pop()
	if e != nil {
		return 0, e
	}
	return Float(math.Float32frombits(uint32(bits))), nil
}

func (s *basicDataStack) PushDouble(v Double) error {
	return s.PushLong(Long(math.Float64bits(float64(v))))
}

func (s *basicDataStack) PopDouble() (Double, error) {
	bits, e := s.PopLong()
	if e != nil {
		return 0, e
	}
	return Double(math.Float64frombits(uint64(bits))), nil
}

// Takes a capacity, in a number of 32-bit integers, and returns a new empty
// stack.
func NewDataStack(capacity uint32) DataStack {
	return &basicDataStack{
		data: make([]int32, 0, capacity),
	}
}

// This interface is used for interacting with a thread's stack storage,
// including the call stack.
type ThreadStack interface {
	Push(v Int) error
	Pop() (Int, error)
	PushLong(v Long) error
	PopLong() (Long, error)
	PushFloat(v Float) error
	PopFloat() (Float, error)
	PushDouble(v Double) error
	PopDouble() (Double, error)
	PushRef(r Object) error
	PopRef() (Object, error)
	// Used to push a return method and instruction index onto the call stack.
	PushFrame(f ReturnInfo) error
	// Used to pop a return method and instruction index from the stack.
	PopFrame() (ReturnInfo, error)
	// Backs up the sizes of the stacks so that they can be restored later.
	GetSizes() StackSizes
	// Restores the stack positions contained in the given call frame. Used
	// when returning from a function.
	RestoreSizes(*StackSizes) error
}

// Tracks stack data for a thread. Implements the stack interface.
type basicStack struct {
	data  DataStack
	refs  ReferenceStack
	calls CallStack
	// Holds whether the previously-pushed value was a reference or a value.
	// This is needed for supporting the STUPID untyped pop instruction.
	IsRef []bool
}

func NewStack() ThreadStack {
	return &basicStack{
		data:  NewDataStack(DefaultDataStackCapacity),
		refs:  NewReferenceStack(DefaultReferenceStackCapacity),
		calls: NewCallStack(DefaultCallStackCapacity),
		IsRef: make([]bool, 0, DefaultDataStackCapacity+
			DefaultReferenceStackCapacity),
	}
}

func (s *basicStack) Push(v Int) error {
	e := s.data.Push(v)
	if e == nil {
		s.IsRef = append(s.IsRef, false)
	}
	return e
}

func (s *basicStack) Pop() (Int, error) {
	v, e := s.data.Pop()
	if e == nil {
		s.IsRef = s.IsRef[0 : len(s.IsRef)-1]
	}
	return v, e
}

func (s *basicStack) PushLong(v Long) error {
	e := s.data.PushLong(v)
	if e == nil {
		s.IsRef = append(s.IsRef, false)
	}
	return e
}

func (s *basicStack) PopLong() (Long, error) {
	v, e := s.data.PopLong()
	if e == nil {
		s.IsRef = s.IsRef[0 : len(s.IsRef)-1]
	}
	return v, e
}

func (s *basicStack) PushFloat(v Float) error {
	e := s.data.PushFloat(v)
	if e == nil {
		s.IsRef = append(s.IsRef, false)
	}
	return e
}

func (s *basicStack) PopFloat() (Float, error) {
	v, e := s.data.PopFloat()
	if e == nil {
		s.IsRef = s.IsRef[0 : len(s.IsRef)-1]
	}
	return v, e
}

func (s *basicStack) PushDouble(v Double) error {
	e := s.data.PushDouble(v)
	if e == nil {
		s.IsRef = append(s.IsRef, false)
	}
	return e
}

func (s *basicStack) PopDouble() (Double, error) {
	v, e := s.data.PopDouble()
	if e == nil {
		s.IsRef = s.IsRef[0 : len(s.IsRef)-1]
	}
	return v, e
}

func (s *basicStack) PushRef(r Object) error {
	e := s.refs.PushRef(r)
	if e == nil {
		s.IsRef = append(s.IsRef, true)
	}
	return e
}

func (s *basicStack) PopRef() (Object, error) {
	r, e := s.refs.PopRef()
	if e == nil {
		s.IsRef = s.IsRef[0 : len(s.IsRef)-1]
	}
	return r, e
}

// Pops and returns an object reference from the stack. Returns an error if the
// reference is nil.
func PopRefNotNull(s ThreadStack) (Object, error) {
	o, e := s.PopRef()
	if e != nil {
		return nil, e
	}
	if o == nil {
		return nil, NullReferenceError("Expected to pop a non-null reference")
	}
	return o, nil
}

func (s *basicStack) PushFrame(f ReturnInfo) error {
	return s.calls.PushFrame(f)
}

func (s *basicStack) PopFrame() (ReturnInfo, error) {
	return s.calls.PopFrame()
}

func (s *basicStack) GetSizes() StackSizes {
	return StackSizes{
		DataStackSize:      s.data.GetSize(),
		ReferenceStackSize: s.refs.GetSize(),
		IsRefStackSize:     len(s.IsRef),
	}
}

func (s *basicStack) RestoreSizes(sizes *StackSizes) error {
	e := s.data.SetSize(sizes.DataStackSize)
	if e != nil {
		return e
	}
	e = s.refs.SetSize(sizes.ReferenceStackSize)
	if e != nil {
		return e
	}
	if sizes.IsRefStackSize > len(s.IsRef) {
		return BadStackSizeError(sizes.IsRefStackSize)
	}
	s.IsRef = s.IsRef[0:sizes.IsRefStackSize]
	return nil
}
