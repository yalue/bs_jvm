package jvm

// This file contains functions for disassembling JVM instructions.

import (
	"fmt"
)

// This error is returned if an unknown/unsupported opcode is encountered.
type UnknownInstructionError uint8

func (e UnknownInstructionError) Error() string {
	return fmt.Sprintf("Unknown/bad JVM opcode: 0x%02x", e)
}

var NotImplementedError = fmt.Errorf("Support not implemented")

// The interface through which JVM opcodes can be inspected or executed.
type JVMInstruction interface {
	// Returns the 8-bit opcode for the instruction
	Raw() uint8
	// Returns additional bytes following the instruction's 8-bit opcode, or
	// nil if the instruction doesn't have such bytes. May be slow for some
	// opcodes.
	OtherBytes() []byte
	// Runs the instruction in the given thread
	Execute(t JVMThread) error
	// Returns the length of the instruction, including the opcode and
	// additional argument bytes.
	Length() uint
	// Returns the disassembly string of the instruction
	String() string
}

// Provices a default implementation of the JVMInstruction interface.
type unknownJVMInstruction struct {
	raw uint8
}

func (n *unknownJVMInstruction) Raw() uint8 {
	return n.raw
}

func (n *unknownJVMInstruction) OtherBytes() []byte {
	return nil
}

func (n *unknownJVMInstruction) Length() uint {
	return 1
}

func (n *unknownJVMInstruction) Execute(t JVMThread) error {
	return UnknownInstructionError(n.raw)
}

func (n *unknownJVMInstruction) String() string {
	return fmt.Sprintf("<unknown instruction 0x%02x>", n.raw)
}

// Like unknownJVMInstruction, but contains an instruction string. Used for
// known instructions which only consist of one byte.
type knownJVMInstruction struct {
	raw  uint8
	name string
}

func (n *knownJVMInstruction) Raw() uint8 {
	return n.raw
}

func (n *knownJVMInstruction) OtherBytes() []byte {
	return nil
}

func (n *knownJVMInstruction) Length() uint {
	return 1
}

func (n *knownJVMInstruction) Execute(t JVMThread) error {
	return fmt.Errorf("Execution not implemented for %s", n.String())
}

func (n *knownJVMInstruction) String() string {
	return n.name
}

// Returns the instruction starting at the given address. Only returns an error
// if the address is invalid. If an invalid/unknown instruction is located at
// the address, then a JVMInstruction will still be returned, but it will
// produce an UnknownInstructionError if executed.
// TODO: Should this behavior be changed? Would it be better to just crash on
// reading?
func GetNextInstruction(m JVMMemory, address uint) (JVMInstruction, error) {
	firstByte, e := m.GetByte(address)
	if e != nil {
		return nil, e
	}
	opcodeInfo := opcodeTable[firstByte]
	// Unknown instruction.
	if opcodeInfo == nil {
		toReturn := &unknownJVMInstruction{
			raw: firstByte,
		}
		return toReturn, nil
	}
	toReturn, e := opcodeInfo.parse(opcodeInfo.opcode, opcodeInfo.name,
		address, m)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing instruction: %s", e)
	}
	return toReturn, nil
}

type nopInstruction struct{ knownJVMInstruction }

func parseNopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := nopInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x00,
			name: name,
		},
	}
	return &toReturn, nil
}

type aconst_nullInstruction struct{ knownJVMInstruction }

func parseAconst_nullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aconst_nullInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x01,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_m1Instruction struct{ knownJVMInstruction }

func parseIconst_m1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_m1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x02,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_0Instruction struct{ knownJVMInstruction }

func parseIconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x03,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_1Instruction struct{ knownJVMInstruction }

func parseIconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x04,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_2Instruction struct{ knownJVMInstruction }

func parseIconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x05,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_3Instruction struct{ knownJVMInstruction }

func parseIconst_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x06,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_4Instruction struct{ knownJVMInstruction }

func parseIconst_4Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_4Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x07,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_5Instruction struct{ knownJVMInstruction }

func parseIconst_5Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_5Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x08,
			name: name,
		},
	}
	return &toReturn, nil
}

type lconst_0Instruction struct{ knownJVMInstruction }

func parseLconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x09,
			name: name,
		},
	}
	return &toReturn, nil
}

type lconst_1Instruction struct{ knownJVMInstruction }

func parseLconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0a,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_0Instruction struct{ knownJVMInstruction }

func parseFconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0b,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_1Instruction struct{ knownJVMInstruction }

func parseFconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0c,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_2Instruction struct{ knownJVMInstruction }

func parseFconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0d,
			name: name,
		},
	}
	return &toReturn, nil
}

type dconst_0Instruction struct{ knownJVMInstruction }

func parseDconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0e,
			name: name,
		},
	}
	return &toReturn, nil
}

type dconst_1Instruction struct{ knownJVMInstruction }

func parseDconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x0f,
			name: name,
		},
	}
	return &toReturn, nil
}

// This covers instructions such as bipush, which have one argument byte past
// the opcode byte.
type singleByteArgumentInstruction struct {
	raw   uint8
	name  string
	value uint8
}

func parseSingleByteArgumentInstruction(opcode uint8, name string,
	address uint, m JVMMemory) (*singleByteArgumentInstruction, error) {
	value, e := m.GetByte(address + 1)
	if e != nil {
		return nil, fmt.Errorf("Failed reading argument byte for %s: %s", name,
			e)
	}
	toReturn := singleByteArgumentInstruction{
		raw: opcode,
		name:  name,
		value: value,
	}
	return &toReturn, nil
}

func (n *singleByteArgumentInstruction) Raw() uint8 {
	return n.raw
}

func (n *singleByteArgumentInstruction) OtherBytes() []byte {
	return []byte{n.value}
}

func (n *singleByteArgumentInstruction) Length() uint {
	return 2
}

func (n *singleByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%02x", n.name, n.value)
}

type bipushInstruction struct{ singleByteArgumentInstruction }

func parseBipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &bipushInstruction{*toReturn}, nil
}

// Any instruction which includes a two-byte value beyond its initial opcode.
type twoByteArgumentInstruction struct {
	raw   uint8
	name  string
	value uint16
}

func (n *twoByteArgumentInstruction) Raw() uint8 {
	return n.raw
}

func (n *twoByteArgumentInstruction) OtherBytes() []byte {
	high := uint8(n.value >> 8)
	low := uint8(n.value & 0xff)
	return []byte{high, low}
}

func (n *twoByteArgumentInstruction) Length() uint {
	return 3
}

func (n *twoByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%04x", n.name, n.value)
}

func parseTwoByteArgumentInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (*twoByteArgumentInstruction, error) {
	value, e := Read16Bits(m, address+1)
	if e != nil {
		return nil, fmt.Errorf("Failed reading argument value for %s: %s",
			name, e)
	}
	toReturn := twoByteArgumentInstruction{
		raw:   opcode,
		name:  name,
		value: value,
	}
	return &toReturn, nil
}

type sipushInstruction struct{ twoByteArgumentInstruction }

func parseSipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &sipushInstruction{*toReturn}, nil
}

type ldcInstruction struct{ singleByteArgumentInstruction }

func parseLdcInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldcInstruction{*toReturn}, nil
}

type ldc_wInstruction struct{ twoByteArgumentInstruction }

func parseLdc_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldc_wInstruction{*toReturn}, nil
}

type ldc2_wInstruction struct{ twoByteArgumentInstruction }

func parseLdc2_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldc2_wInstruction{*toReturn}, nil
}

type iloadInstruction struct{ singleByteArgumentInstruction }

func parseIloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &iloadInstruction{*toReturn}, nil
}

type lloadInstruction struct{ singleByteArgumentInstruction }

func parseLloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &lloadInstruction{*toReturn}, nil
}

type floadInstruction struct{ singleByteArgumentInstruction }

func parseFloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &floadInstruction{*toReturn}, nil
}

type dloadInstruction struct{ singleByteArgumentInstruction }

func parseDloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &dloadInstruction{*toReturn}, nil
}

type aloadInstruction struct{ singleByteArgumentInstruction }

func parseAloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &aloadInstruction{*toReturn}, nil
}

type iload_0Instruction struct{ knownJVMInstruction }

func parseIload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1a,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_1Instruction struct{ knownJVMInstruction }

func parseIload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1b,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_2Instruction struct{ knownJVMInstruction }

func parseIload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1c,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_3Instruction struct{ knownJVMInstruction }

func parseIload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1d,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_0Instruction struct{ knownJVMInstruction }

func parseLload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1e,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_1Instruction struct{ knownJVMInstruction }

func parseLload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x1f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_2Instruction struct{ knownJVMInstruction }

func parseLload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x20,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_3Instruction struct{ knownJVMInstruction }

func parseLload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x21,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_0Instruction struct{ knownJVMInstruction }

func parseFload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x22,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_1Instruction struct{ knownJVMInstruction }

func parseFload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x23,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_2Instruction struct{ knownJVMInstruction }

func parseFload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x24,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_3Instruction struct{ knownJVMInstruction }

func parseFload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x25,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_0Instruction struct{ knownJVMInstruction }

func parseDload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x26,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_1Instruction struct{ knownJVMInstruction }

func parseDload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x27,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_2Instruction struct{ knownJVMInstruction }

func parseDload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x28,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_3Instruction struct{ knownJVMInstruction }

func parseDload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x29,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_0Instruction struct{ knownJVMInstruction }

func parseAload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2a,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_1Instruction struct{ knownJVMInstruction }

func parseAload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2b,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_2Instruction struct{ knownJVMInstruction }

func parseAload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2c,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_3Instruction struct{ knownJVMInstruction }

func parseAload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2d,
			name: name,
		},
	}
	return &toReturn, nil
}

type ialoadInstruction struct{ knownJVMInstruction }

func parseIaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ialoadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2e,
			name: name,
		},
	}
	return &toReturn, nil
}

type laloadInstruction struct{ knownJVMInstruction }

func parseLaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x2f,
			name: name,
		},
	}
	return &toReturn, nil
}

type faloadInstruction struct{ knownJVMInstruction }

func parseFaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x30,
			name: name,
		},
	}
	return &toReturn, nil
}

type daloadInstruction struct{ knownJVMInstruction }

func parseDaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x31,
			name: name,
		},
	}
	return &toReturn, nil
}

type aaloadInstruction struct{ knownJVMInstruction }

func parseAaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aaloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x32,
			name: name,
		},
	}
	return &toReturn, nil
}

type baloadInstruction struct{ knownJVMInstruction }

func parseBaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := baloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x33,
			name: name,
		},
	}
	return &toReturn, nil
}

type caloadInstruction struct{ knownJVMInstruction }

func parseCaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := caloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x34,
			name: name,
		},
	}
	return &toReturn, nil
}

type saloadInstruction struct{ knownJVMInstruction }

func parseSaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := saloadInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x35,
			name: name,
		},
	}
	return &toReturn, nil
}

type istoreInstruction struct{ singleByteArgumentInstruction }

func parseIstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &istoreInstruction{*toReturn}, nil
}

type lstoreInstruction struct{ singleByteArgumentInstruction }

func parseLstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &lstoreInstruction{*toReturn}, nil
}

type fstoreInstruction struct{ singleByteArgumentInstruction }

func parseFstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &fstoreInstruction{*toReturn}, nil
}

type dstoreInstruction struct{ singleByteArgumentInstruction }

func parseDstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &dstoreInstruction{*toReturn}, nil
}

type astoreInstruction struct{ singleByteArgumentInstruction }

func parseAstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &astoreInstruction{*toReturn}, nil
}

type istore_0Instruction struct{ knownJVMInstruction }

func parseIstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x3b,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_1Instruction struct{ knownJVMInstruction }

func parseIstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x3c,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_2Instruction struct{ knownJVMInstruction }

func parseIstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x3d,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_3Instruction struct{ knownJVMInstruction }

func parseIstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x3e,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_0Instruction struct{ knownJVMInstruction }

func parseLstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x3f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_1Instruction struct{ knownJVMInstruction }

func parseLstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x40,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_2Instruction struct{ knownJVMInstruction }

func parseLstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x41,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_3Instruction struct{ knownJVMInstruction }

func parseLstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x42,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_0Instruction struct{ knownJVMInstruction }

func parseFstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x43,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_1Instruction struct{ knownJVMInstruction }

func parseFstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x44,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_2Instruction struct{ knownJVMInstruction }

func parseFstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x45,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_3Instruction struct{ knownJVMInstruction }

func parseFstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x46,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_0Instruction struct{ knownJVMInstruction }

func parseDstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x47,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_1Instruction struct{ knownJVMInstruction }

func parseDstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x48,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_2Instruction struct{ knownJVMInstruction }

func parseDstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x49,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_3Instruction struct{ knownJVMInstruction }

func parseDstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4a,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_0Instruction struct{ knownJVMInstruction }

func parseAstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_0Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4b,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_1Instruction struct{ knownJVMInstruction }

func parseAstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4c,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_2Instruction struct{ knownJVMInstruction }

func parseAstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4d,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_3Instruction struct{ knownJVMInstruction }

func parseAstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_3Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4e,
			name: name,
		},
	}
	return &toReturn, nil
}

type iastoreInstruction struct{ knownJVMInstruction }

func parseIastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x4f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lastoreInstruction struct{ knownJVMInstruction }

func parseLastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x50,
			name: name,
		},
	}
	return &toReturn, nil
}

type fastoreInstruction struct{ knownJVMInstruction }

func parseFastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x51,
			name: name,
		},
	}
	return &toReturn, nil
}

type dastoreInstruction struct{ knownJVMInstruction }

func parseDastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x52,
			name: name,
		},
	}
	return &toReturn, nil
}

type aastoreInstruction struct{ knownJVMInstruction }

func parseAastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x53,
			name: name,
		},
	}
	return &toReturn, nil
}

type bastoreInstruction struct{ knownJVMInstruction }

func parseBastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := bastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x54,
			name: name,
		},
	}
	return &toReturn, nil
}

type castoreInstruction struct{ knownJVMInstruction }

func parseCastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := castoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x55,
			name: name,
		},
	}
	return &toReturn, nil
}

type sastoreInstruction struct{ knownJVMInstruction }

func parseSastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := sastoreInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x56,
			name: name,
		},
	}
	return &toReturn, nil
}

type popInstruction struct{ knownJVMInstruction }

func parsePopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := popInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x57,
			name: name,
		},
	}
	return &toReturn, nil
}

type pop2Instruction struct{ knownJVMInstruction }

func parsePop2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := pop2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x58,
			name: name,
		},
	}
	return &toReturn, nil
}

type dupInstruction struct{ knownJVMInstruction }

func parseDupInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dupInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x59,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup_x1Instruction struct{ knownJVMInstruction }

func parseDup_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5a,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup_x2Instruction struct{ knownJVMInstruction }

func parseDup_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5b,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2Instruction struct{ knownJVMInstruction }

func parseDup2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5c,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2_x1Instruction struct{ knownJVMInstruction }

func parseDup2_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5d,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2_x2Instruction struct{ knownJVMInstruction }

func parseDup2_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5e,
			name: name,
		},
	}
	return &toReturn, nil
}

type swapInstruction struct{ knownJVMInstruction }

func parseSwapInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := swapInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x5f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iaddInstruction struct{ knownJVMInstruction }

func parseIaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iaddInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x60,
			name: name,
		},
	}
	return &toReturn, nil
}

type laddInstruction struct{ knownJVMInstruction }

func parseLaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laddInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x61,
			name: name,
		},
	}
	return &toReturn, nil
}

type faddInstruction struct{ knownJVMInstruction }

func parseFaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faddInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x62,
			name: name,
		},
	}
	return &toReturn, nil
}

type daddInstruction struct{ knownJVMInstruction }

func parseDaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daddInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x63,
			name: name,
		},
	}
	return &toReturn, nil
}

type isubInstruction struct{ knownJVMInstruction }

func parseIsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := isubInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x64,
			name: name,
		},
	}
	return &toReturn, nil
}

type lsubInstruction struct{ knownJVMInstruction }

func parseLsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lsubInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x65,
			name: name,
		},
	}
	return &toReturn, nil
}

type fsubInstruction struct{ knownJVMInstruction }

func parseFsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fsubInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x66,
			name: name,
		},
	}
	return &toReturn, nil
}

type dsubInstruction struct{ knownJVMInstruction }

func parseDsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dsubInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x67,
			name: name,
		},
	}
	return &toReturn, nil
}

type imulInstruction struct{ knownJVMInstruction }

func parseImulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := imulInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x68,
			name: name,
		},
	}
	return &toReturn, nil
}

type lmulInstruction struct{ knownJVMInstruction }

func parseLmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lmulInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x69,
			name: name,
		},
	}
	return &toReturn, nil
}

type fmulInstruction struct{ knownJVMInstruction }

func parseFmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fmulInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6a,
			name: name,
		},
	}
	return &toReturn, nil
}

type dmulInstruction struct{ knownJVMInstruction }

func parseDmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dmulInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6b,
			name: name,
		},
	}
	return &toReturn, nil
}

type idivInstruction struct{ knownJVMInstruction }

func parseIdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := idivInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6c,
			name: name,
		},
	}
	return &toReturn, nil
}

type ldivInstruction struct{ knownJVMInstruction }

func parseLdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ldivInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6d,
			name: name,
		},
	}
	return &toReturn, nil
}

type fdivInstruction struct{ knownJVMInstruction }

func parseFdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fdivInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6e,
			name: name,
		},
	}
	return &toReturn, nil
}

type ddivInstruction struct{ knownJVMInstruction }

func parseDdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ddivInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x6f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iremInstruction struct{ knownJVMInstruction }

func parseIremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iremInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x70,
			name: name,
		},
	}
	return &toReturn, nil
}

type lremInstruction struct{ knownJVMInstruction }

func parseLremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lremInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x71,
			name: name,
		},
	}
	return &toReturn, nil
}

type fremInstruction struct{ knownJVMInstruction }

func parseFremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fremInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x72,
			name: name,
		},
	}
	return &toReturn, nil
}

type dremInstruction struct{ knownJVMInstruction }

func parseDremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dremInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x73,
			name: name,
		},
	}
	return &toReturn, nil
}

type inegInstruction struct{ knownJVMInstruction }

func parseInegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := inegInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x74,
			name: name,
		},
	}
	return &toReturn, nil
}

type lnegInstruction struct{ knownJVMInstruction }

func parseLnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lnegInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x75,
			name: name,
		},
	}
	return &toReturn, nil
}

type fnegInstruction struct{ knownJVMInstruction }

func parseFnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fnegInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x76,
			name: name,
		},
	}
	return &toReturn, nil
}

type dnegInstruction struct{ knownJVMInstruction }

func parseDnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dnegInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x77,
			name: name,
		},
	}
	return &toReturn, nil
}

type ishlInstruction struct{ knownJVMInstruction }

func parseIshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishlInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x78,
			name: name,
		},
	}
	return &toReturn, nil
}

type lshlInstruction struct{ knownJVMInstruction }

func parseLshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshlInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x79,
			name: name,
		},
	}
	return &toReturn, nil
}

type ishrInstruction struct{ knownJVMInstruction }

func parseIshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishrInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7a,
			name: name,
		},
	}
	return &toReturn, nil
}

type lshrInstruction struct{ knownJVMInstruction }

func parseLshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshrInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7b,
			name: name,
		},
	}
	return &toReturn, nil
}

type iushrInstruction struct{ knownJVMInstruction }

func parseIushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iushrInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7c,
			name: name,
		},
	}
	return &toReturn, nil
}

type lushrInstruction struct{ knownJVMInstruction }

func parseLushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lushrInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7d,
			name: name,
		},
	}
	return &toReturn, nil
}

type iandInstruction struct{ knownJVMInstruction }

func parseIandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iandInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7e,
			name: name,
		},
	}
	return &toReturn, nil
}

type landInstruction struct{ knownJVMInstruction }

func parseLandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := landInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x7f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iorInstruction struct{ knownJVMInstruction }

func parseIorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iorInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x80,
			name: name,
		},
	}
	return &toReturn, nil
}

type lorInstruction struct{ knownJVMInstruction }

func parseLorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lorInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x81,
			name: name,
		},
	}
	return &toReturn, nil
}

type ixorInstruction struct{ knownJVMInstruction }

func parseIxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ixorInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x82,
			name: name,
		},
	}
	return &toReturn, nil
}

type lxorInstruction struct{ knownJVMInstruction }

func parseLxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lxorInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x83,
			name: name,
		},
	}
	return &toReturn, nil
}

// The iinc instruction is a fairly unique format, so it gets its own struct
type iincInstruction struct {
	offset uint8
	value  uint8
}

func (n *iincInstruction) Raw() uint8 {
	return 0x84
}

func (n *iincInstruction) OtherBytes() []byte {
	return []byte{n.offset, n.value}
}

func (n *iincInstruction) Length() uint {
	return 3
}

func (n *iincInstruction) String() string {
	return fmt.Sprintf("iinc 0x%02x 0x%02x", n.offset, n.value)
}

func parseIincInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	offset, e := m.GetByte(address + 1)
	if e != nil {
		return nil, fmt.Errorf("Failed getting iinc offset: %s", e)
	}
	value, e := m.GetByte(address + 2)
	if e != nil {
		return nil, fmt.Errorf("Failed getting iinc value: %s", e)
	}
	toReturn := iincInstruction{
		offset: offset,
		value:  value,
	}
	return &toReturn, nil
}

type i2lInstruction struct{ knownJVMInstruction }

func parseI2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2lInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x85,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2fInstruction struct{ knownJVMInstruction }

func parseI2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2fInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x86,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2dInstruction struct{ knownJVMInstruction }

func parseI2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2dInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x87,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2iInstruction struct{ knownJVMInstruction }

func parseL2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2iInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x88,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2fInstruction struct{ knownJVMInstruction }

func parseL2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2fInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x89,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2dInstruction struct{ knownJVMInstruction }

func parseL2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2dInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8a,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2iInstruction struct{ knownJVMInstruction }

func parseF2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2iInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8b,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2lInstruction struct{ knownJVMInstruction }

func parseF2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2lInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8c,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2dInstruction struct{ knownJVMInstruction }

func parseF2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2dInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8d,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2iInstruction struct{ knownJVMInstruction }

func parseD2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2iInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8e,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2lInstruction struct{ knownJVMInstruction }

func parseD2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2lInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x8f,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2fInstruction struct{ knownJVMInstruction }

func parseD2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2fInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x90,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2bInstruction struct{ knownJVMInstruction }

func parseI2bInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2bInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x91,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2cInstruction struct{ knownJVMInstruction }

func parseI2cInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2cInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x92,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2sInstruction struct{ knownJVMInstruction }

func parseI2sInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2sInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x93,
			name: name,
		},
	}
	return &toReturn, nil
}

type lcmpInstruction struct{ knownJVMInstruction }

func parseLcmpInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lcmpInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x94,
			name: name,
		},
	}
	return &toReturn, nil
}

type fcmplInstruction struct{ knownJVMInstruction }

func parseFcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmplInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x95,
			name: name,
		},
	}
	return &toReturn, nil
}

type fcmpgInstruction struct{ knownJVMInstruction }

func parseFcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmpgInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x96,
			name: name,
		},
	}
	return &toReturn, nil
}

type dcmplInstruction struct{ knownJVMInstruction }

func parseDcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmplInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x97,
			name: name,
		},
	}
	return &toReturn, nil
}

type dcmpgInstruction struct{ knownJVMInstruction }

func parseDcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmpgInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0x98,
			name: name,
		},
	}
	return &toReturn, nil
}

type ifeqInstruction struct{ twoByteArgumentInstruction }

func parseIfeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifeqInstruction{*toReturn}, nil
}

type ifneInstruction struct{ twoByteArgumentInstruction }

func parseIfneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifneInstruction{*toReturn}, nil
}

type ifltInstruction struct{ twoByteArgumentInstruction }

func parseIfltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifltInstruction{*toReturn}, nil
}

type ifgeInstruction struct{ twoByteArgumentInstruction }

func parseIfgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifgeInstruction{*toReturn}, nil
}

type ifgtInstruction struct{ twoByteArgumentInstruction }

func parseIfgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifgtInstruction{*toReturn}, nil
}

type ifleInstruction struct{ twoByteArgumentInstruction }

func parseIfleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifleInstruction{*toReturn}, nil
}

type if_icmpeqInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpeqInstruction{*toReturn}, nil
}

type if_icmpneInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpneInstruction{*toReturn}, nil
}

type if_icmpltInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpltInstruction{*toReturn}, nil
}

type if_icmpgeInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpgeInstruction{*toReturn}, nil
}

type if_icmpgtInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpgtInstruction{*toReturn}, nil
}

type if_icmpleInstruction struct{ twoByteArgumentInstruction }

func parseIf_icmpleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpleInstruction{*toReturn}, nil
}

type if_acmpeqInstruction struct{ twoByteArgumentInstruction }

func parseIf_acmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_acmpeqInstruction{*toReturn}, nil
}

type if_acmpneInstruction struct{ twoByteArgumentInstruction }

func parseIf_acmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_acmpneInstruction{*toReturn}, nil
}

type gotoInstruction struct{ twoByteArgumentInstruction }

func parseGotoInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &gotoInstruction{*toReturn}, nil
}

type jsrInstruction struct{ twoByteArgumentInstruction }

func parseJsrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &jsrInstruction{*toReturn}, nil
}

type retInstruction struct{ singleByteArgumentInstruction }

func parseRetInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &retInstruction{*toReturn}, nil
}

// Probably the hairiest opcode in Java; this contains a list of potential
// jump targets.
type tableswitchInstruction struct {
	// This may be nil, if the start of the high offset was already 4-byte
	// aligned. It will contain at most 3 bytes.
	skippedBytes  []byte
	defaultOffset uint32
	lowOffset     uint32
	highOffset    uint32
	offsets       []uint32
}

func (n *tableswitchInstruction) Raw() uint8 {
	return 0xaa
}

func (n *tableswitchInstruction) Length() uint {
	// 12 bytes for high, low and default offsets, plus one byte for the
	// opcode, 4 per offset in the list, and up to 3 skipped padding bytes.
	// This is called by OtherBytes() to allocate a buffer, so it must not
	// depend on OtherBytes().
	return uint(len(n.skippedBytes)) + uint(len(n.offsets)*4) + 13
}

func (n *tableswitchInstruction) OtherBytes() []byte {
	toReturn := make([]byte, n.Length()-1)
	offset := 0
	// Use this inner function for convenience, and allowing us to avoid
	// encoding/binary.
	// TODO: Test this!!
	writeValueToBuffer := func(value uint32) {
		toReturn[offset] = uint8(value >> 24)
		toReturn[offset+1] = uint8(value >> 16)
		toReturn[offset+2] = uint8(value >> 8)
		toReturn[offset+3] = uint8(value)
		offset += 4
	}
	if n.skippedBytes != nil {
		copy(toReturn, n.skippedBytes)
		offset += len(n.skippedBytes)
	}
	writeValueToBuffer(n.highOffset)
	writeValueToBuffer(n.lowOffset)
	writeValueToBuffer(n.defaultOffset)
	for _, v := range n.offsets {
		writeValueToBuffer(v)
	}
	return toReturn
}

func (n *tableswitchInstruction) String() string {
	return fmt.Sprintf("tableswitch 0x%08x-0x%08x (default 0x%08x)",
		n.lowOffset, n.highOffset, n.defaultOffset)
}

func parseTableswitchInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	var e error
	var toReturn tableswitchInstruction
	currentOffset := address + 1
	// First, read up to 3 bytes to get to a 4-byte aligned address
	paddingBytes := address % 4
	if paddingBytes > 0 {
		toReturn.skippedBytes = make([]byte, paddingBytes)
		for i := range toReturn.skippedBytes {
			toReturn.skippedBytes[i], e = m.GetByte(currentOffset)
			if e != nil {
				return nil, fmt.Errorf("Couldn't align tableswitch: %s", e)
			}
			currentOffset++
		}
	}
	toReturn.defaultOffset, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch default: %s", e)
	}
	currentOffset += 4
	toReturn.lowOffset, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch low offset: %s", e)
	}
	currentOffset += 4
	toReturn.highOffset, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch high offset: %s", e)
	}
	currentOffset += 4
	if toReturn.highOffset < toReturn.lowOffset {
		return nil, fmt.Errorf("Tableswitch offset range invalid")
	}
	offsetsCount := toReturn.highOffset - toReturn.lowOffset + 1
	toReturn.offsets = make([]uint32, offsetsCount)
	for i := range toReturn.offsets {
		toReturn.offsets[i], e = Read32Bits(m, currentOffset)
		if e != nil {
			return nil, fmt.Errorf("Failed reading tableswitch offset: %s", e)
		}
		currentOffset += 4
	}
	return &toReturn, nil
}

// A single entry in the lookupswitch instruction's table.
type lookupswitchPair struct {
	match  int32
	offset uint32
}

// A very similar structure to the tableswitch instruction.
type lookupswitchInstruction struct {
	// For 4-byte alignment, like tableswitch
	skippedBytes  []byte
	defaultOffset uint32
	pairs         []lookupswitchPair
}

func (n *lookupswitchInstruction) Raw() uint8 {
	return 0xab
}

func (n *lookupswitchInstruction) Length() uint {
	return uint(len(n.skippedBytes)) + uint(len(n.pairs)*8) + 9
}

func (n *lookupswitchInstruction) OtherBytes() []byte {
	toReturn := make([]byte, n.Length())
	offset := 0
	appendValue := func(v uint32) {
		toReturn[offset] = uint8(v >> 24)
		toReturn[offset+1] = uint8(v >> 16)
		toReturn[offset+2] = uint8(v >> 8)
		toReturn[offset+3] = uint8(v)
		offset += 4
	}
	if n.skippedBytes != nil {
		copy(toReturn, n.skippedBytes)
		offset += len(n.skippedBytes)
	}
	appendValue(n.defaultOffset)
	appendValue(uint32(len(n.pairs)))
	for i := range n.pairs {
		appendValue(uint32(n.pairs[i].match))
		appendValue(n.pairs[i].offset)
	}
	return toReturn
}

func (n *lookupswitchInstruction) String() string {
	return fmt.Sprintf("lookupswitch [%d possible] (default offset 0x%08x)",
		len(n.pairs), n.defaultOffset)
}

func parseLookupswitchInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	var e error
	var toReturn lookupswitchInstruction
	currentOffset := address + 1
	// Skip padding as in tableswitch
	paddingBytes := address % 4
	if paddingBytes > 0 {
		toReturn.skippedBytes = make([]byte, paddingBytes)
		for i := range toReturn.skippedBytes {
			toReturn.skippedBytes[i], e = m.GetByte(currentOffset)
			if e != nil {
				return nil, fmt.Errorf("Couldn't align lookupswitch: %s", e)
			}
			currentOffset++
		}
	}
	toReturn.defaultOffset, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read lookupswitch default: %s", e)
	}
	currentOffset += 4
	pairsCount, e := Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read lookupswitch size: %s", e)
	}
	currentOffset += 4
	toReturn.pairs = make([]lookupswitchPair, pairsCount)
	// 0 pairs is technically legal.
	if pairsCount == 0 {
		return &toReturn, nil
	}
	var tmp uint32
	for i := range toReturn.pairs {
		tmp, e = Read32Bits(m, currentOffset)
		if e != nil {
			return nil, fmt.Errorf("Couldn't read lookupswitch match: %s", e)
		}
		toReturn.pairs[i].match = int32(tmp)
		currentOffset += 4
		toReturn.pairs[i].offset, e = Read32Bits(m, currentOffset)
		if e != nil {
			return nil, fmt.Errorf("Couldn't read lookupswitch offset: %s", e)
		}
		currentOffset += 4
	}
	// Finally, verify that the values are in sorted order.
	prevMatch := toReturn.pairs[0].match
	var thisMatch int32
	for i := range toReturn.pairs {
		if i == 0 {
			continue
		}
		thisMatch = toReturn.pairs[i].match
		if prevMatch >= thisMatch {
			return nil, fmt.Errorf("lookupswitch table not sorted")
		}
		prevMatch = thisMatch
	}
	return &toReturn, nil
}

type ireturnInstruction struct{ knownJVMInstruction }

func parseIreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ireturnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xac,
			name: name,
		},
	}
	return &toReturn, nil
}

type lreturnInstruction struct{ knownJVMInstruction }

func parseLreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lreturnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xad,
			name: name,
		},
	}
	return &toReturn, nil
}

type freturnInstruction struct{ knownJVMInstruction }

func parseFreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := freturnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xae,
			name: name,
		},
	}
	return &toReturn, nil
}

type dreturnInstruction struct{ knownJVMInstruction }

func parseDreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dreturnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xaf,
			name: name,
		},
	}
	return &toReturn, nil
}

type areturnInstruction struct{ knownJVMInstruction }

func parseAreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := areturnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xb0,
			name: name,
		},
	}
	return &toReturn, nil
}

type returnInstruction struct{ knownJVMInstruction }

func parseReturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := returnInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xb1,
			name: name,
		},
	}
	return &toReturn, nil
}

type getstaticInstruction struct{ twoByteArgumentInstruction }

func parseGetstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &getstaticInstruction{*toReturn}, nil
}

type putstaticInstruction struct{ twoByteArgumentInstruction }

func parsePutstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &putstaticInstruction{*toReturn}, nil
}

type getfieldInstruction struct{ twoByteArgumentInstruction }

func parseGetfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &getfieldInstruction{*toReturn}, nil
}

type putfieldInstruction struct{ twoByteArgumentInstruction }

func parsePutfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &putfieldInstruction{*toReturn}, nil
}

type invokevirtualInstruction struct{ twoByteArgumentInstruction }

func parseInvokevirtualInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokevirtualInstruction{*toReturn}, nil
}

type invokespecialInstruction struct{ twoByteArgumentInstruction }

func parseInvokespecialInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokespecialInstruction{*toReturn}, nil
}

type invokestaticInstruction struct{ twoByteArgumentInstruction }

func parseInvokestaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokestaticInstruction{*toReturn}, nil
}

type invokeinterfaceInstruction struct {
	twoByteArgumentInstruction
	count uint8
}

// The invokeinterface instruction contains a single 0-byte at the end.
func (n *invokeinterfaceInstruction) OtherBytes() []byte {
	toReturn := make([]byte, 5)
	copy(toReturn, (&(n.twoByteArgumentInstruction)).OtherBytes())
	toReturn[3] = n.count
	return toReturn
}

func (n *invokeinterfaceInstruction) Length() uint {
	return 5
}

func parseInvokeinterfaceInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	tmp, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	count, e := m.GetByte(address + 3)
	if e != nil {
		return nil, fmt.Errorf("Failed getting invokeinterface count: %s", e)
	}
	toReturn := invokeinterfaceInstruction{
		twoByteArgumentInstruction: *tmp,
		count: count,
	}
	return &toReturn, nil
}

type invokedynamicInstruction struct{ twoByteArgumentInstruction }

// The invokedynamic instruction contains two 0-bytes following the 16-bit
// index.
func (n *invokedynamicInstruction) OtherBytes() []byte {
	toReturn := make([]byte, 5)
	copy(toReturn, (&n.twoByteArgumentInstruction).OtherBytes())
	return toReturn
}

func (n *invokedynamicInstruction) Length() uint {
	return 5
}

func parseInvokedynamicInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokedynamicInstruction{*toReturn}, nil
}

type newInstruction struct{ twoByteArgumentInstruction }

func parseNewInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &newInstruction{*toReturn}, nil
}

type newarrayInstruction struct{ singleByteArgumentInstruction }

func parseNewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &newarrayInstruction{*toReturn}, nil
}

type anewarrayInstruction struct{ twoByteArgumentInstruction }

func parseAnewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &anewarrayInstruction{*toReturn}, nil
}

type arraylengthInstruction struct{ knownJVMInstruction }

func parseArraylengthInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := arraylengthInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xbe,
			name: name,
		},
	}
	return &toReturn, nil
}

type athrowInstruction struct{ knownJVMInstruction }

func parseAthrowInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := athrowInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xbf,
			name: name,
		},
	}
	return &toReturn, nil
}

type checkcastInstruction struct{ twoByteArgumentInstruction }

func parseCheckcastInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &checkcastInstruction{*toReturn}, nil
}

type instanceofInstruction struct{ twoByteArgumentInstruction }

func parseInstanceofInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &instanceofInstruction{*toReturn}, nil
}

type monitorenterInstruction struct{ knownJVMInstruction }

func parseMonitorenterInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorenterInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xc2,
			name: name,
		},
	}
	return &toReturn, nil
}

type monitorexitInstruction struct{ knownJVMInstruction }

func parseMonitorexitInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorexitInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xc3,
			name: name,
		},
	}
	return &toReturn, nil
}

// Holds the opcode and argument affected by the wide instruction opcode.
type wideInstruction struct {
	// This opcode *must* be one of iload, fload, aload, lload, dload, istore,
	// fstore, astore, lstore, dstore, or ret. This will be checked by
	// parseWideInstruction
	opcode   uint8
	argument uint16
}

func (n *wideInstruction) Raw() uint8 {
	return 0xc4
}

func (n *wideInstruction) OtherBytes() []byte {
	return []byte{n.opcode, uint8(n.argument >> 8), uint8(n.argument)}
}

func (n *wideInstruction) Length() uint {
	return 4
}

func (n *wideInstruction) String() string {
	return fmt.Sprintf("wide %s 0x%04x", opcodeTable[n.opcode].name,
		n.argument)
}

// The wide iinc instruction has an additional two-byte argument in addition to
// the wide opcode.
type wideIincInstruction struct {
	index uint16
	value uint16
}

func (n *wideIincInstruction) Raw() uint8 {
	return 0xc4
}

func (n *wideIincInstruction) OtherBytes() []byte {
	// The iinc instruction has opcode 0x84
	return []byte{
		0x84,
		uint8(n.index >> 8),
		uint8(n.index),
		uint8(n.value >> 8),
		uint8(n.value),
	}
}

func (n *wideIincInstruction) Length() uint {
	return 6
}

func (n *wideIincInstruction) String() string {
	return fmt.Sprintf("wide iinc 0x%04x 0x%04x", n.index, n.value)
}

// This only parses "wide iinc ..." instructions, and will be called by
// parseWideInstruction
func parseWideIincInstruction(address uint, m JVMMemory) (
	JVMInstruction, error) {
	index, e := Read16Bits(m, address+2)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read index for wide iinc: %s", e)
	}
	value, e := Read16Bits(m, address+4)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read value for wide iinc: %s", e)
	}
	toReturn := wideIincInstruction{
		index: index,
		value: value,
	}
	return &toReturn, nil
}

func parseWideInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	opcode, e := m.GetByte(address + 1)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get wide instruction opcode: %s", e)
	}
	// Check that the opcode is one of iload, fload, aload, lload, dload,
	// istore, fstore, astore, lstore, dstore, or ret. iinc is also handled
	// in a separate function returning a different struct.
	switch {
	case opcode == 0x84:
		return parseWideIincInstruction(address, m)
	case (opcode >= 0x15) && (opcode <= 0x19):
		// The opcode is one of the load instructions.
		break
	case (opcode >= 0x36) && (opcode <= 0x3a):
		// The opcode is one of the store instructions.
	default:
		return nil, fmt.Errorf("Invalid wide instruction opcode: 0x%02x",
			opcode)
	}
	value, e := Read16Bits(m, address+2)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read wide instruction argument: %s",
			e)
	}
	toReturn := wideInstruction{
		opcode:   opcode,
		argument: value,
	}
	return &toReturn, nil
}

type multianewarrayInstruction struct {
	typeIndex  uint16
	dimensions uint8
}

func (n *multianewarrayInstruction) Raw() uint8 {
	return 0xc5
}

func (n *multianewarrayInstruction) OtherBytes() []byte {
	return []byte{uint8(n.typeIndex >> 8), uint8(n.typeIndex),
		n.dimensions}
}

func (n *multianewarrayInstruction) Length() uint {
	return 4
}

func (n *multianewarrayInstruction) String() string {
	return fmt.Sprintf("multianewarray 0x%04x %d", n.typeIndex, n.dimensions)
}

func parseMultianewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	typeIndex, e := Read16Bits(m, address+1)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get multianewarray type index: %s", e)
	}
	dimensions, e := m.GetByte(address + 3)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get multianewarray dimensions: %s", e)
	}
	toReturn := multianewarrayInstruction{
		typeIndex:  typeIndex,
		dimensions: dimensions,
	}
	return &toReturn, nil
}

type ifnullInstruction struct{ twoByteArgumentInstruction }

func parseIfnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifnullInstruction{*toReturn}, nil
}

type ifnonnullInstruction struct{ twoByteArgumentInstruction }

func parseIfnonnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifnonnullInstruction{*toReturn}, nil
}

// Can be any instruction which uses a single four-byte argument, such as jsr_w
// or goto_w.
type fourByteArgumentInstruction struct {
	raw   uint8
	name  string
	value uint32
}

func (n *fourByteArgumentInstruction) Raw() uint8 {
	return n.raw
}

func (n *fourByteArgumentInstruction) OtherBytes() []byte {
	return []byte{uint8(n.value >> 24), uint8(n.value >> 16),
		uint8(n.value >> 8), uint8(n.value)}
}

func (n *fourByteArgumentInstruction) Length() uint {
	return 5
}

func (n *fourByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%08x", n.name, n.value)
}

func parseFourByteArgumentInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (*fourByteArgumentInstruction, error) {
	value, e := Read32Bits(m, address+1)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read argument for %s: %s", name, e)
	}
	toReturn := fourByteArgumentInstruction{
		raw:   opcode,
		name:  name,
		value: value,
	}
	return &toReturn, nil
}

type goto_wInstruction struct{ fourByteArgumentInstruction }

func parseGoto_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &goto_wInstruction{*toReturn}, nil
}

type jsr_wInstruction struct{ fourByteArgumentInstruction }

func parseJsr_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &jsr_wInstruction{*toReturn}, nil
}

type breakpointInstruction struct{ knownJVMInstruction }

func parseBreakpointInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := breakpointInstruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xca,
			name: name,
		},
	}
	return &toReturn, nil
}

type impdep1Instruction struct{ knownJVMInstruction }

func parseImpdep1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep1Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xfe,
			name: name,
		},
	}
	return &toReturn, nil
}

type impdep2Instruction struct{ knownJVMInstruction }

func parseImpdep2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep2Instruction{
		knownJVMInstruction: knownJVMInstruction{
			raw:  0xff,
			name: name,
		},
	}
	return &toReturn, nil
}
