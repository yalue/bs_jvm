package bs_jvm

import (
	"testing"
)

func TestLoadClass(t *testing.T) {
	jvm := NewJVM()
	if len(jvm.Classes) != 0 {
		t.Logf("A new JVM shouldn't start with classes loaded.\n")
		t.FailNow()
	}
	e := jvm.LoadClass(getTestClassFile(t))
	if e != nil {
		t.Logf("Error loading class: %s\n", e)
		t.FailNow()
	}
	for name, class := range jvm.Classes {
		t.Logf("Methods in loaded class: %s\n", name)
		for methodName := range class.Methods {
			t.Logf("  %s\n", methodName)
		}
	}
}
