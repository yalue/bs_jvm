package class_file

// This file contains functions used for parsing the StackMapTable attribute.
// This is a particularly annoying attribute to parse, so it gets its own file.
// The most important function in this file is ParseStackMapTableAttribute.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// The tag indicating the type of a stack map frame
type StackMapFrameType uint8

func (t StackMapFrameType) String() string {
	switch {
	case t <= 63:
		return "same"
	case t <= 127:
		return "same locals, 1 stack item"
	case t == 247:
		return "same locals, 1 stack item extended"
	case (t >= 248) && (t <= 250):
		return "chop"
	case t == 251:
		return "same frame extended"
	case (t >= 252) && (t <= 254):
		return "append"
	case t == 255:
		return "full frame"
	}
	return fmt.Sprintf("Unknown stack map frame type %d", t)
}

// The tag indicating the type of a verification type info struct.
type VerificationTypeInfoTag uint8

func (t VerificationTypeInfoTag) String() string {
	switch t {
	case 0:
		return "top"
	case 1:
		return "integer"
	case 2:
		return "float"
	case 3:
		return "double"
	case 4:
		return "long"
	case 5:
		return "null"
	case 6:
		return "uninitialized this"
	case 7:
		return "object"
	case 8:
		return "uninitialized"
	}
	return fmt.Sprintf("Invalid verification type info tag: %d", t)
}

type VerificationTypeInfo struct {
	Tag VerificationTypeInfoTag
	// This is only set for one of the two tag types that use an extra 16-bit
	// field (in which case the OtherFieldValid() method will return true).
	Other uint16
}

// This will return true if v's Other field contains a necessary value for the
// given tag type.
func (v VerificationTypeInfo) OtherFieldValid() bool {
	switch v.Tag {
	case 7, 8:
		return true
	}
	return false
}

// Parses and returns a single verification type info structure.
func parseVerificationTypeInfo(data io.Reader) (VerificationTypeInfo, error) {
	var tag VerificationTypeInfoTag
	var toReturn VerificationTypeInfo
	e := binary.Read(data, binary.BigEndian, &tag)
	if e != nil {
		return toReturn, fmt.Errorf("Couldn't read type info tag: %s", e)
	}
	if tag > 8 {
		return toReturn, fmt.Errorf("Invalid verification type info tag: %d",
			tag)
	}
	if tag < 7 {
		toReturn.Tag = tag
		return toReturn, nil
	}
	var other uint16
	e = binary.Read(data, binary.BigEndian, &other)
	if e != nil {
		return toReturn, fmt.Errorf("Couldn't read verification type info: %s",
			e)
	}
	toReturn.Other = other
	return toReturn, nil
}

// A generic type for stack map frames. Type assersions may be used to convert
// this to more detailed types where necessary.
type StackMapFrame interface {
	FrameType() StackMapFrameType
	OffsetDelta() uint16
	String() string
}

// This type contains common data for all stack map frame types, and provides
// the FrameType method
type basicStackMapFrame struct {
	tag StackMapFrameType
	// This will not be set for types which infer the offset delta from the
	// StackMapFrameType tag.
	offsetDelta uint16
}

func (f *basicStackMapFrame) FrameType() StackMapFrameType {
	return f.tag
}

// This function will be overridden by stack map frame types for which the
// offset delta is inferred from the StackMapFrameType tag.
func (f *basicStackMapFrame) OffsetDelta() uint16 {
	return f.offsetDelta
}

func (f *basicStackMapFrame) String() string {
	offsetDelta := f.offsetDelta
	if f.tag < 64 {
		offsetDelta = uint16(f.tag)
	} else if f.tag < 128 {
		offsetDelta = uint16(f.tag) - 64
	}
	return fmt.Sprintf("%s, offset delta %d", f.tag, offsetDelta)
}

// Holds "same"-type stack map frames.
type SameStackMapFrame struct {
	basicStackMapFrame
}

func (f *SameStackMapFrame) OffsetDelta() uint16 {
	return uint16(f.FrameType())
}

// Holds "same locals, one stack item" stack map frames.
type OneItemStackMapFrame struct {
	basicStackMapFrame
	Info VerificationTypeInfo
}

func (f *OneItemStackMapFrame) OffsetDelta() uint16 {
	return uint16(f.FrameType() - 64)
}

// Holds the "same locals, one stack item, extended" stack map frame.
type OneItemStackMapFrameExtended struct {
	basicStackMapFrame
	Info VerificationTypeInfo
}

// Holds a "chop" stack map frame.
type ChopStackMapFrame struct {
	basicStackMapFrame
}

// Holds a "same frame extended" stack map frame.
type SameStackMapFrameExtended struct {
	basicStackMapFrame
}

// Holds an "append" stack map frame
type AppendStackMapFrame struct {
	basicStackMapFrame
	Locals []VerificationTypeInfo
}

// Holds a "full" stack map frame, supporting a larger number of local and
// stack variables.
type FullStackMapFrame struct {
	basicStackMapFrame
	Locals []VerificationTypeInfo
	Stack  []VerificationTypeInfo
}

// This method is used to save some typing and error message formatting.
// Expects to be at the beginning of a 16-bit offset delta field.
func readOffsetDelta(data io.Reader) (uint16, error) {
	var toReturn uint16
	e := binary.Read(data, binary.BigEndian, &toReturn)
	if e != nil {
		return 0, fmt.Errorf("Failed reading stack map frame offset delta: %s",
			e)
	}
	return toReturn, nil
}

func parseStackMapFrame(data io.Reader) (StackMapFrame, error) {
	var tag StackMapFrameType
	e := binary.Read(data, binary.BigEndian, &tag)
	if e != nil {
		return nil, fmt.Errorf("Failed reading stack map frame tag: %s", e)
	}
	if tag <= 63 {
		var toReturn SameStackMapFrame
		toReturn.tag = tag
		return &toReturn, nil
	}
	if tag <= 127 {
		info, e := parseVerificationTypeInfo(data)
		if e != nil {
			return nil, e
		}
		var toReturn OneItemStackMapFrame
		toReturn.tag = tag
		toReturn.Info = info
		return &toReturn, nil
	}
	if tag <= 246 {
		return nil, fmt.Errorf("Invalid stack map frame tag: %d", tag)
	}
	if tag == 247 {
		offset, e := readOffsetDelta(data)
		if e != nil {
			return nil, e
		}
		info, e := parseVerificationTypeInfo(data)
		if e != nil {
			return nil, e
		}
		var toReturn OneItemStackMapFrameExtended
		toReturn.tag = tag
		toReturn.offsetDelta = offset
		toReturn.Info = info
		return &toReturn, nil
	}
	if tag <= 250 {
		offset, e := readOffsetDelta(data)
		if e != nil {
			return nil, e
		}
		var toReturn ChopStackMapFrame
		toReturn.tag = tag
		toReturn.offsetDelta = offset
		return &toReturn, nil
	}
	if tag == 251 {
		offset, e := readOffsetDelta(data)
		if e != nil {
			return nil, e
		}
		var toReturn SameStackMapFrameExtended
		toReturn.tag = tag
		toReturn.offsetDelta = offset
		return &toReturn, nil
	}
	if tag <= 254 {
		offset, e := readOffsetDelta(data)
		if e != nil {
			return nil, e
		}
		var toReturn AppendStackMapFrame
		toReturn.tag = tag
		toReturn.offsetDelta = offset
		toReturn.Locals = make([]VerificationTypeInfo, tag-251)
		for i := range toReturn.Locals {
			toReturn.Locals[i], e = parseVerificationTypeInfo(data)
			if e != nil {
				return nil, e
			}
		}
		return &toReturn, nil
	}
	// Sanity check
	if tag != 255 {
		return nil, fmt.Errorf("Unhandled bad tag %d parsing stack map frame",
			tag)
	}
	var toReturn FullStackMapFrame
	offset, e := readOffsetDelta(data)
	if e != nil {
		return nil, e
	}
	toReturn.tag = tag
	toReturn.offsetDelta = offset
	var tmp uint16
	// Read the number of locals
	e = binary.Read(data, binary.BigEndian, &tmp)
	if e != nil {
		return nil, fmt.Errorf("Failed reading number of locals: %s", e)
	}
	toReturn.Locals = make([]VerificationTypeInfo, tmp)
	for i := range toReturn.Locals {
		toReturn.Locals[i], e = parseVerificationTypeInfo(data)
		if e != nil {
			return nil, e
		}
	}
	e = binary.Read(data, binary.BigEndian, &tmp)
	if e != nil {
		return nil, fmt.Errorf("Failed reading nunmber of stack vars: %s", e)
	}
	toReturn.Stack = make([]VerificationTypeInfo, tmp)
	for i := range toReturn.Stack {
		toReturn.Stack[i], e = parseVerificationTypeInfo(data)
		if e != nil {
			return nil, e
		}
	}
	return &toReturn, nil
}

// Parses a stack map table attribute, returning a slice, in order, of the
// stack map frames it contains. Returns an error if one occurs.
func ParseStackMapTableAttribute(a *Attribute) ([]StackMapFrame, error) {
	if string(a.Name) != "StackMapTable" {
		return nil, fmt.Errorf("Expected a stack map table attribute.")
	}
	var entryCount uint16
	data := bytes.NewReader(a.Info)
	e := binary.Read(data, binary.BigEndian, &entryCount)
	if e != nil {
		return nil, fmt.Errorf("Couldn't read stack map frame count: %s", e)
	}
	entries := make([]StackMapFrame, entryCount)
	for i := range entries {
		entries[i], e = parseStackMapFrame(data)
		if e != nil {
			return nil, e
		}
	}
	return entries, nil
}
