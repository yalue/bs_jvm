// This is a simple command-line tool for viewing information contained in a
// single Java class file.
package main

import (
	"flag"
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/builtin_classes"
	"os"
)

func NewJVMWithBuiltins() (*bs_jvm.JVM, error) {
	j := bs_jvm.NewJVM()
	builtins, e := builtin_classes.GetBuiltinClasses(j)
	if e != nil {
		return nil, fmt.Errorf("Failed getting builtin classes: %w", e)
	}
	for _, class := range builtins {
		j.Classes[string(class.Name)] = class
	}
	return j, nil
}

func run() int {
	var filename string
	flag.StringVar(&filename, "filename", "",
		"The name of the class file to view.")
	flag.Parse()
	if filename == "" {
		fmt.Println("Invalid arguments. Run with -help for more information.")
		return 1
	}
	jvm, e := NewJVMWithBuiltins()
	if e != nil {
		fmt.Printf("Failed initializing JVM: %s\n", e)
		return 1
	}
	className, e := jvm.LoadClassFromFile(filename)
	if e != nil {
		fmt.Printf("Failed loading class: %s\n", e)
		return 1
	}
	var offset uint
	fmt.Printf("Methods in class %s:\n", className)
	class := jvm.Classes[className]
	for key, method := range class.Methods {
		e = method.Optimize()
		if e != nil {
			fmt.Printf("Unable to resolve instructions in method %s: %s\n",
				key, e)
		}
		offset = 0
		fmt.Printf("  Method %s:\n", key)
		for _, n := range method.Instructions {
			fmt.Printf("    0x%08x %s\n", offset, n)
			offset += n.Length()
		}
	}
	return 0
}

func main() {
	// Idiom to allow defer statements in the main routine
	os.Exit(run())
}
