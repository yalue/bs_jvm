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

// Holds information about an ordinary annotation.
type Annotation struct {
	NameIndex         uint16
	ElementValuePairs []ElementValuePair
}

func parseSingleAnnotation(data io.Reader) (*Annotation, error) {
	var nameIndex uint16
	var count uint16
	e := binary.Read(data, binary.BigEndian, &nameIndex)
	if e != nil {
		return nil, fmt.Errorf("Failed reading annotation name index: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading size of annotation: %s", e)
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

// TODO: Parse type annotations (ugh!)
