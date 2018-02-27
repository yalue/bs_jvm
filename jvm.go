// This package defines a JVM library for executing class and JAR files.
package jvm

import (
	"fmt"
	"github.com/yalue/jvm/class_file"
)

type ThreadStack interface {
}

// Holds the state of a single JVM thread.
type Thread interface {
	// The method that the thread is currently executing.
	CurrentMethod() *Method
	// The containing JVM.
	ParentJVM() *JVM
	// The program counter into the current method.
	PC() uint
}

// Holds state of the entire JVM, including threads, class files, etc.
type JVM struct {
	// A list of threads in the JVM.
	Threads []Thread
	// Maps class names to all loaded classes.
	Classes map[string]*Class
}

// Holds a parsed JVM method.
type Method struct {
	// A reference to the parent JVM.
	ParentJVM *JVM
	// The class in which the method was defined.
	ContainingClass *class_file.Class
	// Contains all parsed functions in the method.
	Instructions []Instruction
	// Contains a mapping between an index into the instruction array and the
	// offset in bytes of a given instruction.
	InstructionOffsets []uint
}

// Parses the given method from the class file into the structure needed by the
// JVM for actual execution. Also carries out pre-optimization.
func GetMethodFromClassFile(class *class_file.Class, methodIndex int) (*Method,
	error) {
	// TODO (next): Implement GetMethodFromClassFile
	return nil, NotImplementedError
}

// Holds a loaded JVM class.
type Class struct {
	Methods map[string]*Method
	File    *class_file.Class
}

// Takes a class loaded by the class_file package and converts it to the Class
// type needed by the JVM.
func NewClassFromClassFile(class *class_file.Class) (*Class, error) {
	toReturn := Class{
		Methods: make(map[string]*Method),
		File:    class,
	}
	var methodName []byte
	var method *Method
	var e error
	for i := range class.Methods {
		methodName = class.Methods[i].Name
		method, e = GetMethodFromClassFile(class, i)
		if e != nil {
			return nil, fmt.Errorf("Failed loading method %s: %s", methodName,
				e)
		}
		toReturn.Methods[string(methodName)] = method
	}
	return &toReturn, nil
}

// Adds the given class file to the JVM so that its code
func (j *JVM) LoadClass(class *class_file.Class) error {
	name, e := class.GetName()
	if e != nil {
		return fmt.Errorf("Failed getting class name: %s", e)
	}
	loadedClass, e := NewClassFromClassFile(class)
	if e != nil {
		return fmt.Errorf("Error loading class %s: %s", name, e)
	}
	j.Classes[string(name)] = loadedClass
	return nil
}
