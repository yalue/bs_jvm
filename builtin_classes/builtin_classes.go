// The builtin_classes package contains the classes implemented in go, used by
// the BS JVM.  To use it, call GetBuiltinClasses to get a list of classes,
// which may be registered with a JVM at initialization time.
package builtin_classes

import (
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/class_file"
)

// Returns a new *bs_jvm.Class instance with the given name, but all fields
// initialized but empty (maps and slices will be allocated, but not filled).
// The "File" field will be nil. Intended to be used as a helper function
// within the builtin_classes package.
func GetEmptyClass(jvm *bs_jvm.JVM, className string) *bs_jvm.Class {
	toReturn := &bs_jvm.Class{
		ParentJVM:         jvm,
		Name:              []byte(className),
		Methods:           make(map[string]*bs_jvm.Method),
		FieldInfo:         make(map[string]*bs_jvm.ClassField),
		StaticFieldValues: make([]bs_jvm.Object, 0, 10),
		StaticFieldTypes:  make([]class_file.FieldType, 0, 10),
		FieldTypes:        make([]class_file.FieldType, 0, 10),
		StaticFieldNames:  make([]string, 0, 10),
		FieldNames:        make([]string, 0, 10),
		File:              nil,
	}
	return toReturn
}

// Adds a static field to the class' list of static fields. Used when setting
// up internal classes.
func AppendStaticField(c *bs_jvm.Class, name string,
	a class_file.FieldAccessFlags, t class_file.FieldType, v bs_jvm.Object) {
	fileField := &class_file.Field{
		Access:     a,
		Name:       []byte(name),
		Descriptor: t,
		Attributes: nil,
	}
	index := len(c.StaticFieldNames)
	classField := &bs_jvm.ClassField{
		FileField: fileField,
		Index:     index,
	}
	c.FieldInfo[name] = classField
	c.StaticFieldNames = append(c.StaticFieldNames, name)
	c.StaticFieldTypes = append(c.StaticFieldTypes, t)
	c.StaticFieldValues = append(c.StaticFieldValues, v)
}

// Adds a native method to the given class. The args, return type, name, access
// flags, and implementation must all be specified.
func AddMethod(c *bs_jvm.Class, name string,
	access class_file.MethodAccessFlags, args []class_file.FieldType,
	returns class_file.FieldType, f bs_jvm.NativeMethod) {
	descriptor := &class_file.MethodDescriptor{
		ReturnType:    returns,
		ArgumentTypes: args,
	}
	tmp := &class_file.Method{
		Access:     access,
		Name:       []byte(name),
		Descriptor: descriptor,
	}
	key := bs_jvm.GetMethodKey(tmp)
	// The remaining uninitialized fields in this struct aren't needed for
	// native implementations.
	method := &bs_jvm.Method{
		ContainingClass: c,
		Types:           descriptor,
		OptimizeDone:    true,
		Native:          f,
	}
	c.Methods[key] = method
}

// Wraps AddMethod, simplifying usage for a public non-static method with a
// single arg and returning void.
func AddSingleArgVoidMethod(c *bs_jvm.Class, name string,
	arg class_file.FieldType, f bs_jvm.NativeMethod) {
	AddMethod(c, name, 1, []class_file.FieldType{arg},
		class_file.PrimitiveFieldType('V'), f)
}

// Returns a list of builtin Class objects, that may be registered with a given
// JVM. Each class' Name field will be set to the fully-qualified name of the
// class that it implements, but class-file-specific information may be unset,
// or nil. (Cases where this may happen will be documented, but be careful.)
// For the most part, all of the important fields in the Class struct *will* be
// set, though. Will return an error if one occurs while initializing a class.
// Requires a reference to the parent JVM, but will not modify its state.
func GetBuiltinClasses(jvm *bs_jvm.JVM) ([]*bs_jvm.Class, error) {
	toReturn := make([]*bs_jvm.Class, 0, 10)
	// Create new builtin classes and add them here as needed.
	tmp, e := GetSystemClass(jvm)
	if e != nil {
		return nil, fmt.Errorf("Failed initializing System class: %w", e)
	}
	toReturn = append(toReturn, tmp)
	tmp, e = GetRandomClass(jvm)
	if e != nil {
		return nil, fmt.Errorf("Failed initializing Random class: %w", e)
	}
	toReturn = append(toReturn, tmp)
	return toReturn, nil
}
