// This is a simple command-line tool for viewing information contained in a
// single Java class file.
package main

import (
	"flag"
	"fmt"
	"github.com/yalue/jvm"
	"github.com/yalue/jvm/class_file"
	"os"
)

// Define a type to trivially implement the jvm.JVMMemory interface
type methodCode []byte

func (m methodCode) GetByte(address uint) (uint8, error) {
	if address >= uint(len(m)) {
		return 0, fmt.Errorf("Can't read code offset %d in a %d-byte method",
			address, len(m))
	}
	return m[address], nil
}

func (m methodCode) SetByte(value uint8, address uint) error {
	return fmt.Errorf("Can't write to method code")
}

// Prints the disassembly for a method's code.
func printDisassembly(codeBytes []byte) error {
	var e error
	var instruction jvm.Instruction
	code := methodCode(codeBytes)
	address := uint(0)
	for address < uint(len(code)) {
		instruction, e = jvm.GetNextInstruction(code, address)
		if e != nil {
			return fmt.Errorf("Failed reading instruction at offset %d: %s",
				address, e)
		}
		fmt.Printf("  0x%08x: %s\n", address, instruction)
		address += instruction.Length()
	}
	return nil
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
	file, e := os.Open(filename)
	if e != nil {
		fmt.Printf("Error opening class file: %s\n", e)
		return 1
	}
	defer file.Close()
	class, e := class_file.ParseClassFile(file)
	if e != nil {
		fmt.Printf("Error parsing class file: %s\n", e)
		return 1
	}
	fmt.Printf("Methods in %s:\n", filename)
	var codeAttribute *class_file.CodeAttribute
	// Display disassembly for each method.
	for i, method := range class.Methods {
		fmt.Printf("Method %d: %s\n", i, method)
		codeAttribute, e = method.GetCodeAttribute(class)
		if e != nil {
			fmt.Printf("  Couldn't get code attribute: %s\n", e)
			continue
		}
		e = printDisassembly(codeAttribute.Code)
		if e != nil {
			fmt.Printf("  Failed disassembling code: %s\n", e)
			continue
		}
	}
	return 0
}

func main() {
	// Idiom to allow defer statements in the main routine
	os.Exit(run())
}
