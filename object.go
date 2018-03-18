package bs_jvm

// This file contains types relating to various JVM objects.

// A JVM object can be either a primitive or a reference type.
type Object interface {
	IsPrimitive() bool
	TypeName() string
	String() string
}

// A reference is generally an object, but should never be a primitive.
type Reference interface {
	Object
}
