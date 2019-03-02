package bs_jvm

// This file contains definitions of the JVM's primitivae data types.

import (
	"strconv"
)

type Byte int8

func (b Byte) String() string {
	return "byte: " + strconv.Itoa(int(b))
}

func (b Byte) TypeName() string {
	return "byte"
}

func (b Byte) IsPrimitive() bool {
	return true
}

type Short int16

func (s Short) String() string {
	return "short: " + strconv.Itoa(int(s))
}

func (s Short) TypeName() string {
	return "short"
}

func (s Short) IsPrimitive() bool {
	return true
}

type Int int32

func (i Int) String() string {
	return "int: " + strconv.Itoa(int(i))
}

func (i Int) TypeName() string {
	return "int"
}

func (i Int) IsPrimitive() bool {
	return true
}

type Long int64

func (l Long) String() string {
	return "long: " + strconv.FormatInt(int64(l), 10)
}

func (l Long) TypeName() string {
	return "long"
}

func (l Long) IsPrimitive() bool {
	return true
}

type Char uint16

func (c Char) String() string {
	return "char: " + strconv.QuoteRuneToASCII(rune(c))
}

func (c Char) TypeName() string {
	return "char"
}

func (c Char) IsPrimitive() bool {
	return true
}

// TODO: Remove the bool type? Where is it used?
type Bool bool

func (b Bool) String() string {
	return "bool: " + strconv.FormatBool(bool(b))
}

func (b Bool) TypeName() string {
	return "bool"
}

func (b Bool) IsPrimitive() bool {
	return true
}

type Float float32

func (f Float) String() string {
	return "float: " + strconv.FormatFloat(float64(f), 'g', 5, 32)
}

func (f Float) TypeName() string {
	return "float"
}

func (f Float) IsPrimitive() bool {
	return true
}

type Double float64

func (d Double) String() string {
	return "double: " + strconv.FormatFloat(float64(d), 'g', 5, 64)
}

func (d Double) TypeName() string {
	return "double"
}

func (d Double) IsPrimitive() bool {
	return true
}
