package bs_jvm

import (
	"github.com/yalue/bs_jvm/class_file"
	"os"
	"testing"
)

func getTestClassFile(t *testing.T) *class_file.Class {
	file, e := os.Open("class_file/test_data/RandomDots.class")
	if e != nil {
		t.Logf("Failed opening the test class file: %s\n", e)
		t.FailNow()
	}
	defer file.Close()
	toReturn, e := class_file.ParseClass(file)
	if e != nil {
		t.Logf("Failed parsing the test class file: %s\n", e)
		t.FailNow()
	}
	return toReturn
}

func getRandomDotMethodCode(t *testing.T, class *class_file.Class) []byte {
	var method *class_file.Method
	for i := range class.Methods {
		if string(class.Methods[i].Name) == "getDot" {
			method = class.Methods[i]
			break
		}
	}
	if method == nil {
		t.Logf("Failed finding getDot() method in the test class file.\n")
		t.FailNow()
		return nil
	}
	codeAttribute, e := method.GetCodeAttribute(class)
	if e != nil {
		t.Logf("Failed getting method code attribute: %s\n", e)
		t.FailNow()
		return nil
	}
	return codeAttribute.Code
}

func TestGetNextInstruction(t *testing.T) {
	var e error
	class := getTestClassFile(t)
	codeBytes := getRandomDotMethodCode(t, class)
	codeMemory := MemoryFromSlice(codeBytes)
	var instruction Instruction
	address := uint(0)
	t.Logf("getDot disassembly:\n")
	for address < uint(len(codeBytes)) {
		instruction, e = GetNextInstruction(codeMemory, address)
		if e != nil {
			t.Logf("Error getting next instruction: %s\n", e)
			t.FailNow()
		}
		t.Logf("0x%x: %s\n", address, instruction)
		address += instruction.Length()
	}
	_, e = GetNextInstruction(codeMemory, address)
	if e == nil {
		t.Logf("Didn't get an error reading an inst. from a bad address.\n")
		t.FailNow()
	}
	t.Logf("Got expected error from GetNextInstruction: %s\n", e)
}
