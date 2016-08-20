package jvm

// This file contains a list of JVM opcode bytes and metadata needed for
// parsing them into JVMInstruction-compatible types.

// Takes an opcode, the instruction name, the address of the opcode, and a
// JVMMemory, then returns a JVMInstruction object.
type parserFunction func(uint8, string, uint, JVMMemory) (JVMInstruction,
	error)

// Metadata to assist in parsing opcodes
type jvmOpcodeInfo struct {
	name   string
	opcode uint8
	parse  parserFunction
}

// TODO: This is an array, not a map! Fix it!
var opcodeTable = map[uint8]*jvmOpcodeInfo{
	&jvmOpcodeInfo{
		name:   "nop",
		opcode: 0x00,
		parse:  parseNopInstruction,
	},
	&jvmOpcodeInfo{
		name:   "aconst_null",
		opcode: 0x01,
		parse:  parseAconst_nullInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_m1",
		opcode: 0x02,
		parse:  parseIconst_m1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_0",
		opcode: 0x03,
		parse:  parseIconst_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_1",
		opcode: 0x04,
		parse:  parseIconst_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_2",
		opcode: 0x05,
		parse:  parseIconst_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_3",
		opcode: 0x06,
		parse:  parseIconst_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_4",
		opcode: 0x07,
		parse:  parseIconst_4Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iconst_5",
		opcode: 0x08,
		parse:  parseIconst_5Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lconst_0",
		opcode: 0x09,
		parse:  parseLconst_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lconst_1",
		opcode: 0x0a,
		parse:  parseLconst_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fconst_0",
		opcode: 0x0b,
		parse:  parseFconst_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fconst_1",
		opcode: 0x0c,
		parse:  parseFconst_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fconst_2",
		opcode: 0x0d,
		parse:  parseFconst_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dconst_0",
		opcode: 0x0e,
		parse:  parseDconst_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dconst_1",
		opcode: 0x0f,
		parse:  parseDconst_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "bipush",
		opcode: 0x10,
		parse:  parseBipushInstruction,
	},
	&jvmOpcodeInfo{
		name:   "sipush",
		opcode: 0x11,
		parse:  parseSipushInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ldc",
		opcode: 0x12,
		parse:  parseLdcInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ldc_w",
		opcode: 0x13,
		parse:  parseLdc_wInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ldc2_w",
		opcode: 0x14,
		parse:  parseLdc2_wInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iload",
		opcode: 0x15,
		parse:  parseIloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lload",
		opcode: 0x16,
		parse:  parseLloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fload",
		opcode: 0x17,
		parse:  parseFloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dload",
		opcode: 0x18,
		parse:  parseDloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "aload",
		opcode: 0x19,
		parse:  parseAloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iload_0",
		opcode: 0x1a,
		parse:  parseIload_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iload_1",
		opcode: 0x1b,
		parse:  parseIload_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iload_2",
		opcode: 0x1c,
		parse:  parseIload_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iload_3",
		opcode: 0x1d,
		parse:  parseIload_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lload_0",
		opcode: 0x1e,
		parse:  parseLload_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lload_1",
		opcode: 0x1f,
		parse:  parseLload_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lload_2",
		opcode: 0x20,
		parse:  parseLload_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lload_3",
		opcode: 0x21,
		parse:  parseLload_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fload_0",
		opcode: 0x22,
		parse:  parseFload_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fload_1",
		opcode: 0x23,
		parse:  parseFload_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fload_2",
		opcode: 0x24,
		parse:  parseFload_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fload_3",
		opcode: 0x25,
		parse:  parseFload_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dload_0",
		opcode: 0x26,
		parse:  parseDload_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dload_1",
		opcode: 0x27,
		parse:  parseDload_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dload_2",
		opcode: 0x28,
		parse:  parseDload_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dload_3",
		opcode: 0x29,
		parse:  parseDload_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "aload_0",
		opcode: 0x2a,
		parse:  parseAload_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "aload_1",
		opcode: 0x2b,
		parse:  parseAload_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "aload_2",
		opcode: 0x2c,
		parse:  parseAload_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "aload_3",
		opcode: 0x2d,
		parse:  parseAload_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iaload",
		opcode: 0x2e,
		parse:  parseIaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "laload",
		opcode: 0x2f,
		parse:  parseLaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "faload",
		opcode: 0x30,
		parse:  parseFaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "daload",
		opcode: 0x31,
		parse:  parseDaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "aaload",
		opcode: 0x32,
		parse:  parseAaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "baload",
		opcode: 0x33,
		parse:  parseBaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "caload",
		opcode: 0x34,
		parse:  parseCaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "saload",
		opcode: 0x35,
		parse:  parseSaloadInstruction,
	},
	&jvmOpcodeInfo{
		name:   "istore",
		opcode: 0x36,
		parse:  parseIstoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lstore",
		opcode: 0x37,
		parse:  parseLstoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fstore",
		opcode: 0x38,
		parse:  parseFstoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dstore",
		opcode: 0x39,
		parse:  parseDstoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "astore",
		opcode: 0x3a,
		parse:  parseAstoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "istore_0",
		opcode: 0x3b,
		parse:  parseIstore_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "istore_1",
		opcode: 0x3c,
		parse:  parseIstore_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "istore_2",
		opcode: 0x3d,
		parse:  parseIstore_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "istore_3",
		opcode: 0x3e,
		parse:  parseIstore_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lstore_0",
		opcode: 0x3f,
		parse:  parseLstore_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lstore_1",
		opcode: 0x40,
		parse:  parseLstore_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lstore_2",
		opcode: 0x41,
		parse:  parseLstore_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "lstore_3",
		opcode: 0x42,
		parse:  parseLstore_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fstore_0",
		opcode: 0x43,
		parse:  parseFstore_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fstore_1",
		opcode: 0x44,
		parse:  parseFstore_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fstore_2",
		opcode: 0x45,
		parse:  parseFstore_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "fstore_3",
		opcode: 0x46,
		parse:  parseFstore_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dstore_0",
		opcode: 0x47,
		parse:  parseDstore_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dstore_1",
		opcode: 0x48,
		parse:  parseDstore_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dstore_2",
		opcode: 0x49,
		parse:  parseDstore_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dstore_3",
		opcode: 0x4a,
		parse:  parseDstore_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "astore_0",
		opcode: 0x4b,
		parse:  parseAstore_0Instruction,
	},
	&jvmOpcodeInfo{
		name:   "astore_1",
		opcode: 0x4c,
		parse:  parseAstore_1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "astore_2",
		opcode: 0x4d,
		parse:  parseAstore_2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "astore_3",
		opcode: 0x4e,
		parse:  parseAstore_3Instruction,
	},
	&jvmOpcodeInfo{
		name:   "iastore",
		opcode: 0x4f,
		parse:  parseIastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lastore",
		opcode: 0x50,
		parse:  parseLastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fastore",
		opcode: 0x51,
		parse:  parseFastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dastore",
		opcode: 0x52,
		parse:  parseDastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "aastore",
		opcode: 0x53,
		parse:  parseAastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "bastore",
		opcode: 0x54,
		parse:  parseBastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "castore",
		opcode: 0x55,
		parse:  parseCastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "sastore",
		opcode: 0x56,
		parse:  parseSastoreInstruction,
	},
	&jvmOpcodeInfo{
		name:   "pop",
		opcode: 0x57,
		parse:  parsePopInstruction,
	},
	&jvmOpcodeInfo{
		name:   "pop2",
		opcode: 0x58,
		parse:  parsePop2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dup",
		opcode: 0x59,
		parse:  parseDupInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dup_x1",
		opcode: 0x5a,
		parse:  parseDup_x1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dup_x2",
		opcode: 0x5b,
		parse:  parseDup_x2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dup2",
		opcode: 0x5c,
		parse:  parseDup2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dup2_x1",
		opcode: 0x5d,
		parse:  parseDup2_x1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "dup2_x2",
		opcode: 0x5e,
		parse:  parseDup2_x2Instruction,
	},
	&jvmOpcodeInfo{
		name:   "swap",
		opcode: 0x5f,
		parse:  parseSwapInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iadd",
		opcode: 0x60,
		parse:  parseIaddInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ladd",
		opcode: 0x61,
		parse:  parseLaddInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fadd",
		opcode: 0x62,
		parse:  parseFaddInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dadd",
		opcode: 0x63,
		parse:  parseDaddInstruction,
	},
	&jvmOpcodeInfo{
		name:   "isub",
		opcode: 0x64,
		parse:  parseIsubInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lsub",
		opcode: 0x65,
		parse:  parseLsubInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fsub",
		opcode: 0x66,
		parse:  parseFsubInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dsub",
		opcode: 0x67,
		parse:  parseDsubInstruction,
	},
	&jvmOpcodeInfo{
		name:   "imul",
		opcode: 0x68,
		parse:  parseImulInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lmul",
		opcode: 0x69,
		parse:  parseLmulInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fmul",
		opcode: 0x6a,
		parse:  parseFmulInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dmul",
		opcode: 0x6b,
		parse:  parseDmulInstruction,
	},
	&jvmOpcodeInfo{
		name:   "idiv",
		opcode: 0x6c,
		parse:  parseIdivInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ldiv",
		opcode: 0x6d,
		parse:  parseLdivInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fdiv",
		opcode: 0x6e,
		parse:  parseFdivInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ddiv",
		opcode: 0x6f,
		parse:  parseDdivInstruction,
	},
	&jvmOpcodeInfo{
		name:   "irem",
		opcode: 0x70,
		parse:  parseIremInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lrem",
		opcode: 0x71,
		parse:  parseLremInstruction,
	},
	&jvmOpcodeInfo{
		name:   "frem",
		opcode: 0x72,
		parse:  parseFremInstruction,
	},
	&jvmOpcodeInfo{
		name:   "drem",
		opcode: 0x73,
		parse:  parseDremInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ineg",
		opcode: 0x74,
		parse:  parseInegInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lneg",
		opcode: 0x75,
		parse:  parseLnegInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fneg",
		opcode: 0x76,
		parse:  parseFnegInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dneg",
		opcode: 0x77,
		parse:  parseDnegInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ishl",
		opcode: 0x78,
		parse:  parseIshlInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lshl",
		opcode: 0x79,
		parse:  parseLshlInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ishr",
		opcode: 0x7a,
		parse:  parseIshrInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lshr",
		opcode: 0x7b,
		parse:  parseLshrInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iushr",
		opcode: 0x7c,
		parse:  parseIushrInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lushr",
		opcode: 0x7d,
		parse:  parseLushrInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iand",
		opcode: 0x7e,
		parse:  parseIandInstruction,
	},
	&jvmOpcodeInfo{
		name:   "land",
		opcode: 0x7f,
		parse:  parseLandInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ior",
		opcode: 0x80,
		parse:  parseIorInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lor",
		opcode: 0x81,
		parse:  parseLorInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ixor",
		opcode: 0x82,
		parse:  parseIxorInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lxor",
		opcode: 0x83,
		parse:  parseLxorInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iinc",
		opcode: 0x84,
		parse:  parseIincInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2l",
		opcode: 0x85,
		parse:  parseI2lInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2f",
		opcode: 0x86,
		parse:  parseI2fInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2d",
		opcode: 0x87,
		parse:  parseI2dInstruction,
	},
	&jvmOpcodeInfo{
		name:   "l2i",
		opcode: 0x88,
		parse:  parseL2iInstruction,
	},
	&jvmOpcodeInfo{
		name:   "l2f",
		opcode: 0x89,
		parse:  parseL2fInstruction,
	},
	&jvmOpcodeInfo{
		name:   "l2d",
		opcode: 0x8a,
		parse:  parseL2dInstruction,
	},
	&jvmOpcodeInfo{
		name:   "f2i",
		opcode: 0x8b,
		parse:  parseF2iInstruction,
	},
	&jvmOpcodeInfo{
		name:   "f2l",
		opcode: 0x8c,
		parse:  parseF2lInstruction,
	},
	&jvmOpcodeInfo{
		name:   "f2d",
		opcode: 0x8d,
		parse:  parseF2dInstruction,
	},
	&jvmOpcodeInfo{
		name:   "d2i",
		opcode: 0x8e,
		parse:  parseD2iInstruction,
	},
	&jvmOpcodeInfo{
		name:   "d2l",
		opcode: 0x8f,
		parse:  parseD2lInstruction,
	},
	&jvmOpcodeInfo{
		name:   "d2f",
		opcode: 0x90,
		parse:  parseD2fInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2b",
		opcode: 0x91,
		parse:  parseI2bInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2c",
		opcode: 0x92,
		parse:  parseI2cInstruction,
	},
	&jvmOpcodeInfo{
		name:   "i2s",
		opcode: 0x93,
		parse:  parseI2sInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lcmp",
		opcode: 0x94,
		parse:  parseLcmpInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fcmpl",
		opcode: 0x95,
		parse:  parseFcmplInstruction,
	},
	&jvmOpcodeInfo{
		name:   "fcmpg",
		opcode: 0x96,
		parse:  parseFcmpgInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dcmpl",
		opcode: 0x97,
		parse:  parseDcmplInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dcmpg",
		opcode: 0x98,
		parse:  parseDcmpgInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifeq",
		opcode: 0x99,
		parse:  parseIfeqInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifne",
		opcode: 0x9a,
		parse:  parseIfneInstruction,
	},
	&jvmOpcodeInfo{
		name:   "iflt",
		opcode: 0x9b,
		parse:  parseIfltInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifge",
		opcode: 0x9c,
		parse:  parseIfgeInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifgt",
		opcode: 0x9d,
		parse:  parseIfgtInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifle",
		opcode: 0x9e,
		parse:  parseIfleInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmpeq",
		opcode: 0x9f,
		parse:  parseIf_icmpeqInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmpne",
		opcode: 0xa0,
		parse:  parseIf_icmpneInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmplt",
		opcode: 0xa1,
		parse:  parseIf_icmpltInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmpge",
		opcode: 0xa2,
		parse:  parseIf_icmpgeInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmpgt",
		opcode: 0xa3,
		parse:  parseIf_icmpgtInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_icmple",
		opcode: 0xa4,
		parse:  parseIf_icmpleInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_acmpeq",
		opcode: 0xa5,
		parse:  parseIf_acmpeqInstruction,
	},
	&jvmOpcodeInfo{
		name:   "if_acmpne",
		opcode: 0xa6,
		parse:  parseIf_acmpneInstruction,
	},
	&jvmOpcodeInfo{
		name:   "goto",
		opcode: 0xa7,
		parse:  parseGotoInstruction,
	},
	&jvmOpcodeInfo{
		name:   "jsr",
		opcode: 0xa8,
		parse:  parseJsrInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ret",
		opcode: 0xa9,
		parse:  parseRetInstruction,
	},
	&jvmOpcodeInfo{
		name:   "tableswitch",
		opcode: 0xaa,
		parse:  parseTableswitchInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lookupswitch",
		opcode: 0xab,
		parse:  parseLookupswitchInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ireturn",
		opcode: 0xac,
		parse:  parseIreturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "lreturn",
		opcode: 0xad,
		parse:  parseLreturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "freturn",
		opcode: 0xae,
		parse:  parseFreturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "dreturn",
		opcode: 0xaf,
		parse:  parseDreturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "areturn",
		opcode: 0xb0,
		parse:  parseAreturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "return",
		opcode: 0xb1,
		parse:  parseReturnInstruction,
	},
	&jvmOpcodeInfo{
		name:   "getstatic",
		opcode: 0xb2,
		parse:  parseGetstaticInstruction,
	},
	&jvmOpcodeInfo{
		name:   "putstatic",
		opcode: 0xb3,
		parse:  parsePutstaticInstruction,
	},
	&jvmOpcodeInfo{
		name:   "getfield",
		opcode: 0xb4,
		parse:  parseGetfieldInstruction,
	},
	&jvmOpcodeInfo{
		name:   "putfield",
		opcode: 0xb5,
		parse:  parsePutfieldInstruction,
	},
	&jvmOpcodeInfo{
		name:   "invokevirtual",
		opcode: 0xb6,
		parse:  parseInvokevirtualInstruction,
	},
	&jvmOpcodeInfo{
		name:   "invokespecial",
		opcode: 0xb7,
		parse:  parseInvokespecialInstruction,
	},
	&jvmOpcodeInfo{
		name:   "invokestatic",
		opcode: 0xb8,
		parse:  parseInvokestaticInstruction,
	},
	&jvmOpcodeInfo{
		name:   "invokeinterface",
		opcode: 0xb9,
		parse:  parseInvokeinterfaceInstruction,
	},
	&jvmOpcodeInfo{
		name:   "invokedynamic",
		opcode: 0xba,
		parse:  parseInvokedynamicInstruction,
	},
	&jvmOpcodeInfo{
		name:   "new",
		opcode: 0xbb,
		parse:  parseNewInstruction,
	},
	&jvmOpcodeInfo{
		name:   "newarray",
		opcode: 0xbc,
		parse:  parseNewarrayInstruction,
	},
	&jvmOpcodeInfo{
		name:   "anewarray",
		opcode: 0xbd,
		parse:  parseAnewarrayInstruction,
	},
	&jvmOpcodeInfo{
		name:   "arraylength",
		opcode: 0xbe,
		parse:  parseArraylengthInstruction,
	},
	&jvmOpcodeInfo{
		name:   "athrow",
		opcode: 0xbf,
		parse:  parseAthrowInstruction,
	},
	&jvmOpcodeInfo{
		name:   "checkcast",
		opcode: 0xc0,
		parse:  parseCheckcastInstruction,
	},
	&jvmOpcodeInfo{
		name:   "instanceof",
		opcode: 0xc1,
		parse:  parseInstanceofInstruction,
	},
	&jvmOpcodeInfo{
		name:   "monitorenter",
		opcode: 0xc2,
		parse:  parseMonitorenterInstruction,
	},
	&jvmOpcodeInfo{
		name:   "monitorexit",
		opcode: 0xc3,
		parse:  parseMonitorexitInstruction,
	},
	&jvmOpcodeInfo{
		name:   "wide",
		opcode: 0xc4,
		parse:  parseWideInstruction,
	},
	&jvmOpcodeInfo{
		name:   "multianewarray",
		opcode: 0xc5,
		parse:  parseMultianewarrayInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifnull",
		opcode: 0xc6,
		parse:  parseIfnullInstruction,
	},
	&jvmOpcodeInfo{
		name:   "ifnonnull",
		opcode: 0xc7,
		parse:  parseIfnonnullInstruction,
	},
	&jvmOpcodeInfo{
		name:   "goto_w",
		opcode: 0xc8,
		parse:  parseGoto_wInstruction,
	},
	&jvmOpcodeInfo{
		name:   "jsr_w",
		opcode: 0xc9,
		parse:  parseJsr_wInstruction,
	},
	&jvmOpcodeInfo{
		name:   "breakpoint",
		opcode: 0xca,
		parse:  parseBreakpointInstruction,
	},
	&jvmOpcodeInfo{
		name:   "impdep1",
		opcode: 0xfe,
		parse:  parseImpdep1Instruction,
	},
	&jvmOpcodeInfo{
		name:   "impdep2",
		opcode: 0xff,
		parse:  parseImpdep2Instruction,
	},
}
