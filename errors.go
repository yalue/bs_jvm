package jvm

// This file contains various error types used by various portions of the
// library. They're all located in a single file for faster reference.

import (
	"fmt"
)

// This error is returned if an unknown/unsupported opcode is encountered.
type UnknownInstructionError uint8

func (e UnknownInstructionError) Error() string {
	return fmt.Sprintf("Unknown/bad JVM opcode: 0x%02x", e)
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
