package jvm

// This file contains definitions used when parsing a class file's constant
// pool.

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Used to identify constant types in the class file's constant table.
type ConstantTag uint8

func (t ConstantTag) String() string {
	switch t {
	case 1:
		return "utf-8 value"
	case 3:
		return "integer"
	case 4:
		return "float"
	case 5:
		return "long"
	case 6:
		return "double"
	case 7:
		return "class"
	case 8:
		return "string"
	case 9:
		return "field"
	case 10:
		return "method"
	case 11:
		return "interface method"
	case 12:
		return "name and type"
	case 15:
		return "method handle"
	case 16:
		return "method type"
	case 18:
		return "InvokeDynamic information"
	}
	return fmt.Sprintf("unknown tag %d", uint8(t))
}

// Returns true if the constant counts as 2 entries in the constant table
// rather than 1.
func (t ConstantTag) CountsDouble() bool {
	switch t {
	case 5, 6:
		return true
	}
	return false
}

// The high-level qualities available for all constants. Specific information
// can be obtained by using type assertions to convert this to a pointer to one
// of the Constant<..> structs.
type ClassFileConstant interface {
	Tag() ConstantTag
	String() string
}

// Represents a class or interface
type ConstantClassInfo struct {
	// The index of a UTF-8 constant containing the class name
	ClassNameIndex uint16
}

func (n *ConstantClassInfo) Tag() ConstantTag {
	return ConstantTag(7)
}

func (n *ConstantClassInfo) String() string {
	return fmt.Sprintf("%s, name index %d", n.Tag(), n.ClassNameIndex)
}

// Contains information about a field in the class
type ConstantFieldInfo struct {
	// The index of a class info constant containing this field.
	ClassIndex uint16
	// The index of a name and type constant for this field.
	NameAndTypeIndex uint16
}

func (n *ConstantFieldInfo) Tag() ConstantTag {
	return ConstantTag(9)
}

func (n *ConstantFieldInfo) String() string {
	return fmt.Sprintf("%s, class index %d, name and type index %d", n.Tag(),
		n.ClassIndex, n.NameAndTypeIndex)
}

// Contains information about a method in a class.
type ConstantMethodInfo struct {
	// The index of a class info constant containing this or method.
	ClassIndex uint16
	// The index of a name and type constant for this method.
	NameAndTypeIndex uint16
}

func (n *ConstantMethodInfo) Tag() ConstantTag {
	return ConstantTag(10)
}

func (n *ConstantMethodInfo) String() string {
	return fmt.Sprintf("%s, class index %d, name and type index %d", n.Tag(),
		n.ClassIndex, n.NameAndTypeIndex)
}

// Contains information about a method in an interface.
type ConstantInterfaceMethodInfo struct {
	// The index of a class (interface) constant containing this method.
	ClassIndex uint16
	// The index of a name and type constant for this method.
	NameAndTypeIndex uint16
}

func (n *ConstantInterfaceMethodInfo) Tag() ConstantTag {
	return ConstantTag(11)
}

func (n *ConstantInterfaceMethodInfo) String() string {
	return fmt.Sprintf("%s, class index %d, name and type index %d", n.Tag(),
		n.ClassIndex, n.NameAndTypeIndex)
}

// Holds information about constant String values.
type ConstantStringInfo struct {
	// The index of the UTF-8 constant containing the string's initial value.
	StringIndex uint16
}

func (n *ConstantStringInfo) Tag() ConstantTag {
	return ConstantTag(8)
}

func (n *ConstantStringInfo) String() string {
	return fmt.Sprintf("%s, utf-8 index %d", n.Tag(), n.StringIndex)
}

// Holds information about a 4-byte integer constant
type ConstantIntegerInfo struct {
	Value int32
}

func (n *ConstantIntegerInfo) Tag() ConstantTag {
	return ConstantTag(3)
}

func (n *ConstantIntegerInfo) String() string {
	return fmt.Sprintf("%s, %d", n.Tag(), n.Value)
}

// Holds information about a 4-byte float constant
type ConstantFloatInfo struct {
	Value float32
}

func (n *ConstantFloatInfo) Tag() ConstantTag {
	return ConstantTag(4)
}

func (n *ConstantFloatInfo) String() string {
	return fmt.Sprintf("%s, %f", n.Tag(), n.Value)
}

// Holds information about an 8-byte integer constant
type ConstantLongInfo struct {
	Value int64
}

func (n *ConstantLongInfo) Tag() ConstantTag {
	return ConstantTag(5)
}

func (n *ConstantLongInfo) String() string {
	return fmt.Sprintf("%s, %d", n.Tag(), n.Value)
}

// Holds information about an 8-byte float constant
type ConstantDoubleInfo struct {
	Value float64
}

func (n *ConstantDoubleInfo) Tag() ConstantTag {
	return ConstantTag(6)
}

func (n *ConstantDoubleInfo) String() string {
	return fmt.Sprintf("%s, %f", n.Tag(), n.Value)
}

// Represents a field or method without referring to the class or interface it
// belongs to.
type ConstantNameAndTypeInfo struct {
	// The index of a UTF-8 constant containing the method or field's name
	NameIndex uint16
	// The index of a UTF-8 constant containing the field or method descriptor
	DescriptorIndex uint16
}

func (n *ConstantNameAndTypeInfo) Tag() ConstantTag {
	return ConstantTag(12)
}

func (n *ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("%s, name index %d, descriptor index %d", n.Tag(),
		n.NameIndex, n.DescriptorIndex)
}

// Contains a UTF-8 string's bytes.
type ConstantUTF8Info struct {
	Bytes []byte
}

func (n *ConstantUTF8Info) Tag() ConstantTag {
	return ConstantTag(1)
}

func (n *ConstantUTF8Info) String() string {
	return fmt.Sprintf("%s, %q", n.Tag(), n.Bytes)
}

// Holds the kind of method handle reference in a method handle constant.
type MethodHandleReferenceKind uint8

func (k MethodHandlReferenceKind) String() string {
	switch k {
	case 1:
		return "get field"
	case 2:
		return "get static"
	case 3:
		return "put field"
	case 4:
		return "put static"
	case 5:
		return "invoke virtual"
	case 6:
		return "invoke static"
	case 7:
		return "invoke special"
	case 8:
		return "new invoke special"
	case 9:
		return "invoke interface"
	}
	return fmt.Sprintf("unkown method handle kind %d", uint8(k))
}

// Represents a method handle used during runtime.
type ConstantMethodHandleInfo struct {
	// Indicates what the method handle will be used for.
	ReferenceKind MethodHandleReferenceKind
	// An index into the constant table; the meaning depends on the reference
	// kind.
	Index uint16
}

func (n *ConstantMethodHandleInfo) Tag() ConstantTag {
	return ConstantTag(15)
}

func (n *ConstantMethodHandleInfo) String() string {
	return fmt.Sprintf("%s, kind = %s, index = %d", n.Tag(), n.ReferenceKind,
		n.Index)
}

// Holds a method's type.
type ConstantMethodTypeInfo struct {
	// An index into the constant table of a UTF-8 descriptor.
	DescriptorIndex uint16
}

func (n *ConstantMethodTypeInfo) Tag() ConstantTag {
	return ConstantTag(16)
}

func (n *ConstantMethodTypeInfo) String() string {
	return fmt.Sprintf("%s, descriptor index = %d", n.Tag(), n.DescriptorIndex)
}

// Used by the invokedynamic instruction.
type ConstantInvokeDynamicInfo struct {
	// An index into the bootstrap method array in the bootstrap methods table
	// (in the class file's attributes).
	BootstrapMethodAttributeIndex uint16
	// An index into the constants of a ConstantNameAndTypeInfo structure.
	NameAndTypeIndex uint16
}

func (n *ConstantInvokeDynamicInfo) Tag() ConstantTag {
	return ConstantTag(18)
}

func (n *ConstantInvokeDynamicInfo) String() string {
	return fmt.Sprintf(
		"%s, bootstrap method attribute index %d, name and type index %d",
		n.Tag(), n.BootstrapMethodAttributeIndex, n.NameAndTypeIndex)
}

// Parses and returns a single class file constant in the table.
func parseSingleClassConstant(data io.Reader) (ClassFileConstant, error) {
	var tag ConstantTag
	var toReturn ClassFileConstant
	e := binary.Read(data, binary.BigEndian, &tag)
	if e != nil {
		return nil, fmt.Errorf("Failed reading constant tag: %s", e)
	}
	switch tag {
	case 1:
		var length uint16
		e = binary.Read(data, binary.BigEndian, &length)
		if e != nil {
			return nil, fmt.Errorf("Failed reading utf-8 length: %s", e)
		}
		utf8Bytes := make([]byte, length)
		e = binary.Read(data, binary.BigEndian, utf8Bytes)
		if e != nil {
			return nil, fmt.Errorf("Failed reading utf-8 constant: %s", e)
		}
		toReturn = &ConstantUTF8Info{
			Bytes: utf8Bytes,
		}
	case 3:
		var value ConstantIntegerInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading integer constant: %s", e)
		}
		toReturn = &value
	case 4:
		var value ConstantFloatInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading float constant: %s", e)
		}
		toReturn = &value
	case 5:
		var value ConstantLongInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading long constant: %s", e)
		}
		toReturn = &value
	case 6:
		var value ConstantDoubleInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading double constant: %s", e)
		}
		toReturn = &value
	case 7:
		var value ConstantClassInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading class constant: %s", e)
		}
		toReturn = &value
	case 8:
		var value ConstantStringInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading string constant: %s", e)
		}
		toReturn = &value
	case 9:
		var value ConstantFieldInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading field constant: %s", e)
		}
		toReturn = &value
	case 10:
		var value ConstantMethodInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading method constant: %s", e)
		}
		toReturn = &value
	case 11:
		var value ConstantInterfaceMethodInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf(
				"Failed reading interface method constant: %s", e)
		}
		toReturn = &value
	case 12:
		var value ConstantNameAndTypeInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading name and type constant: %s",
				e)
		}
		toReturn = &value
	case 15:
		var value ConstantMethodHandleInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading method handle constant: %s",
				e)
		}
		toReturn = &value
	case 16:
		var value ConstantMethodTypeInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf("Failed reading method type constant: %s",
				e)
		}
	case 18:
		var value ConstantInvokeDynamicInfo
		e = binary.Read(data, binary.BigEndian, &value)
		if e != nil {
			return nil, fmt.Errorf(
				"Failed reading invokedynamic information constant: %s", e)
		}
	default:
		return nil, fmt.Errorf("Unknown class file constant: %s", tag)
	}
	return toReturn, nil
}

// Assumes the data reader is at the start of the ConstantPoolInfo table in the
// class file. Fills in the ConstantPoolInfo struct in the ClassFile struct.
func parseClassConstantsTable(data io.Reader, class *ClassFile,
	count uint16) error {
	var e error
	var constant ClassFileConstant
	// Note that since long and double constants are counted twice, this
	// may actually be longer than needed...
	constants := make([]ClassFileConstant, 0, count)
	remaining := int(count)
	for remaining > 0 {
		constant, e = parseSingleClassConstant(data)
		if e != nil {
			return e
		}
		remaining--
		if constant.Tag().CountsDouble() {
			remaining--
		}
		constants = append(constants, constant)
	}
	// This should only be possible if we encounter a long or double constant
	// which wasn't accounted for properly in the count.
	if remaining < 0 {
		return fmt.Errorf("Invalid class file constant count: %d", count)
	}
	// Allocate a new slice that contains the actual number of constants rather
	// than the amount contained in "count".
	class.Constants = make([]ClassFileConstant, len(constants))
	copy(class.Constants, constants)
	return nil
}
