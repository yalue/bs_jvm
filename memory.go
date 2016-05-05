package jvm

// This file contains functions and types related to JVM memory management.

import (
	"fmt"
)

type InvalidAddressError uint

func (e InvalidAddressError) Error() string {
	return fmt.Sprintf("Invalid address: 0x%x", e)
}

// A generic interface for accessing JVM memories. Returns an error if an
// address is invalid.
type JVMMemory interface {
	GetByte(address uint) (uint8, error)
	SetByte(value uint8, address uint) error
}

type basicJVMMemory struct {
	memory []byte
}

func (m *basicJVMMemory) GetByte(address uint) (uint8, error) {
	if address > uint(len(m.memory)) {
		return 0, InvalidAddressError(address)
	}
	return m.memory[address], nil
}

func (m *basicJVMMemroy) SetByte(value uint8, address uint) error {
	if address > uint(len(m.memory)) {
		return 0, InvalidAddressError(address)
	}
	m.memory[address] = value
	return nil
}

// Returns a basic implementation of the JVM memory struct, addressed starting
// at 0, and containing the given number of bytes.
func NewJVMMemory(size uint) JVMMemory {
	return &basicJVMMemory{
		memory: make([]byte, size),
	}
}
