package bs_jvm

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
	"math"
	"sort"
)

// This file contains functions for executing individual JVM instructions.

func (n *nopInstruction) Execute(t *Thread) error {
	return nil
}

func (n *aconst_nullInstruction) Execute(t *Thread) error {
	return t.Stack.PushRef(nil)
}

func (n *iconst_m1Instruction) Execute(t *Thread) error {
	return t.Stack.Push(-1)
}

func (n *iconst_0Instruction) Execute(t *Thread) error {
	return t.Stack.Push(0)
}

func (n *iconst_1Instruction) Execute(t *Thread) error {
	return t.Stack.Push(1)
}

func (n *iconst_2Instruction) Execute(t *Thread) error {
	return t.Stack.Push(2)
}

func (n *iconst_3Instruction) Execute(t *Thread) error {
	return t.Stack.Push(3)
}

func (n *iconst_4Instruction) Execute(t *Thread) error {
	return t.Stack.Push(4)
}

func (n *iconst_5Instruction) Execute(t *Thread) error {
	return t.Stack.Push(5)
}

func (n *lconst_0Instruction) Execute(t *Thread) error {
	return t.Stack.PushLong(0)
}

func (n *lconst_1Instruction) Execute(t *Thread) error {
	return t.Stack.PushLong(1)
}

func (n *fconst_0Instruction) Execute(t *Thread) error {
	return t.Stack.PushFloat(0.0)
}

func (n *fconst_1Instruction) Execute(t *Thread) error {
	return t.Stack.PushFloat(1.0)
}

func (n *fconst_2Instruction) Execute(t *Thread) error {
	return t.Stack.PushFloat(2.0)
}

func (n *dconst_0Instruction) Execute(t *Thread) error {
	return t.Stack.PushDouble(0.0)
}

func (n *dconst_1Instruction) Execute(t *Thread) error {
	return t.Stack.PushDouble(1.0)
}

func (n *bipushInstruction) Execute(t *Thread) error {
	return t.Stack.Push(Int(int8(n.value)))
}

func (n *sipushInstruction) Execute(t *Thread) error {
	return t.Stack.Push(Int(int16(n.value)))
}

func (n *ldcInstruction) Execute(t *Thread) error {
	if n.isPrimitive {
		return t.Stack.Push(n.primitiveValue)
	}
	return t.Stack.PushRef(n.reference)
}

func (n *ldc_wInstruction) Execute(t *Thread) error {
	if n.isPrimitive {
		return t.Stack.Push(n.primitiveValue)
	}
	return t.Stack.PushRef(n.reference)
}

func (n *ldc2_wInstruction) Execute(t *Thread) error {
	return t.Stack.PushLong(n.primitiveValue)
}

// Reads an int from the local variables and returns it. Returns an error if
// the given index doesn't contain an int.
func getLocalInt(t *Thread, index int) (Int, error) {
	if index >= len(t.LocalVariables) {
		return 0, BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	v, ok := o.(Int)
	if !ok {
		return 0, TypeError(fmt.Sprintf("Expected to read a local int, got %s",
			o.TypeName()))
	}
	return v, nil
}

// Pushes an int from the local variable array onto the stack.
func loadLocalInt(t *Thread, index int) error {
	v, e := getLocalInt(t, index)
	if e != nil {
		return e
	}
	return t.Stack.Push(v)
}

func (n *iloadInstruction) Execute(t *Thread) error {
	return loadLocalInt(t, int(n.value))
}

// Pushes a long from the local variable array onto the stack.
func loadLocalLong(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	v, ok := o.(Long)
	if !ok {
		return TypeError(fmt.Sprintf("Expected to load a long, got %s",
			o.TypeName()))
	}
	return t.Stack.PushLong(v)
}

func (n *lloadInstruction) Execute(t *Thread) error {
	return loadLocalLong(t, int(n.value))
}

// Pushes a float from the local variable array onto the stack.
func loadLocalFloat(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	v, ok := o.(Float)
	if !ok {
		return TypeError(fmt.Sprintf("Expected to load a float, got %s",
			o.TypeName()))
	}
	return t.Stack.PushFloat(v)
}

func (n *floadInstruction) Execute(t *Thread) error {
	return loadLocalFloat(t, int(n.value))
}

// Pushes a double from the local variable array onto the stack
func loadLocalDouble(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	v, ok := o.(Double)
	if !ok {
		return TypeError(fmt.Sprintf("Expected to load a double, got %s",
			o.TypeName()))
	}
	return t.Stack.PushDouble(v)
}

func (n *dloadInstruction) Execute(t *Thread) error {
	return loadLocalDouble(t, int(n.value))
}

func loadLocalReference(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	if o.IsPrimitive() {
		return TypeError(fmt.Sprintf("Expected to load a reference, got %s",
			o.TypeName()))
	}
	return t.Stack.PushRef(o)
}

func (n *aloadInstruction) Execute(t *Thread) error {
	return loadLocalReference(t, int(n.value))
}

func (n *iload_0Instruction) Execute(t *Thread) error {
	return loadLocalInt(t, 0)
}

func (n *iload_1Instruction) Execute(t *Thread) error {
	return loadLocalInt(t, 1)
}

func (n *iload_2Instruction) Execute(t *Thread) error {
	return loadLocalInt(t, 2)
}

func (n *iload_3Instruction) Execute(t *Thread) error {
	return loadLocalInt(t, 3)
}

func (n *lload_0Instruction) Execute(t *Thread) error {
	return loadLocalLong(t, 0)
}

func (n *lload_1Instruction) Execute(t *Thread) error {
	return loadLocalLong(t, 1)
}

func (n *lload_2Instruction) Execute(t *Thread) error {
	return loadLocalLong(t, 2)
}

func (n *lload_3Instruction) Execute(t *Thread) error {
	return loadLocalLong(t, 3)
}

func (n *fload_0Instruction) Execute(t *Thread) error {
	return loadLocalFloat(t, 0)
}

func (n *fload_1Instruction) Execute(t *Thread) error {
	return loadLocalFloat(t, 1)
}

func (n *fload_2Instruction) Execute(t *Thread) error {
	return loadLocalFloat(t, 2)
}

func (n *fload_3Instruction) Execute(t *Thread) error {
	return loadLocalFloat(t, 3)
}

func (n *dload_0Instruction) Execute(t *Thread) error {
	return loadLocalDouble(t, 0)
}

func (n *dload_1Instruction) Execute(t *Thread) error {
	return loadLocalDouble(t, 1)
}

func (n *dload_2Instruction) Execute(t *Thread) error {
	return loadLocalDouble(t, 2)
}

func (n *dload_3Instruction) Execute(t *Thread) error {
	return loadLocalDouble(t, 3)
}

func (n *aload_0Instruction) Execute(t *Thread) error {
	return loadLocalReference(t, 0)
}

func (n *aload_1Instruction) Execute(t *Thread) error {
	return loadLocalReference(t, 1)
}

func (n *aload_2Instruction) Execute(t *Thread) error {
	return loadLocalReference(t, 2)
}

func (n *aload_3Instruction) Execute(t *Thread) error {
	return loadLocalReference(t, 3)
}

func (n *ialoadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(IntArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected an int array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.Push(a[i])
}

func (n *laloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(LongArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a long array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a[i])
}

func (n *faloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(FloatArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a float array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(a[i])
}

func (n *daloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(DoubleArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a double array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(a[i])
}

func (n *aaloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ReferenceArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a reference array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.PushRef(a[i])
}

func (n *baloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ByteArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a byte array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(a[i]))
}

func (n *caloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(CharArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a char array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(uint32(a[i])))
}

func (n *saloadInstruction) Execute(t *Thread) error {
	i, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ShortArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a short array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(i, len(a))
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(a[i]))
}

// Pushes an int from the local variable array onto the stack.
func storeLocalInt(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	t.LocalVariables[index] = v
	return nil
}

func (n *istoreInstruction) Execute(t *Thread) error {
	return storeLocalInt(t, int(n.value))
}

func storeLocalLong(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	v, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	t.LocalVariables[index] = v
	return nil
}

func (n *lstoreInstruction) Execute(t *Thread) error {
	return storeLocalLong(t, int(n.value))
}

func storeLocalFloat(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	v, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	t.LocalVariables[index] = v
	return nil
}

func (n *fstoreInstruction) Execute(t *Thread) error {
	return storeLocalFloat(t, int(n.value))
}

func storeLocalDouble(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	v, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	t.LocalVariables[index] = v
	return nil
}

func (n *dstoreInstruction) Execute(t *Thread) error {
	return storeLocalDouble(t, int(n.value))
}

func storeLocalRef(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	v, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	t.LocalVariables[index] = v
	return nil
}

func (n *astoreInstruction) Execute(t *Thread) error {
	return storeLocalRef(t, int(n.value))
}

func (n *istore_0Instruction) Execute(t *Thread) error {
	return storeLocalInt(t, 0)
}

func (n *istore_1Instruction) Execute(t *Thread) error {
	return storeLocalInt(t, 1)
}

func (n *istore_2Instruction) Execute(t *Thread) error {
	return storeLocalInt(t, 2)
}

func (n *istore_3Instruction) Execute(t *Thread) error {
	return storeLocalInt(t, 3)
}

func (n *lstore_0Instruction) Execute(t *Thread) error {
	return storeLocalLong(t, 0)
}

func (n *lstore_1Instruction) Execute(t *Thread) error {
	return storeLocalLong(t, 1)
}

func (n *lstore_2Instruction) Execute(t *Thread) error {
	return storeLocalLong(t, 2)
}

func (n *lstore_3Instruction) Execute(t *Thread) error {
	return storeLocalLong(t, 3)
}

func (n *fstore_0Instruction) Execute(t *Thread) error {
	return storeLocalFloat(t, 0)
}

func (n *fstore_1Instruction) Execute(t *Thread) error {
	return storeLocalFloat(t, 1)
}

func (n *fstore_2Instruction) Execute(t *Thread) error {
	return storeLocalFloat(t, 2)
}

func (n *fstore_3Instruction) Execute(t *Thread) error {
	return storeLocalFloat(t, 3)
}

func (n *dstore_0Instruction) Execute(t *Thread) error {
	return storeLocalDouble(t, 0)
}

func (n *dstore_1Instruction) Execute(t *Thread) error {
	return storeLocalDouble(t, 1)
}

func (n *dstore_2Instruction) Execute(t *Thread) error {
	return storeLocalDouble(t, 2)
}

func (n *dstore_3Instruction) Execute(t *Thread) error {
	return storeLocalDouble(t, 3)
}

func (n *astore_0Instruction) Execute(t *Thread) error {
	return storeLocalRef(t, 0)
}

func (n *astore_1Instruction) Execute(t *Thread) error {
	return storeLocalRef(t, 1)
}

func (n *astore_2Instruction) Execute(t *Thread) error {
	return storeLocalRef(t, 2)
}

func (n *astore_3Instruction) Execute(t *Thread) error {
	return storeLocalRef(t, 3)
}

// A short utility function to check whether an integer index is in the range
// of a given array size. Returns an IndexOutOfBoundsError if the index is
// invalid.
func checkArrayIndex(index Int, arrayLength int) error {
	if index < 0 {
		return IndexOutOfBoundsError(index)
	}
	if int(index) < arrayLength {
		return nil
	}
	return IndexOutOfBoundsError(index)
}

func (n *iastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(IntArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected an int array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = value
	return nil
}

func (n *lastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(LongArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a long array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = value
	return nil
}

func (n *fastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(FloatArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a float array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = value
	return nil
}

func (n *dastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(DoubleArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a double array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = value
	return nil
}

func (n *aastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ReferenceArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a reference array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = value
	return nil
}

func (n *bastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ByteArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a byte array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = Byte(value)
	return nil
}

func (n *castoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(CharArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a char array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = Char(value)
	return nil
}

func (n *sastoreInstruction) Execute(t *Thread) error {
	value, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	index, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	o, e := PopRefNotNull(t.Stack)
	if e != nil {
		return e
	}
	a, ok := o.(ShortArray)
	if !ok {
		return TypeError(fmt.Sprintf("Expected a short array, got %s",
			o.TypeName()))
	}
	e = checkArrayIndex(index, len(a))
	if e != nil {
		return e
	}
	a[index] = Short(value)
	return nil
}

func (n *popInstruction) Execute(t *Thread) error {
	_, e := t.Stack.PopUnconditional()
	return e
}

func (n *pop2Instruction) Execute(t *Thread) error {
	_, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	_, e = t.Stack.PopUnconditional()
	return e
}

// Pushes multiple values onto the stack, returning an error if any one of the
// pushes returned an error. Values are pushed in the order they are listed, so
// the first value is pushed first, etc.
func pushMultiUnconditional(stack ThreadStack, values ...Object) error {
	var e error
	for _, v := range values {
		e = stack.PushUnconditional(v)
		if e != nil {
			return e
		}
	}
	return nil
}

func (n *dupInstruction) Execute(t *Thread) error {
	o, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	return pushMultiUnconditional(t.Stack, o, o)
}

func (n *dup_x1Instruction) Execute(t *Thread) error {
	top, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	second, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	return pushMultiUnconditional(t.Stack, top, second, top)
}

func (n *dup_x2Instruction) Execute(t *Thread) error {
	var o [3]Object
	var e error
	for i := range o {
		o[i], e = t.Stack.PopUnconditional()
		if e != nil {
			return e
		}
	}
	return pushMultiUnconditional(t.Stack, o[0], o[2], o[1], o[0])
}

func (n *dup2Instruction) Execute(t *Thread) error {
	top, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	second, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	return pushMultiUnconditional(t.Stack, second, top, second, top)
}

func (n *dup2_x1Instruction) Execute(t *Thread) error {
	var o [3]Object
	var e error
	for i := range o {
		o[i], e = t.Stack.PopUnconditional()
		if e != nil {
			return e
		}
	}
	return pushMultiUnconditional(t.Stack, o[1], o[0], o[2], o[1], o[0])
}

func (n *dup2_x2Instruction) Execute(t *Thread) error {
	var o [4]Object
	var e error
	for i := range o {
		o[i], e = t.Stack.PopUnconditional()
		if e != nil {
			return e
		}
	}
	return pushMultiUnconditional(t.Stack, o[1], o[0], o[3], o[2], o[1], o[0])
}

func (n *swapInstruction) Execute(t *Thread) error {
	top, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	second, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	return pushMultiUnconditional(t.Stack, top, second)
}

// Pops and returns 2 ints from the stack. The first return value was the top.
func pop2Int(s ThreadStack) (Int, Int, error) {
	a, e := s.Pop()
	if e != nil {
		return 0, 0, e
	}
	b, e := s.Pop()
	if e != nil {
		return 0, 0, e
	}
	return a, b, nil
}

// Like pop2Int, but for longs.
func pop2Long(s ThreadStack) (Long, Long, error) {
	a, e := s.PopLong()
	if e != nil {
		return 0, 0, e
	}
	b, e := s.PopLong()
	if e != nil {
		return 0, 0, e
	}
	return a, b, nil
}

// Like pop2Int, but for floats.
func pop2Float(s ThreadStack) (Float, Float, error) {
	a, e := s.PopFloat()
	if e != nil {
		return 0, 0, e
	}
	b, e := s.PopFloat()
	if e != nil {
		return 0, 0, e
	}
	return a, b, nil
}

// Like pop2Int, but for doubles.
func pop2Double(s ThreadStack) (Double, Double, error) {
	a, e := s.PopDouble()
	if e != nil {
		return 0, 0, e
	}
	b, e := s.PopDouble()
	if e != nil {
		return 0, 0, e
	}
	return a, b, nil
}

func (n *iaddInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(a + b)
}

func (n *laddInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a + b)
}

func (n *faddInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(a + b)
}

func (n *daddInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(a + b)
}

func (n *isubInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(b - a)
}

func (n *lsubInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(b - a)
}

func (n *fsubInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(b - a)
}

func (n *dsubInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(b - a)
}

func (n *imulInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(a * b)
}

func (n *lmulInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a * b)
}

func (n *fmulInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(a * b)
}

func (n *dmulInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(a * b)
}

func (n *idivInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if a == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.Push(b / a)
}

func (n *ldivInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	if a == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.PushLong(b / a)
}

func (n *fdivInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(b / a)
}

func (n *ddivInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(b / a)
}

func (n *iremInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if a == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.Push(b % a)
}

func (n *lremInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	if a == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.PushLong(b % a)
}

// This is the same as the IEEE 754 remainder, but using truncation rather than
// rounding.
func javaRemainder(a, b float64) float64 {
	return math.Remainder(math.Trunc(a), math.Trunc(b))
}

func (n *fremInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	// This is required behavior according to the JVM spec.
	if int64(a) == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.PushFloat(Float(javaRemainder(float64(b), float64(a))))
}

func (n *dremInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	if int64(a) == 0 {
		return ArithmeticError("Division by zero")
	}
	return t.Stack.PushDouble(Double(javaRemainder(float64(b), float64(a))))
}

func (n *inegInstruction) Execute(t *Thread) error {
	a, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.Push(-a)
}

func (n *lnegInstruction) Execute(t *Thread) error {
	a, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(-a)
}

func (n *fnegInstruction) Execute(t *Thread) error {
	a, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(-a)
}

func (n *dnegInstruction) Execute(t *Thread) error {
	a, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(-a)
}

func (n *ishlInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(b << uint(a&0x1f))
}

func (n *lshlInstruction) Execute(t *Thread) error {
	a, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	b, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(b << uint(a&0x3f))
}

func (n *ishrInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(b >> uint(a&0x1f))
}

func (n *lshrInstruction) Execute(t *Thread) error {
	a, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	b, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(b >> uint(a&0x3f))
}

func (n *iushrInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(uint32(b) >> uint32(a&0x1f)))
}

func (n *lushrInstruction) Execute(t *Thread) error {
	a, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	b, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(Long(uint64(b) >> uint64(a&0x3f)))
}

func (n *iandInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(a & b)
}

func (n *landInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a & b)
}

func (n *iorInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(a | b)
}

func (n *lorInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a | b)
}

func (n *ixorInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.Push(a ^ b)
}

func (n *lxorInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	return t.Stack.PushLong(a ^ b)
}

func (n *iincInstruction) Execute(t *Thread) error {
	v, e := getLocalInt(t, int(n.index))
	if e != nil {
		return e
	}
	t.LocalVariables[n.index] = v + Int(n.value)
	return nil
}

func (n *i2lInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(Long(v))
}

func (n *i2fInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(Float(v))
}

func (n *i2dInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(Double(v))
}

func (n *l2iInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(v))
}

func (n *l2fInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(Float(v))
}

func (n *l2dInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(Double(v))
}

func (n *f2iInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(v))
}

func (n *f2lInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(Long(v))
}

func (n *f2dInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(Double(v))
}

func (n *d2iInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(v))
}

func (n *d2lInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(Long(v))
}

func (n *d2fInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(Float(v))
}

func (n *i2bInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(Byte(v)))
}

func (n *i2cInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(Char(v)))
}

func (n *i2sInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	return t.Stack.Push(Int(Short(v)))
}

func (n *lcmpInstruction) Execute(t *Thread) error {
	a, b, e := pop2Long(t.Stack)
	if e != nil {
		return e
	}
	v := Int(0)
	if b > a {
		v = 1
	} else if b < a {
		v = -1
	}
	return t.Stack.Push(v)
}

func (n *fcmplInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	if math.IsNaN(float64(a)) || math.IsNaN(float64(b)) {
		return t.Stack.Push(-1)
	}
	v := Int(0)
	if b > a {
		v = 1
	} else if b < a {
		v = -1
	}
	return t.Stack.Push(v)
}

func (n *fcmpgInstruction) Execute(t *Thread) error {
	a, b, e := pop2Float(t.Stack)
	if e != nil {
		return e
	}
	if math.IsNaN(float64(a)) || math.IsNaN(float64(b)) {
		return t.Stack.Push(1)
	}
	v := Int(0)
	if b > a {
		v = 1
	} else if b < a {
		v = -1
	}
	return t.Stack.Push(v)
}

func (n *dcmplInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	if math.IsNaN(float64(a)) || math.IsNaN(float64(b)) {
		return t.Stack.Push(-1)
	}
	v := Int(0)
	if b > a {
		v = 1
	} else if b < a {
		v = -1
	}
	return t.Stack.Push(v)
}

func (n *dcmpgInstruction) Execute(t *Thread) error {
	a, b, e := pop2Double(t.Stack)
	if e != nil {
		return e
	}
	if math.IsNaN(float64(a)) || math.IsNaN(float64(b)) {
		return t.Stack.Push(1)
	}
	v := Int(0)
	if b > a {
		v = 1
	} else if b < a {
		v = -1
	}
	return t.Stack.Push(v)
}

func (n *ifeqInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v == 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *ifneInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v != 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *ifltInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v < 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *ifgeInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v >= 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *ifgtInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v > 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *ifleInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if v <= 0 {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpeqInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if a == b {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpneInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if a != b {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpltInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if b < a {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpgeInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if b >= a {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpgtInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if b > a {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_icmpleInstruction) Execute(t *Thread) error {
	a, b, e := pop2Int(t.Stack)
	if e != nil {
		return e
	}
	if b <= a {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_acmpeqInstruction) Execute(t *Thread) error {
	a, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	b, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	if a == b {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *if_acmpneInstruction) Execute(t *Thread) error {
	a, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	b, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	if a != b {
		t.InstructionIndex = n.nextIndex
		t.WasBranch = true
	}
	return nil
}

func (n *gotoInstruction) Execute(t *Thread) error {
	t.InstructionIndex = n.nextIndex
	t.WasBranch = true
	return nil
}

func (n *jsrInstruction) Execute(t *Thread) error {
	e := t.Stack.Push(Int(n.returnIndex))
	if e != nil {
		return e
	}
	t.InstructionIndex = n.nextIndex
	t.WasBranch = true
	return nil
}

func (n *retInstruction) Execute(t *Thread) error {
	// Remember that our "return address" type is an int corresponding to an
	// instruction *index*.
	returnIndex, e := getLocalInt(t, int(n.value))
	if e != nil {
		return e
	}
	t.InstructionIndex = uint(returnIndex)
	t.WasBranch = true
	return nil
}

func (n *tableswitchInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	if (uint32(v) < n.lowIndex) || (uint32(v) > n.highIndex) {
		t.InstructionIndex = n.defaultIndex
		t.WasBranch = true
		return nil
	}
	t.InstructionIndex = n.indices[uint32(v)-n.lowIndex]
	t.WasBranch = true
	return nil
}

func (n *lookupswitchInstruction) Execute(t *Thread) error {
	v, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	i := sort.Search(len(n.pairs), func(i int) bool {
		return int32(v) >= n.pairs[i].match
	})
	if (i >= len(n.pairs)) || (n.pairs[i].match != int32(v)) {
		t.InstructionIndex = n.defaultIndex
		t.WasBranch = true
		return nil
	}
	t.InstructionIndex = n.indices[i]
	t.WasBranch = true
	return nil
}

func (n *ireturnInstruction) Execute(t *Thread) error {
	returnValue, e := t.Stack.Pop()
	if e != nil {
		return e
	}
	// We should have already checked that this type is OK during the optimize
	// pass.
	r := t.CurrentMethod.Types.ReturnType.(class_file.PrimitiveFieldType)
	// Convert the value popped off the stack if the return value of the method
	// is something shorter, such as a byte, char, boolean, or short.
	switch r {
	case 'B':
		returnValue &= 0xff
	case 'C':
		returnValue = Int(Char(returnValue))
	case 'S':
		returnValue = Int(Short(returnValue))
	case 'Z':
		// This is how the JVM spec requires converting an int to a boolean.
		returnValue &= 1
	}
	// This may return a ThreadExitedError if this was the initial method in
	// a thread.
	e = t.Return()
	if e != nil {
		return e
	}
	return t.Stack.Push(returnValue)
}

func (n *lreturnInstruction) Execute(t *Thread) error {
	returnValue, e := t.Stack.PopLong()
	if e != nil {
		return e
	}
	// Optimize() already should have verified that this method returns a long.
	e = t.Return()
	if e != nil {
		return e
	}
	return t.Stack.PushLong(returnValue)
}

func (n *freturnInstruction) Execute(t *Thread) error {
	returnValue, e := t.Stack.PopFloat()
	if e != nil {
		return e
	}
	// The expected return type was already checked by Optimize()
	e = t.Return()
	if e != nil {
		return e
	}
	return t.Stack.PushFloat(returnValue)
}

func (n *dreturnInstruction) Execute(t *Thread) error {
	returnValue, e := t.Stack.PopDouble()
	if e != nil {
		return e
	}
	// The expected return type was already checked by Optimize()
	e = t.Return()
	if e != nil {
		return e
	}
	return t.Stack.PushDouble(returnValue)
}

func (n *areturnInstruction) Execute(t *Thread) error {
	returnValue, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	// The expected return type was already (somewhat) checked by Optimize()
	// NOTE: Should we verify the returned type is what we expect here? I.e.,
	// make sure class names and/or array dimensions match up?
	e = t.Return()
	if e != nil {
		return e
	}
	return t.Stack.PushRef(returnValue)
}

func (n *returnInstruction) Execute(t *Thread) error {
	return t.Return()
}

func (n *getstaticInstruction) Execute(t *Thread) error {
	v := n.class.StaticFieldValues[n.index]
	return t.Stack.PushUnconditional(v)
}

func (n *putstaticInstruction) Execute(t *Thread) error {
	// We'll first look up the type that's stored in the field in order to pop
	// the right type from the stack.
	targetValue := n.class.StaticFieldValues[n.index]

	// First, if this isn't a primitive it must be a reference, so we'll pop a
	// reference off the stack and store it.
	if !targetValue.IsPrimitive() {
		newValue, e := t.Stack.PopRef()
		if e != nil {
			return e
		}
		e = AssignmentOK(newValue, targetValue)
		if e != nil {
			return TypeError(fmt.Sprintf("Trying to assign incompatible type "+
				"to static field: %s", e))
		}
		n.class.StaticFieldValues[n.index] = newValue
		return nil
	}

	// Now that we know the value was a primitive we will need to pop the right
	// type of primitive off the stack.
	var newValue PrimitiveType
	var e error

	// We only care about floats, longs, and doubles. By default, we pop an
	// int, since that's the smallest integral primitive that can be pushed
	// onto the stack.
	switch targetValue.(type) {
	case Double:
		newValue, e = t.Stack.PopDouble()
	case Float:
		newValue, e = t.Stack.PopFloat()
	case Long:
		newValue, e = t.Stack.PopLong()
	default:
		newValue, e = t.Stack.Pop()
	}
	if e != nil {
		return e
	}

	// No matter what we popped from the stack, this will allow us to convert
	// it to the correct type before storing it.
	tmp := targetValue.(PrimitiveType)
	toStore := tmp.ConvertFrom(newValue)
	n.class.StaticFieldValues[n.index] = toStore
	return nil
}

func (n *getfieldInstruction) Execute(t *Thread) error {
	v, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	instance, ok := v.(*ClassInstance)
	if !ok {
		return TypeError(fmt.Sprintf("Didn't get a valid object for getfield."+
			" Instead got %s", v.String()))
	}
	fieldName := string(n.fieldReference.Field.Name)
	instance, fieldIndex, e := instance.ResolveField(fieldName)
	if e != nil {
		return fmt.Errorf("Couldn't resolve field %s in %s: %s", fieldName,
			instance.String(), e)
	}
	return t.Stack.PushUnconditional(instance.FieldValues[fieldIndex])
}

func (n *putfieldInstruction) Execute(t *Thread) error {
	// Unfortunately, unlike with putstatic, we can't just look at the type
	// that's in the field, because the value to store in the field needs to
	// be popped of the stack first (and we don't know its size yet). So, we'll
	// need to actually look at the type of the FieldOrMethodReference to
	// figure out what to pop.
	nameAndType := n.fieldReference.Field
	fieldType, e := class_file.ParseFieldType(nameAndType.Type)
	if e != nil {
		return fmt.Errorf("Invalid field type descriptor: %s", e)
	}
	var toStore Object
	primitiveFieldType, isPrimitive :=
		fieldType.(class_file.PrimitiveFieldType)
	if !isPrimitive {
		toStore, e = t.Stack.PopRef()
	} else {
		switch primitiveFieldType {
		case 'B', 'C', 'I', 'S', 'Z':
			// Bytes, chars, ints, shorts, and bools are all stored as "ints"
			// on the stack.
			toStore, e = t.Stack.Pop()
		case 'F':
			toStore, e = t.Stack.PopFloat()
		case 'D':
			toStore, e = t.Stack.PopDouble()
		case 'L':
			toStore, e = t.Stack.PopLong()
		default:
			return fmt.Errorf("Unknown primitive field type: %s",
				primitiveFieldType)
			// Afterwards, check the error, then check AssignmentOK
		}
	}
	if e != nil {
		// This may be a stack-empty error; we've already checked for an
		// invalid primitiveFieldType.
		return e
	}

	// Now that we've popped the value to store from the stack, we can get the
	// object reference and figure out where to store it.
	tmp, e := t.Stack.PopRef()
	if e != nil {
		return e
	}
	classInstance, ok := tmp.(*ClassInstance)
	if !ok {
		return fmt.Errorf("Expected to pop a class instance from the stack, "+
			"got %s instead", tmp.String())
	}
	classInstance, fieldIndex, e := classInstance.ResolveField(
		string(nameAndType.Name))
	if e != nil {
		return fmt.Errorf("Couldn't resolve field %s.%s: %s",
			classInstance.TypeName(), nameAndType.Name, e)
	}

	// Finally! Check that it's okay to assign the value to the one in the
	// field, and actually update the field.
	e = AssignmentOK(toStore, classInstance.FieldValues[fieldIndex])
	if e != nil {
		return TypeError(fmt.Sprintf("Trying to assign incompatible type to "+
			"%s.%s: %s", classInstance.TypeName(), nameAndType.Name, e))
	}
	classInstance.FieldValues[fieldIndex] = toStore
	return nil
}

func (n *invokevirtualInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *invokespecialInstruction) Execute(t *Thread) error {
	return t.Call(n.method)
}

func (n *invokestaticInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *invokeinterfaceInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *invokedynamicInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *newInstruction) Execute(t *Thread) error {
	instance, e := n.class.CreateInstance()
	if e != nil {
		return fmt.Errorf("new %s failed: %w", n.class.Name, e)
	}
	return t.Stack.PushRef(instance)
}

func (n *newarrayInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *anewarrayInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *arraylengthInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *athrowInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *checkcastInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *instanceofInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *monitorenterInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *monitorexitInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *wideInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *wideIincInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *multianewarrayInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifnullInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifnonnullInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *goto_wInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *jsr_wInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *breakpointInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *impdep1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *impdep2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}
