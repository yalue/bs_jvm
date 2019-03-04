package bs_jvm

import (
	"fmt"
	"math"
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

func (n *dupInstruction) Execute(t *Thread) error {
	o, e := t.Stack.PopUnconditional()
	if e != nil {
		return e
	}
	e = t.Stack.PushUnconditional(o)
	if e != nil {
		return e
	}
	return t.Stack.PushUnconditional(o)
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
	// TODO (next): Implement iflt instruction, including Optimize(...)
	return NotImplementedError
}

func (n *ifgeInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifgtInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifleInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpeqInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpneInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpltInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpgeInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpgtInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_icmpleInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_acmpeqInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *if_acmpneInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *gotoInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *jsrInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *retInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *tableswitchInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lookupswitchInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ireturnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lreturnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *freturnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dreturnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *areturnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *returnInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *getstaticInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *putstaticInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *getfieldInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *putfieldInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *invokevirtualInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *invokespecialInstruction) Execute(t *Thread) error {
	return NotImplementedError
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
	return NotImplementedError
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
