package class_file

import (
	"fmt"
)

// This file contains code used for parsing field or method descriptors, so
// that type information can be accessed.

// This is the basic type used by descriptors in JVM class files.
type FieldType interface {
	// Returns the type's name, as a string.
	String() string
}

// This is used to represent a primitive type, read from a descriptor.
// Implements the FieldType interface.
type PrimitiveFieldType byte

func (t PrimitiveFieldType) String() string {
	switch t {
	case 'B':
		return "byte"
	case 'C':
		return "char"
	case 'D':
		return "double"
	case 'F':
		return "float"
	case 'I':
		return "int"
	case 'J':
		return "long"
	case 'S':
		return "short"
	case 'V':
		return "void"
	case 'Z':
		return "boolean"
	}
	return fmt.Sprintf("Unknown primitive descriptor type: 0x%02x", byte(t))
}

// This is used to represent a reference to an instance of a class, read from
// a descriptor. Implements the FieldType interface.
type ClassInstanceType string

func (t ClassInstanceType) String() string {
	return string(t)
}

// This is used to represent a reference to an instance of an array, read from
// a descriptor. Implements the FieldType interface.
type ArrayType struct {
	// The number of dimensions in the array.
	Dimensions uint8
	// The type of object or primitive the array contains.
	ContentType FieldType
}

func (t *ArrayType) String() string {
	tmp := ""
	for i := uint8(0); i < t.Dimensions; i++ {
		tmp += "[]"
	}
	return fmt.Sprintf("%s%s", t.ContentType.String(), tmp)
}

// Parses a field descriptor referring to an instance of a class. Returns a
// FieldType and the remaining descriptor bytes, or an error if one occurs.
func parseClassInstanceDescriptor(descriptor []byte) (FieldType, []byte,
	error) {
	endIndex := -1
	// Class reference descriptors take the form L<class name>;
	for i := range descriptor {
		if descriptor[i] == ';' {
			endIndex = i
			break
		}
	}
	if endIndex <= 0 {
		return nil, nil, fmt.Errorf("Invalid descriptor: %s", descriptor)
	}
	toReturn := ClassInstanceType(string(descriptor[1:endIndex]))
	return toReturn, descriptor[endIndex+1:], nil
}

// Parses a field descriptor referring to an array. Returns a FieldType and the
// remaining descriptor bytes, or an error if one occurs.
func parseArrayDescriptor(descriptor []byte) (FieldType, []byte, error) {
	dimensions := 0
	typeStartIndex := -1
	for i := range descriptor {
		if descriptor[i] != '[' {
			typeStartIndex = i
			break
		}
		dimensions++
	}
	if dimensions > 255 {
		return nil, nil, fmt.Errorf(
			"Too many dimensions in array descriptor: %d", dimensions)
	}
	t, remaining, e := parseFieldTypeInternal(descriptor[typeStartIndex:],
		false)
	if e != nil {
		return nil, nil, e
	}
	toReturn := &ArrayType{
		Dimensions:  uint8(dimensions),
		ContentType: t,
	}
	return toReturn, remaining, nil
}

// Parses a field descriptor and returns a FieldType and the remaining
// bytes in the descriptor, or an error if one occurs.
func parseFieldTypeInternal(descriptor []byte, allowVoid bool) (FieldType,
	[]byte, error) {
	if len(descriptor) == 0 {
		return nil, nil, fmt.Errorf("Empty descriptor string")
	}
	switch descriptor[0] {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		return PrimitiveFieldType(descriptor[0]), descriptor[1:], nil
	case '[':
		return parseArrayDescriptor(descriptor)
	case 'L':
		return parseClassInstanceDescriptor(descriptor)
	case 'V':
		if !allowVoid {
			return nil, nil, fmt.Errorf("A void type descriptor is not " +
				"allowed here")
		}
		return PrimitiveFieldType('V'), descriptor[1:], nil
	}
	return nil, nil, fmt.Errorf("Invalid field descriptor: %s", descriptor)
}

// Parses a single field descriptor, returning a FieldType instance or an error
// if the descriptor string was invalid.
func ParseFieldType(descriptor []byte) (FieldType, error) {
	toReturn, _, e := parseFieldTypeInternal(descriptor, false)
	if e != nil {
		return nil, e
	}
	// NOTE: Should we check for trailing data after the descriptor?
	return toReturn, nil
}

// This contains parsed information from a method descriptor.
type MethodDescriptor struct {
	ArgumentTypes []FieldType
	ReturnType    FieldType
}

// Returns a list of the argument types, as a comma-separated string.
func (d *MethodDescriptor) ArgumentsString() string {
	toReturn := ""
	for i, v := range d.ArgumentTypes {
		if i != 0 {
			toReturn += ", "
		}
		toReturn += v.String()
	}
	return toReturn
}

// Returns the method's return type, as a string.
func (d *MethodDescriptor) ReturnString() string {
	return d.ReturnType.String()
}

func ParseMethodDescriptor(descriptor []byte) (*MethodDescriptor, error) {
	var e error
	if (len(descriptor) == 0) || (descriptor[0] != '(') {
		return nil, fmt.Errorf("Invalid method descriptor")
	}
	descriptor = descriptor[1:]
	var argument FieldType
	argumentTypes := make([]FieldType, 0, 4)
	for {
		if len(descriptor) == 0 {
			return nil, fmt.Errorf("Invalid method descriptor: missing \")\"")
		}
		if descriptor[0] == ')' {
			descriptor = descriptor[1:]
			break
		}
		argument, descriptor, e = parseFieldTypeInternal(descriptor, false)
		if e != nil {
			return nil, fmt.Errorf("Bad method argument type: %s", e)
		}
		argumentTypes = append(argumentTypes, argument)
	}
	returnType, descriptor, e := parseFieldTypeInternal(descriptor, true)
	if e != nil {
		return nil, fmt.Errorf("Bad method return type: %s", e)
	}
	// NOTE: Should we check for trailing data after the descriptor?
	toReturn := MethodDescriptor{
		ArgumentTypes: argumentTypes,
		ReturnType:    returnType,
	}
	return &toReturn, nil
}
