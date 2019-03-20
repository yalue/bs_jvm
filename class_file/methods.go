package class_file

// This file contains information related to parsing field information
// structures in class files. Almost everything here is identical to the
// parsing of field structures.

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Holds information about a method's permissions
type MethodAccessFlags uint16

func (f MethodAccessFlags) String() string {
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
		toReturn += "synchronized "
	}
	if (f & 0x0040) != 0 {
		toReturn += "bridge "
	}
	if (f & 0x0080) != 0 {
		toReturn += "varargs "
	}
	if (f & 0x0100) != 0 {
		toReturn += "native "
	}
	if (f & 0x0400) != 0 {
		toReturn += "abstract "
	}
	if (f & 0x0800) != 0 {
		toReturn += "strict "
	}
	if (f & 0x1000) != 0 {
		toReturn += "synthetic "
	}
	return strings.TrimRight(toReturn, " ")
}

// Contains information about a single method in the class file.
type Method struct {
	// Access permissions and properties, e.g. "public static"
	Access MethodAccessFlags
	// The UTF-8 string containing the method's name
	Name []byte
	// Contains information about the method's arguments and return type.
	Descriptor *MethodDescriptor
	// A table of attributes for this specific method
	Attributes []*Attribute
}

func (m *Method) String() string {
	return fmt.Sprintf("%s %s %s(%s), %d attribute(s)", m.Access,
		m.Descriptor.ReturnString(), m.Name, m.Descriptor.ArgumentsString(),
		len(m.Attributes))
}

// Returns this method's code attribute, which must exist by the JVM spec.
// Returns an error if one occurs. This will scan the method's attributes table
// and parse the Code attribute, so this shouldn't be called frequently to
// preserve performance.
func (m *Method) GetCodeAttribute(class *Class) (*CodeAttribute, error) {
	found := false
	var e error
	var codeAttribute *CodeAttribute
	for _, attribute := range m.Attributes {
		if string(attribute.Name) != "Code" {
			continue
		}
		codeAttribute, e = ParseCodeAttribute(attribute, class)
		if e != nil {
			return nil, fmt.Errorf("Invalid code attribute: %s", e)
		}
		found = true
		break
	}
	if !found {
		return nil, fmt.Errorf("The method was missing a Code attribute")
	}
	return codeAttribute, nil
}

// Parses a single method structure.
func (c *Class) parseSingleMethod(data io.Reader) (*Method, error) {
	var toReturn Method
	e := binary.Read(data, binary.BigEndian, &(toReturn.Access))
	if e != nil {
		return nil, fmt.Errorf("Failed reading method access flags: %s", e)
	}
	var index uint16
	e = binary.Read(data, binary.BigEndian, &index)
	if e != nil {
		return nil, fmt.Errorf("Failed reading method name index: %s", e)
	}
	toReturn.Name, e = c.GetUTF8Constant(index)
	if e != nil {
		return nil, fmt.Errorf("Invalid method name: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &index)
	if e != nil {
		return nil, fmt.Errorf("Failed reading method descriptor index: %s", e)
	}
	descriptorBytes, e := c.GetUTF8Constant(index)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get method descriptor string: %s", e)
	}
	toReturn.Descriptor, e = ParseMethodDescriptor(descriptorBytes)
	if e != nil {
		return nil, fmt.Errorf("Bad descriptor for method %s: %s",
			toReturn.Name, e)
	}
	var attributeCount uint16
	e = binary.Read(data, binary.BigEndian, &attributeCount)
	if e != nil {
		return nil, fmt.Errorf("Failed reading method attribute count: %s", e)
	}
	attributes, e := c.parseAttributesTable(data, attributeCount)
	if e != nil {
		return nil, fmt.Errorf("Failed reading method attribute table: %s", e)
	}
	toReturn.Attributes = attributes
	return &toReturn, nil
}

// Assumes input is directly before a table of Method structures. Parses and
// returns the methods.
func (c *Class) parseMethodTable(data io.Reader, count uint16) ([]*Method,
	error) {
	var e error
	var method *Method
	methods := make([]*Method, count)
	for i := 0; i < int(count); i++ {
		method, e = c.parseSingleMethod(data)
		if e != nil {
			return nil, e
		}
		methods[i] = method
	}
	return methods, nil
}
