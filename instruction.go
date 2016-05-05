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
	// nil if the instruction doesn't have such bytes.
	OtherBytes() []byte
	// Runs the instruction in the given thread
	Execute(t JVMThread) error
	// Returns the disassembly string of the instruction
	String() string
}

// The size, in bytes, of the given JVMInstruction. This is the amount the pc
// should be advanced to get to the next instruction.
func JVMInstructionLength(n JVMInstruction) uint {
	return uint(1 + len(n.OtherBytes))
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

func (n *basicJVMInstruction) Execute(t JVMThread) error {
	return UnknownInstructionError(n.raw)
}

func (n *basicJVMInstruction) String() string {
	return fmt.Sprintf("<unknown instruction 0x%08x>", n.raw)
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
func GetNextInstruction(m JVMMemory, address uint) (JVMInstruction, error) {
	firstByte, e := m.GetByte(address)
	if e != nil {
		return e
	}
	opcodeInfo := opcodeTable[firstByte]
	// Unknown instruction.
	if opcodeInfo == nil {
		toReturn := &basicJVMInstruction{
			raw: firstByte,
		}
		return toReturn, nil
	}
}

type nopInstruction knownJVMInstruction

func parseNopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := nopInstruction{
		basicJVMInstruction{
			raw: 0x00,
		},
		name: "nop",
	}
	return &toReturn, nil
}

type aconst_nullInstruction knownJVMInstruction

func parseAconst_nullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aconst_nullInstruction{
		basicJVMInstruction{
			raw: 0x01,
		},
		name: "aconst_null",
	}
	return &toReturn, nil
}

type iconst_m1Instruction knownJVMInstruction

func parseIconst_m1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_m1Instruction{
		basicJVMInstruction{
			raw: 0x02,
		},
		name: "iconst_m1",
	}
	return &toReturn, nil
}

type iconst_0Instruction knownJVMInstruction

func parseIconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_0Instruction{
		basicJVMInstruction{
			raw: 0x03,
		},
		name: "iconst_0",
	}
	return &toReturn, nil
}

type iconst_1Instruction knownJVMInstruction

func parseIconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_1Instruction{
		basicJVMInstruction{
			raw: 0x04,
		},
		name: "iconst_1",
	}
	return &toReturn, nil
}

type iconst_2Instruction knownJVMInstruction

func parseIconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_2Instruction{
		basicJVMInstruction{
			raw: 0x05,
		},
		name: "iconst_2",
	}
	return &toReturn, nil
}

type iconst_3Instruction knownJVMInstruction

func parseIconst_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_3Instruction{
		basicJVMInstruction{
			raw: 0x06,
		},
		name: "iconst_3",
	}
	return &toReturn, nil
}

type iconst_4Instruction knownJVMInstruction

func parseIconst_4Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_4Instruction{
		basicJVMInstruction{
			raw: 0x07,
		},
		name: "iconst_4",
	}
	return &toReturn, nil
}

type iconst_5Instruction knownJVMInstruction

func parseIconst_5Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iconst_5Instruction{
		basicJVMInstruction{
			raw: 0x08,
		},
		name: "iconst_5",
	}
	return &toReturn, nil
}

type lconst_0Instruction knownJVMInstruction

func parseLconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_0Instruction{
		basicJVMInstruction{
			raw: 0x09,
		},
		name: "lconst_0",
	}
	return &toReturn, nil
}

type lconst_1Instruction knownJVMInstruction

func parseLconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lconst_1Instruction{
		basicJVMInstruction{
			raw: 0x0a,
		},
		name: "lconst_1",
	}
	return &toReturn, nil
}

type fconst_0Instruction knownJVMInstruction

func parseFconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_0Instruction{
		basicJVMInstruction{
			raw: 0x0b,
		},
		name: "fconst_0",
	}
	return &toReturn, nil
}

type fconst_1Instruction knownJVMInstruction

func parseFconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_1Instruction{
		basicJVMInstruction{
			raw: 0x0c,
		},
		name: "fconst_1",
	}
	return &toReturn, nil
}

type fconst_2Instruction knownJVMInstruction

func parseFconst_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fconst_2Instruction{
		basicJVMInstruction{
			raw: 0x0d,
		},
		name: "fconst_2",
	}
	return &toReturn, nil
}

type dconst_0Instruction knownJVMInstruction

func parseDconst_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_0Instruction{
		basicJVMInstruction{
			raw: 0x0e,
		},
		name: "dconst_0",
	}
	return &toReturn, nil
}

type dconst_1Instruction knownJVMInstruction

func parseDconst_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dconst_1Instruction{
		basicJVMInstruction{
			raw: 0x0f,
		},
		name: "dconst_1",
	}
	return &toReturn, nil
}

func parseBipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseSipushInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLdcInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLdc_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLdc2_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseFloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseDloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseAloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type iload_0Instruction knownJVMInstruction

func parseIload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_0Instruction{
		basicJVMInstruction{
			raw: 0x1a,
		},
		name: "iload_0",
	}
	return &toReturn, nil
}

type iload_1Instruction knownJVMInstruction

func parseIload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_1Instruction{
		basicJVMInstruction{
			raw: 0x1b,
		},
		name: "iload_1",
	}
	return &toReturn, nil
}

type iload_2Instruction knownJVMInstruction

func parseIload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_2Instruction{
		basicJVMInstruction{
			raw: 0x1c,
		},
		name: "iload_2",
	}
	return &toReturn, nil
}

type iload_3Instruction knownJVMInstruction

func parseIload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iload_3Instruction{
		basicJVMInstruction{
			raw: 0x1d,
		},
		name: "iload_3",
	}
	return &toReturn, nil
}

type lload_0Instruction knownJVMInstruction

func parseLload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_0Instruction{
		basicJVMInstruction{
			raw: 0x1e,
		},
		name: "lload_0",
	}
	return &toReturn, nil
}

type lload_1Instruction knownJVMInstruction

func parseLload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_1Instruction{
		basicJVMInstruction{
			raw: 0x1f,
		},
		name: "lload_1",
	}
	return &toReturn, nil
}

type lload_2Instruction knownJVMInstruction

func parseLload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_2Instruction{
		basicJVMInstruction{
			raw: 0x20,
		},
		name: "lload_2",
	}
	return &toReturn, nil
}

type lload_3Instruction knownJVMInstruction

func parseLload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lload_3Instruction{
		basicJVMInstruction{
			raw: 0x21,
		},
		name: "lload_3",
	}
	return &toReturn, nil
}

type fload_0Instruction knownJVMInstruction

func parseFload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_0Instruction{
		basicJVMInstruction{
			raw: 0x22,
		},
		name: "fload_0",
	}
	return &toReturn, nil
}

type fload_1Instruction knownJVMInstruction

func parseFload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_1Instruction{
		basicJVMInstruction{
			raw: 0x23,
		},
		name: "fload_1",
	}
	return &toReturn, nil
}

type fload_2Instruction knownJVMInstruction

func parseFload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_2Instruction{
		basicJVMInstruction{
			raw: 0x24,
		},
		name: "fload_2",
	}
	return &toReturn, nil
}

type fload_3Instruction knownJVMInstruction

func parseFload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fload_3Instruction{
		basicJVMInstruction{
			raw: 0x25,
		},
		name: "fload_3",
	}
	return &toReturn, nil
}

type dload_0Instruction knownJVMInstruction

func parseDload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_0Instruction{
		basicJVMInstruction{
			raw: 0x26,
		},
		name: "dload_0",
	}
	return &toReturn, nil
}

type dload_1Instruction knownJVMInstruction

func parseDload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_1Instruction{
		basicJVMInstruction{
			raw: 0x27,
		},
		name: "dload_1",
	}
	return &toReturn, nil
}

type dload_2Instruction knownJVMInstruction

func parseDload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_2Instruction{
		basicJVMInstruction{
			raw: 0x28,
		},
		name: "dload_2",
	}
	return &toReturn, nil
}

type dload_3Instruction knownJVMInstruction

func parseDload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dload_3Instruction{
		basicJVMInstruction{
			raw: 0x29,
		},
		name: "dload_3",
	}
	return &toReturn, nil
}

type aload_0Instruction knownJVMInstruction

func parseAload_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_0Instruction{
		basicJVMInstruction{
			raw: 0x2a,
		},
		name: "aload_0",
	}
	return &toReturn, nil
}

type aload_1Instruction knownJVMInstruction

func parseAload_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_1Instruction{
		basicJVMInstruction{
			raw: 0x2b,
		},
		name: "aload_1",
	}
	return &toReturn, nil
}

type aload_2Instruction knownJVMInstruction

func parseAload_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_2Instruction{
		basicJVMInstruction{
			raw: 0x2c,
		},
		name: "aload_2",
	}
	return &toReturn, nil
}

type aload_3Instruction knownJVMInstruction

func parseAload_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aload_3Instruction{
		basicJVMInstruction{
			raw: 0x2d,
		},
		name: "aload_3",
	}
	return &toReturn, nil
}

type ialoadInstruction knownJVMInstruction

func parseIaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ialoadInstruction{
		basicJVMInstruction{
			raw: 0x2e,
		},
		name: "iaload",
	}
	return &toReturn, nil
}

type laloadInstruction knownJVMInstruction

func parseLaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laloadInstruction{
		basicJVMInstruction{
			raw: 0x2f,
		},
		name: "laload",
	}
	return &toReturn, nil
}

type faloadInstruction knownJVMInstruction

func parseFaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faloadInstruction{
		basicJVMInstruction{
			raw: 0x30,
		},
		name: "faload",
	}
	return &toReturn, nil
}

type daloadInstruction knownJVMInstruction

func parseDaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daloadInstruction{
		basicJVMInstruction{
			raw: 0x31,
		},
		name: "daload",
	}
	return &toReturn, nil
}

type aaloadInstruction knownJVMInstruction

func parseAaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aaloadInstruction{
		basicJVMInstruction{
			raw: 0x32,
		},
		name: "aaload",
	}
	return &toReturn, nil
}

type baloadInstruction knownJVMInstruction

func parseBaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := baloadInstruction{
		basicJVMInstruction{
			raw: 0x33,
		},
		name: "baload",
	}
	return &toReturn, nil
}

type caloadInstruction knownJVMInstruction

func parseCaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := caloadInstruction{
		basicJVMInstruction{
			raw: 0x34,
		},
		name: "caload",
	}
	return &toReturn, nil
}

type saloadInstruction knownJVMInstruction

func parseSaloadInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := saloadInstruction{
		basicJVMInstruction{
			raw: 0x35,
		},
		name: "saload",
	}
	return &toReturn, nil
}

func parseIstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseFstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseDstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseAstoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type istore_0Instruction knownJVMInstruction

func parseIstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_0Instruction{
		basicJVMInstruction{
			raw: 0x3b,
		},
		name: "istore_0",
	}
	return &toReturn, nil
}

type istore_1Instruction knownJVMInstruction

func parseIstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_1Instruction{
		basicJVMInstruction{
			raw: 0x3c,
		},
		name: "istore_1",
	}
	return &toReturn, nil
}

type istore_2Instruction knownJVMInstruction

func parseIstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_2Instruction{
		basicJVMInstruction{
			raw: 0x3d,
		},
		name: "istore_2",
	}
	return &toReturn, nil
}

type istore_3Instruction knownJVMInstruction

func parseIstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := istore_3Instruction{
		basicJVMInstruction{
			raw: 0x3e,
		},
		name: "istore_3",
	}
	return &toReturn, nil
}

type lstore_0Instruction knownJVMInstruction

func parseLstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_0Instruction{
		basicJVMInstruction{
			raw: 0x3f,
		},
		name: "lstore_0",
	}
	return &toReturn, nil
}

type lstore_1Instruction knownJVMInstruction

func parseLstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_1Instruction{
		basicJVMInstruction{
			raw: 0x40,
		},
		name: "lstore_1",
	}
	return &toReturn, nil
}

type lstore_2Instruction knownJVMInstruction

func parseLstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_2Instruction{
		basicJVMInstruction{
			raw: 0x41,
		},
		name: "lstore_2",
	}
	return &toReturn, nil
}

type lstore_3Instruction knownJVMInstruction

func parseLstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lstore_3Instruction{
		basicJVMInstruction{
			raw: 0x42,
		},
		name: "lstore_3",
	}
	return &toReturn, nil
}

type fstore_0Instruction knownJVMInstruction

func parseFstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_0Instruction{
		basicJVMInstruction{
			raw: 0x43,
		},
		name: "fstore_0",
	}
	return &toReturn, nil
}

type fstore_1Instruction knownJVMInstruction

func parseFstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_1Instruction{
		basicJVMInstruction{
			raw: 0x44,
		},
		name: "fstore_1",
	}
	return &toReturn, nil
}

type fstore_2Instruction knownJVMInstruction

func parseFstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_2Instruction{
		basicJVMInstruction{
			raw: 0x45,
		},
		name: "fstore_2",
	}
	return &toReturn, nil
}

type fstore_3Instruction knownJVMInstruction

func parseFstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fstore_3Instruction{
		basicJVMInstruction{
			raw: 0x46,
		},
		name: "fstore_3",
	}
	return &toReturn, nil
}

type dstore_0Instruction knownJVMInstruction

func parseDstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_0Instruction{
		basicJVMInstruction{
			raw: 0x47,
		},
		name: "dstore_0",
	}
	return &toReturn, nil
}

type dstore_1Instruction knownJVMInstruction

func parseDstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_1Instruction{
		basicJVMInstruction{
			raw: 0x48,
		},
		name: "dstore_1",
	}
	return &toReturn, nil
}

type dstore_2Instruction knownJVMInstruction

func parseDstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_2Instruction{
		basicJVMInstruction{
			raw: 0x49,
		},
		name: "dstore_2",
	}
	return &toReturn, nil
}

type dstore_3Instruction knownJVMInstruction

func parseDstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dstore_3Instruction{
		basicJVMInstruction{
			raw: 0x4a,
		},
		name: "dstore_3",
	}
	return &toReturn, nil
}

type astore_0Instruction knownJVMInstruction

func parseAstore_0Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_0Instruction{
		basicJVMInstruction{
			raw: 0x4b,
		},
		name: "astore_0",
	}
	return &toReturn, nil
}

type astore_1Instruction knownJVMInstruction

func parseAstore_1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_1Instruction{
		basicJVMInstruction{
			raw: 0x4c,
		},
		name: "astore_1",
	}
	return &toReturn, nil
}

type astore_2Instruction knownJVMInstruction

func parseAstore_2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_2Instruction{
		basicJVMInstruction{
			raw: 0x4d,
		},
		name: "astore_2",
	}
	return &toReturn, nil
}

type astore_3Instruction knownJVMInstruction

func parseAstore_3Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := astore_3Instruction{
		basicJVMInstruction{
			raw: 0x4e,
		},
		name: "astore_3",
	}
	return &toReturn, nil
}

type iastoreInstruction knownJVMInstruction

func parseIastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iastoreInstruction{
		basicJVMInstruction{
			raw: 0x4f,
		},
		name: "iastore",
	}
	return &toReturn, nil
}

type lastoreInstruction knownJVMInstruction

func parseLastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lastoreInstruction{
		basicJVMInstruction{
			raw: 0x50,
		},
		name: "lastore",
	}
	return &toReturn, nil
}

type fastoreInstruction knownJVMInstruction

func parseFastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fastoreInstruction{
		basicJVMInstruction{
			raw: 0x51,
		},
		name: "fastore",
	}
	return &toReturn, nil
}

type dastoreInstruction knownJVMInstruction

func parseDastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dastoreInstruction{
		basicJVMInstruction{
			raw: 0x52,
		},
		name: "dastore",
	}
	return &toReturn, nil
}

type aastoreInstruction knownJVMInstruction

func parseAastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := aastoreInstruction{
		basicJVMInstruction{
			raw: 0x53,
		},
		name: "aastore",
	}
	return &toReturn, nil
}

type bastoreInstruction knownJVMInstruction

func parseBastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := bastoreInstruction{
		basicJVMInstruction{
			raw: 0x54,
		},
		name: "bastore",
	}
	return &toReturn, nil
}

type castoreInstruction knownJVMInstruction

func parseCastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := castoreInstruction{
		basicJVMInstruction{
			raw: 0x55,
		},
		name: "castore",
	}
	return &toReturn, nil
}

type sastoreInstruction knownJVMInstruction

func parseSastoreInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := sastoreInstruction{
		basicJVMInstruction{
			raw: 0x56,
		},
		name: "sastore",
	}
	return &toReturn, nil
}

type popInstruction knownJVMInstruction

func parsePopInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := popInstruction{
		basicJVMInstruction{
			raw: 0x57,
		},
		name: "pop",
	}
	return &toReturn, nil
}

type pop2Instruction knownJVMInstruction

func parsePop2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := pop2Instruction{
		basicJVMInstruction{
			raw: 0x58,
		},
		name: "pop2",
	}
	return &toReturn, nil
}

type dupInstruction knownJVMInstruction

func parseDupInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dupInstruction{
		basicJVMInstruction{
			raw: 0x59,
		},
		name: "dup",
	}
	return &toReturn, nil
}

type dup_x1Instruction knownJVMInstruction

func parseDup_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x1Instruction{
		basicJVMInstruction{
			raw: 0x5a,
		},
		name: "dup_x1",
	}
	return &toReturn, nil
}

type dup_x2Instruction knownJVMInstruction

func parseDup_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup_x2Instruction{
		basicJVMInstruction{
			raw: 0x5b,
		},
		name: "dup_x2",
	}
	return &toReturn, nil
}

type dup2Instruction knownJVMInstruction

func parseDup2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2Instruction{
		basicJVMInstruction{
			raw: 0x5c,
		},
		name: "dup2",
	}
	return &toReturn, nil
}

type dup2_x1Instruction knownJVMInstruction

func parseDup2_x1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x1Instruction{
		basicJVMInstruction{
			raw: 0x5d,
		},
		name: "dup2_x1",
	}
	return &toReturn, nil
}

type dup2_x2Instruction knownJVMInstruction

func parseDup2_x2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dup2_x2Instruction{
		basicJVMInstruction{
			raw: 0x5e,
		},
		name: "dup2_x2",
	}
	return &toReturn, nil
}

type swapInstruction knownJVMInstruction

func parseSwapInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := swapInstruction{
		basicJVMInstruction{
			raw: 0x5f,
		},
		name: "swap",
	}
	return &toReturn, nil
}

type iaddInstruction knownJVMInstruction

func parseIaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iaddInstruction{
		basicJVMInstruction{
			raw: 0x60,
		},
		name: "iadd",
	}
	return &toReturn, nil
}

type laddInstruction knownJVMInstruction

func parseLaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := laddInstruction{
		basicJVMInstruction{
			raw: 0x61,
		},
		name: "ladd",
	}
	return &toReturn, nil
}

type faddInstruction knownJVMInstruction

func parseFaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := faddInstruction{
		basicJVMInstruction{
			raw: 0x62,
		},
		name: "fadd",
	}
	return &toReturn, nil
}

type daddInstruction knownJVMInstruction

func parseDaddInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := daddInstruction{
		basicJVMInstruction{
			raw: 0x63,
		},
		name: "dadd",
	}
	return &toReturn, nil
}

type isubInstruction knownJVMInstruction

func parseIsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := isubInstruction{
		basicJVMInstruction{
			raw: 0x64,
		},
		name: "isub",
	}
	return &toReturn, nil
}

type lsubInstruction knownJVMInstruction

func parseLsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lsubInstruction{
		basicJVMInstruction{
			raw: 0x65,
		},
		name: "lsub",
	}
	return &toReturn, nil
}

type fsubInstruction knownJVMInstruction

func parseFsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fsubInstruction{
		basicJVMInstruction{
			raw: 0x66,
		},
		name: "fsub",
	}
	return &toReturn, nil
}

type dsubInstruction knownJVMInstruction

func parseDsubInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dsubInstruction{
		basicJVMInstruction{
			raw: 0x67,
		},
		name: "dsub",
	}
	return &toReturn, nil
}

type imulInstruction knownJVMInstruction

func parseImulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := imulInstruction{
		basicJVMInstruction{
			raw: 0x68,
		},
		name: "imul",
	}
	return &toReturn, nil
}

type lmulInstruction knownJVMInstruction

func parseLmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lmulInstruction{
		basicJVMInstruction{
			raw: 0x69,
		},
		name: "lmul",
	}
	return &toReturn, nil
}

type fmulInstruction knownJVMInstruction

func parseFmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fmulInstruction{
		basicJVMInstruction{
			raw: 0x6a,
		},
		name: "fmul",
	}
	return &toReturn, nil
}

type dmulInstruction knownJVMInstruction

func parseDmulInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dmulInstruction{
		basicJVMInstruction{
			raw: 0x6b,
		},
		name: "dmul",
	}
	return &toReturn, nil
}

type idivInstruction knownJVMInstruction

func parseIdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := idivInstruction{
		basicJVMInstruction{
			raw: 0x6c,
		},
		name: "idiv",
	}
	return &toReturn, nil
}

type ldivInstruction knownJVMInstruction

func parseLdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ldivInstruction{
		basicJVMInstruction{
			raw: 0x6d,
		},
		name: "ldiv",
	}
	return &toReturn, nil
}

type fdivInstruction knownJVMInstruction

func parseFdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fdivInstruction{
		basicJVMInstruction{
			raw: 0x6e,
		},
		name: "fdiv",
	}
	return &toReturn, nil
}

type ddivInstruction knownJVMInstruction

func parseDdivInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ddivInstruction{
		basicJVMInstruction{
			raw: 0x6f,
		},
		name: "ddiv",
	}
	return &toReturn, nil
}

type iremInstruction knownJVMInstruction

func parseIremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iremInstruction{
		basicJVMInstruction{
			raw: 0x70,
		},
		name: "irem",
	}
	return &toReturn, nil
}

type lremInstruction knownJVMInstruction

func parseLremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lremInstruction{
		basicJVMInstruction{
			raw: 0x71,
		},
		name: "lrem",
	}
	return &toReturn, nil
}

type fremInstruction knownJVMInstruction

func parseFremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fremInstruction{
		basicJVMInstruction{
			raw: 0x72,
		},
		name: "frem",
	}
	return &toReturn, nil
}

type dremInstruction knownJVMInstruction

func parseDremInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dremInstruction{
		basicJVMInstruction{
			raw: 0x73,
		},
		name: "drem",
	}
	return &toReturn, nil
}

type inegInstruction knownJVMInstruction

func parseInegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := inegInstruction{
		basicJVMInstruction{
			raw: 0x74,
		},
		name: "ineg",
	}
	return &toReturn, nil
}

type lnegInstruction knownJVMInstruction

func parseLnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lnegInstruction{
		basicJVMInstruction{
			raw: 0x75,
		},
		name: "lneg",
	}
	return &toReturn, nil
}

type fnegInstruction knownJVMInstruction

func parseFnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fnegInstruction{
		basicJVMInstruction{
			raw: 0x76,
		},
		name: "fneg",
	}
	return &toReturn, nil
}

type dnegInstruction knownJVMInstruction

func parseDnegInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dnegInstruction{
		basicJVMInstruction{
			raw: 0x77,
		},
		name: "dneg",
	}
	return &toReturn, nil
}

type ishlInstruction knownJVMInstruction

func parseIshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishlInstruction{
		basicJVMInstruction{
			raw: 0x78,
		},
		name: "ishl",
	}
	return &toReturn, nil
}

type lshlInstruction knownJVMInstruction

func parseLshlInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshlInstruction{
		basicJVMInstruction{
			raw: 0x79,
		},
		name: "lshl",
	}
	return &toReturn, nil
}

type ishrInstruction knownJVMInstruction

func parseIshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ishrInstruction{
		basicJVMInstruction{
			raw: 0x7a,
		},
		name: "ishr",
	}
	return &toReturn, nil
}

type lshrInstruction knownJVMInstruction

func parseLshrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lshrInstruction{
		basicJVMInstruction{
			raw: 0x7b,
		},
		name: "lshr",
	}
	return &toReturn, nil
}

type iushrInstruction knownJVMInstruction

func parseIushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iushrInstruction{
		basicJVMInstruction{
			raw: 0x7c,
		},
		name: "iushr",
	}
	return &toReturn, nil
}

type lushrInstruction knownJVMInstruction

func parseLushrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lushrInstruction{
		basicJVMInstruction{
			raw: 0x7d,
		},
		name: "lushr",
	}
	return &toReturn, nil
}

type iandInstruction knownJVMInstruction

func parseIandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iandInstruction{
		basicJVMInstruction{
			raw: 0x7e,
		},
		name: "iand",
	}
	return &toReturn, nil
}

type landInstruction knownJVMInstruction

func parseLandInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := landInstruction{
		basicJVMInstruction{
			raw: 0x7f,
		},
		name: "land",
	}
	return &toReturn, nil
}

type iorInstruction knownJVMInstruction

func parseIorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := iorInstruction{
		basicJVMInstruction{
			raw: 0x80,
		},
		name: "ior",
	}
	return &toReturn, nil
}

type lorInstruction knownJVMInstruction

func parseLorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lorInstruction{
		basicJVMInstruction{
			raw: 0x81,
		},
		name: "lor",
	}
	return &toReturn, nil
}

type ixorInstruction knownJVMInstruction

func parseIxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ixorInstruction{
		basicJVMInstruction{
			raw: 0x82,
		},
		name: "ixor",
	}
	return &toReturn, nil
}

type lxorInstruction knownJVMInstruction

func parseLxorInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lxorInstruction{
		basicJVMInstruction{
			raw: 0x83,
		},
		name: "lxor",
	}
	return &toReturn, nil
}

func parseIincInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type i2lInstruction knownJVMInstruction

func parseI2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2lInstruction{
		basicJVMInstruction{
			raw: 0x85,
		},
		name: "i2l",
	}
	return &toReturn, nil
}

type i2fInstruction knownJVMInstruction

func parseI2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2fInstruction{
		basicJVMInstruction{
			raw: 0x86,
		},
		name: "i2f",
	}
	return &toReturn, nil
}

type i2dInstruction knownJVMInstruction

func parseI2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2dInstruction{
		basicJVMInstruction{
			raw: 0x87,
		},
		name: "i2d",
	}
	return &toReturn, nil
}

type l2iInstruction knownJVMInstruction

func parseL2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2iInstruction{
		basicJVMInstruction{
			raw: 0x88,
		},
		name: "l2i",
	}
	return &toReturn, nil
}

type l2fInstruction knownJVMInstruction

func parseL2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2fInstruction{
		basicJVMInstruction{
			raw: 0x89,
		},
		name: "l2f",
	}
	return &toReturn, nil
}

type l2dInstruction knownJVMInstruction

func parseL2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := l2dInstruction{
		basicJVMInstruction{
			raw: 0x8a,
		},
		name: "l2d",
	}
	return &toReturn, nil
}

type f2iInstruction knownJVMInstruction

func parseF2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2iInstruction{
		basicJVMInstruction{
			raw: 0x8b,
		},
		name: "f2i",
	}
	return &toReturn, nil
}

type f2lInstruction knownJVMInstruction

func parseF2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2lInstruction{
		basicJVMInstruction{
			raw: 0x8c,
		},
		name: "f2l",
	}
	return &toReturn, nil
}

type f2dInstruction knownJVMInstruction

func parseF2dInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := f2dInstruction{
		basicJVMInstruction{
			raw: 0x8d,
		},
		name: "f2d",
	}
	return &toReturn, nil
}

type d2iInstruction knownJVMInstruction

func parseD2iInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2iInstruction{
		basicJVMInstruction{
			raw: 0x8e,
		},
		name: "d2i",
	}
	return &toReturn, nil
}

type d2lInstruction knownJVMInstruction

func parseD2lInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2lInstruction{
		basicJVMInstruction{
			raw: 0x8f,
		},
		name: "d2l",
	}
	return &toReturn, nil
}

type d2fInstruction knownJVMInstruction

func parseD2fInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := d2fInstruction{
		basicJVMInstruction{
			raw: 0x90,
		},
		name: "d2f",
	}
	return &toReturn, nil
}

type i2bInstruction knownJVMInstruction

func parseI2bInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2bInstruction{
		basicJVMInstruction{
			raw: 0x91,
		},
		name: "i2b",
	}
	return &toReturn, nil
}

type i2cInstruction knownJVMInstruction

func parseI2cInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2cInstruction{
		basicJVMInstruction{
			raw: 0x92,
		},
		name: "i2c",
	}
	return &toReturn, nil
}

type i2sInstruction knownJVMInstruction

func parseI2sInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := i2sInstruction{
		basicJVMInstruction{
			raw: 0x93,
		},
		name: "i2s",
	}
	return &toReturn, nil
}

type lcmpInstruction knownJVMInstruction

func parseLcmpInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lcmpInstruction{
		basicJVMInstruction{
			raw: 0x94,
		},
		name: "lcmp",
	}
	return &toReturn, nil
}

type fcmplInstruction knownJVMInstruction

func parseFcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmplInstruction{
		basicJVMInstruction{
			raw: 0x95,
		},
		name: "fcmpl",
	}
	return &toReturn, nil
}

type fcmpgInstruction knownJVMInstruction

func parseFcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := fcmpgInstruction{
		basicJVMInstruction{
			raw: 0x96,
		},
		name: "fcmpg",
	}
	return &toReturn, nil
}

type dcmplInstruction knownJVMInstruction

func parseDcmplInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmplInstruction{
		basicJVMInstruction{
			raw: 0x97,
		},
		name: "dcmpl",
	}
	return &toReturn, nil
}

type dcmpgInstruction knownJVMInstruction

func parseDcmpgInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dcmpgInstruction{
		basicJVMInstruction{
			raw: 0x98,
		},
		name: "dcmpg",
	}
	return &toReturn, nil
}

func parseIfeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpltInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpgeInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpgtInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_icmpleInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_acmpeqInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIf_acmpneInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseGotoInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseJsrInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseRetInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseTableswitchInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseLookupswitchInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type ireturnInstruction knownJVMInstruction

func parseIreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := ireturnInstruction{
		basicJVMInstruction{
			raw: 0xac,
		},
		name: "ireturn",
	}
	return &toReturn, nil
}

type lreturnInstruction knownJVMInstruction

func parseLreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := lreturnInstruction{
		basicJVMInstruction{
			raw: 0xad,
		},
		name: "lreturn",
	}
	return &toReturn, nil
}

type freturnInstruction knownJVMInstruction

func parseFreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := freturnInstruction{
		basicJVMInstruction{
			raw: 0xae,
		},
		name: "freturn",
	}
	return &toReturn, nil
}

type dreturnInstruction knownJVMInstruction

func parseDreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := dreturnInstruction{
		basicJVMInstruction{
			raw: 0xaf,
		},
		name: "dreturn",
	}
	return &toReturn, nil
}

type areturnInstruction knownJVMInstruction

func parseAreturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := areturnInstruction{
		basicJVMInstruction{
			raw: 0xb0,
		},
		name: "areturn",
	}
	return &toReturn, nil
}

type returnInstruction knownJVMInstruction

func parseReturnInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := returnInstruction{
		basicJVMInstruction{
			raw: 0xb1,
		},
		name: "return",
	}
	return &toReturn, nil
}

func parseGetstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parsePutstaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseGetfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parsePutfieldInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInvokevirtualInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInvokespecialInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInvokestaticInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInvokeinterfaceInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInvokedynamicInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseNewInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseNewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseAnewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type arraylengthInstruction knownJVMInstruction

func parseArraylengthInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := arraylengthInstruction{
		basicJVMInstruction{
			raw: 0xbe,
		},
		name: "arraylength",
	}
	return &toReturn, nil
}

type athrowInstruction knownJVMInstruction

func parseAthrowInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := athrowInstruction{
		basicJVMInstruction{
			raw: 0xbf,
		},
		name: "athrow",
	}
	return &toReturn, nil
}

func parseCheckcastInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseInstanceofInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type monitorenterInstruction knownJVMInstruction

func parseMonitorenterInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorenterInstruction{
		basicJVMInstruction{
			raw: 0xc2,
		},
		name: "monitorenter",
	}
	return &toReturn, nil
}

type monitorexitInstruction knownJVMInstruction

func parseMonitorexitInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := monitorexitInstruction{
		basicJVMInstruction{
			raw: 0xc3,
		},
		name: "monitorexit",
	}
	return &toReturn, nil
}

func parseWideInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseMultianewarrayInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseIfnonnullInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseGoto_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

func parseJsr_wInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	return nil, NotImplementedError
}

type breakpointInstruction knownJVMInstruction

func parseBreakpointInstruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := breakpointInstruction{
		basicJVMInstruction{
			raw: 0xca,
		},
		name: "breakpoint",
	}
	return &toReturn, nil
}

type impdep1Instruction knownJVMInstruction

func parseImpdep1Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep1Instruction{
		basicJVMInstruction{
			raw: 0xfe,
		},
		name: "impdep1",
	}
	return &toReturn, nil
}

type impdep2Instruction knownJVMInstruction

func parseImpdep2Instruction(opcode uint8, name string, address uint,
	m JVMMemory) (JVMInstruction, error) {
	toReturn := impdep2Instruction{
		basicJVMInstruction{
			raw: 0xff,
		},
		name: "impdep2",
	}
	return &toReturn, nil
}
