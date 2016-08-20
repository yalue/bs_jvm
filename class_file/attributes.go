package class_file

// This file contains definitions used when parsing a class file's attribute
// structures.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Reads a uint16 at the start of the given slice. The slice must contain at
// least 2 bytes.
func readUint16BigEndian(data []byte) uint16 {
	return (uint16(data[0]) << 8) | uint16(data[1])
}

// Attributes are used to contain data in several structures within a class
// file. May be further parsed by other functions.
type Attribute struct {
	// A UTF-8 string representing the attribute's name.
	Name []byte
	// The actual bytes of the attribute.
	Info []byte
}

func (a *Attribute) String() string {
	return fmt.Sprintf("Attribute name %s, info bytes %d", a.Name, len(a.Info))
}

// This attribute only contains a reference to a single constant.
type ConstantValueAttribute struct {
	Value Constant
}

func (a *Attribute) ToConstantValueAttribute(c *ClassFile) (
	*ConstantValueAttribute, error) {
	if len(a.Info) != 2 {
		return nil, fmt.Errorf("Constant value attributes must be 2 bytes.")
	}
	index := readUint16BigEndian(a.Info)
	value, e := c.GetConstant(index)
	if e != nil {
		return nil, fmt.Errorf("Invalid constant value attribute: %s", e)
	}
	return &ConstantValueAttribute{Value: value}, nil
}

// Contains parsed data from a code attribute
type CodeAttribute struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	Attributes     []*Attribute
}

func ParseCodeAttribute(a *Attribute, c *ClassFile) (*CodeAttribute, error) {
	var toReturn CodeAttribute
	var e error
	data := bytes.NewReader(a.Info)
	e = binary.Read(data, binary.BigEndian, &(toReturn.MaxStack))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code max stack: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &(toReturn.MaxLocals))
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code max locals: %s", e)
	}
	var codeLength uint32
	e = binary.Read(data, binary.BigEndian, &codeLength)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code length: %s", e)
	}
	// Even though code length is 4 bytes, the spec limits it to these values.
	if (codeLength == 0) || (codeLength > 0xffff) {
		return nil, fmt.Errorf("Invalid code length: %d", codeLength)
	}
	code := make([]byte, codeLength)
	e = binary.Read(data, binary.BigEndian, code)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code: %s", e)
	}
	toReturn.Code = code
	var count uint16
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read exception table length: %s", e)
	}
	toReturn.ExceptionTable, e = parseExceptionTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Error reading exception table: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code attribute count: %s", e)
	}
	toReturn.Attributes, e = c.parseAttributesTable(data, count)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read code attributes: %s", e)
	}
	return &toReturn, nil
}

// Parses an Exceptions attribute. Returns an error if one occurs, otherwise
// returns a slice of exception table indices.
func ParseExceptionsAttribute(a *Attribute) ([]uint16, error) {
	if string(a.Name) != "Exceptions" {
		return nil, fmt.Errorf("Expected an exceptions attribute.")
	}
	data := bytes.NewReader(a.Info)
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading number of exceptions: %s", e)
	}
	toReturn := make([]uint16, count)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading exception indices: %s", e)
	}
	return toReturn, nil
}

// Contains parsed inner class information from an InnerClasses attribute.
type InnerClass struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags ClassAccessFlags
}

// Parses an InnerClasses attribute, returning a slice of InnerClass structs.
func ParseInnerClassesAttribute(a *Attribute) ([]InnerClass, error) {
	if string(a.Name) != "InnerClasses" {
		return nil, fmt.Errorf("Expected an InnerClasses attribute")
	}
	data := bytes.NewReader(a.Info)
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading number of InnerClasses: %s", e)
	}
	toReturn := make([]InnerClass, count)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading InnerClasses table: %s", e)
	}
	return toReturn, nil
}

// Returns the class index and method index, respectively, contained in an
// EnclosingMethod attribute.
func ParseEnclosingMethodAttribute(a *Attribute) (uint16, uint16, error) {
	if string(a.Name) != "EnclosingMethod" {
		return nil, fmt.Errorf("Expected an EnclosingMethod attribute")
	}
	data := bytes.NewReader(a.Info)
	var classIndex, methodIndex uint16
	e := binary.Read(data, binary.BigEndian, &classIndex)
	if e != nil {
		return 0, 0, fmt.Errorf("Failed reading class index: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &methodIndex)
	if e != nil {
		return 0, 0, fmt.Errorf("Failed reading method index: %s", e)
	}
	return classIndex, methodIndex, nil
}

// Returns the signature index contained in a signature attribute.
func ParseSignatureAttribute(a *Attribute) (uint16, error) {
	if string(a.Name) != "Signature" {
		return 0, fmt.Errorf("Expected a signature attribute")
	}
	var toReturn uint16
	data := bytes.NewReader(a.Info)
	e := binary.Read(data, binary.BigEndian, &toReturn)
	if e != nil {
		return 0, fmt.Errorf("Failed reading signature index: %s", e)
	}
	return toReturn, nil
}

// Returns the source file index contained in a source file attribute.
func ParseSignatureAttribute(a *Attribute) (uint16, error) {
	if string(a.Name) != "SourceFile" {
		return 0, fmt.Errorf("Expected a source file attribute")
	}
	var toReturn uint16
	data := bytes.NewReader(a.Info)
	e := binary.Read(data, binary.BigEndian, &toReturn)
	if e != nil {
		return 0, fmt.Errorf("Failed reading source file index: %s", e)
	}
	return toReturn, nil
}

// Holds a single entry from a line number table
type LineNumberEntry struct {
	StartPC    uint16
	LineNumber uint16
}

// Parses a line number table attribute into a slice of line number entries.
func ParseLineNumberTableAttribute(a *Attribute) ([]LineNumberEntry, error) {
	if string(a.Name) != "LineNumberTable" {
		return nil, fmt.Errorf("Expected a line number table attribute")
	}
	var count uint16
	data := bytes.NewReader(a.Info)
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading size of line number table: %s",
			e)
	}
	toReturn := make([]LineNumberEntry, count)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading the line number table: %s", e)
	}
	return toReturn, nil
}

// Holds a single entry from a local variable table
type LocalVariableEntry struct {
	StartPC         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

// Parses a local variable table attribute, returning a slice of entries.
func ParseLocalVariableTableAttribute(a *Attribute) ([]LocalVariableEntry,
	error) {
	if string(a.Name) != "LocalVariableTable" {
		return nil, fmt.Errorf("Expected a local variable table attribute")
	}
	data := bytes.NewReader(a.Info)
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf("Failed reading size of local var table: %s", e)
	}
	toReturn := make([]LocalVariableEntry, count)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading local variable table: %s", e)
	}
	return toReturn, nil
}

// Holds a single entry from a local variable type table
type LocalVariableTypeEntry struct {
	StartPC        uint16
	Length         uint16
	NameIndex      uint16
	SignatureIndex uint16
	Index          uint16
}

// Parses a local variable type table attribute, returning a slice of entries.
func ParseLocalVariableTypeTableAttribute(a *Attribute) (
	[]LocalVariableTypeEntry, e) {
	if string(a.Name) != "LocalVariableTypeTable" {
		return nil, fmt.Errorf("Expected a local variable type table")
	}
	data := bytes.NewReader(a.Info)
	var count uint16
	e := binary.Read(data, binary.BigEndian, &count)
	if e != nil {
		return nil, fmt.Errorf(
			"Couldn't read local variable type table size: %s", e)
	}
	toReturn := make([]LocalVariableTypeEntry, count)
	e = binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, fmt.Errorf("Failed reading local variable type table: %s",
			e)
	}
	return toReturn, nil
}

// Assumes the data reader is at the start of a class file attribute struct.
// Parses and returns the struct, or an error if one occurs.
func (c *ClassFile) parseSingleAttribute(data io.Reader) (*Attribute, error) {
	var toReturn Attribute
	var index uint16
	e := binary.Read(data, binary.BigEndian, &index)
	if e != nil {
		return nil, fmt.Errorf("Failed reading attribute name index: %s", e)
	}
	toReturn.Name, e = c.GetUTF8Constant(index)
	if e != nil {
		return nil, fmt.Errorf("Invalid attribute name: %s", e)
	}
	var infoLength uint32
	e = binary.Read(data, binary.BigEndian, &infoLength)
	if e != nil {
		return nil, fmt.Errorf("Failed reading attribute length: %s", e)
	}
	info := make([]byte, infoLength)
	e = binary.Read(data, binary.BigEndian, info)
	if e != nil {
		return nil, fmt.Errorf("Failed reading attribute info: %s", e)
	}
	toReturn.Info = info
	return &toReturn, nil
}

// Assumes the data input is at the start of an attribute table in the class
// file. Reads and parses the attributes in the table.
func (c *ClassFile) parseAttributesTable(data io.Reader,
	count uint16) ([]*Attribute, error) {
	var e error
	var attribute *Attribute
	attributes := make([]*Attribute, count)
	for i := 0; i < int(count); i++ {
		attribute, e = c.parseSingleAttribute(data)
		if e != nil {
			return nil, e
		}
		attributes[i] = attribute
	}
	return attributes, nil
}

// TODO: Add parsing for all remaining attribute types:
// - RuntimeVisibleTypeAnnotations
// - RuntimeInvisibleTypeAnnotations
// - AnnotationDefault
// - BootstrapMethods
// - MethodParameters
