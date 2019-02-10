package bs_jvm

// This file contains functions for executing individual JVM instructions.

func (n *nopInstruction) Execute(t *Thread) error {
	return nil
}

func (n *aconst_nullInstruction) Execute(t *Thread) error {
	return t.References.Push(nil)
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
	return t.References.Push(n.reference)
}

func (n *ldc_wInstruction) Execute(t *Thread) error {
	if n.isPrimitive {
		return t.Stack.Push(n.primitiveValue)
	}
	return t.References.Push(n.reference)
}

func (n *ldc2_wInstruction) Execute(t *Thread) error {
	return t.Stack.PushLong(n.primitiveValue)
}

// Pushes an int from the local variable array onto the stack.
func loadLocalInt(t *Thread, index int) error {
	if index >= len(t.LocalVariables) {
		return BadLocalVariableError(index)
	}
	o := t.LocalVariables[index]
	v, ok := o.(Int)
	if !ok {
		return TypeError("Expected to load an int")
	}
	return t.Stack.Push(v)
}

func (n *iloadInstruction) Execute(t *Thread) error {
	return loadLocalInt(t, int(n.value))
}

func (n *lloadInstruction) Execute(t *Thread) error {
	if int(n.value) >= len(t.LocalVariables) {
		return BadLocalVariableError(n.value)
	}
	o := t.LocalVariables[n.value]
	v, ok := o.(Long)
	if !ok {
		return TypeError("Expected to load a long")
	}
	return t.Stack.PushLong(v)
}

func (n *floadInstruction) Execute(t *Thread) error {
	if int(n.value) >= len(t.LocalVariables) {
		return BadLocalVariableError(n.value)
	}
	o := t.LocalVariables[n.value]
	v, ok := o.(Float)
	if !ok {
		return TypeError("Expected to load a float")
	}
	return t.Stack.PushFloat(v)
}

func (n *dloadInstruction) Execute(t *Thread) error {
	if int(n.value) >= len(t.LocalVariables) {
		return BadLocalVariableError(n.value)
	}
	o := t.LocalVariables[n.value]
	v, ok := o.(Double)
	if !ok {
		return TypeError("Expected to load a double")
	}
	return t.Stack.PushDouble(v)
}

func (n *aloadInstruction) Execute(t *Thread) error {
	if int(n.value) >= len(t.LocalVariables) {
		return BadLocalVariableError(n.value)
	}
	o := t.LocalVariables[n.value]
	if o.IsPrimitive() {
		return TypeError("Expected to load a reference")
	}
	return t.References.Push(o)
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
	// TODO (next): Implement lload_0. Start by implementing something like
	// loadLocalInt for longs.
	return NotImplementedError
}

func (n *lload_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lload_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lload_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fload_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fload_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fload_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fload_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dload_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dload_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dload_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dload_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aload_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aload_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aload_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aload_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ialoadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *laloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *faloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *daloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aaloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *baloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *caloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *saloadInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *istoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lstoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fstoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dstoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *astoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *istore_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *istore_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *istore_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *istore_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lstore_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lstore_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lstore_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lstore_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fstore_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fstore_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fstore_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fstore_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dstore_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dstore_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dstore_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dstore_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *astore_0Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *astore_1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *astore_2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *astore_3Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *aastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *bastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *castoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *sastoreInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *popInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *pop2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dupInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dup_x1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dup_x2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dup2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dup2_x1Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dup2_x2Instruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *swapInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iaddInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *laddInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *faddInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *daddInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *isubInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lsubInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fsubInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dsubInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *imulInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lmulInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fmulInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dmulInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *idivInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ldivInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fdivInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ddivInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iremInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lremInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fremInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dremInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *inegInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lnegInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fnegInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dnegInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ishlInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lshlInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ishrInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lshrInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iushrInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lushrInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iandInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *landInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iorInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lorInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ixorInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lxorInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *iincInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2lInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2fInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2dInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *l2iInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *l2fInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *l2dInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *f2iInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *f2lInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *f2dInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *d2iInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *d2lInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *d2fInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2bInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2cInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *i2sInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *lcmpInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fcmplInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *fcmpgInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dcmplInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *dcmpgInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifeqInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifneInstruction) Execute(t *Thread) error {
	return NotImplementedError
}

func (n *ifltInstruction) Execute(t *Thread) error {
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
