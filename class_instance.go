package bs_jvm

// This file contains code specific to instances of a class.

// An instance of a class. One of the non-primitive reference types. Expected
// to be created using Class.CreateInstance.
type ClassInstance struct {
	// The class this is an instance of.
	C *Class
	// The non-static fields of this class. Get indices into this using
	// C.FieldInfo[fieldName].Index.
	FieldValues []Object
	// Used by builtin classes to refer to Go information. Otherwise, should be
	// nil.
	NativeData interface{}
	// TODO: Needs to also keep track of superclass fields. Probably need to
	// add a "superclass" ClassInstance reference.
}

func (o *ClassInstance) IsPrimitive() bool {
	return false
}

func (o *ClassInstance) TypeName() string {
	return string(o.C.Name)
}

func (o *ClassInstance) String() string {
	return "instance of " + string(o.C.Name)
}

// Like Class.ResolveStaticField, but used for non-static fields of a class.
// The named field must NOT be static in order for this to work. May return
// a ClassInstance for a superclass. The returned int is an index into the
// returned ClassInstance's FieldValues array. Returns an error if the field
// can't be resolved.
// NOTE: Make this work with static fields, too?
func (o *ClassInstance) ResolveField(name string) (*ClassInstance, int,
	error) {
	info := o.C.FieldInfo[name]
	// TODO: Actually look up fields in superclasses, etc.
	if info == nil {
		return nil, 0, FieldError("Could not find field " + name)
	}
	if info.FileField.Access.IsStatic() {
		return nil, 0, FieldError("Field " + name + " is static")
	}
	return o, info.Index, nil
}
