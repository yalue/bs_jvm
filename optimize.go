package bs_jvm

// This file contains instruction-specific Optimize function definitions.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
	"math"
)

func (n *ldcInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	constant, e := m.ContainingClass.File.GetConstant(uint16(n.value))
	if e != nil {
		return e
	}
	var reference Reference
	// Primitive types can be converted now, but references to other objects
	// can be handled later--for now just read them from the class file.
	switch v := constant.(type) {
	case *class_file.ConstantIntegerInfo:
		n.isPrimitive = true
		n.primitiveValue = Int(v.Value)
	case *class_file.ConstantFloatInfo:
		n.isPrimitive = true
		n.primitiveValue = Int(math.Float32bits(v.Value))
	case *class_file.ConstantStringInfo, *class_file.ConstantMethodHandleInfo,
		*class_file.ConstantMethodTypeInfo, *class_file.ConstantClassInfo:
		n.isPrimitive = false
		reference, e = ConvertConstantToObject(m.ContainingClass, constant)
		if e != nil {
			return e
		}
		n.reference = reference
	default:
		// ldc only allows the types of constants listed above.
		return TypeError(fmt.Sprintf("Invalid ldc constant: %s", constant))
	}
	return nil
}
