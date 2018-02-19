// This package defines a JVM library for executing class and JAR files.
package jvm

// Holds state of the entire JVM, including threads, class files, etc.
type JVM struct {
}

// Holds a parsed JVM method.
type JVMMethod struct {
	// A reference to the parent JVM
	ParentJVM *JVM
	// Contains all parsed functions in the method.
	Instructions []JVMInstruction
}
