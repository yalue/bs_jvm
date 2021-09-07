package bs_jvm

// This file contains definitions of internal constant object types, and code
// to convert the class_file constant format to the internal representation.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
)

// Used to refer to a string literal (object).
type StringObject string

func (s *StringObject) IsPrimitive() bool {
	return false
}

func (s *StringObject) TypeName() string {
	return "String"
}

func (s *StringObject) Value() string {
	return string(*s)
}

func (s *StringObject) String() string {
	return fmt.Sprintf("%q", string(*s))
}

// Holds a method type descriptor, which is a UTF-8 string. Implements the
// Object interface.
type MethodType string

func (t *MethodType) IsPrimitive() bool {
	return false
}

func (t *MethodType) TypeName() string {
	return "method type descriptor"
}

func (t *MethodType) String() string {
	return "type descriptor: " + string(*t)
}

// Holds the resolved results from a name and type constant.
type NameAndTypeInfo struct {
	// A UTF-8 string constaining the name of the method or field.
	Name []byte
	// A UTF-8 string containing the type descriptor of the method or field.
	Type []byte
}

func (n *NameAndTypeInfo) String() string {
	return "name: " + string(n.Name) + ", type: " + string(n.Type)
}

// Takes a name and type info constant and reads the name and type into a
// NameAndTypeInfo struct.
func ResolveNameAndTypeInfoConstant(class *Class,
	info *class_file.ConstantNameAndTypeInfo) (*NameAndTypeInfo, error) {
	name, e := class.File.GetUTF8Constant(info.NameIndex)
	if e != nil {
		return nil, fmt.Errorf("Failed getting field/method constant name: %s",
			e)
	}
	typeDescriptor, e := class.File.GetUTF8Constant(info.DescriptorIndex)
	if e != nil {
		return nil, fmt.Errorf("Failed getting field/method constant type: %s",
			e)
	}
	toReturn := NameAndTypeInfo{
		Name: name,
		Type: typeDescriptor,
	}
	return &toReturn, nil
}

// A handle to a field reference constant.
type FieldOrMethodReference struct {
	// The class the constant referred to.
	C *Class
	// The name and type of the field or method in the class.
	Field *NameAndTypeInfo
}

func (h *FieldOrMethodReference) String() string {
	return h.C.String() + ", " + h.Field.String()
}

func (h *FieldOrMethodReference) IsPrimitive() bool {
	return false
}

func (h *FieldOrMethodReference) TypeName() string {
	return "field or method reference constant"
}

type GetFieldMethodHandle struct {
	FieldOrMethodReference
}

func (h *GetFieldMethodHandle) TypeName() string {
	return "get field method handle"
}

type GetStaticMethodHandle struct {
	FieldOrMethodReference
}

func (h *GetStaticMethodHandle) TypeName() string {
	return "get static method handle"
}

type PutFieldMethodHandle struct {
	FieldOrMethodReference
}

func (h *PutFieldMethodHandle) TypeName() string {
	return "put field method handle"
}

type PutStaticMethodHandle struct {
	FieldOrMethodReference
}

func (h *PutStaticMethodHandle) TypeName() string {
	return "put static method handle"
}

type InvokeVirtualMethodHandle struct {
	FieldOrMethodReference
}

func (h *InvokeVirtualMethodHandle) TypeName() string {
	return "invoke virtual method handle"
}

type InvokeStaticMethodHandle struct {
	FieldOrMethodReference
}

func (h *InvokeStaticMethodHandle) TypeName() string {
	return "invoke static method handle"
}

type InvokeSpecialMethodHandle struct {
	FieldOrMethodReference
}

func (h *InvokeSpecialMethodHandle) TypeName() string {
	return "invoke special method handle"
}

type NewInvokeSpecialMethodHandle struct {
	FieldOrMethodReference
}

func (h *NewInvokeSpecialMethodHandle) TypeName() string {
	return "new invoke special method handle"
}

type InvokeInterfaceMethodHandle struct {
	FieldOrMethodReference
}

func (h *InvokeInterfaceMethodHandle) TypeName() string {
	return "invoke interface method handle"
}

// Takes a field, method reference, or interface method reference constant.
func convertFieldOrMethodRefConstantToObject(class *Class,
	info class_file.Constant) (*FieldOrMethodReference, error) {
	// First, verify that the constant is one of the types we expect, and grab
	// the appropriate fields in the meantime.
	var classIndex, nameAndTypeIndex uint16
	switch v := info.(type) {
	case *class_file.ConstantFieldInfo:
		classIndex = v.ClassIndex
		nameAndTypeIndex = v.NameAndTypeIndex
	case *class_file.ConstantMethodInfo:
		classIndex = v.ClassIndex
		nameAndTypeIndex = v.NameAndTypeIndex
	case *class_file.ConstantInterfaceMethodInfo:
		classIndex = v.ClassIndex
		nameAndTypeIndex = v.NameAndTypeIndex
	default:
		return nil, fmt.Errorf("Not a field or method ref constant: %s", info)
	}
	// First, get the referenced class.
	tmp, e := class.File.GetConstant(classIndex)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get class index for field info: %s",
			e)
	}
	classInfo, ok := tmp.(*class_file.ConstantClassInfo)
	if !ok {
		return nil, fmt.Errorf("Got bad class info constant for field info")
	}
	className, e := class.File.GetUTF8Constant(classInfo.NameIndex)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get class name for field info: %s", e)
	}
	// TODO: May need to "load" classes here if the referenced class isn't
	// loaded yet.
	fieldClass, e := class.ParentJVM.GetClass(string(className))
	if e != nil {
		return nil, e
	}
	// Now that we know the class that contains the field, get the name and
	// type of the field.
	tmp, e = class.File.GetConstant(nameAndTypeIndex)
	if e != nil {
		return nil, fmt.Errorf("Couldn't get field name and type constant: %s",
			e)
	}
	nameAndTypeInfo, ok := tmp.(*class_file.ConstantNameAndTypeInfo)
	if !ok {
		return nil, fmt.Errorf("Got bad name and type constant for field: %s",
			e)
	}
	parsedNameAndType, e := ResolveNameAndTypeInfoConstant(class,
		nameAndTypeInfo)
	if e != nil {
		return nil, fmt.Errorf("Error getting name and type for field: %s", e)
	}
	toReturn := FieldOrMethodReference{
		C:     fieldClass,
		Field: parsedNameAndType,
	}
	return &toReturn, nil
}

func convertMethodHandleInfoToObject(class *Class,
	info *class_file.ConstantMethodHandleInfo) (Object, error) {
	var toReturn Object
	var e error
	fieldOrMethodConstant, e := class.File.GetConstant(info.Index)
	if e != nil {
		return nil, fmt.Errorf("Failed looking up method handle's referred "+
			"constant: %s", e)
	}
	// NOTE: The spec actually has further restrictions on the constants these
	// can point to, but I am not checking them for now. I'll change that in
	// the future if it leads to problems.
	fieldOrMethod, e := convertFieldOrMethodRefConstantToObject(class,
		fieldOrMethodConstant)
	if e != nil {
		return nil, e
	}
	k := info.ReferenceKind
	switch k {
	case 1:
		toReturn = &GetFieldMethodHandle{
			*fieldOrMethod,
		}
	case 2:
		toReturn = &GetStaticMethodHandle{
			*fieldOrMethod,
		}
	case 3:
		toReturn = &PutFieldMethodHandle{
			*fieldOrMethod,
		}
	case 4:
		toReturn = &PutStaticMethodHandle{
			*fieldOrMethod,
		}
	case 5:
		toReturn = &InvokeVirtualMethodHandle{
			*fieldOrMethod,
		}
	case 6:
		toReturn = &InvokeStaticMethodHandle{
			*fieldOrMethod,
		}
	case 7:
		toReturn = &InvokeSpecialMethodHandle{
			*fieldOrMethod,
		}
	case 8:
		toReturn = &NewInvokeSpecialMethodHandle{
			*fieldOrMethod,
		}
	case 9:
		toReturn = &InvokeInterfaceMethodHandle{
			*fieldOrMethod,
		}
	default:
		return nil, fmt.Errorf("Invalid method handle reference kind: %s", k)
	}
	// Additional checks on method names for certain types.
	methodName := string(fieldOrMethod.Field.Name)
	switch k {
	case 5, 6, 7, 9:
		if (methodName == "<init>") || (methodName == "<cinit>") {
			return nil, fmt.Errorf("%s method handle can't use method %s", k,
				methodName)
		}
	case 8:
		if methodName != "<init>" {
			return nil, fmt.Errorf("%s method handle can't use method %s", k,
				methodName)
		}
	}
	return toReturn, nil
}

// Converts a class file constant to an object for use by the JVM.
func ConvertConstantToObject(class *Class,
	constant class_file.Constant) (Object, error) {
	switch v := constant.(type) {
	case *class_file.ConstantIntegerInfo:
		return Int(v.Value), nil
	case *class_file.ConstantFloatInfo:
		return Float(v.Value), nil
	case *class_file.ConstantLongInfo:
		return Long(v.Value), nil
	case *class_file.ConstantDoubleInfo:
		return Double(v.Value), nil
	case *class_file.ConstantStringInfo:
		stringValue, e := class.File.GetUTF8Constant(v.StringIndex)
		if e != nil {
			return nil, fmt.Errorf("Failed getting string constant: %s", e)
		}
		tmp := StringObject(stringValue)
		return &tmp, nil
	case *class_file.ConstantClassInfo:
		className, e := class.File.GetUTF8Constant(v.NameIndex)
		if e != nil {
			return nil, fmt.Errorf("Failed getting class name: %s", e)
		}
		return class.ParentJVM.GetClass(string(className))
	case *class_file.ConstantMethodTypeInfo:
		descriptor, e := class.File.GetUTF8Constant(v.DescriptorIndex)
		if e != nil {
			return nil, fmt.Errorf("Failed getting method type descriptor: %s",
				e)
		}
		tmp := MethodType(descriptor)
		return &tmp, nil
	case *class_file.ConstantMethodHandleInfo:
		return convertMethodHandleInfoToObject(class, v)
	case *class_file.ConstantFieldInfo:
		return convertFieldOrMethodRefConstantToObject(class, constant)
	}
	return nil, fmt.Errorf("Object conversion for constant %s not implemented",
		constant)
}
