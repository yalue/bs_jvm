package bs_jvm

// This file contains instruction-specific Optimize function definitions.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
	"math"
)

// Converts a constant to a value that can be pushed by ldc or ldc_w. Returns
// the primitive value or reference to push. Returns true if the returned value
// is a primitive.
func constantToLdcInfo(class *Class, c class_file.Constant) (Int, Object, bool,
	error) {
	var primitive Int
	var reference Object = nil
	isPrimitive := false
	var e error
	// Primitive types can be converted now, but references to other objects
	// can be handled later--for now just read them from the class file.
	switch v := c.(type) {
	case *class_file.ConstantIntegerInfo:
		primitive = Int(v.Value)
		reference = Int(v.Value)
		isPrimitive = true
	case *class_file.ConstantFloatInfo:
		primitive = Int(math.Float32bits(v.Value))
		reference = Float(v.Value)
		isPrimitive = true
	case *class_file.ConstantStringInfo, *class_file.ConstantMethodHandleInfo,
		*class_file.ConstantMethodTypeInfo, *class_file.ConstantClassInfo:
		reference, e = ConvertConstantToObject(class, c)
		if e != nil {
			return 0, nil, false, e
		}
	default:
		// ldc only allows the types of constants listed above.
		return 0, nil, false, TypeError(fmt.Sprintf("Invalid ldc constant: %s",
			c))
	}
	return primitive, reference, isPrimitive, nil
}

func (n *ldcInstruction) Optimize(m *Method, offset uint,
	instructionIndices map[uint]int) error {
	constant, e := m.ContainingClass.File.GetConstant(uint16(n.value))
	if e != nil {
		return e
	}
	primitive, reference, isPrimitive, e := constantToLdcInfo(
		m.ContainingClass, constant)
	if e != nil {
		return e
	}
	n.isPrimitive = isPrimitive
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
	primitive, reference, isPrimitive, e := constantToLdcInfo(
		m.ContainingClass, constant)
	if e != nil {
		return e
	}
	n.isPrimitive = isPrimitive
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

// This just lets us type-check the method's return type during the optimize
// pass rather than at runtime.
func (n *ireturnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	t, ok := m.Types.ReturnType.(class_file.PrimitiveFieldType)
	if !ok {
		return TypeError("Encountered ireturn in a function that doesn't " +
			"return a primitive")
	}
	switch t {
	case 'Z', 'B', 'C', 'S', 'I':
		return nil
	}
	return TypeError(fmt.Sprintf("Encountered ireturn in a function that "+
		"returns a %s", t.String()))
}

// Similar to ireturn's Optimize(...), this just makes sure that the lreturn
// is being called from a method that actually returns a long.
func (n *lreturnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	t, ok := m.Types.ReturnType.(class_file.PrimitiveFieldType)
	if !ok {
		return TypeError("Encountered lreturn in a function that doesn't " +
			"return a primitive")
	}
	if t == 'J' {
		return nil
	}
	return TypeError(fmt.Sprintf("Encountered lreturn in a function that "+
		"returns a %s", t.String()))
}

// Similar to ireturn's Optimize(...), this just checks that a method using
// an freturn instruction is supposed to return a float.
func (n *freturnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	t, ok := m.Types.ReturnType.(class_file.PrimitiveFieldType)
	if !ok {
		return TypeError("Encountered freturn in a function that doesn't " +
			"return a primitive")
	}
	if t == 'F' {
		return nil
	}
	return TypeError(fmt.Sprintf("Encountered freturn in a function that "+
		"returns a %s", t.String()))
}

// Similar to ireturn's Optimize(...), this just checks that a method using
// a dreturn instruction is supposed to return a double.
func (n *dreturnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	t, ok := m.Types.ReturnType.(class_file.PrimitiveFieldType)
	if !ok {
		return TypeError("Encountered dreturn in a function that doesn't " +
			"return a primitive")
	}
	if t == 'D' {
		return nil
	}
	return TypeError(fmt.Sprintf("Encountered dreturn in a function that "+
		"returns a %s", t.String()))
}

// Similar to ireturn's Optimize(...) but checks that a method using areturn
// actually returns a reference rather than a primitive.
func (n *areturnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	// Only arrays and object references should be returned using areturn.
	switch m.Types.ReturnType.(type) {
	case class_file.ClassInstanceType, *class_file.ArrayType:
		return nil
	}
	return TypeError(fmt.Sprintf("Encountered areturn in a function that "+
		"returns a %s", m.Types.ReturnType.String()))
}

// Similar to the Optimize(...) functions for other return instructions-- but
// in this case just ensures the method has a return type of void.
func (n *returnInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	t, ok := m.Types.ReturnType.(class_file.PrimitiveFieldType)
	if !ok || (t != 'V') {
		return TypeError("Encountered a return instruction in a function that" +
			" doesn't return void")
	}
	return nil
}

// Takes a class and an index into that class' constant pool. The index must
// specify a field info constant (*class_file.ConstantFieldInfo). Returns the
// resolved constant.
func resolveStaticFieldConstant(currentClass *Class, index uint16) (
	*FieldOrMethodReference, error) {
	classFile := currentClass.File
	constant, e := classFile.GetConstant(index)
	if e != nil {
		return nil, FieldError(fmt.Sprintf("Couldn't get field info "+
			"constant: %s", e))
	}
	// We'll do this check here, because ConvertConstantToObject will also work
	// with a Method constant, which we want to make sure we didn't get.
	// (Though such a thing should never happen in practice...)
	_, ok := constant.(*class_file.ConstantFieldInfo)
	if !ok {
		return nil, FieldError(fmt.Sprintf("Didn't get a field info "+
			"constant, but instead got: %s", constant.String()))
	}
	tmp, e := ConvertConstantToObject(currentClass, constant)
	if e != nil {
		return nil, FieldError(fmt.Sprintf("Error processing field info "+
			"constant: %s", e))
	}
	fieldRef, ok := tmp.(*FieldOrMethodReference)
	if !ok {
		return nil, FieldError(fmt.Sprintf("Didn't get expected field "+
			"reference object, instead got: %s", tmp.String()))
	}
	return fieldRef, nil
}

// Figures out the class and field to get, and makes sure the field is static.
func (n *getstaticInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	fieldInfo, e := resolveStaticFieldConstant(m.ContainingClass, n.value)
	if e != nil {
		return fmt.Errorf("Failed resolving field for getstatic "+
			"instruction: %s", e)
	}
	var index int
	fieldName := string(fieldInfo.Field.Name)
	// Note that resolving the field may change the target class (if it's
	// defined in a superclass, for example)
	targetClass, index, e := fieldInfo.C.ResolveStaticField(fieldName)
	if e != nil {
		return fmt.Errorf("Couldn't resolve static field %s in class %s: %s",
			fieldName, fieldInfo.C.Name, e)
	}
	n.class = targetClass
	n.index = index
	return nil
}

func (n *putstaticInstruction) Optimize(m *Method, offset uint,
	indices map[uint]int) error {
	// Basically the same as for getstatic.
	fieldInfo, e := resolveStaticFieldConstant(m.ContainingClass, n.value)
	if e != nil {
		return fmt.Errorf("Failed resolving field for putstatic "+
			"instruction: %s", e)
	}
	var index int
	fieldName := string(fieldInfo.Field.Name)
	targetClass, index, e := fieldInfo.C.ResolveStaticField(fieldName)
	if e != nil {
		return fmt.Errorf("Couldn't resolve static field %s in class %s: %s",
			fieldName, fieldInfo.C.Name, e)
	}
	n.class = targetClass
	n.index = index
	return nil
}
