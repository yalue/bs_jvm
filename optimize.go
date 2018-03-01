package bs_jvm

// This file contains instruction-specific Optimize function definitions.

func (n *ldcInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	// TODO (next): Optimize the ldc instruction
	//  - Get constant from constant pool
	//  - Store the constant in the ldcInstruction struct.
	return NotImplementedError
}
