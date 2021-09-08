package bs_jvm

// This file contains functions for disassembling JVM instructions.

import (
	"fmt"
)

// The interface through which JVM opcodes can be inspected or executed.
type Instruction interface {
	// Returns the 8-bit opcode for the instruction
	Raw() uint8
	// Returns additional bytes following the instruction's 8-bit opcode, or
	// nil if the instruction doesn't have such bytes. May be slow for some
	// opcodes.
	OtherBytes() []byte
	// This takes a reference to the current pre-parsed method, the
	// instruction's offset (in bytes) into the method code, and a map of
	// instruction offsets to instruction indices. This must be called during
	// an optimization pass before execution.
	Optimize(m *Method, offset uint, instructionIndices map[uint]int) error
	// Runs the instruction in the given thread
	Execute(t *Thread) error
	// Returns the length of the instruction, including the opcode and
	// additional argument bytes.
	Length() uint
	// Returns the disassembly string of the instruction
	String() string
}

// Provices a default implementation of the Instruction interface.
type unknownInstruction struct {
	raw uint8
}

func (n *unknownInstruction) Raw() uint8 {
	return n.raw
}

func (n *unknownInstruction) OtherBytes() []byte {
	return nil
}

func (n *unknownInstruction) Length() uint {
	return 1
}

func (n *unknownInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *unknownInstruction) Execute(t *Thread) error {
	return UnknownInstructionError(n.raw)
}

func (n *unknownInstruction) String() string {
	return fmt.Sprintf("<unknown instruction 0x%02x>", n.raw)
}

// Like unknownInstruction, but contains an instruction string. Used for
// known instructions which only consist of one byte.
type knownInstruction struct {
	raw  uint8
	name string
}

func (n *knownInstruction) Raw() uint8 {
	return n.raw
}

func (n *knownInstruction) OtherBytes() []byte {
	return nil
}

func (n *knownInstruction) Length() uint {
	return 1
}

func (n *knownInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *knownInstruction) Execute(t *Thread) error {
	return fmt.Errorf("Execution not implemented for %s", n.String())
}

func (n *knownInstruction) String() string {
	return n.name
}

// Returns the instruction starting at the given address. Returns an error if
// the address is invalid. If an invalid/unknown instruction is located at the
// address, then a Instruction will still be returned, but it will produce
// an UnknownInstructionError if executed.
func GetNextInstruction(m Memory, address uint) (Instruction, error) {
	firstByte, e := m.GetByte(address)
	if e != nil {
		return nil, e
	}
	opcodeInfo := opcodeTable[firstByte]
	// Unknown instruction.
	if opcodeInfo == nil {
		toReturn := &unknownInstruction{
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

type nopInstruction struct{ knownInstruction }

func parseNopInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := nopInstruction{
		knownInstruction: knownInstruction{
			raw:  0x00,
			name: name,
		},
	}
	return &toReturn, nil
}

type aconst_nullInstruction struct{ knownInstruction }

func parseAconst_nullInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aconst_nullInstruction{
		knownInstruction: knownInstruction{
			raw:  0x01,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_m1Instruction struct{ knownInstruction }

func parseIconst_m1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_m1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x02,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_0Instruction struct{ knownInstruction }

func parseIconst_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x03,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_1Instruction struct{ knownInstruction }

func parseIconst_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x04,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_2Instruction struct{ knownInstruction }

func parseIconst_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x05,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_3Instruction struct{ knownInstruction }

func parseIconst_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x06,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_4Instruction struct{ knownInstruction }

func parseIconst_4Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_4Instruction{
		knownInstruction: knownInstruction{
			raw:  0x07,
			name: name,
		},
	}
	return &toReturn, nil
}

type iconst_5Instruction struct{ knownInstruction }

func parseIconst_5Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iconst_5Instruction{
		knownInstruction: knownInstruction{
			raw:  0x08,
			name: name,
		},
	}
	return &toReturn, nil
}

type lconst_0Instruction struct{ knownInstruction }

func parseLconst_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lconst_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x09,
			name: name,
		},
	}
	return &toReturn, nil
}

type lconst_1Instruction struct{ knownInstruction }

func parseLconst_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lconst_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x0a,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_0Instruction struct{ knownInstruction }

func parseFconst_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fconst_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x0b,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_1Instruction struct{ knownInstruction }

func parseFconst_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fconst_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x0c,
			name: name,
		},
	}
	return &toReturn, nil
}

type fconst_2Instruction struct{ knownInstruction }

func parseFconst_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fconst_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x0d,
			name: name,
		},
	}
	return &toReturn, nil
}

type dconst_0Instruction struct{ knownInstruction }

func parseDconst_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dconst_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x0e,
			name: name,
		},
	}
	return &toReturn, nil
}

type dconst_1Instruction struct{ knownInstruction }

func parseDconst_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dconst_1Instruction{
		knownInstruction: knownInstruction{
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
	address uint, m Memory) (*singleByteArgumentInstruction, error) {
	value, e := m.GetByte(address + 1)
	if e != nil {
		return nil, fmt.Errorf("Failed reading argument byte for %s: %s", name,
			e)
	}
	toReturn := singleByteArgumentInstruction{
		raw:   opcode,
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

func (n *singleByteArgumentInstruction) Optimize(m *Method,
	offset uint, instructionIndices map[uint]int) error {
	return nil
}

func (n *singleByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%02x", n.name, n.value)
}

type bipushInstruction struct{ singleByteArgumentInstruction }

func parseBipushInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
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

func (n *twoByteArgumentInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *twoByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%04x", n.name, n.value)
}

func parseTwoByteArgumentInstruction(opcode uint8, name string, address uint,
	m Memory) (*twoByteArgumentInstruction, error) {
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
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &sipushInstruction{*toReturn}, nil
}

type ldcInstruction struct {
	singleByteArgumentInstruction
	// True if the LDC constant was an int or float.
	isPrimitive bool
	// This will contain the value to push, if isPrimitive was true. If the
	// constant is a float, this will be the bits of the float.
	primitiveValue Int
	// This will be the reference to push, if isPrimitive was false. If
	// isPrimitive was true, this will still be set, but to the primitive
	// reference.
	reference Object
}

func parseLdcInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldcInstruction{singleByteArgumentInstruction: *toReturn}, nil
}

func (n *ldcInstruction) String() string {
	if n.reference != nil {
		return fmt.Sprintf("ldc %s", n.reference)
	}
	// This will be the case if the instruction hasn't been optimized yet.
	return fmt.Sprintf("ldc 0x%02x", n.value)
}

type ldc_wInstruction struct {
	twoByteArgumentInstruction
	// All of these are the same as for ldcInstruction
	isPrimitive    bool
	primitiveValue Int
	reference      Object
}

func parseLdc_wInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldc_wInstruction{twoByteArgumentInstruction: *toReturn}, nil
}

func (n *ldc_wInstruction) String() string {
	if n.reference != nil {
		return fmt.Sprintf("ldc_w %s", n.reference)
	}
	return fmt.Sprintf("ldc_w 0x%04x", n.value)
}

type ldc2_wInstruction struct {
	twoByteArgumentInstruction
	// Once again, this will contain the bits of the double-precision number
	primitiveValue Long
	// This will be the primitive as an object, mostly so that a string can be
	// formatted nicely.
	reference Object
}

func parseLdc2_wInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ldc2_wInstruction{twoByteArgumentInstruction: *toReturn}, nil
}

func (n *ldc2_wInstruction) String() string {
	if n.reference != nil {
		return fmt.Sprintf("ldc2_w %s", n.reference)
	}
	return fmt.Sprintf("ldc2_w 0x%04x", n.value)
}

type iloadInstruction struct{ singleByteArgumentInstruction }

func parseIloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &iloadInstruction{*toReturn}, nil
}

type lloadInstruction struct{ singleByteArgumentInstruction }

func parseLloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &lloadInstruction{*toReturn}, nil
}

type floadInstruction struct{ singleByteArgumentInstruction }

func parseFloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &floadInstruction{*toReturn}, nil
}

type dloadInstruction struct{ singleByteArgumentInstruction }

func parseDloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &dloadInstruction{*toReturn}, nil
}

type aloadInstruction struct{ singleByteArgumentInstruction }

func parseAloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &aloadInstruction{*toReturn}, nil
}

type iload_0Instruction struct{ knownInstruction }

func parseIload_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iload_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1a,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_1Instruction struct{ knownInstruction }

func parseIload_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iload_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1b,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_2Instruction struct{ knownInstruction }

func parseIload_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iload_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1c,
			name: name,
		},
	}
	return &toReturn, nil
}

type iload_3Instruction struct{ knownInstruction }

func parseIload_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iload_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1d,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_0Instruction struct{ knownInstruction }

func parseLload_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lload_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1e,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_1Instruction struct{ knownInstruction }

func parseLload_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lload_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x1f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_2Instruction struct{ knownInstruction }

func parseLload_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lload_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x20,
			name: name,
		},
	}
	return &toReturn, nil
}

type lload_3Instruction struct{ knownInstruction }

func parseLload_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lload_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x21,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_0Instruction struct{ knownInstruction }

func parseFload_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fload_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x22,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_1Instruction struct{ knownInstruction }

func parseFload_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fload_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x23,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_2Instruction struct{ knownInstruction }

func parseFload_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fload_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x24,
			name: name,
		},
	}
	return &toReturn, nil
}

type fload_3Instruction struct{ knownInstruction }

func parseFload_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fload_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x25,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_0Instruction struct{ knownInstruction }

func parseDload_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dload_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x26,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_1Instruction struct{ knownInstruction }

func parseDload_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dload_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x27,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_2Instruction struct{ knownInstruction }

func parseDload_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dload_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x28,
			name: name,
		},
	}
	return &toReturn, nil
}

type dload_3Instruction struct{ knownInstruction }

func parseDload_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dload_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x29,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_0Instruction struct{ knownInstruction }

func parseAload_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aload_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x2a,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_1Instruction struct{ knownInstruction }

func parseAload_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aload_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x2b,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_2Instruction struct{ knownInstruction }

func parseAload_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aload_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x2c,
			name: name,
		},
	}
	return &toReturn, nil
}

type aload_3Instruction struct{ knownInstruction }

func parseAload_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aload_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x2d,
			name: name,
		},
	}
	return &toReturn, nil
}

type ialoadInstruction struct{ knownInstruction }

func parseIaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ialoadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x2e,
			name: name,
		},
	}
	return &toReturn, nil
}

type laloadInstruction struct{ knownInstruction }

func parseLaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := laloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x2f,
			name: name,
		},
	}
	return &toReturn, nil
}

type faloadInstruction struct{ knownInstruction }

func parseFaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := faloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x30,
			name: name,
		},
	}
	return &toReturn, nil
}

type daloadInstruction struct{ knownInstruction }

func parseDaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := daloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x31,
			name: name,
		},
	}
	return &toReturn, nil
}

type aaloadInstruction struct{ knownInstruction }

func parseAaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aaloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x32,
			name: name,
		},
	}
	return &toReturn, nil
}

type baloadInstruction struct{ knownInstruction }

func parseBaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := baloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x33,
			name: name,
		},
	}
	return &toReturn, nil
}

type caloadInstruction struct{ knownInstruction }

func parseCaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := caloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x34,
			name: name,
		},
	}
	return &toReturn, nil
}

type saloadInstruction struct{ knownInstruction }

func parseSaloadInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := saloadInstruction{
		knownInstruction: knownInstruction{
			raw:  0x35,
			name: name,
		},
	}
	return &toReturn, nil
}

type istoreInstruction struct{ singleByteArgumentInstruction }

func parseIstoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &istoreInstruction{*toReturn}, nil
}

type lstoreInstruction struct{ singleByteArgumentInstruction }

func parseLstoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &lstoreInstruction{*toReturn}, nil
}

type fstoreInstruction struct{ singleByteArgumentInstruction }

func parseFstoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &fstoreInstruction{*toReturn}, nil
}

type dstoreInstruction struct{ singleByteArgumentInstruction }

func parseDstoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &dstoreInstruction{*toReturn}, nil
}

type astoreInstruction struct{ singleByteArgumentInstruction }

func parseAstoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &astoreInstruction{*toReturn}, nil
}

type istore_0Instruction struct{ knownInstruction }

func parseIstore_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := istore_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x3b,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_1Instruction struct{ knownInstruction }

func parseIstore_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := istore_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x3c,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_2Instruction struct{ knownInstruction }

func parseIstore_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := istore_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x3d,
			name: name,
		},
	}
	return &toReturn, nil
}

type istore_3Instruction struct{ knownInstruction }

func parseIstore_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := istore_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x3e,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_0Instruction struct{ knownInstruction }

func parseLstore_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lstore_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x3f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_1Instruction struct{ knownInstruction }

func parseLstore_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lstore_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x40,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_2Instruction struct{ knownInstruction }

func parseLstore_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lstore_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x41,
			name: name,
		},
	}
	return &toReturn, nil
}

type lstore_3Instruction struct{ knownInstruction }

func parseLstore_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lstore_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x42,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_0Instruction struct{ knownInstruction }

func parseFstore_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fstore_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x43,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_1Instruction struct{ knownInstruction }

func parseFstore_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fstore_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x44,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_2Instruction struct{ knownInstruction }

func parseFstore_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fstore_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x45,
			name: name,
		},
	}
	return &toReturn, nil
}

type fstore_3Instruction struct{ knownInstruction }

func parseFstore_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fstore_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x46,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_0Instruction struct{ knownInstruction }

func parseDstore_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dstore_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x47,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_1Instruction struct{ knownInstruction }

func parseDstore_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dstore_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x48,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_2Instruction struct{ knownInstruction }

func parseDstore_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dstore_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x49,
			name: name,
		},
	}
	return &toReturn, nil
}

type dstore_3Instruction struct{ knownInstruction }

func parseDstore_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dstore_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x4a,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_0Instruction struct{ knownInstruction }

func parseAstore_0Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := astore_0Instruction{
		knownInstruction: knownInstruction{
			raw:  0x4b,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_1Instruction struct{ knownInstruction }

func parseAstore_1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := astore_1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x4c,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_2Instruction struct{ knownInstruction }

func parseAstore_2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := astore_2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x4d,
			name: name,
		},
	}
	return &toReturn, nil
}

type astore_3Instruction struct{ knownInstruction }

func parseAstore_3Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := astore_3Instruction{
		knownInstruction: knownInstruction{
			raw:  0x4e,
			name: name,
		},
	}
	return &toReturn, nil
}

type iastoreInstruction struct{ knownInstruction }

func parseIastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x4f,
			name: name,
		},
	}
	return &toReturn, nil
}

type lastoreInstruction struct{ knownInstruction }

func parseLastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x50,
			name: name,
		},
	}
	return &toReturn, nil
}

type fastoreInstruction struct{ knownInstruction }

func parseFastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x51,
			name: name,
		},
	}
	return &toReturn, nil
}

type dastoreInstruction struct{ knownInstruction }

func parseDastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x52,
			name: name,
		},
	}
	return &toReturn, nil
}

type aastoreInstruction struct{ knownInstruction }

func parseAastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := aastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x53,
			name: name,
		},
	}
	return &toReturn, nil
}

type bastoreInstruction struct{ knownInstruction }

func parseBastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := bastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x54,
			name: name,
		},
	}
	return &toReturn, nil
}

type castoreInstruction struct{ knownInstruction }

func parseCastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := castoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x55,
			name: name,
		},
	}
	return &toReturn, nil
}

type sastoreInstruction struct{ knownInstruction }

func parseSastoreInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := sastoreInstruction{
		knownInstruction: knownInstruction{
			raw:  0x56,
			name: name,
		},
	}
	return &toReturn, nil
}

type popInstruction struct{ knownInstruction }

func parsePopInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := popInstruction{
		knownInstruction: knownInstruction{
			raw:  0x57,
			name: name,
		},
	}
	return &toReturn, nil
}

type pop2Instruction struct{ knownInstruction }

func parsePop2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := pop2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x58,
			name: name,
		},
	}
	return &toReturn, nil
}

type dupInstruction struct{ knownInstruction }

func parseDupInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dupInstruction{
		knownInstruction: knownInstruction{
			raw:  0x59,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup_x1Instruction struct{ knownInstruction }

func parseDup_x1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dup_x1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x5a,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup_x2Instruction struct{ knownInstruction }

func parseDup_x2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dup_x2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x5b,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2Instruction struct{ knownInstruction }

func parseDup2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dup2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x5c,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2_x1Instruction struct{ knownInstruction }

func parseDup2_x1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dup2_x1Instruction{
		knownInstruction: knownInstruction{
			raw:  0x5d,
			name: name,
		},
	}
	return &toReturn, nil
}

type dup2_x2Instruction struct{ knownInstruction }

func parseDup2_x2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dup2_x2Instruction{
		knownInstruction: knownInstruction{
			raw:  0x5e,
			name: name,
		},
	}
	return &toReturn, nil
}

type swapInstruction struct{ knownInstruction }

func parseSwapInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := swapInstruction{
		knownInstruction: knownInstruction{
			raw:  0x5f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iaddInstruction struct{ knownInstruction }

func parseIaddInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iaddInstruction{
		knownInstruction: knownInstruction{
			raw:  0x60,
			name: name,
		},
	}
	return &toReturn, nil
}

type laddInstruction struct{ knownInstruction }

func parseLaddInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := laddInstruction{
		knownInstruction: knownInstruction{
			raw:  0x61,
			name: name,
		},
	}
	return &toReturn, nil
}

type faddInstruction struct{ knownInstruction }

func parseFaddInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := faddInstruction{
		knownInstruction: knownInstruction{
			raw:  0x62,
			name: name,
		},
	}
	return &toReturn, nil
}

type daddInstruction struct{ knownInstruction }

func parseDaddInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := daddInstruction{
		knownInstruction: knownInstruction{
			raw:  0x63,
			name: name,
		},
	}
	return &toReturn, nil
}

type isubInstruction struct{ knownInstruction }

func parseIsubInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := isubInstruction{
		knownInstruction: knownInstruction{
			raw:  0x64,
			name: name,
		},
	}
	return &toReturn, nil
}

type lsubInstruction struct{ knownInstruction }

func parseLsubInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lsubInstruction{
		knownInstruction: knownInstruction{
			raw:  0x65,
			name: name,
		},
	}
	return &toReturn, nil
}

type fsubInstruction struct{ knownInstruction }

func parseFsubInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fsubInstruction{
		knownInstruction: knownInstruction{
			raw:  0x66,
			name: name,
		},
	}
	return &toReturn, nil
}

type dsubInstruction struct{ knownInstruction }

func parseDsubInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dsubInstruction{
		knownInstruction: knownInstruction{
			raw:  0x67,
			name: name,
		},
	}
	return &toReturn, nil
}

type imulInstruction struct{ knownInstruction }

func parseImulInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := imulInstruction{
		knownInstruction: knownInstruction{
			raw:  0x68,
			name: name,
		},
	}
	return &toReturn, nil
}

type lmulInstruction struct{ knownInstruction }

func parseLmulInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lmulInstruction{
		knownInstruction: knownInstruction{
			raw:  0x69,
			name: name,
		},
	}
	return &toReturn, nil
}

type fmulInstruction struct{ knownInstruction }

func parseFmulInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fmulInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6a,
			name: name,
		},
	}
	return &toReturn, nil
}

type dmulInstruction struct{ knownInstruction }

func parseDmulInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dmulInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6b,
			name: name,
		},
	}
	return &toReturn, nil
}

type idivInstruction struct{ knownInstruction }

func parseIdivInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := idivInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6c,
			name: name,
		},
	}
	return &toReturn, nil
}

type ldivInstruction struct{ knownInstruction }

func parseLdivInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ldivInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6d,
			name: name,
		},
	}
	return &toReturn, nil
}

type fdivInstruction struct{ knownInstruction }

func parseFdivInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fdivInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6e,
			name: name,
		},
	}
	return &toReturn, nil
}

type ddivInstruction struct{ knownInstruction }

func parseDdivInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ddivInstruction{
		knownInstruction: knownInstruction{
			raw:  0x6f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iremInstruction struct{ knownInstruction }

func parseIremInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iremInstruction{
		knownInstruction: knownInstruction{
			raw:  0x70,
			name: name,
		},
	}
	return &toReturn, nil
}

type lremInstruction struct{ knownInstruction }

func parseLremInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lremInstruction{
		knownInstruction: knownInstruction{
			raw:  0x71,
			name: name,
		},
	}
	return &toReturn, nil
}

type fremInstruction struct{ knownInstruction }

func parseFremInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fremInstruction{
		knownInstruction: knownInstruction{
			raw:  0x72,
			name: name,
		},
	}
	return &toReturn, nil
}

type dremInstruction struct{ knownInstruction }

func parseDremInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dremInstruction{
		knownInstruction: knownInstruction{
			raw:  0x73,
			name: name,
		},
	}
	return &toReturn, nil
}

type inegInstruction struct{ knownInstruction }

func parseInegInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := inegInstruction{
		knownInstruction: knownInstruction{
			raw:  0x74,
			name: name,
		},
	}
	return &toReturn, nil
}

type lnegInstruction struct{ knownInstruction }

func parseLnegInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lnegInstruction{
		knownInstruction: knownInstruction{
			raw:  0x75,
			name: name,
		},
	}
	return &toReturn, nil
}

type fnegInstruction struct{ knownInstruction }

func parseFnegInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fnegInstruction{
		knownInstruction: knownInstruction{
			raw:  0x76,
			name: name,
		},
	}
	return &toReturn, nil
}

type dnegInstruction struct{ knownInstruction }

func parseDnegInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dnegInstruction{
		knownInstruction: knownInstruction{
			raw:  0x77,
			name: name,
		},
	}
	return &toReturn, nil
}

type ishlInstruction struct{ knownInstruction }

func parseIshlInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ishlInstruction{
		knownInstruction: knownInstruction{
			raw:  0x78,
			name: name,
		},
	}
	return &toReturn, nil
}

type lshlInstruction struct{ knownInstruction }

func parseLshlInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lshlInstruction{
		knownInstruction: knownInstruction{
			raw:  0x79,
			name: name,
		},
	}
	return &toReturn, nil
}

type ishrInstruction struct{ knownInstruction }

func parseIshrInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ishrInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7a,
			name: name,
		},
	}
	return &toReturn, nil
}

type lshrInstruction struct{ knownInstruction }

func parseLshrInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lshrInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7b,
			name: name,
		},
	}
	return &toReturn, nil
}

type iushrInstruction struct{ knownInstruction }

func parseIushrInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iushrInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7c,
			name: name,
		},
	}
	return &toReturn, nil
}

type lushrInstruction struct{ knownInstruction }

func parseLushrInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lushrInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7d,
			name: name,
		},
	}
	return &toReturn, nil
}

type iandInstruction struct{ knownInstruction }

func parseIandInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iandInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7e,
			name: name,
		},
	}
	return &toReturn, nil
}

type landInstruction struct{ knownInstruction }

func parseLandInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := landInstruction{
		knownInstruction: knownInstruction{
			raw:  0x7f,
			name: name,
		},
	}
	return &toReturn, nil
}

type iorInstruction struct{ knownInstruction }

func parseIorInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := iorInstruction{
		knownInstruction: knownInstruction{
			raw:  0x80,
			name: name,
		},
	}
	return &toReturn, nil
}

type lorInstruction struct{ knownInstruction }

func parseLorInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lorInstruction{
		knownInstruction: knownInstruction{
			raw:  0x81,
			name: name,
		},
	}
	return &toReturn, nil
}

type ixorInstruction struct{ knownInstruction }

func parseIxorInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ixorInstruction{
		knownInstruction: knownInstruction{
			raw:  0x82,
			name: name,
		},
	}
	return &toReturn, nil
}

type lxorInstruction struct{ knownInstruction }

func parseLxorInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lxorInstruction{
		knownInstruction: knownInstruction{
			raw:  0x83,
			name: name,
		},
	}
	return &toReturn, nil
}

// The iinc instruction is a fairly unique format, so it gets its own struct
type iincInstruction struct {
	index uint8
	value uint8
}

func (n *iincInstruction) Raw() uint8 {
	return 0x84
}

func (n *iincInstruction) OtherBytes() []byte {
	return []byte{n.index, n.value}
}

func (n *iincInstruction) Length() uint {
	return 3
}

func (n *iincInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *iincInstruction) String() string {
	return fmt.Sprintf("iinc 0x%02x 0x%02x", n.index, n.value)
}

func parseIincInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	index, e := m.GetByte(address + 1)
	if e != nil {
		return nil, fmt.Errorf("Failed getting iinc offset: %s", e)
	}
	value, e := m.GetByte(address + 2)
	if e != nil {
		return nil, fmt.Errorf("Failed getting iinc value: %s", e)
	}
	toReturn := iincInstruction{
		index: index,
		value: value,
	}
	return &toReturn, nil
}

type i2lInstruction struct{ knownInstruction }

func parseI2lInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2lInstruction{
		knownInstruction: knownInstruction{
			raw:  0x85,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2fInstruction struct{ knownInstruction }

func parseI2fInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2fInstruction{
		knownInstruction: knownInstruction{
			raw:  0x86,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2dInstruction struct{ knownInstruction }

func parseI2dInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2dInstruction{
		knownInstruction: knownInstruction{
			raw:  0x87,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2iInstruction struct{ knownInstruction }

func parseL2iInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := l2iInstruction{
		knownInstruction: knownInstruction{
			raw:  0x88,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2fInstruction struct{ knownInstruction }

func parseL2fInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := l2fInstruction{
		knownInstruction: knownInstruction{
			raw:  0x89,
			name: name,
		},
	}
	return &toReturn, nil
}

type l2dInstruction struct{ knownInstruction }

func parseL2dInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := l2dInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8a,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2iInstruction struct{ knownInstruction }

func parseF2iInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := f2iInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8b,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2lInstruction struct{ knownInstruction }

func parseF2lInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := f2lInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8c,
			name: name,
		},
	}
	return &toReturn, nil
}

type f2dInstruction struct{ knownInstruction }

func parseF2dInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := f2dInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8d,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2iInstruction struct{ knownInstruction }

func parseD2iInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := d2iInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8e,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2lInstruction struct{ knownInstruction }

func parseD2lInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := d2lInstruction{
		knownInstruction: knownInstruction{
			raw:  0x8f,
			name: name,
		},
	}
	return &toReturn, nil
}

type d2fInstruction struct{ knownInstruction }

func parseD2fInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := d2fInstruction{
		knownInstruction: knownInstruction{
			raw:  0x90,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2bInstruction struct{ knownInstruction }

func parseI2bInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2bInstruction{
		knownInstruction: knownInstruction{
			raw:  0x91,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2cInstruction struct{ knownInstruction }

func parseI2cInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2cInstruction{
		knownInstruction: knownInstruction{
			raw:  0x92,
			name: name,
		},
	}
	return &toReturn, nil
}

type i2sInstruction struct{ knownInstruction }

func parseI2sInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := i2sInstruction{
		knownInstruction: knownInstruction{
			raw:  0x93,
			name: name,
		},
	}
	return &toReturn, nil
}

type lcmpInstruction struct{ knownInstruction }

func parseLcmpInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lcmpInstruction{
		knownInstruction: knownInstruction{
			raw:  0x94,
			name: name,
		},
	}
	return &toReturn, nil
}

type fcmplInstruction struct{ knownInstruction }

func parseFcmplInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fcmplInstruction{
		knownInstruction: knownInstruction{
			raw:  0x95,
			name: name,
		},
	}
	return &toReturn, nil
}

type fcmpgInstruction struct{ knownInstruction }

func parseFcmpgInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := fcmpgInstruction{
		knownInstruction: knownInstruction{
			raw:  0x96,
			name: name,
		},
	}
	return &toReturn, nil
}

type dcmplInstruction struct{ knownInstruction }

func parseDcmplInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dcmplInstruction{
		knownInstruction: knownInstruction{
			raw:  0x97,
			name: name,
		},
	}
	return &toReturn, nil
}

type dcmpgInstruction struct{ knownInstruction }

func parseDcmpgInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dcmpgInstruction{
		knownInstruction: knownInstruction{
			raw:  0x98,
			name: name,
		},
	}
	return &toReturn, nil
}

type ifeqInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfeqInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifeqInstruction{*toReturn, 0}, nil
}

type ifneInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfneInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifneInstruction{*toReturn, 0}, nil
}

type ifltInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfltInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifltInstruction{*toReturn, 0}, nil
}

type ifgeInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfgeInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifgeInstruction{*toReturn, 0}, nil
}

type ifgtInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfgtInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifgtInstruction{*toReturn, 0}, nil
}

type ifleInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIfleInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifleInstruction{*toReturn, 0}, nil
}

type if_icmpeqInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpeqInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpeqInstruction{*toReturn, 0}, nil
}

type if_icmpneInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpneInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpneInstruction{*toReturn, 0}, nil
}

type if_icmpltInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpltInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpltInstruction{*toReturn, 0}, nil
}

type if_icmpgeInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpgeInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpgeInstruction{*toReturn, 0}, nil
}

type if_icmpgtInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpgtInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpgtInstruction{*toReturn, 0}, nil
}

type if_icmpleInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_icmpleInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_icmpleInstruction{*toReturn, 0}, nil
}

type if_acmpeqInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_acmpeqInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_acmpeqInstruction{*toReturn, 0}, nil
}

type if_acmpneInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseIf_acmpneInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &if_acmpneInstruction{*toReturn, 0}, nil
}

type gotoInstruction struct {
	twoByteArgumentInstruction
	nextIndex uint
}

func parseGotoInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &gotoInstruction{*toReturn, 0}, nil
}

type jsrInstruction struct {
	twoByteArgumentInstruction
	// This is the instruction index of the subroutine start.
	nextIndex uint
	// Our "return address" type is an instruction *index*, not a byte offset.
	returnIndex int
}

func parseJsrInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &jsrInstruction{*toReturn, 0, 0}, nil
}

type retInstruction struct{ singleByteArgumentInstruction }

func parseRetInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
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
	skippedCount  uint8
	defaultOffset uint32
	lowIndex      uint32
	highIndex     uint32
	offsets       []uint32
	// The same as defaultOffset and offsets, but converted to instruction
	// indices within the current method.
	defaultIndex uint
	indices      []uint
}

func (n *tableswitchInstruction) Raw() uint8 {
	return 0xaa
}

func (n *tableswitchInstruction) Length() uint {
	// 12 bytes for high, low and default offsets, plus one byte for the
	// opcode, 4 per offset in the list, and up to 3 skipped padding bytes.
	// This is called by OtherBytes() to allocate a buffer, so it must not
	// depend on OtherBytes().
	return uint(n.skippedCount) + uint(len(n.offsets)*4) + 13
}

func (n *tableswitchInstruction) OtherBytes() []byte {
	toReturn := make([]byte, n.Length()-1)
	offset := 0
	// Use this inner function for convenience, and allowing us to avoid
	// encoding/binary.
	// TODO: Test OtherBytes() for tableswitchInstruction
	writeValueToBuffer := func(value uint32) {
		toReturn[offset] = uint8(value >> 24)
		toReturn[offset+1] = uint8(value >> 16)
		toReturn[offset+2] = uint8(value >> 8)
		toReturn[offset+3] = uint8(value)
		offset += 4
	}
	for i := uint8(0); i < n.skippedCount; i++ {
		toReturn[offset] = 0
		offset++
	}
	writeValueToBuffer(n.highIndex)
	writeValueToBuffer(n.lowIndex)
	writeValueToBuffer(n.defaultOffset)
	for _, v := range n.offsets {
		writeValueToBuffer(v)
	}
	return toReturn
}

func (n *tableswitchInstruction) String() string {
	return fmt.Sprintf("tableswitch 0x%08x-0x%08x (default 0x%08x)",
		n.lowIndex, n.highIndex, n.defaultOffset)
}

func parseTableswitchInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	var e error
	var toReturn tableswitchInstruction
	currentOffset := address + 1
	// First, read up to 3 bytes to get to a 4-byte aligned address
	toReturn.skippedCount = uint8(currentOffset % 4)
	currentOffset += uint(toReturn.skippedCount)
	toReturn.defaultOffset, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch default: %s", e)
	}
	currentOffset += 4
	toReturn.lowIndex, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch low offset: %s", e)
	}
	currentOffset += 4
	toReturn.highIndex, e = Read32Bits(m, currentOffset)
	if e != nil {
		return nil, fmt.Errorf("Failed reading tableswitch high offset: %s", e)
	}
	currentOffset += 4
	if toReturn.highIndex < toReturn.lowIndex {
		return nil, fmt.Errorf("Tableswitch offset range invalid")
	}
	offsetsCount := toReturn.highIndex - toReturn.lowIndex + 1
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
	// Holds instruction indices rather than offsets. The indices array is in
	// the same order as the pairs array, but doesn't contain the value.
	defaultIndex uint
	indices      []uint
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
	m Memory) (Instruction, error) {
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

type ireturnInstruction struct{ knownInstruction }

func parseIreturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := ireturnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xac,
			name: name,
		},
	}
	return &toReturn, nil
}

type lreturnInstruction struct{ knownInstruction }

func parseLreturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := lreturnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xad,
			name: name,
		},
	}
	return &toReturn, nil
}

type freturnInstruction struct{ knownInstruction }

func parseFreturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := freturnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xae,
			name: name,
		},
	}
	return &toReturn, nil
}

type dreturnInstruction struct{ knownInstruction }

func parseDreturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := dreturnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xaf,
			name: name,
		},
	}
	return &toReturn, nil
}

type areturnInstruction struct{ knownInstruction }

func parseAreturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := areturnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xb0,
			name: name,
		},
	}
	return &toReturn, nil
}

type returnInstruction struct{ knownInstruction }

func parseReturnInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := returnInstruction{
		knownInstruction: knownInstruction{
			raw:  0xb1,
			name: name,
		},
	}
	return &toReturn, nil
}

type getstaticInstruction struct {
	twoByteArgumentInstruction
	// The class containing the static field to be accessed.
	class *Class
	// The index into the class' StaticFieldValues array.
	index int
}

func (n *getstaticInstruction) String() string {
	if n.class == nil {
		return fmt.Sprintf("getstatic %d", n.value)
	}
	fieldName := n.class.StaticFieldNames[n.index]
	return fmt.Sprintf("getstatic %s.%s", n.class.Name, fieldName)
}

func parseGetstaticInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &getstaticInstruction{*toReturn, nil, 0}, nil
}

type putstaticInstruction struct {
	twoByteArgumentInstruction
	// The class containing the static field to be accessed.
	class *Class
	// The index into the class' StaticFieldValues array.
	index int
}

func parsePutstaticInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &putstaticInstruction{*toReturn, nil, 0}, nil
}

func (n *putstaticInstruction) String() string {
	if n.class == nil {
		return fmt.Sprintf("putstatic %d", n.value)
	}
	fieldName := n.class.StaticFieldNames[n.index]
	return fmt.Sprintf("putstatic %s.%s", n.class.Name, fieldName)
}

type getfieldInstruction struct {
	twoByteArgumentInstruction
	// Unlike getstatic, we can't figure out a field's index until runtime,
	// when we know what object is actually on the stack. (Technically we
	// *could* try, but doing so would rely on better type checking so it will
	// be easier this way for now.)
	fieldReference *FieldOrMethodReference
}

func parseGetfieldInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &getfieldInstruction{*toReturn, nil}, nil
}

func (n *getfieldInstruction) String() string {
	if n.fieldReference != nil {
		return "getfield " + string(n.fieldReference.Field.Name)
	}
	return fmt.Sprintf("getfield %d", n.value)
}

type putfieldInstruction struct {
	twoByteArgumentInstruction
	fieldReference *FieldOrMethodReference
}

func parsePutfieldInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &putfieldInstruction{*toReturn, nil}, nil
}

func (n *putfieldInstruction) String() string {
	if n.fieldReference != nil {
		return "getfield " + string(n.fieldReference.Field.Name)
	}
	return fmt.Sprintf("getfield %d", n.value)
}

type invokevirtualInstruction struct{ twoByteArgumentInstruction }

func parseInvokevirtualInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokevirtualInstruction{*toReturn}, nil
}

type invokespecialInstruction struct {
	twoByteArgumentInstruction
	// The method to be invoked.
	method *Method
}

func parseInvokespecialInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokespecialInstruction{
		twoByteArgumentInstruction: *toReturn,
		method:                     nil,
	}, nil
}

func (n *invokespecialInstruction) String() string {
	if n.method == nil {
		return fmt.Sprintf("invokespecial %d", n.value)
	}
	m := n.method
	return fmt.Sprintf("invokespecial %s %s.%s(%s)",
		m.Types.ReturnString(), n.method.ContainingClass.Name, m.Name,
		m.Types.ArgumentsString())
}

type invokestaticInstruction struct{ twoByteArgumentInstruction }

func parseInvokestaticInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
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
	m Memory) (Instruction, error) {
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
		count:                      count,
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
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &invokedynamicInstruction{*toReturn}, nil
}

type newInstruction struct {
	twoByteArgumentInstruction
	// The class to instantiate
	class *Class
}

func parseNewInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &newInstruction{
		twoByteArgumentInstruction: *toReturn,
		class:                      nil,
	}, nil
}

func (n *newInstruction) String() string {
	if n.class == nil {
		return fmt.Sprintf("new %d", n.value)
	}
	return fmt.Sprintf("new %s", n.class.Name)
}

type newarrayInstruction struct{ singleByteArgumentInstruction }

func parseNewarrayInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseSingleByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &newarrayInstruction{*toReturn}, nil
}

type anewarrayInstruction struct{ twoByteArgumentInstruction }

func parseAnewarrayInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &anewarrayInstruction{*toReturn}, nil
}

type arraylengthInstruction struct{ knownInstruction }

func parseArraylengthInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := arraylengthInstruction{
		knownInstruction: knownInstruction{
			raw:  0xbe,
			name: name,
		},
	}
	return &toReturn, nil
}

type athrowInstruction struct{ knownInstruction }

func parseAthrowInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := athrowInstruction{
		knownInstruction: knownInstruction{
			raw:  0xbf,
			name: name,
		},
	}
	return &toReturn, nil
}

type checkcastInstruction struct{ twoByteArgumentInstruction }

func parseCheckcastInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &checkcastInstruction{*toReturn}, nil
}

type instanceofInstruction struct{ twoByteArgumentInstruction }

func parseInstanceofInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &instanceofInstruction{*toReturn}, nil
}

type monitorenterInstruction struct{ knownInstruction }

func parseMonitorenterInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := monitorenterInstruction{
		knownInstruction: knownInstruction{
			raw:  0xc2,
			name: name,
		},
	}
	return &toReturn, nil
}

type monitorexitInstruction struct{ knownInstruction }

func parseMonitorexitInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := monitorexitInstruction{
		knownInstruction: knownInstruction{
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

func (n *wideInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
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

func (n *wideIincInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *wideIincInstruction) String() string {
	return fmt.Sprintf("wide iinc 0x%04x 0x%04x", n.index, n.value)
}

// This only parses "wide iinc ..." instructions, and will be called by
// parseWideInstruction
func parseWideIincInstruction(address uint, m Memory) (
	Instruction, error) {
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
	m Memory) (Instruction, error) {
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

func (n *multianewarrayInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *multianewarrayInstruction) String() string {
	return fmt.Sprintf("multianewarray 0x%04x %d", n.typeIndex, n.dimensions)
}

func parseMultianewarrayInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
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
	m Memory) (Instruction, error) {
	toReturn, e := parseTwoByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &ifnullInstruction{*toReturn}, nil
}

type ifnonnullInstruction struct{ twoByteArgumentInstruction }

func parseIfnonnullInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
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

func (n *fourByteArgumentInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	return nil
}

func (n *fourByteArgumentInstruction) String() string {
	return fmt.Sprintf("%s 0x%08x", n.name, n.value)
}

func parseFourByteArgumentInstruction(opcode uint8, name string, address uint,
	m Memory) (*fourByteArgumentInstruction, error) {
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
	m Memory) (Instruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &goto_wInstruction{*toReturn}, nil
}

type jsr_wInstruction struct{ fourByteArgumentInstruction }

func parseJsr_wInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn, e := parseFourByteArgumentInstruction(opcode, name, address, m)
	if e != nil {
		return nil, e
	}
	return &jsr_wInstruction{*toReturn}, nil
}

type breakpointInstruction struct{ knownInstruction }

func parseBreakpointInstruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := breakpointInstruction{
		knownInstruction: knownInstruction{
			raw:  0xca,
			name: name,
		},
	}
	return &toReturn, nil
}

type impdep1Instruction struct{ knownInstruction }

func parseImpdep1Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := impdep1Instruction{
		knownInstruction: knownInstruction{
			raw:  0xfe,
			name: name,
		},
	}
	return &toReturn, nil
}

type impdep2Instruction struct{ knownInstruction }

func parseImpdep2Instruction(opcode uint8, name string, address uint,
	m Memory) (Instruction, error) {
	toReturn := impdep2Instruction{
		knownInstruction: knownInstruction{
			raw:  0xff,
			name: name,
		},
	}
	return &toReturn, nil
}
