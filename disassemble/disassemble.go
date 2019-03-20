// This is a simple command-line tool for viewing information contained in a
// single Java class file.
package main

import (
	"flag"
	"fmt"
	"github.com/yalue/bs_jvm"
	"os"
)

func run() int {
	var filename string
	flag.StringVar(&filename, "filename", "",
		"The name of the class file to view.")
	flag.Parse()
	if filename == "" {
		fmt.Println("Invalid arguments. Run with -help for more information.")
		return 1
	}
	jvm := bs_jvm.NewJVM()
	className, e := jvm.LoadClassFromFile(filename)
	if e != nil {
		fmt.Printf("Failed loading class: %s\n", e)
		return 1
	}
	var offset uint
	fmt.Printf("Methods in class %s:\n", className)
	class := jvm.Classes[className]
	for name, method := range class.Methods {
		offset = 0
		fmt.Printf("  Method %s %s(%s):\n", method.Types.ReturnString(), name,
			method.Types.ArgumentsString())
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
