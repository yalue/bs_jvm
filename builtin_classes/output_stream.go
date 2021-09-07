package builtin_classes

// This file contains code implementing java.io.OutputStream-related classes.
import (
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/class_file"
	"io"
)

// An initialized version of our builtin PrintStream class.
var printStreamClass *bs_jvm.Class

// This holds internal data for the builtin PrintStream class.
type internalPrintStream struct {
	// The io.Writer to which we are writing data.
	w io.Writer
	// Holds an error if one occurred.
	lastError error
}

// Pops an instance of the PrintStream class from the thread's stack. Returns
// an error if it couldn't be popped, or wasn't an instance of the correct
// class. Also makes sure the private data is an io.Writer.
func popPrintStreamInstance(t *bs_jvm.Thread) (*bs_jvm.ClassInstance, error) {
	tmp, e := t.Stack.PopRef()
	if e != nil {
		return nil, fmt.Errorf("Failed popping PrintStream instance: %w", e)
	}
	instance, ok := tmp.(*bs_jvm.ClassInstance)
	if !ok {
		return nil, bs_jvm.TypeError("Didn't get class instance")
	}
	if instance.C != printStreamClass {
		return nil, bs_jvm.TypeError("Didn't get PrintStream instance")
	}
	_, ok = instance.NativeData.(*internalPrintStream)
	if !ok {
		return nil, fmt.Errorf("Internal error: didn't get expected " +
			"internalPrintStream data")
	}
	return instance, nil
}

// Implements the "print" method for a single char.
func printCharMethod(t *bs_jvm.Thread) error {
	instance, e := popPrintStreamInstance(t)
	if e != nil {
		return fmt.Errorf("print(char) failed: %w", e)
	}
	toPrint, e := t.Stack.Pop()
	if e != nil {
		return fmt.Errorf("print(char) failed popping char to print: %w", e)
	}
	p := instance.NativeData.(*internalPrintStream)
	_, p.lastError = fmt.Fprintf(p.w, "%c", toPrint)
	// The printStream class shouldn't raise any exceptions, but just set the
	// internal flag if necessary.
	return nil
}

// Implements the "println" method for a single string.
func printlnStringMethod(t *bs_jvm.Thread) error {
	instance, e := popPrintStreamInstance(t)
	if e != nil {
		return fmt.Errorf("println(String) failed: %w", e)
	}
	tmp, e := t.Stack.PopRef()
	if e != nil {
		return fmt.Errorf("println(String) failed popping String: %w", e)
	}
	toPrint, ok := tmp.(*bs_jvm.StringObject)
	if !ok {
		return bs_jvm.TypeError("Didn't get String instance")
	}
	p := instance.NativeData.(*internalPrintStream)
	_, p.lastError = fmt.Fprintf(p.w, "%s\n", toPrint.Value())
	return nil
}

// Adds the given function, f, to the class' Methods list. The method must
// be void, with a single arg with the given type.
func addSingleArgVoid(c *bs_jvm.Class, name string, arg class_file.FieldType,
	f bs_jvm.NativeImplementation) {
	descriptor := &class_file.MethodDescriptor{
		// void
		ReturnType:    class_file.PrimitiveFieldType('V'),
		ArgumentTypes: []class_file.FieldType{arg},
	}
	tmp := &class_file.Method{
		// public
		Access:     1,
		Name:       []byte(name),
		Descriptor: descriptor,
	}
	key := bs_jvm.GetMethodKey(tmp)

	// The remaining uninitialized fields in this struct aren't needed for
	// native implementations.
	method := &bs_jvm.Method{
		ContainingClass: printStreamClass,
		Types:           descriptor,
		OptimizeDone:    true,
		Native:          f,
	}
	c.Methods[key] = method
}

// Returns a BS-JVM class implementing java/io/PrintStream. If a class has
// already been initialized, returns the existing copy.
func GetPrintStreamClass(jvm *bs_jvm.JVM) (*bs_jvm.Class, error) {
	if printStreamClass != nil {
		return printStreamClass, nil
	}
	toReturn := GetEmptyClass(jvm, "java/io/PrintStream")
	// printStreamClass must be set before adding the methods.
	printStreamClass = toReturn

	// The print method with a single char arg.
	addSingleArgVoid(toReturn, "print", class_file.PrimitiveFieldType('C'),
		printCharMethod)
	// println with a single String arg
	addSingleArgVoid(toReturn, "println",
		class_file.ClassInstanceType("java/lang/String"), printlnStringMethod)
	// TODO: Continue implementing the PrintStream builtin class
	//  - checkError, clearError, print, printf, println, etc.
	return toReturn, nil
}
