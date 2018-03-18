package bs_jvm

// Not to be confused with the class_file package, this file contains code
// relating to the internal definition and states of classes.

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
)

// Holds a loaded JVM class. Implements the Reference interface, too, for
// references to class definitions.
type Class struct {
	ParentJVM *JVM
	Name      []byte
	Methods   map[string]*Method
	File      *class_file.Class
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

// Returns the named method from the class. Returns a MethodNotFoundError if
// the method isn't found.
func (c *Class) GetMethod(name string) (*Method, error) {
	toReturn := c.Methods[name]
	if toReturn == nil {
		return nil, MethodNotFoundError(name)
	}
	return toReturn, nil
}

// Takes a class loaded by the class_file package and converts it to the Class
// type needed by the JVM. Does *not* modify the state of the JVM.
func NewClass(j *JVM, class *class_file.Class) (*Class, error) {
	className, e := class.GetName()
	if e != nil {
		return nil, fmt.Errorf("Error getting class name: %s", e)
	}
	toReturn := Class{
		ParentJVM: j,
		Name:      className,
		Methods:   make(map[string]*Method),
		File:      class,
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
	return &toReturn, nil
}
