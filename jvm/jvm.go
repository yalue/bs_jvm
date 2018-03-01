// This is the main command-line executable used to launch a program using the
// BS-JVM.
package main

import (
	"flag"
	"github.com/yalue/bs_jvm"
	"log"
	"os"
)

func run() int {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Printf("Usage: ./jvm [OPTIONS] <file to run>\n")
		return 1
	}
	filename := flag.Arg(0)
	j := bs_jvm.NewJVM()
	e := j.StartMainClass(filename)
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
