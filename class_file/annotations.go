package class_file

// This file contains code used for parsing annotations. There are is enough
// variety in the forms of annotation attributes that putting this in its own
// file helps keep the (relatively simpler) parsing in attributes.go from
// getting overshadowed.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// The 1-character tag indicating the type of an element value.
type ElementValueTag uint8

func (t ElementValueTag) String() string {
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
	case 'Z':
		return "boolean"
	case 's':
		return "string"
	case 'e':
		return "enum"
	case 'c':
		return "class"
	case '@':
		return "annotation"
	case '[':
		return "array"
	}
	return fmt.Sprintf("unknown element value tag %d", t)
}

// Refers to a single ElementValue. To obtain more detailed information from
// some values, type assertions should be used to convert this into a pointer
// to a struct with the concrete values.
type ElementValue interface {
	// Returns the ElementValue's 1-character tag.
	Tag() ElementValueTag
	// Returns the class info or constant value index, so that class and
	// constant element values can be read without having to use type
	// assertions. Returns 0 for enum, annotation, and array values.
	Index() uint16
}

// Holds element value tags and indices for types that use only a single index.
type basicElementValue struct {
	tag   ElementValueTag
	index uint16
}

func (v *basicElementValue) Tag() ElementValueTag {
	return v.tag
}

func (v *basicElementValue) Index() uint16 {
	return v.index
}

// Holds data for an enum element value
type EnumElementValue struct {
	basicElementValue
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

// Represents an annotation element value--which contains a nested annotation.
type AnnotationElementValue struct {
	basicElementValue
	Value *Annotation
}

// Represents an array element value, which contains a nested list of further
// element values.
type ArrayElementValue struct {
	basicElementValue
	Values []ElementValue
}

// Consumes and returns a single ElementValue. Returns an error if one occurs.
func parseElementValue(data io.Reader) (ElementValue, error) {
	var tag ElementValueTag
	e := binary.Read(data, binary.BigEndian, &tag)
	if e != nil {
		return nil, fmt.Errorf("Failed reading element value tag: %s", e)
	}
	switch tag {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 's', 'c':
		var index uint16
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading element value index: %s", e)
		}
		var toReturn basicElementValue
		toReturn.tag = tag
		toReturn.index = index
		return &toReturn, nil
	case 'e', '@', '[':
		break
	default:
		return nil, fmt.Errorf("Unknown element value tag: %d", tag)
	}
	if tag == 'e' {
		var typeName, constName uint16
		e = binary.Read(data, binary.BigEndian, &typeName)
		if e != nil {
			return nil, fmt.Errorf("Failed reading enum type name: %s", e)
		}
		e = binary.Read(data, binary.BigEndian, &constName)
		if e != nil {
			return nil, fmt.Errorf("Failed reading enum const name: %s", e)
		}
		var toReturn EnumElementValue
		toReturn.tag = tag
		toReturn.TypeNameIndex = typeName
		toReturn.ConstNameIndex = constName
		return &toReturn, nil
	}
	if tag == '@' {
		nestedAnnotation, e := parseSingleAnnotation(data)
		if e != nil {
			return nil, e
		}
		var toReturn AnnotationElementValue
		toReturn.tag = tag
		toReturn.Value = nestedAnnotation
		return &toReturn, nil
	}
	// Finally, parse an array element value.
	var count uint16
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading size of array elem. value: %s",
			e)
	}
	values := make([]ElementValue, count)
	for i := range values {
		values[i], e = parseElementValue(data)
		if e != nil {
			return nil, e
		}
	}
	var toReturn ArrayElementValue
	toReturn.tag = tag
	toReturn.Values = values
	return &toReturn, nil
}

// Holds a single element-value pair from an annotation.
type ElementValuePair struct {
	ElementNameIndex uint16
	Value            ElementValue
}

// Parses a 16-bit length of an element-value pairs table, followed by the
// table itself. Returns a slice of the table's entries.
func parseElementValuePairsTable(data io.Reader) ([]ElementValuePair, error) {
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read # of element-value pairs: %s", e)
	}
	pairs := make([]ElementValuePair, count)
	for i := range pairs {
		e = binary.Read(data, binary.BigEndian, &(pairs[i].ElementNameIndex))
		if e != nil {
			return nil, fmt.Errorf("Couldn't read element name: %s", e)
		}
		pairs[i].Value, e = parseElementValue(data)
		if e != nil {
			return nil, fmt.Errorf("Failed parsing element value: %s", e)
		}
	}
	return pairs, nil
}

// Holds information about an ordinary annotation.
type Annotation struct {
	NameIndex         uint16
	ElementValuePairs []ElementValuePair
}

func parseSingleAnnotation(data io.Reader) (*Annotation, error) {
	var nameIndex uint16
	e := binary.Read(data, binary.BigEndian, &nameIndex)
	if e != nil {
		return nil, fmt.Errorf("Failed reading annotation name index: %s", e)
	}
	pairs, e := parseElementValuePairsTable(data)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing element-value pairs: %s", e)
	}
	var toReturn Annotation
	toReturn.NameIndex = nameIndex
	toReturn.ElementValuePairs = pairs
	return &toReturn, nil
}

// Parses a uint16 count of annotations, followed by the annotations themselves
func parseAnnotationGroup(data io.Reader) ([]*Annotation, error) {
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading annotation count: %s", e)
	}
	toReturn := make([]*Annotation, count)
	for i := range toReturn {
		toReturn[i], e = parseSingleAnnotation(data)
		if e != nil {
			return nil, fmt.Errorf("Failed parsing annotation: %s", e)
		}
	}
	return toReturn, nil
}

// This can be used to parse both visible and invisible runtime annotation
// attributes.
func ParseRuntimeAnnotationsAttribute(a *Attribute) ([]*Annotation, error) {
	switch string(a.Name) {
	case "RuntimeVisibleAnnotations", "RuntimeInvisibleAnnotations":
		break
	default:
		return nil, fmt.Errorf("Expected a runtime annotations attribute")
	}
	data := bytes.NewReader(a.Info)
	return parseAnnotationGroup(data)
}

// Parses a RuntimeVisibleParameterAnnotations or a
// RuntimeInvisibleParameterAnnotations attribute. Returns a slice of slices of
// annotations--1 per parameter.
func ParseParameterAnnotationsAttribute(a *Attribute) ([][]*Annotation,
	error) {
	switch string(a.Name) {
	case "RuntimeVisibleParameterAnnotations",
		"RuntimeInvisibleParameterAnnotations":
		break
	default:
		return nil, fmt.Errorf("Expected a parameter annotations attribute")
	}
	data := bytes.NewReader(a.Info)
	var parameterCount uint8
	e := binary.Read(data, binary.BigEndian, &parameterCount)
	if e != nil {
		return nil, fmt.Errorf("Failed reading paramter count: %s", e)
	}
	toReturn := make([][]*Annotation, parameterCount)
	for i := range toReturn {
		toReturn[i], e = parseAnnotationGroup(data)
		if e != nil {
			return nil, fmt.Errorf("Failed parsing param %d annotations: %s",
				i, e)
		}
	}
	return toReturn, nil
}

// The first field in type annotations--a single byte indicating the type of
// target on which the annotation appears.
type TargetType uint8

// A single element in a type annotations Type Path.
type TypePathElement struct {
	TypePathKind      uint8
	TypeArgumentIndex uint8
}

// Reads a type path from the given input stream.
func parseTypePath(data io.Reader) ([]TypePathElement, error) {
	var length uint8
	e := binary.Read(data, binary.BigEndian, &length)
	if e != nil {
		return nil, fmt.Errorf("Failed reading type path length: %s", e)
	}
	toReturn := make([]TypePathElement, length)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading type path: %s", e)
	}
	return toReturn, nil
}

// A generic type annotation interface. Must be cast to a specific struct type
// to access fields of the target_info member.
type TypeAnnotation interface {
	Target() TargetType
	TypePath() []TypePathElement
	TypeIndex() uint16
	ElementValuePairs() []ElementValuePair
}

// Contains values common to all type annotations, and fulfills the
// TypeAnnotation interface.
type basicTypeAnnotation struct {
	target            TargetType
	typePath          []TypePathElement
	typeIndex         uint16
	elementValuePairs []ElementValuePair
}

func (a *basicTypeAnnotation) Target() TargetType {
	return a.target
}

func (a *basicTypeAnnotation) TypePath() []TypePathElement {
	return a.typePath
}

func (a *basicTypeAnnotation) TypeIndex() uint16 {
	return a.typeIndex
}

func (a *basicTypeAnnotation) ElementValuePairs() []ElementValuePair {
	return a.elementValuePairs
}

// This struct covers all type annotations with a target_info field containing
// a single member. If target_info is a 1-byte value, then the value will be
// zero-extended in this struct.
type SingleFieldTypeAnnotation struct {
	basicTypeAnnotation
	Data uint16
}

// This is used for type annotations with type_parameter_bound_target
// target_info fields.
type TypeParameterBoundAnnotation struct {
	basicTypeAnnotation
	TypeParameterIndex uint8
	BoundIndex         uint8
}

// This is used for type annotations with type_argument_target type_info
// fields.
type TypeArgumentAnnotation struct {
	basicTypeAnnotation
	Offset            uint16
	TypeArgumentIndex uint8
}

// The form of a single entry in a local variable type annotation table.
type LocalVariableTypeAnnotationEntry struct {
	StartPC uint16
	Length  uint16
	Index   uint16
}

type LocalVariableTypeAnnotation struct {
	basicTypeAnnotation
	Table []LocalVariableTypeAnnotationEntry
}

// Assuming the target_info field has already been parsed, this will parse the
// remaining fields of a type annotation. Takes a pointer to a basic type
// annotation struct, which will be filled in with the parsed data. Returns an
// error if one occurs.
func parsePostTargetInfoTypeAnnotation(data io.Reader,
	a *basicTypeAnnotation) error {
	var e error
	a.typePath, e = parseTypePath(data)
	if e != nil {
		return e
	}
	var typeIndex uint16
	e = binary.Read(data, binary.BigEndian, &typeIndex)
	if e != nil {
		return fmt.Errorf("Failed reading type index: %s", e)
	}
	a.typeIndex = typeIndex
	a.elementValuePairs, e = parseElementValuePairsTable(data)
	if e != nil {
		return fmt.Errorf("Failed parsing element-value pairs: %s", e)
	}
	return nil
}

func parseSingleTypeAnnotation(data io.Reader) (TypeAnnotation, error) {
	var tag TargetType
	e := binary.Read(data, binary.BigEndian, &tag)
	if e != nil {
		return nil, fmt.Errorf("Failed reading type annotation tag: %s", e)
	}
	switch tag {
	case 0, 1, 0x16:
		// Will either be a type parameter or formal parameter index.
		var index uint8
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		var toReturn SingleFieldTypeAnnotation
		toReturn.target = tag
		toReturn.Data = uint16(index)
		e = parsePostTargetInfoTypeAnnotation(data,
			&(toReturn.basicTypeAnnotation))
		if e != nil {
			return nil, e
		}
		return &toReturn, nil
	case 0x13, 0x14, 0x15:
		// target_info is an empty_target, so we don't need to parse anything
		// extra here.
		var toReturn basicTypeAnnotation
		toReturn.target = tag
		e = parsePostTargetInfoTypeAnnotation(data, &toReturn)
		if e != nil {
			return nil, e
		}
		return &toReturn, nil
	case 0x10, 0x17, 0x42, 0x43, 0x44, 0x45, 0x46:
		// target_info is either a supertype, throws, or catch index, or an
		// offset.
		var index uint16
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		var toReturn SingleFieldTypeAnnotation
		toReturn.target = tag
		toReturn.Data = index
		e = parsePostTargetInfoTypeAnnotation(data,
			&(toReturn.basicTypeAnnotation))
		if e != nil {
			return nil, e
		}
		return &toReturn, nil
	case 0x11, 0x12:
		// target_info is a type_parameter_bound_target struct.
		var toReturn TypeParameterBoundAnnotation
		toReturn.target = tag
		var index uint8
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		toReturn.TypeParameterIndex = index
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		toReturn.BoundIndex = index
		e = parsePostTargetInfoTypeAnnotation(data,
			&(toReturn.basicTypeAnnotation))
		if e != nil {
			return nil, e
		}
		return &toReturn, nil
	case 0x47, 0x48, 0x49, 0x4a, 0x4b:
		// target_info is a type_parameter_target struct.
		var toReturn TypeArgumentAnnotation
		toReturn.target = tag
		var offset uint16
		var index uint8
		e = binary.Read(data, binary.BigEndian, &offset)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		e = binary.Read(data, binary.BigEndian, &index)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		toReturn.Offset = offset
		toReturn.TypeArgumentIndex = index
		e = parsePostTargetInfoTypeAnnotation(data,
			&(toReturn.basicTypeAnnotation))
		if e != nil {
			return nil, e
		}
	case 0x40, 0x41:
		// target_info is a localvar_target struct.
		var toReturn LocalVariableTypeAnnotation
		toReturn.target = tag
		var count uint16
		e = binary.Read(data, binary.BigEndian, &count)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		table := make([]LocalVariableTypeAnnotationEntry, count)
		e = binary.Read(data, binary.BigEndian, table)
		if e != nil {
			return nil, fmt.Errorf("Failed reading target_info: %s", e)
		}
		toReturn.Table = table
		e = parsePostTargetInfoTypeAnnotation(data,
			&(toReturn.basicTypeAnnotation))
		if e != nil {
			return nil, e
		}
		return &toReturn, nil
	}
	return nil, fmt.Errorf("Unknown type annotation target type: %d", tag)
}

// This can be used to parse both visible and invisible type annotation
// attributes.
func ParseTypeAnnotationsAttribute(a *Attribute) ([]TypeAnnotation, error) {
	switch string(a.Name) {
	case "RuntimeVisibleTypeAnnotations", "RuntimeInvisibleTypeAnnotations":
		break
	default:
		return nil, fmt.Errorf("Expected a type annotations attribute")
	}
	data := bytes.NewReader(a.Info)
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading number of type annotations: %s",
			e)
	}
	toReturn := make([]TypeAnnotation, count)
	for i := range toReturn {
		toReturn[i], e = parseSingleTypeAnnotation(data)
		if e != nil {
			return nil, fmt.Errorf("Failed parsing type annotation: %s", e)
		}
	}
	return toReturn, nil
}

// Parses and returns the ElementValue contained in an AnnotationDefault
// attribute.
func ParseAnnotationDefaultAttribute(a *Attribute) (ElementValue, error) {
	if string(a.Name) != "AnnotationDefault" {
		return nil, fmt.Errorf("Expected an AnnotationDefault attribute")
	}
	data := bytes.NewReader(a.Info)
	return parseElementValue(data)
}
