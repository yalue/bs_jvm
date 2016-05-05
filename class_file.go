package jvm

// This file contains functions and types needed for parsing class files.

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Holds information about a class' permissions or type--including whether it's
// public or private.
type AccessFlags uint16

func (f AccessFlags) String() string {
	toReturn := ""
	if (f & 0x0001) != 0 {
		toReturn += "public "
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
	Constants    []ClassFileConstant
	// Permissions of the class or interface, such as public or private
	AcessFlags uint16
	// ThisClass, SuperClass and Interfaces, are all indices into the constant
	// table.
	ThisClass  uint16
	SuperClass uint16
	Interfaces []uint16
	Fields     []FieldInfo
	Methods    []MethodInfo
	Attributes []AttributeInfo
}

// Parses a class file; returns an error if the file is unparsable or invalid.
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
	e = parseClassConstantsTable(data, &toReturn, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing constant pool: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.AccessFlags))
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
	e = parseFieldTable(data, &toReturn, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing fields: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't parse the number of methods: %s", e)
	}
	e = parseMethodTable(data, &toReturn, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing methods: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't parse the attribute count: %s", e)
	}
	e = parseAttributesTable(data, &toReturn, count)
	if e != nil {
		return nil, fmt.Errorf("Failed parsing attributes: %s", e)
	}
	return &toReturn, nil
}
