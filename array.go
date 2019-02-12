package bs_jvm

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
