package bs_jvm

// This file contains types relating to various JVM objects.

import (
	"github.com/yalue/bs_jvm/class_file"
)

// A JVM object can be either a primitive or a reference type.
type Object interface {
	IsPrimitive() bool
	TypeName() string
	String() string
}

// A "null" object in java, used as a placeholder for uninitialized objects.
type NullObject struct {
	// The type destcriptor of the object, if it's known. Will be nil if it
	// isn't known for some reason.
	ExpectedType class_file.FieldType
}

func (o *NullObject) IsPrimitive() bool {
	return false
}

func (o *NullObject) TypeName() string {
	return "null"
}

func (o *NullObject) String() string {
	if o.ExpectedType == nil {
		return "null, unknown type"
	}
	return "null, instance of type " + o.ExpectedType.String()
}
