// This is the main command-line executable used to launch a program using the
// BS-JVM.
package main

import (
	"flag"
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/builtin_classes"
	"log"
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
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Printf("Usage: ./jvm [OPTIONS] <file to run>\n")
		return 1
	}
	filename := flag.Arg(0)
	j, e := NewJVMWithBuiltins()
	if e != nil {
		log.Printf("Failed initializing JVM: %s\n")
		return 1
	}

	// Now actually run the loaded class.
	e = j.StartMainClass(filename)
	if e != nil {
		log.Printf("Error running main class: %s\n", e)
		return 1
	}
	e = j.WaitForAllThreads()
	if e != nil {
		log.Printf("JVM exited with an error: %s\n", e)
		return 1
	}
	return 0
}

func main() {
	log.SetFlags(0)
	os.Exit(run())
}
