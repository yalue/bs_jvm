package builtin_classes

// This file contains the code implementing the java/lang/System class. Call
// GetSystemClass to get an instance of it. (It's already included in
// GetBuiltinClasses).
import (
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/class_file"
	"os"
)

// Returns a BS-JVM class implementing java/lang/System.
func GetSystemClass(jvm *bs_jvm.JVM) (*bs_jvm.Class, error) {
	ps, e := GetPrintStreamClass(jvm)
	if e != nil {
		return nil, fmt.Errorf("Failed getting PrintStream class: %w", e)
	}
	// For an instance of the builtin PrintStream class, we don't need to use
	// any fields; we'll instead just set the output to os.Stdout.
	out := &bs_jvm.ClassInstance{
		C: ps,
		NativeData: &internalPrintStream{
			w:         os.Stdout,
			lastError: nil,
		},
	}
	toReturn := GetEmptyClass(jvm, "java/lang/System")
	publicStatic := class_file.FieldAccessFlags(1 | 8)
	AppendStaticField(toReturn, "out", publicStatic,
		class_file.ClassInstanceType("java/lang/PrintStream"), out)
	// TODO (next): Start populating java/lang/System. Start with printf,
	// probably (or whatever's needed for the RandomDots test).
	//  - Will need some way to represent a native method.
	return toReturn, nil
}
