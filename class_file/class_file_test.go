package class_file

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func getTestClassFile(t *testing.T) io.Reader {
	content, e := ioutil.ReadFile("test_data/RandomDots.class")
	if e != nil {
		t.Logf("Failed reading the test class file: %s\n", e)
		t.FailNow()
	}
	return bytes.NewReader(content)
}

func getParsedClassFile(t *testing.T) *ClassFile {
	content := getTestClassFile(t)
	toReturn, e := ParseClassFile(content)
	if e != nil {
		t.Logf("Failed parsing the class file: %s\n", e)
		t.FailNow()
	}
	return toReturn
}

func TestParseClassFile(t *testing.T) {
	class := getParsedClassFile(t)
	if len(class.Fields) != 3 {
		t.Logf("Expected 3 fields, got %d\n", len(class.Fields))
		t.Fail()
	}
	if len(class.Methods) != 3 {
		t.Logf("Expected 3 methods, got %d\n", len(class.Methods))
		t.Fail()
	}
	for i, c := range class.Constants {
		t.Logf("Constant %d: %s\n", i, c)
	}
	for i, f := range class.Fields {
		t.Logf("Field %d: %s\n", i, f)
	}
	for i, m := range class.Methods {
		t.Logf("Method %d: %s\n", i, m)
	}
}

func TestParseCodeAttributes(t *testing.T) {
	class := getParsedClassFile(t)
	var codeAttribute *Attribute
	var parsedCodeAttribute *CodeAttribute
	var e error
	for _, m := range class.Methods {
		codeAttribute = nil
		for _, a := range m.Attributes {
			if string(a.Name) != "Code" {
				continue
			}
			codeAttribute = a
			break
		}
		if codeAttribute == nil {
			t.Logf("Couldn't find code attribute for method %s\n", m)
			t.FailNow()
		}
		parsedCodeAttribute, e = ParseCodeAttribute(codeAttribute, class)
		if e != nil {
			t.Logf("Failed parsing code attribute: %s\n", e)
			t.FailNow()
		}
		t.Logf("Attributes for %s's code:\n", m)
		for _, a := range parsedCodeAttribute.Attributes {
			t.Logf("  %s\n", a.Name)
		}
	}
}

func TestParseStackMapFrameAttributes(t *testing.T) {
	class := getParsedClassFile(t)
	var code *CodeAttribute
	var e error
	var frames []StackMapFrame
	for _, m := range class.Methods {
		if string(m.Name) != "main" {
			continue
		}
		for _, a := range m.Attributes {
			if string(a.Name) != "Code" {
				continue
			}
			code, e = ParseCodeAttribute(a, class)
			if e != nil {
				t.Logf("Failed parsing main's code attribute: %s\n", e)
				t.FailNow()
			}
		}
		if code != nil {
			break
		}
	}
	if code == nil {
		t.Logf("Couldn't find code attribute for main method.\n")
		t.FailNow()
	}
	for _, a := range code.Attributes {
		if string(a.Name) != "StackMapTable" {
			continue
		}
		frames, e = ParseStackMapTableAttribute(a)
		if e != nil {
			t.Logf("Failed parsing main's stack map table: %s", e)
			t.FailNow()
		}
	}
	t.Logf("Stack map frame table for the main method:\n")
	for _, f := range frames {
		t.Logf("  %s\n", f)
	}
}

// TODO: Add a test for annotations.
