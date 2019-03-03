package bs_jvm

// This file contains various error types used by various portions of the
// library. They're all located in a single file for faster reference.

import (
	"fmt"
)

// This error is returned if an unknown/unsupported opcode is encountered.
type UnknownInstructionError uint8

func (e UnknownInstructionError) Error() string {
	return fmt.Sprintf("Unknown/bad JVM opcode: 0x%02x", uint8(e))
}

// This error is returned when a feature is invoked that has not yet been
// implemented in the JVM.
var NotImplementedError = fmt.Errorf("Support not implemented")

// This error is returned when a memory reference fails due to an address
// being out of range or otherwise invalid.
type InvalidAddressError uint

func (e InvalidAddressError) Error() string {
	return fmt.Sprintf("Invalid address: 0x%x", uint(e))
}

// This is returned if a stack grows beyond its capacity.
var StackOverflowError = fmt.Errorf("Stack overflow")

// This is returned if an attempt to read from an empty stack occurs.
var StackEmptyError = fmt.Errorf("Stack empty")

// This is returned if a class lookup fails. Contains the class name that was
// requested but not known.
type ClassNotFoundError string

func (e ClassNotFoundError) Error() string {
	return fmt.Sprintf("Class not found: %s", string(e))
}

// This is returned if a method lookup fails. Consists of the name of the
// requested method.
type MethodNotFoundError string

func (e MethodNotFoundError) Error() string {
	return fmt.Sprintf("Method not found: %s", string(e))
}

// This will be returned when a thread exits, either explicity or by allowing
// its initial method to return. It should not usually indicate a problem.
var ThreadExitedError = fmt.Errorf("Thread exited")

// This is returned when attempting to set an invalid stack size.
type BadStackSizeError int

func (e BadStackSizeError) Error() string {
	return fmt.Sprintf("Attempted to set a bad stack size: %d", int(e))
}

// This is returned if an attempt to operate on invalid data is detected during
// instruction execution or an optimization pass.
type TypeError string

func (e TypeError) Error() string {
	return fmt.Sprintf("Type error: %s", string(e))
}

// This is returned when attempting to access an invalid local variable.
type BadLocalVariableError int

func (e BadLocalVariableError) Error() string {
	return fmt.Sprintf("Attempted to access an invalid local variable at "+
		"index %d", int(e))
}

type IndexOutOfBoundsError uint64

func (e IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("Index out of bounds: %d", int(e))
}

// This is returned if an object reference is expected, but nil is found.
type NullReferenceError string

func (e NullReferenceError) Error() string {
	return fmt.Sprintf("Null reference error: %s", string(e))
}


// This is usually returned if an instruction attempts to do something that
// requires dividing by zero.
type ArithmeticError string

func (e ArithmeticError) Error() string {
	return fmt.Sprintf("Arithmetic error: %s", string(e))
}
