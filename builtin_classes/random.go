package builtin_classes

// This file contains code implementing java.util.Random
import (
	"fmt"
	"github.com/yalue/bs_jvm"
	"github.com/yalue/bs_jvm/class_file"
	"math/rand"
	"sync"
)

// An initialized version of the builtin Random class.
var randomClass *bs_jvm.Class

// Holds internal state for the Random class.
type internalRandom struct {
	// The underlying go RNG we'll use to provide random numbers.
	source *rand.Rand
	// Java's Random must be thread-safe, so we'll use this lock when accessing
	// the source.
	mutex sync.Mutex
}

// Pops an instance of the builtin Random class. Returns an error if the
// value couldn't be popped or wasn't an instance of the correct class.
func popRandomInstance(t *bs_jvm.Thread) (*bs_jvm.ClassInstance, error) {
	tmp, e := t.Stack.PopRef()
	if e != nil {
		return nil, fmt.Errorf("Failed popping Random instance: %w", e)
	}
	instance, ok := tmp.(*bs_jvm.ClassInstance)
	if !ok {
		return nil, bs_jvm.TypeError("Didn't get class instance")
	}
	if instance.C != randomClass {
		return nil, bs_jvm.TypeError("Didn't get Random instance")
	}
	_, ok = instance.NativeData.(*internalRandom)
	if !ok {
		return nil, fmt.Errorf("Internal error: didn't get expected " +
			"internalRandom data")
	}
	return instance, nil
}

// Implements the nextInt(int) method
func nextIntWithBoundMethod(t *bs_jvm.Thread) error {
	instance, e := popRandomInstance(t)
	if e != nil {
		return e
	}
	bound, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	// Should throw an illegal argument exception.
	if bound <= 0 {
		return bs_jvm.IllegalArgumentError("nextInt(int) requires a positive " +
			"argument")
	}
	r := instance.NativeData.(*internalRandom)
	r.mutex.Lock()
	toReturn := bs_jvm.Int(r.source.Int31n(int32(bound)))
	r.mutex.Unlock()
	return t.Stack.Push(toReturn)
}

// Implements the nextInt() method
func nextIntMethod(t *bs_jvm.Thread) error {
	instance, e := popRandomInstance(t)
	if e != nil {
		return e
	}
	r := instance.NativeData.(*internalRandom)
	r.mutex.Lock()
	// Java's nextInt can return negative or positive values, so we'll take
	// Go's uint64 rahter than its 32-bit versions, which only return positive.
	toReturn := bs_jvm.Int(r.source.Uint64())
	r.mutex.Unlock()
	return t.Stack.Push(toReturn)
}

// Returns a BS-JVM class implementing java/util/Random. If a class has already
// been initialized, returns the existing copy.
func GetRandomClass(jvm *bs_jvm.JVM) (*bs_jvm.Class, error) {
	if randomClass != nil {
		return randomClass, nil
	}
	toReturn := GetEmptyClass(jvm, "java/util/Random")
	AddMethod(toReturn, "nextInt", 1,
		[]class_file.FieldType{class_file.PrimitiveFieldType('I')},
		class_file.PrimitiveFieldType('I'), nextIntWithBoundMethod)
	AddMethod(toReturn, "nextInt", 1, []class_file.FieldType{},
		class_file.PrimitiveFieldType('I'), nextIntMethod)
	// TODO: Continue java/util/Random
	//  - constructor
	randomClass = toReturn
	return toReturn, nil
}
