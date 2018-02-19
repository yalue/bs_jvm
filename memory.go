package jvm

// This file contains functions and types related to JVM memory management.

import (
	"fmt"
)

type InvalidAddressError uint

func (e InvalidAddressError) Error() string {
	return fmt.Sprintf("Invalid address: 0x%x", uint(e))
}

// A generic interface for accessing JVM memories. Returns an error if an
// address is invalid.
type JVMMemory interface {
	GetByte(address uint) (uint8, error)
	SetByte(value uint8, address uint) error
}

// Reads and returns a big-endian 16-bit integer at the given address.
func Read16Bits(m JVMMemory, address uint) (uint16, error) {
	high, e := m.GetByte(address)
	if e != nil {
		return 0, e
	}
	low, e := m.GetByte(address + 1)
	if e != nil {
		return 0, e
	}
	return (uint16(high) << 8) | uint16(low), nil
}

// Reads and returns a big-endian 32-bit integer at the given address.
func Read32Bits(m JVMMemory, address uint) (uint32, error) {
	high, e := Read16Bits(m, address)
	if e != nil {
		return 0, e
	}
	low, e := Read16Bits(m, address+2)
	if e != nil {
		return 0, e
	}
	return (uint32(high) << 16) | uint32(low), nil
}

type basicJVMMemory struct {
	memory []byte
}

func (m *basicJVMMemory) GetByte(address uint) (uint8, error) {
	if address >= uint(len(m.memory)) {
		return 0, InvalidAddressError(address)
	}
	return m.memory[address], nil
}

func (m *basicJVMMemory) SetByte(value uint8, address uint) error {
	if address >= uint(len(m.memory)) {
		return InvalidAddressError(address)
	}
	m.memory[address] = value
	return nil
}

// Returns a JVM memory-compatible wrapper around a byte slice, where the first
// byte is at address 0.
func JVMMemoryFromSlice(data []byte) JVMMemory {
	return &basicJVMMemory{
		memory: data,
	}
}

// Returns a basic implementation of the JVM memory struct, addressed starting
// at 0, and containing the given number of bytes.
func NewJVMMemory(size uint) JVMMemory {
	return &basicJVMMemory{
		memory: make([]byte, size),
	}
}
