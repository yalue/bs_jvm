package bs_jvm

// Not to be confused with the class_file package, this file contains code
// relating to the internal definition and states of classes.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
)

// Holds metadata used by the JVM when accessing fields of a class, static or
// otherwise.
type ClassField struct {
	// A reference to the field info parsed directly from the class file.
	FileField *class_file.Field
	// The index of the field in either the Class.StaticFieldValues array or in
	// the ClassInstance.FieldValues array. (Used to get the value the field
	// holds at runtime.) Figure out which array this applies to use by calling
	// FileField.Access.IsStatic().
	Index int
}

// Holds a loaded JVM class. Implements the Object interface, too, for
// references to class definitions.
type Class struct {
	ParentJVM *JVM
	Name      []byte
	Methods   map[string]*Method
	// Maps field names to metadata about the field.  For example, this can be
	// used to look up a field's index in the StaticFields array, by checking
	// the "Index" member of the FieldInfo struct.
	FieldInfo map[string]*ClassField
	// The values of the static fields of this class. Get the index into this
	// array by checkig FieldInfo[fieldName].
	StaticFieldValues []Object
	// Holds the types, as specified in the class file, of the static fields in
	// this class. The indices into this array are the same as the indices in
	// the StaticFieldValues array. Used for type checking and initializing
	// each field to a matching type.
	StaticFieldTypes []class_file.FieldType
	// Like StaticFieldTypes, but corresponding to entries in the FieldValues
	// array in *instances* of this class.
	FieldTypes []class_file.FieldType
	File       *class_file.Class
}

func (c *Class) String() string {
	return "class: " + string(c.Name)
}

func (c *Class) IsPrimitive() bool {
	return false
}

func (c *Class) TypeName() string {
	return "class"
}

// Resolves the named static field. Returns the class containing the field, and
// the index of the field in the class' StaticFieldValues array. (This needs to
// return a different class instance in case the field is in a superclass or
// interface.) Returns an error if the field can't be resolved.
func (c *Class) ResolveStaticField(name string) (*Class, int, error) {
	info := c.FieldInfo[name]
	// TODO: Actually look up fields in superclasses, etc.
	if info == nil {
		return nil, 0, FieldError("Could not find field " + name)
	}
	if !info.FileField.Access.IsStatic() {
		return nil, 0, FieldError("Field " + name + " is not static")
	}
	return c, info.Index, nil
}

// Returns the named method from the class. Returns a MethodNotFoundError if
// the method isn't found.
func (c *Class) GetMethod(name string) (*Method, error) {
	toReturn := c.Methods[name]
	if toReturn == nil {
		return nil, MethodNotFoundError(name)
	}
	return toReturn, nil
}

// Gets the default "zero" value object for the given field type. Returns an
// error if the FieldType is invalid.
func getDefaultFieldValue(t class_file.FieldType) (Object, error) {
	primitiveType, isPrimitive := t.(class_file.PrimitiveFieldType)
	if isPrimitive {
		switch primitiveType {
		case 'B':
			return Byte(0), nil
		case 'C':
			return Char(0), nil
		case 'D':
			return Double(0), nil
		case 'F':
			return Float(0), nil
		case 'I':
			return Int(0), nil
		case 'J':
			return Long(0), nil
		case 'S':
			return Short(0), nil
		case 'Z':
			// NOTE: We're using a zero byte for a false bool. Probably remove
			// the "bool" primitive type entirely?
			return Byte(0), nil
		}
		return nil, fmt.Errorf("Bad primitive type: %s", primitiveType)
	}
	// All other types are just going to be initialized with "null", but with
	// expected type information.
	return &NullObject{
		ExpectedType: t,
	}, nil
}

// Fills in default values in a list of field values, based on "zero" values of
// the types given by the FieldType array. Returns an error if one occurs.
func getDefaultFieldValues(fieldValues []Object,
	types []class_file.FieldType) error {
	for i := range types {
		v, e := getDefaultFieldValue(types[i])
		if e != nil {
			return e
		}
		fieldValues[i] = v
	}
	return nil
}

// Instantiates an object of this class. Doesn't do any initialization besides
// setting fields to zero or null.
func (c *Class) CreateInstance() (*ClassInstance, error) {
	fieldValues := make([]Object, len(c.FieldTypes))
	e := getDefaultFieldValues(fieldValues, c.FieldTypes)
	if e != nil {
		return nil, fmt.Errorf("Couldn't initialize object fields: %s", e)
	}
	return &ClassInstance{
		C:           c,
		FieldValues: fieldValues,
	}, nil
}

// Iterates over the class' field information, initializes the
// StaticFieldValues, FieldCount, and FieldInfo members of the Class struct.
func (c *Class) getFieldInfo() error {
	c.FieldInfo = make(map[string]*ClassField)
	staticCount := 0
	nonStaticCount := 0
	for _, f := range c.File.Fields {
		isStatic := f.Access.IsStatic()
		name := string(f.Name)
		if isStatic {
			c.FieldInfo[name] = &ClassField{
				FileField: f,
				Index:     staticCount,
			}
			staticCount++
		} else {
			c.FieldInfo[name] = &ClassField{
				FileField: f,
				Index:     nonStaticCount,
			}
			nonStaticCount++
		}
	}
	c.StaticFieldValues = make([]Object, staticCount)

	// Also allocate and populate the lists of field types.
	c.StaticFieldTypes = make([]class_file.FieldType, staticCount)
	c.FieldTypes = make([]class_file.FieldType, nonStaticCount)
	for _, f := range c.FieldInfo {
		if f.FileField.Access.IsStatic() {
			c.StaticFieldTypes[f.Index] = f.FileField.Descriptor
		} else {
			c.FieldTypes[f.Index] = f.FileField.Descriptor
		}
	}
	return nil
}

// Takes a class loaded by the class_file package and converts it to the Class
// type needed by the JVM. Does *not* modify the state of the JVM.
func NewClass(j *JVM, class *class_file.Class) (*Class, error) {
	className, e := class.GetName()
	if e != nil {
		return nil, fmt.Errorf("Error getting class name: %s", e)
	}
	toReturn := Class{
		ParentJVM:         j,
		Name:              className,
		Methods:           make(map[string]*Method),
		FieldInfo:         nil,
		StaticFieldValues: nil,
		FieldTypes:        nil,
		StaticFieldTypes:  nil,
		File:              class,
	}
	var methodName []byte
	var method *Method
	for i := range class.Methods {
		methodName = class.Methods[i].Name
		method, e = j.NewMethod(&toReturn, i)
		if e != nil {
			return nil, fmt.Errorf("Failed loading method %s: %s", methodName,
				e)
		}
		toReturn.Methods[string(methodName)] = method
	}
	e = (&toReturn).getFieldInfo()
	if e != nil {
		return nil, fmt.Errorf("Failed populating field info: %s", e)
	}
	e = getDefaultFieldValues(toReturn.StaticFieldValues,
		toReturn.StaticFieldTypes)
	if e != nil {
		return nil, fmt.Errorf("Failed setting default static fields: %s", e)
	}
	return &toReturn, nil
}
