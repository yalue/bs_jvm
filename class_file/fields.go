package class_file

// This file contains information related to parsing field information
// structures in class files.

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Holds information about a field's permissions
type FieldAccessFlags uint16

func (f FieldAccessFlags) String() string {
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
	if (f & 0x0040) != 0 {
		toReturn += "volatile "
	}
	if (f & 0x0080) != 0 {
		toReturn += "transient "
	}
	if (f & 0x1000) != 0 {
		toReturn += "synthetic "
	}
	if (f & 0x4000) != 0 {
		toReturn += "enum "
	}
	return strings.TrimRight(toReturn, " ")
}

// Contains information about a single field in the class file.
type Field struct {
	// Access permissions and properties, e.g. "public static"
	Access FieldAccessFlags
	// The UTF-8 string containing the field's name
	Name []byte
	// Contains the type of the field.
	Descriptor FieldType
	// A table of attributes for this specific field
	Attributes []*Attribute
}

func (f *Field) String() string {
	return fmt.Sprintf("%s field, name %s, type %s, %d attributes",
		f.Access, f.Name, f.Descriptor, len(f.Attributes))
}

// Parses a single field structure.
func (c *Class) parseSingleField(data io.Reader) (*Field, error) {
	var toReturn Field
	e := binary.Read(data, binary.BigEndian, &(toReturn.Access))
	if e != nil {
		return nil, fmt.Errorf("Failed reading field access flags: %s", e)
	}
	var index uint16
	e = binary.Read(data, binary.BigEndian, &index)
	if e != nil {
		return nil, fmt.Errorf("Failed reading field name index: %s", e)
	}
	toReturn.Name, e = c.GetUTF8Constant(index)
	if e != nil {
		return nil, fmt.Errorf("Invalid field name: %s", e)
	}
	e = binary.Read(data, binary.BigEndian, &index)
	if e != nil {
		return nil, fmt.Errorf("Failed reading field descriptor index: %s", e)
	}
	descriptorBytes, e := c.GetUTF8Constant(index)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get field descriptor string: %s", e)
	}
	toReturn.Descriptor, e = ParseFieldType(descriptorBytes)
	if e != nil {
		return nil, fmt.Errorf("Error getting type for field %s: %s",
			toReturn.Name, e)
	}
	var attributeCount uint16
	e = binary.Read(data, binary.BigEndian, &attributeCount)
	if e != nil {
		return nil, fmt.Errorf("Failed reading field attribute count: %s", e)
	}
	attributes, e := c.parseAttributesTable(data, attributeCount)
	if e != nil {
		return nil, fmt.Errorf("Failed reading field attribute table: %s", e)
	}
	toReturn.Attributes = attributes
	return &toReturn, nil
}

// Assumes input is directly before a table of Field structures. Parses and
// returns the fields.
func (c *Class) parseFieldTable(data io.Reader, count uint16) ([]*Field,
	error) {
	var e error
	var field *Field
	fields := make([]*Field, count)
	for i := 0; i < int(count); i++ {
		field, e = c.parseSingleField(data)
		if e != nil {
			return nil, e
		}
		fields[i] = field
	}
	return fields, nil
}
