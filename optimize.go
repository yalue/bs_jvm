package bs_jvm

// This file contains instruction-specific Optimize function definitions.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
	"math"
)

// Converts a constant to a value that can be pushed by ldc or ldc_w. Returns
// the primitive value or reference to push. The constant is a primitive if and
// only if the returned Object is nil.
func constantToLdcInfo(class *Class, c class_file.Constant) (Int, Object,
	error) {
	var primitive Int
	var reference Object = nil
	var e error
	// Primitive types can be converted now, but references to other objects
	// can be handled later--for now just read them from the class file.
	switch v := c.(type) {
	case *class_file.ConstantIntegerInfo:
		primitive = Int(v.Value)
		reference = Int(v.Value)
	case *class_file.ConstantFloatInfo:
		primitive = Int(math.Float32bits(v.Value))
		reference = Float(v.Value)
	case *class_file.ConstantStringInfo, *class_file.ConstantMethodHandleInfo,
		*class_file.ConstantMethodTypeInfo, *class_file.ConstantClassInfo:
		reference, e = ConvertConstantToObject(class, c)
		if e != nil {
			return 0, nil, e
		}
	default:
		// ldc only allows the types of constants listed above.
		return 0, nil, TypeError(fmt.Sprintf("Invalid ldc constant: %s", c))
	}
	return primitive, reference, nil
}

func (n *ldcInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	constant, e := m.ContainingClass.File.GetConstant(uint16(n.value))
	if e != nil {
		return e
	}
	primitive, reference, e := constantToLdcInfo(m.ContainingClass, constant)
	if e != nil {
		return e
	}
	n.isPrimitive = reference == nil
	n.primitiveValue = primitive
	n.reference = reference
	return nil
}

func (n *ldc_wInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	constant, e := m.ContainingClass.File.GetConstant(n.value)
	if e != nil {
		return e
	}
	primitive, reference, e := constantToLdcInfo(m.ContainingClass, constant)
	if e != nil {
		return e
	}
	n.isPrimitive = reference == nil
	n.primitiveValue = primitive
	n.reference = reference
	return nil
}

func (n *ldc2_wInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	constant, e := m.ContainingClass.File.GetConstant(n.value)
	if e != nil {
		return e
	}
	switch v := constant.(type) {
	case *class_file.ConstantLongInfo:
		n.primitiveValue = Long(v.Value)
		n.reference = Long(v.Value)
	case *class_file.ConstantDoubleInfo:
		n.primitiveValue = Long(math.Float64bits(v.Value))
		n.reference = Double(v.Value)
	default:
		return TypeError(fmt.Sprintf("Invalid ldc2_w constant: %s", constant))
	}
	return nil
}

// Takes an instruction's offset and a signed offset relative to the
// instruction, and returns the index of the instruction at the relative
// offset. Returns an appropriate error if one occurs, e.g., if the offset
// doesn't correspond to the start of an instruction.
func getRelativeIndex(startOffset uint, relativeOffset int64,
	instructionIndices map[uint]int) (uint, error) {
	newOffset := int64(startOffset) + relativeOffset
	if newOffset < 0 {
		return 0, InvalidAddressError(newOffset)
	}
	nextIndex, ok := instructionIndices[uint(newOffset)]
	if !ok {
		return 0, InvalidAddressError(newOffset)
	}
	return uint(nextIndex), nil
}

func (n *ifeqInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)),
		instructionIndices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *ifneInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)),
		instructionIndices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *ifltInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *ifgeInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *ifgtInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *ifleInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpeqInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpneInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpltInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpgeInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpgtInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_icmpleInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_acmpeqInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *if_acmpneInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *gotoInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	return nil
}

func (n *jsrInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	nextIndex, e := getRelativeIndex(offset, int64(int16(n.value)), indices)
	if e != nil {
		return e
	}
	n.nextIndex = nextIndex
	// If the return address is somehow invalid, we'll just catch it at
	// runtime whenever the subroutine returns.
	n.returnIndex = indices[offset] + 1
	return nil
}

func (n *tableswitchInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	var e error
	n.defaultIndex, e = getRelativeIndex(offset, int64(int32(n.defaultOffset)),
		indices)
	if e != nil {
		return e
	}
	n.indices = make([]uint, len(n.offsets))
	for i := range n.indices {
		n.indices[i], e = getRelativeIndex(offset, int64(int32(n.offsets[i])),
			indices)
		if e != nil {
			return e
		}
	}
	return nil
}

func (n *lookupswitchInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	var e error
	n.defaultIndex, e = getRelativeIndex(offset, int64(int32(n.defaultOffset)),
		indices)
	if e != nil {
		return e
	}
	n.indices = make([]uint, len(n.pairs))
	for i := range n.pairs {
		n.indices[i], e = getRelativeIndex(offset,
			int64(int32(n.pairs[i].offset)), indices)
		if e != nil {
			return e
		}
	}
	return nil
}
