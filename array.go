package bs_jvm

import (
	"fmt"
	"strconv"
)

// This file contains types relating to managing JVM arrays.

// Implements the Object interface for arrays of Ints.
type IntArray []Int

func (n IntArray) IsPrimitive() bool {
	return false
}

func (n IntArray) TypeName() string {
	return "int[]"
}

func (n IntArray) String() string {
	s := "["
	for i, v := range n {
		s += string(int32(v))
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// Implements the Object interface for arrays of Longs.
type LongArray []Long

func (n LongArray) IsPrimitive() bool {
	return false
}

func (n LongArray) TypeName() string {
	return "long[]"
}

func (n LongArray) String() string {
	s := "["
	for i, v := range n {
		s += string(int64(v))
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// This implements the Object interface for arrays of Floats.
type FloatArray []Float

func (n FloatArray) IsPrimitive() bool {
	return false
}

func (n FloatArray) TypeName() string {
	return "float[]"
}

func (n FloatArray) String() string {
	s := "["
	for i, v := range n {
		s += strconv.FormatFloat(float64(v), 'g', 5, 32)
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// This implements the Object interface for arrays of Doubles.
type DoubleArray []Double

func (n DoubleArray) IsPrimitive() bool {
	return false
}

func (n DoubleArray) TypeName() string {
	return "double[]"
}

func (n DoubleArray) String() string {
	s := "["
	for i, v := range n {
		s += strconv.FormatFloat(float64(v), 'g', 5, 64)
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// This implements the Object interface for arrays of references.
type ReferenceArray []Object

func (n ReferenceArray) IsPrimitive() bool {
	return false
}

func (n ReferenceArray) TypeName() string {
	return "Object[]"
}

func (n ReferenceArray) String() string {
	s := "["
	for i, v := range n {
		s += v.String()
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// This implements the Object interface for arrays of bytes.
type ByteArray []Byte

func (n ByteArray) IsPrimitive() bool {
	return false
}

func (n ByteArray) TypeName() string {
	return "byte[]"
}

func (n ByteArray) String() string {
	s := "["
	for i, v := range n {
		s += string(int8(v))
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}

// This implements the Object interface for arrays of chars.
type CharArray []Char

func (n CharArray) IsPrimitive() bool {
	return false
}

func (n CharArray) TypeName() string {
	return "char[]"
}

func (n CharArray) String() string {
	s := ""
	for _, v := range n {
		s += string(rune(v))
	}
	return fmt.Sprintf("%q", s)
}

// This implements the Object interface for arrays of shorts.
type ShortArray []Short

func (n ShortArray) IsPrimitive() bool {
	return false
}

func (n ShortArray) TypeName() string {
	return "short[]"
}

func (n ShortArray) String() string {
	s := "["
	for i, v := range n {
		s += string(int16(v))
		if i < (len(n) - 1) {
			s += ","
		}
	}
	s += "]"
	return s
}
