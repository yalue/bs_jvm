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
	showTrace := false
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("   %s [OPTIONS] <file to run>\n", os.Args[0])
		fmt.Printf("[OPTIONS] are one or more of:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&showTrace, "show_trace", false, "If true, prints a trace "+
		"of all executed instructions to stdout.")
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Printf("Usage: ./jvm [OPTIONS] <file to run>\n")
		log.Printf("Run with \"--help\" for more information.\n")
		return 1
	}
	filename := flag.Arg(0)
	j, e := NewJVMWithBuiltins()
	if e != nil {
		log.Printf("Failed initializing JVM: %s\n")
		return 1
	}
	if showTrace {
		j.TraceSink = os.Stdout
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
