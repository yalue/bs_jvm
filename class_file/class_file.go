// This package contains utilities for parsing JVM class files.
package class_file

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Holds information about a class' permissions or type--including whether it's
// public or private.
type ClassAccessFlags uint16

func (f ClassAccessFlags) String() string {
	toReturn := ""
	if (f & 0x0001) != 0 {
		toReturn += "public "
	}
	if (f & 0x0002) != 0 {
		toReturn += "private "
	}
	if (f & 0x0004) != 0 {
		toReturn += "protected "
	}
	if (f & 0x0008) != 0 {
		toReturn += "static "
	}
	if (f & 0x0010) != 0 {
		toReturn += "final "
	}
	if (f & 0x0020) != 0 {
		toReturn += "super "
	}
	if (f & 0x0200) != 0 {
		toReturn += "interface "
	}
	if (f & 0x0400) != 0 {
		toReturn += "abstract "
	}
	if (f & 0x1000) != 0 {
		toReturn += "synthetic "
	}
	if (f & 0x2000) != 0 {
		toReturn += "annotation "
	}
	if (f & 0x4000) != 0 {
		toReturn += "enum "
	}
	return strings.TrimRight(toReturn, " ")
}

// Holds relevant data from a parsed class file.
type ClassFile struct {
	MinorVersion uint16
	MajorVersion uint16
	Constants    []Constant
	// Permissions of the class or interface, such as public or private
	Access ClassAccessFlags
	// ThisClass, SuperClass and Interfaces are all indices into the constant
	// table.
	ThisClass  uint16
	SuperClass uint16
	Interfaces []uint16
	Fields     []*Field
	Methods    []*Method
	Attributes []*Attribute
}

// Returns the constant with the given index, or an error if the index is
// invalid.
func (c *ClassFile) GetConstant(index uint16) (Constant, error) {
	if index == 0 {
		return nil, fmt.Errorf("Constant indices must be greater than 0")
	}
	if int(index) > len(c.Constants) {
		return nil, fmt.Errorf("Invalid constant index: %d", index)
	}
	toReturn := c.Constants[index]
	if toReturn == nil {
		return nil, fmt.Errorf("Constant index %d is invald", index)
	}
	return toReturn, nil
}

// Returns the constant at the given index, but only if it as a UTF-8 string.
// Returns an error in any other case.
func (c *ClassFile) GetUTF8Constant(index uint16) ([]byte, error) {
	value, e := c.GetConstant(index)
	if e != nil {
		return nil, e
	}
	toReturn, ok := value.(*ConstantUTF8Info)
	if !ok {
		return nil, fmt.Errorf("Constant %s is not a UTF-8 constant", value)
	}
	return toReturn.Bytes, nil
}

// Parses a class file; returns an error if the file is not valid.
func ParseClassFile(data io.Reader) (*ClassFile, error) {
	var toReturn ClassFile
	var magic uint32
	e := binary.Read(data, binary.BigEndian, &magic)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read class file magic number: %s", e)
	}
	if magic != 0xcafebabe {
		return nil, fmt.Errorf("Invalid class file magic: 0x%08x", magic)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.MinorVersion))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read minor version: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.MajorVersion))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read major version: %s", e)
	}
	var count uint16
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read the constant pool count: %s", e)
	}
	constants, e := parseConstantsTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing constant pool: %s", e)
	}
	toReturn.Constants = constants
	e = binary.Read(data, binary.BigEndian, &(toReturn.Access))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read the class' access flags: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.ThisClass))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read this class' info index: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.SuperClass))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read the superclass' info: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read the number of interfaces: %s", e)
	}
	interfaces := make([]uint16, count)
	e = binary.Read(data, binary.BigEndian, interfaces)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read the interfaces list: %s", e)
	}
	toReturn.Interfaces = interfaces
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't parse the number of fields: %s", e)
	}
	fields, e := (&toReturn).parseFieldTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing fields: %s", e)
	}
	toReturn.Fields = fields
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't parse the number of methods: %s", e)
	}
	methods, e := (&toReturn).parseMethodTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing methods: %s", e)
	}
	toReturn.Methods = methods
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't parse the attribute count: %s", e)
	}
	attributes, e := (&toReturn).parseAttributesTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing attributes: %s", e)
	}
	toReturn.Attributes = attributes
	return &toReturn, nil
}
