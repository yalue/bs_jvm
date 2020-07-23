package bs_jvm

// This file contains definitions of the JVM's primitivae data types.

import (
	"strconv"
)

// A special interface implemented only by primitive types, to allow converting
// between types of primitives without needing to check the type of both the
// "source" primitive and the "destination" it is overwriting.
type PrimitiveType interface {
	// PrimitiveType instances must also satisfy the Object interface.
	Object
	// Converts the primitive to an int. Bools will be 1 = true, 0 = false.
	IntValue() int64
	// Converts the primitive's value to a float. Bools will be equal to 1.
	FloatValue() float64
	// Converts the given PrimitiveType value to the same type as the receiver.
	ConvertFrom(v PrimitiveType) PrimitiveType
}

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

func (b Byte) IntValue() int64 {
	return int64(b)
}

func (b Byte) FloatValue() float64 {
	return float64(b)
}

func (b Byte) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Byte(v.IntValue())
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

func (s Short) IntValue() int64 {
	return int64(s)
}

func (s Short) FloatValue() float64 {
	return float64(s)
}

func (s Short) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Short(v.IntValue())
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

func (i Int) IntValue() int64 {
	return int64(i)
}

func (i Int) FloatValue() float64 {
	return float64(i)
}

func (i Int) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Int(v.IntValue())
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

func (l Long) IntValue() int64 {
	return int64(l)
}

func (l Long) FloatValue() float64 {
	return float64(l)
}

func (l Long) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Long(v.IntValue())
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

func (c Char) IntValue() int64 {
	return int64(c)
}

func (c Char) FloatValue() float64 {
	return float64(c)
}

func (c Char) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Char(v.IntValue())
}

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

func (b Bool) IntValue() int64 {
	if b {
		return int64(1)
	}
	return int64(0)
}

func (b Bool) FloatValue() float64 {
	return float64(b.IntValue())
}

func (b Bool) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Bool((v.IntValue() & 1) != 0)
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

func (f Float) IntValue() int64 {
	return int64(f)
}

func (f Float) FloatValue() float64 {
	return float64(f)
}

func (f Float) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Float(v.FloatValue())
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

func (d Double) IntValue() int64 {
	return int64(d)
}

func (d Double) FloatValue() float64 {
	return float64(d)
}

func (d Double) ConvertFrom(v PrimitiveType) PrimitiveType {
	return Double(v.FloatValue())
}
