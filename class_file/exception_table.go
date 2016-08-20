package class_file

// This file contains code related to parsing exceptions/the exception table
// in code attributes.

import (
	"encoding/binary"
	"io"
)

// A single entry in a code attribute's exception table.
type ExceptionTableEntry struct {
	StartPC   uint16
	EndPC     uint16
	HandlerPC uint16
	CatchType uint16
}

func parseExceptionTable(data io.Reader, count uint16) ([]ExceptionTableEntry,
	error) {
	toReturn := make([]ExceptionTableEntry, count)
	e := binary.Read(data, binary.BigEndian, toReturn)
	if e != nil {
		return nil, e
	}
	return toReturn, nil
}
