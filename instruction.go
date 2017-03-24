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
type basicJVMInstruction struct {
	raw uint8
}

func (n *basicJVMInstruction) Raw() uint8 {
	return n.raw
}

func (n *basicJVMInstruction) OtherBytes() []byte {
	return nil
}

func (n *basicJVMInstruction) Length() uint {
	return 1
}

func (n *basicJVMInstruction) Execute(t JVMThread) error {
	return UnknownInstructionError(n.raw)
}

func (n *basicJVMInstruction) String() string {
	return fmt.Sprintf("<unknown instruction 0x%02x>", n.raw)
}

// Like basicJVMInstruction, but contains an instruction string. Used for
// known instructions which only consist of one byte.
type knownJVMInstruction struct {
	basicJVMInstruction
	name string
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
		toReturn := &basicJVMInstruction{
			raw: firstByte,
		}
		return toReturn, nil
	}
	return opcodeInfo.parse(opcodeInfo.opcode, opcodeInfo.name, address, m)
}

type nopInstruction knownJVMInstruction

func parseNopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := nopInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x00,
		},
		name: name,
	}
	return &toReturn, nil
}

type aconst_nullInstruction knownJVMInstruction

func parseAconst_nullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aconst_nullInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x01,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_m1Instruction knownJVMInstruction

func parseIconst_m1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_m1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x02,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_0Instruction knownJVMInstruction

func parseIconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x03,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_1Instruction knownJVMInstruction

func parseIconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x04,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_2Instruction knownJVMInstruction

func parseIconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x05,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_3Instruction knownJVMInstruction

func parseIconst_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x06,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_4Instruction knownJVMInstruction

func parseIconst_4Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_4Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x07,
		},
		name: name,
	}
	return &toReturn, nil
}

type iconst_5Instruction knownJVMInstruction

func parseIconst_5Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_5Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x08,
		},
		name: name,
	}
	return &toReturn, nil
}

type lconst_0Instruction knownJVMInstruction

func parseLconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x09,
		},
		name: name,
	}
	return &toReturn, nil
}

type lconst_1Instruction knownJVMInstruction

func parseLconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0a,
		},
		name: name,
	}
	return &toReturn, nil
}

type fconst_0Instruction knownJVMInstruction

func parseFconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0b,
		},
		name: name,
	}
	return &toReturn, nil
}

type fconst_1Instruction knownJVMInstruction

func parseFconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0c,
		},
		name: name,
	}
	return &toReturn, nil
}

type fconst_2Instruction knownJVMInstruction

func parseFconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0d,
		},
		name: name,
	}
	return &toReturn, nil
}

type dconst_0Instruction knownJVMInstruction

func parseDconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0e,
		},
		name: name,
	}
	return &toReturn, nil
}

type dconst_1Instruction knownJVMInstruction

func parseDconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x0f,
		},
		name: name,
	}
	return &toReturn, nil
}

// This covers instructions such as bipush, which have one argument byte past
// the opcode byte.
type singleByteArgumentInstruction struct {
	basicJVMInstruction
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
		basicJVMInstruction: basicJVMInstruction{
			raw: opcode,
		},
		name:  name,
		value: value,
	}
	return &toReturn, nil
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

type bipushInstruction singleByteArgumentInstruction

func parseBipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*bipushInstruction)(toReturn), nil
}

// Any instruction which includes a two-byte value beyond its initial opcode.
type twoByteArgumentInstruction struct {
	basicJVMInstruction
	name  string
	value uint16
}

func (n *twoByteArgumentInstruction) Otherbytes() []byte {
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
		basicJVMInstruction: basicJVMInstruction{
			raw: opcode,
		},
		name:  name,
		value: value,
	}
	return &toReturn, nil
}

type sipushInstruction twoByteArgumentInstruction

func parseSipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*sipushInstruction)(toReturn), nil
}

type ldcInstruction singleByteArgumentInstruction

func parseLdcInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ldcInstruction)(toReturn), nil
}

type ldc_wInstruction twoByteArgumentInstruction

func parseLdc_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ldc_wInstruction)(toReturn), nil
}

type ldc2_wInstruction twoByteArgumentInstruction

func parseLdc2_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ldc2_wInstruction)(toReturn), nil
}

type iloadInstruction singleByteArgumentInstruction

func parseIloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*iloadInstruction)(toReturn), nil
}

type lloadInstruction singleByteArgumentInstruction

func parseLloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*lloadInstruction)(toReturn), nil
}

type floadInstruction singleByteArgumentInstruction

func parseFloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*floadInstruction)(toReturn), nil
}

type dloadInstruction singleByteArgumentInstruction

func parseDloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*dloadInstruction)(toReturn), nil
}

type aloadInstruction singleByteArgumentInstruction

func parseAloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*aloadInstruction)(toReturn), nil
}

type iload_0Instruction knownJVMInstruction

func parseIload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1a,
		},
		name: name,
	}
	return &toReturn, nil
}

type iload_1Instruction knownJVMInstruction

func parseIload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1b,
		},
		name: name,
	}
	return &toReturn, nil
}

type iload_2Instruction knownJVMInstruction

func parseIload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1c,
		},
		name: name,
	}
	return &toReturn, nil
}

type iload_3Instruction knownJVMInstruction

func parseIload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1d,
		},
		name: name,
	}
	return &toReturn, nil
}

type lload_0Instruction knownJVMInstruction

func parseLload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1e,
		},
		name: name,
	}
	return &toReturn, nil
}

type lload_1Instruction knownJVMInstruction

func parseLload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x1f,
		},
		name: name,
	}
	return &toReturn, nil
}

type lload_2Instruction knownJVMInstruction

func parseLload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x20,
		},
		name: name,
	}
	return &toReturn, nil
}

type lload_3Instruction knownJVMInstruction

func parseLload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x21,
		},
		name: name,
	}
	return &toReturn, nil
}

type fload_0Instruction knownJVMInstruction

func parseFload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x22,
		},
		name: name,
	}
	return &toReturn, nil
}

type fload_1Instruction knownJVMInstruction

func parseFload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x23,
		},
		name: name,
	}
	return &toReturn, nil
}

type fload_2Instruction knownJVMInstruction

func parseFload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x24,
		},
		name: name,
	}
	return &toReturn, nil
}

type fload_3Instruction knownJVMInstruction

func parseFload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x25,
		},
		name: name,
	}
	return &toReturn, nil
}

type dload_0Instruction knownJVMInstruction

func parseDload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x26,
		},
		name: name,
	}
	return &toReturn, nil
}

type dload_1Instruction knownJVMInstruction

func parseDload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x27,
		},
		name: name,
	}
	return &toReturn, nil
}

type dload_2Instruction knownJVMInstruction

func parseDload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x28,
		},
		name: name,
	}
	return &toReturn, nil
}

type dload_3Instruction knownJVMInstruction

func parseDload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x29,
		},
		name: name,
	}
	return &toReturn, nil
}

type aload_0Instruction knownJVMInstruction

func parseAload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2a,
		},
		name: name,
	}
	return &toReturn, nil
}

type aload_1Instruction knownJVMInstruction

func parseAload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2b,
		},
		name: name,
	}
	return &toReturn, nil
}

type aload_2Instruction knownJVMInstruction

func parseAload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2c,
		},
		name: name,
	}
	return &toReturn, nil
}

type aload_3Instruction knownJVMInstruction

func parseAload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2d,
		},
		name: name,
	}
	return &toReturn, nil
}

type ialoadInstruction knownJVMInstruction

func parseIaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ialoadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2e,
		},
		name: name,
	}
	return &toReturn, nil
}

type laloadInstruction knownJVMInstruction

func parseLaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x2f,
		},
		name: name,
	}
	return &toReturn, nil
}

type faloadInstruction knownJVMInstruction

func parseFaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x30,
		},
		name: name,
	}
	return &toReturn, nil
}

type daloadInstruction knownJVMInstruction

func parseDaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x31,
		},
		name: name,
	}
	return &toReturn, nil
}

type aaloadInstruction knownJVMInstruction

func parseAaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aaloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x32,
		},
		name: name,
	}
	return &toReturn, nil
}

type baloadInstruction knownJVMInstruction

func parseBaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := baloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x33,
		},
		name: name,
	}
	return &toReturn, nil
}

type caloadInstruction knownJVMInstruction

func parseCaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := caloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x34,
		},
		name: name,
	}
	return &toReturn, nil
}

type saloadInstruction knownJVMInstruction

func parseSaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := saloadInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x35,
		},
		name: name,
	}
	return &toReturn, nil
}

type istoreInstruction singleByteArgumentInstruction

func parseIstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*istoreInstruction)(toReturn), nil
}

type lstoreInstruction singleByteArgumentInstruction

func parseLstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*lstoreInstruction)(toReturn), nil
}

type fstoreInstruction singleByteArgumentInstruction

func parseFstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*fstoreInstruction)(toReturn), nil
}

type dstoreInstruction singleByteArgumentInstruction

func parseDstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*dstoreInstruction)(toReturn), nil
}

type astoreInstruction singleByteArgumentInstruction

func parseAstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*astoreInstruction)(toReturn), nil
}

type istore_0Instruction knownJVMInstruction

func parseIstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x3b,
		},
		name: name,
	}
	return &toReturn, nil
}

type istore_1Instruction knownJVMInstruction

func parseIstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x3c,
		},
		name: name,
	}
	return &toReturn, nil
}

type istore_2Instruction knownJVMInstruction

func parseIstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x3d,
		},
		name: name,
	}
	return &toReturn, nil
}

type istore_3Instruction knownJVMInstruction

func parseIstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x3e,
		},
		name: name,
	}
	return &toReturn, nil
}

type lstore_0Instruction knownJVMInstruction

func parseLstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x3f,
		},
		name: name,
	}
	return &toReturn, nil
}

type lstore_1Instruction knownJVMInstruction

func parseLstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x40,
		},
		name: name,
	}
	return &toReturn, nil
}

type lstore_2Instruction knownJVMInstruction

func parseLstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x41,
		},
		name: name,
	}
	return &toReturn, nil
}

type lstore_3Instruction knownJVMInstruction

func parseLstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x42,
		},
		name: name,
	}
	return &toReturn, nil
}

type fstore_0Instruction knownJVMInstruction

func parseFstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x43,
		},
		name: name,
	}
	return &toReturn, nil
}

type fstore_1Instruction knownJVMInstruction

func parseFstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x44,
		},
		name: name,
	}
	return &toReturn, nil
}

type fstore_2Instruction knownJVMInstruction

func parseFstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x45,
		},
		name: name,
	}
	return &toReturn, nil
}

type fstore_3Instruction knownJVMInstruction

func parseFstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x46,
		},
		name: name,
	}
	return &toReturn, nil
}

type dstore_0Instruction knownJVMInstruction

func parseDstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x47,
		},
		name: name,
	}
	return &toReturn, nil
}

type dstore_1Instruction knownJVMInstruction

func parseDstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x48,
		},
		name: name,
	}
	return &toReturn, nil
}

type dstore_2Instruction knownJVMInstruction

func parseDstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x49,
		},
		name: name,
	}
	return &toReturn, nil
}

type dstore_3Instruction knownJVMInstruction

func parseDstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4a,
		},
		name: name,
	}
	return &toReturn, nil
}

type astore_0Instruction knownJVMInstruction

func parseAstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_0Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4b,
		},
		name: name,
	}
	return &toReturn, nil
}

type astore_1Instruction knownJVMInstruction

func parseAstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4c,
		},
		name: name,
	}
	return &toReturn, nil
}

type astore_2Instruction knownJVMInstruction

func parseAstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4d,
		},
		name: name,
	}
	return &toReturn, nil
}

type astore_3Instruction knownJVMInstruction

func parseAstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_3Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4e,
		},
		name: name,
	}
	return &toReturn, nil
}

type iastoreInstruction knownJVMInstruction

func parseIastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x4f,
		},
		name: name,
	}
	return &toReturn, nil
}

type lastoreInstruction knownJVMInstruction

func parseLastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x50,
		},
		name: name,
	}
	return &toReturn, nil
}

type fastoreInstruction knownJVMInstruction

func parseFastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x51,
		},
		name: name,
	}
	return &toReturn, nil
}

type dastoreInstruction knownJVMInstruction

func parseDastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x52,
		},
		name: name,
	}
	return &toReturn, nil
}

type aastoreInstruction knownJVMInstruction

func parseAastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x53,
		},
		name: name,
	}
	return &toReturn, nil
}

type bastoreInstruction knownJVMInstruction

func parseBastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := bastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x54,
		},
		name: name,
	}
	return &toReturn, nil
}

type castoreInstruction knownJVMInstruction

func parseCastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := castoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x55,
		},
		name: name,
	}
	return &toReturn, nil
}

type sastoreInstruction knownJVMInstruction

func parseSastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := sastoreInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x56,
		},
		name: name,
	}
	return &toReturn, nil
}

type popInstruction knownJVMInstruction

func parsePopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := popInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x57,
		},
		name: name,
	}
	return &toReturn, nil
}

type pop2Instruction knownJVMInstruction

func parsePop2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := pop2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x58,
		},
		name: name,
	}
	return &toReturn, nil
}

type dupInstruction knownJVMInstruction

func parseDupInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dupInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x59,
		},
		name: name,
	}
	return &toReturn, nil
}

type dup_x1Instruction knownJVMInstruction

func parseDup_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5a,
		},
		name: name,
	}
	return &toReturn, nil
}

type dup_x2Instruction knownJVMInstruction

func parseDup_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5b,
		},
		name: name,
	}
	return &toReturn, nil
}

type dup2Instruction knownJVMInstruction

func parseDup2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5c,
		},
		name: name,
	}
	return &toReturn, nil
}

type dup2_x1Instruction knownJVMInstruction

func parseDup2_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5d,
		},
		name: name,
	}
	return &toReturn, nil
}

type dup2_x2Instruction knownJVMInstruction

func parseDup2_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5e,
		},
		name: name,
	}
	return &toReturn, nil
}

type swapInstruction knownJVMInstruction

func parseSwapInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := swapInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x5f,
		},
		name: name,
	}
	return &toReturn, nil
}

type iaddInstruction knownJVMInstruction

func parseIaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iaddInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x60,
		},
		name: name,
	}
	return &toReturn, nil
}

type laddInstruction knownJVMInstruction

func parseLaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laddInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x61,
		},
		name: name,
	}
	return &toReturn, nil
}

type faddInstruction knownJVMInstruction

func parseFaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faddInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x62,
		},
		name: name,
	}
	return &toReturn, nil
}

type daddInstruction knownJVMInstruction

func parseDaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daddInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x63,
		},
		name: name,
	}
	return &toReturn, nil
}

type isubInstruction knownJVMInstruction

func parseIsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := isubInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x64,
		},
		name: name,
	}
	return &toReturn, nil
}

type lsubInstruction knownJVMInstruction

func parseLsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lsubInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x65,
		},
		name: name,
	}
	return &toReturn, nil
}

type fsubInstruction knownJVMInstruction

func parseFsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fsubInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x66,
		},
		name: name,
	}
	return &toReturn, nil
}

type dsubInstruction knownJVMInstruction

func parseDsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dsubInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x67,
		},
		name: name,
	}
	return &toReturn, nil
}

type imulInstruction knownJVMInstruction

func parseImulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := imulInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x68,
		},
		name: name,
	}
	return &toReturn, nil
}

type lmulInstruction knownJVMInstruction

func parseLmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lmulInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x69,
		},
		name: name,
	}
	return &toReturn, nil
}

type fmulInstruction knownJVMInstruction

func parseFmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fmulInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6a,
		},
		name: name,
	}
	return &toReturn, nil
}

type dmulInstruction knownJVMInstruction

func parseDmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dmulInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6b,
		},
		name: name,
	}
	return &toReturn, nil
}

type idivInstruction knownJVMInstruction

func parseIdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := idivInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6c,
		},
		name: name,
	}
	return &toReturn, nil
}

type ldivInstruction knownJVMInstruction

func parseLdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ldivInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6d,
		},
		name: name,
	}
	return &toReturn, nil
}

type fdivInstruction knownJVMInstruction

func parseFdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fdivInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6e,
		},
		name: name,
	}
	return &toReturn, nil
}

type ddivInstruction knownJVMInstruction

func parseDdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ddivInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x6f,
		},
		name: name,
	}
	return &toReturn, nil
}

type iremInstruction knownJVMInstruction

func parseIremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iremInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x70,
		},
		name: name,
	}
	return &toReturn, nil
}

type lremInstruction knownJVMInstruction

func parseLremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lremInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x71,
		},
		name: name,
	}
	return &toReturn, nil
}

type fremInstruction knownJVMInstruction

func parseFremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fremInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x72,
		},
		name: name,
	}
	return &toReturn, nil
}

type dremInstruction knownJVMInstruction

func parseDremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dremInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x73,
		},
		name: name,
	}
	return &toReturn, nil
}

type inegInstruction knownJVMInstruction

func parseInegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := inegInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x74,
		},
		name: name,
	}
	return &toReturn, nil
}

type lnegInstruction knownJVMInstruction

func parseLnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lnegInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x75,
		},
		name: name,
	}
	return &toReturn, nil
}

type fnegInstruction knownJVMInstruction

func parseFnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fnegInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x76,
		},
		name: name,
	}
	return &toReturn, nil
}

type dnegInstruction knownJVMInstruction

func parseDnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dnegInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x77,
		},
		name: name,
	}
	return &toReturn, nil
}

type ishlInstruction knownJVMInstruction

func parseIshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishlInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x78,
		},
		name: name,
	}
	return &toReturn, nil
}

type lshlInstruction knownJVMInstruction

func parseLshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshlInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x79,
		},
		name: name,
	}
	return &toReturn, nil
}

type ishrInstruction knownJVMInstruction

func parseIshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishrInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7a,
		},
		name: name,
	}
	return &toReturn, nil
}

type lshrInstruction knownJVMInstruction

func parseLshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshrInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7b,
		},
		name: name,
	}
	return &toReturn, nil
}

type iushrInstruction knownJVMInstruction

func parseIushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iushrInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7c,
		},
		name: name,
	}
	return &toReturn, nil
}

type lushrInstruction knownJVMInstruction

func parseLushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lushrInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7d,
		},
		name: name,
	}
	return &toReturn, nil
}

type iandInstruction knownJVMInstruction

func parseIandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iandInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7e,
		},
		name: name,
	}
	return &toReturn, nil
}

type landInstruction knownJVMInstruction

func parseLandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := landInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x7f,
		},
		name: name,
	}
	return &toReturn, nil
}

type iorInstruction knownJVMInstruction

func parseIorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iorInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x80,
		},
		name: name,
	}
	return &toReturn, nil
}

type lorInstruction knownJVMInstruction

func parseLorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lorInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x81,
		},
		name: name,
	}
	return &toReturn, nil
}

type ixorInstruction knownJVMInstruction

func parseIxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ixorInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x82,
		},
		name: name,
	}
	return &toReturn, nil
}

type lxorInstruction knownJVMInstruction

func parseLxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lxorInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x83,
		},
		name: name,
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

type i2lInstruction knownJVMInstruction

func parseI2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2lInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x85,
		},
		name: name,
	}
	return &toReturn, nil
}

type i2fInstruction knownJVMInstruction

func parseI2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2fInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x86,
		},
		name: name,
	}
	return &toReturn, nil
}

type i2dInstruction knownJVMInstruction

func parseI2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2dInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x87,
		},
		name: name,
	}
	return &toReturn, nil
}

type l2iInstruction knownJVMInstruction

func parseL2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2iInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x88,
		},
		name: name,
	}
	return &toReturn, nil
}

type l2fInstruction knownJVMInstruction

func parseL2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2fInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x89,
		},
		name: name,
	}
	return &toReturn, nil
}

type l2dInstruction knownJVMInstruction

func parseL2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2dInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8a,
		},
		name: name,
	}
	return &toReturn, nil
}

type f2iInstruction knownJVMInstruction

func parseF2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2iInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8b,
		},
		name: name,
	}
	return &toReturn, nil
}

type f2lInstruction knownJVMInstruction

func parseF2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2lInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8c,
		},
		name: name,
	}
	return &toReturn, nil
}

type f2dInstruction knownJVMInstruction

func parseF2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2dInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8d,
		},
		name: name,
	}
	return &toReturn, nil
}

type d2iInstruction knownJVMInstruction

func parseD2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2iInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8e,
		},
		name: name,
	}
	return &toReturn, nil
}

type d2lInstruction knownJVMInstruction

func parseD2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2lInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x8f,
		},
		name: name,
	}
	return &toReturn, nil
}

type d2fInstruction knownJVMInstruction

func parseD2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2fInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x90,
		},
		name: name,
	}
	return &toReturn, nil
}

type i2bInstruction knownJVMInstruction

func parseI2bInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2bInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x91,
		},
		name: name,
	}
	return &toReturn, nil
}

type i2cInstruction knownJVMInstruction

func parseI2cInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2cInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x92,
		},
		name: name,
	}
	return &toReturn, nil
}

type i2sInstruction knownJVMInstruction

func parseI2sInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2sInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x93,
		},
		name: name,
	}
	return &toReturn, nil
}

type lcmpInstruction knownJVMInstruction

func parseLcmpInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lcmpInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x94,
		},
		name: name,
	}
	return &toReturn, nil
}

type fcmplInstruction knownJVMInstruction

func parseFcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmplInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x95,
		},
		name: name,
	}
	return &toReturn, nil
}

type fcmpgInstruction knownJVMInstruction

func parseFcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmpgInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x96,
		},
		name: name,
	}
	return &toReturn, nil
}

type dcmplInstruction knownJVMInstruction

func parseDcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmplInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x97,
		},
		name: name,
	}
	return &toReturn, nil
}

type dcmpgInstruction knownJVMInstruction

func parseDcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmpgInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0x98,
		},
		name: name,
	}
	return &toReturn, nil
}

type ifeqInstruction twoByteArgumentInstruction

func parseIfeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifeqInstruction)(toReturn), nil
}

type ifneInstruction twoByteArgumentInstruction

func parseIfneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifneInstruction)(toReturn), nil
}

type ifltInstruction twoByteArgumentInstruction

func parseIfltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifltInstruction)(toReturn), nil
}

type ifgeInstruction twoByteArgumentInstruction

func parseIfgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifgeInstruction)(toReturn), nil
}

type ifgtInstruction twoByteArgumentInstruction

func parseIfgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifgtInstruction)(toReturn), nil
}

type ifleInstruction twoByteArgumentInstruction

func parseIfleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifleInstruction)(toReturn), nil
}

type if_icmpeqInstruction twoByteArgumentInstruction

func parseIf_icmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpeqInstruction)(toReturn), nil
}

type if_icmpneInstruction twoByteArgumentInstruction

func parseIf_icmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpneInstruction)(toReturn), nil
}

type if_icmpltInstruction twoByteArgumentInstruction

func parseIf_icmpltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpltInstruction)(toReturn), nil
}

type if_icmpgeInstruction twoByteArgumentInstruction

func parseIf_icmpgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpgeInstruction)(toReturn), nil
}

type if_icmpgtInstruction twoByteArgumentInstruction

func parseIf_icmpgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpgtInstruction)(toReturn), nil
}

type if_icmpleInstruction twoByteArgumentInstruction

func parseIf_icmpleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_icmpleInstruction)(toReturn), nil
}

type if_acmpeqInstruction twoByteArgumentInstruction

func parseIf_acmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_acmpeqInstruction)(toReturn), nil
}

type if_acmpneInstruction twoByteArgumentInstruction

func parseIf_acmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*if_acmpneInstruction)(toReturn), nil
}

type gotoInstruction twoByteArgumentInstruction

func parseGotoInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*gotoInstruction)(toReturn), nil
}

type jsrInstruction twoByteArgumentInstruction

func parseJsrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*jsrInstruction)(toReturn), nil
}

type retInstruction singleByteArgumentInstruction

func parseRetInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*retInstruction)(toReturn), nil
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

type ireturnInstruction knownJVMInstruction

func parseIreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ireturnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xac,
		},
		name: name,
	}
	return &toReturn, nil
}

type lreturnInstruction knownJVMInstruction

func parseLreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lreturnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xad,
		},
		name: name,
	}
	return &toReturn, nil
}

type freturnInstruction knownJVMInstruction

func parseFreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := freturnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xae,
		},
		name: name,
	}
	return &toReturn, nil
}

type dreturnInstruction knownJVMInstruction

func parseDreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dreturnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xaf,
		},
		name: name,
	}
	return &toReturn, nil
}

type areturnInstruction knownJVMInstruction

func parseAreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := areturnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xb0,
		},
		name: name,
	}
	return &toReturn, nil
}

type returnInstruction knownJVMInstruction

func parseReturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := returnInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xb1,
		},
		name: name,
	}
	return &toReturn, nil
}

type getstaticInstruction twoByteArgumentInstruction

func parseGetstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*getstaticInstruction)(toReturn), nil
}

type putstaticInstruction twoByteArgumentInstruction

func parsePutstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*putstaticInstruction)(toReturn), nil
}

type getfieldInstruction twoByteArgumentInstruction

func parseGetfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*getfieldInstruction)(toReturn), nil
}

type putfieldInstruction twoByteArgumentInstruction

func parsePutfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*putfieldInstruction)(toReturn), nil
}

type invokevirtualInstruction twoByteArgumentInstruction

func parseInvokevirtualInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*invokevirtualInstruction)(toReturn), nil
}

type invokespecialInstruction twoByteArgumentInstruction

func parseInvokespecialInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*invokespecialInstruction)(toReturn), nil
}

type invokestaticInstruction twoByteArgumentInstruction

func parseInvokestaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*invokestaticInstruction)(toReturn), nil
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

type invokedynamicInstruction twoByteArgumentInstruction

// The invokedynamic instruction contains two 0-bytes following the 16-bit
// index.
func (n *invokedynamicInstruction) OtherBytes() []byte {
	toReturn := make([]byte, 5)
	copy(toReturn, (*twoByteArgumentInstruction)(n).OtherBytes())
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
	return (*invokedynamicInstruction)(toReturn), nil
}

type newInstruction twoByteArgumentInstruction

func parseNewInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*newInstruction)(toReturn), nil
}

type newarrayInstruction singleByteArgumentInstruction

func parseNewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*newarrayInstruction)(toReturn), nil
}

type anewarrayInstruction twoByteArgumentInstruction

func parseAnewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*anewarrayInstruction)(toReturn), nil
}

type arraylengthInstruction knownJVMInstruction

func parseArraylengthInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := arraylengthInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xbe,
		},
		name: name,
	}
	return &toReturn, nil
}

type athrowInstruction knownJVMInstruction

func parseAthrowInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := athrowInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xbf,
		},
		name: name,
	}
	return &toReturn, nil
}

type checkcastInstruction twoByteArgumentInstruction

func parseCheckcastInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*checkcastInstruction)(toReturn), nil
}

type instanceofInstruction twoByteArgumentInstruction

func parseInstanceofInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*instanceofInstruction)(toReturn), nil
}

type monitorenterInstruction knownJVMInstruction

func parseMonitorenterInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorenterInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xc2,
		},
		name: name,
	}
	return &toReturn, nil
}

type monitorexitInstruction knownJVMInstruction

func parseMonitorexitInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorexitInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xc3,
		},
		name: name,
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

type ifnullInstruction twoByteArgumentInstruction

func parseIfnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifnullInstruction)(toReturn), nil
}

type ifnonnullInstruction twoByteArgumentInstruction

func parseIfnonnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*ifnonnullInstruction)(toReturn), nil
}

// Can be any instruction which uses a single four-byte argument, such as jsr_w
// or goto_w.
type fourByteArgumentInstruction struct {
	basicJVMInstruction
	name  string
	value uint32
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
		basicJVMInstruction: basicJVMInstruction{
			raw: opcode,
		},
		name:  name,
		value: value,
	}
	return &toReturn, nil
}

type goto_wInstruction fourByteArgumentInstruction

func parseGoto_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*goto_wInstruction)(toReturn), nil
}

type jsr_wInstruction fourByteArgumentInstruction

func parseJsr_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return (*jsr_wInstruction)(toReturn), nil
}

type breakpointInstruction knownJVMInstruction

func parseBreakpointInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := breakpointInstruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xca,
		},
		name: name,
	}
	return &toReturn, nil
}

type impdep1Instruction knownJVMInstruction

func parseImpdep1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep1Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xfe,
		},
		name: name,
	}
	return &toReturn, nil
}

type impdep2Instruction knownJVMInstruction

func parseImpdep2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep2Instruction{
		basicJVMInstruction: basicJVMInstruction{
			raw: 0xff,
		},
		name: name,
	}
	return &toReturn, nil
}
