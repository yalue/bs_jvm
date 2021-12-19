// BS-JVM: The Blinding Speed JVM.
//
// A JVM library for the Go programming language.
package bs_jvm

import (
	"fmt"
	"github.com/yalue/bs_jvm/class_file"
	"os"
	"sync"
)

// Holds the state of a single JVM thread.
type Thread struct {
	// The method that the thread is currently executing.
	CurrentMethod *Method
	// A pointer to the JVM running this thread.
	ParentJVM *JVM
	// The index of the current instruction in the method.
	InstructionIndex uint
	// This will be true if the last executed instruction performed a branch.
	WasBranch bool
	// The stack for this thread
	Stack ThreadStack
	// The list of local variables, starting with arguments.
	LocalVariables []Object
	// A channel that will contain the thread exit reason when the thread has
	// finished.
	threadComplete chan error
	// Set this to a non-nil value to force the thread to exit before the next
	// instruction. If not set by an external reason, this will be set when a
	// thread exits normally.
	ThreadExitReason error
	// The index into the JVM's list of active threads. ONLY ACCESS THIS
	// (INCLUDING JUST FOR READS) WHILE HOLDING THE PARENT JVM THREAD LIST
	// LOCK.
	threadIndex int
}

// This method will cause a thread to start running. The thread will run
// asynchronously, so this function only returns an error if the thread failed
// to start.
func (t *Thread) Run() error {
	go func() {
		var e error
		var n Instruction
		for e == nil {
			if t.ThreadExitReason != nil {
				t.threadComplete <- t.ThreadExitReason
				close(t.threadComplete)
				return
			}
			if t.InstructionIndex >= uint(len(t.CurrentMethod.Instructions)) {
				e = fmt.Errorf("Invalid instruction index: %d",
					t.InstructionIndex)
				break
			}
			t.WasBranch = false
			n = t.CurrentMethod.Instructions[t.InstructionIndex]
			fmt.Printf("Running instruction: %s\n", n.String())
			e = n.Execute(t)
			if !t.WasBranch {
				// Go to the next instruction in the sequence if we didn't
				// encounter a branch.
				t.InstructionIndex++
			}
		}
		t.ThreadExitReason = e
		t.threadComplete <- e
		close(t.threadComplete)
	}()
	return nil
}

// This will return when the thread is complete. Returns the reason the thread
// exited (will return ThreadExitedError on a normal exit, rather than nil.
func (t *Thread) WaitForCompletion() error {
	// End when either a value is received or the channel was closed.
	_, ok := <-t.threadComplete
	// Only one waiter can possibly receive the "ok" value, so it will be
	// responsible for cleaning up the list of threads.
	if ok {
		t.ParentJVM.lockThreadList()
		// Swap the thread at the end of the list into this thread's position.
		currentThreadCount := len(t.ParentJVM.threads)
		t.ParentJVM.threads[currentThreadCount-1].threadIndex = t.threadIndex
		t.ParentJVM.threads[t.threadIndex] =
			t.ParentJVM.threads[currentThreadCount-1]
		t.ParentJVM.threads = t.ParentJVM.threads[0:(currentThreadCount - 1)]
		t.ParentJVM.unlockThreadList()
	}
	return t.ThreadExitReason
}

// This function ends the thread, passing the given error to the ThreadComplete
// channel. This should only ever be called once, otherwise odd behavior may
// occur.
func (t *Thread) EndThread(e error) {
	t.ThreadExitReason = e
}

// To be used when the current instruction is calling a method. The returned
// frame should be "restored" when the called method returns.
func (t *Thread) GetReturnInfo() ReturnInfo {
	return ReturnInfo{
		Method:         t.CurrentMethod,
		ReturnIndex:    t.InstructionIndex + 1,
		StackState:     t.Stack.GetSizes(),
		LocalVariables: t.LocalVariables,
	}
}

// Restores the given method frame. Does not modify r. Used when returning.
func (t *Thread) RestoreReturnInfo(r *ReturnInfo) error {
	t.CurrentMethod = r.Method
	t.InstructionIndex = r.ReturnIndex
	e := t.Stack.RestoreSizes(&(r.StackState))
	if e != nil {
		return e
	}
	t.LocalVariables = r.LocalVariables
	return nil
}

// Carries out a method call, including pushing the return location. Returns an
// error if one occurs. Expects the instruction index to point at the
// instruction causing the call.
func (t *Thread) Call(method *Method) error {
	// First, check for a native implementation; there's no further action if
	// we're just calling something native.
	if method.Native != nil {
		return method.Native(t)
	}
	if (t.InstructionIndex + 1) >= uint(len(t.CurrentMethod.Instructions)) {
		return fmt.Errorf("Invalid return address (inst. index %d)",
			t.InstructionIndex)
	}
	e := t.Stack.PushFrame(t.GetReturnInfo())
	if e != nil {
		return e
	}
	// TODO: Optimize local variable allocation so we don't have an allocation
	// per method invocation. IDEA: Each thread maintains a simple "stack" of
	// local variables, grown if needed. When popping local variables, etc, we
	// simply set the slice. When calling, we just use the next slice, assuming
	// it's big enough. Will simply increase capacity when needed.
	// TODO: Initialize local variables with argument contents.
	t.LocalVariables = make([]Object, method.MaxLocals)
	t.CurrentMethod = method
	t.InstructionIndex = 0
	return nil
}

// Carries out a method return, popping a return location. If the thread's
// initial method returns in the thread, this ends the thread and returns nil.
func (t *Thread) Return() error {
	returnInfo, e := t.Stack.PopFrame()
	if e == StackEmptyError {
		t.EndThread(ThreadExitedError)
		return ThreadExitedError
	}
	if e != nil {
		return e
	}
	// TODO: Write the return value onto the stack?
	return t.RestoreReturnInfo(&returnInfo)
}

// Holds state of the entire JVM, including threads, class files, etc.
type JVM struct {
	// A list of threads in the JVM.
	threads []*Thread
	// This lock is acquired whenever the list of active threads must be
	// modified.
	threadsLock sync.Mutex
	// Maps class names to all loaded classes.
	Classes map[string]*Class
}

// Returns a new, uninitialized, JVM instance.
func NewJVM() *JVM {
	return &JVM{
		threads: make([]*Thread, 0, 1),
		Classes: make(map[string]*Class),
	}
}

// This is a function type that is used for method implementations written
// in Go.
type NativeMethod func(t *Thread) error

// Holds a parsed JVM method.
type Method struct {
	// The class in which the method was defined.
	ContainingClass *Class
	// The name of the method. Mostly for debugging purposes.
	Name string
	// The argument and return types of the method.
	Types *class_file.MethodDescriptor
	// The number of local variables used by the method, more or less. Note
	// that doubles and longs will be counted twice here, which will currently
	// waste a bit of space in our implementation... oh well.
	MaxLocals int
	// Contains all parsed functions in the method.
	Instructions []Instruction
	// The raw binary of the function's code.
	CodeBytes []byte
	// This will be true if the "Optimize" pass is done. Must be done before
	// calling the method.
	OptimizeDone bool
	// This can be used for Go-implemented methods, but otherwise must be nil.
	// If this is non-nil, most of the other fields of the Method struct may be
	// nil, so check this first when invoking a method.
	Native NativeMethod
}

// Parses the given method from the class file into the structure needed by the
// JVM for actual execution. Does *not* modify the state of the JVM. The
// returned Method's Instructions slice will *not* be populated until the
// Method's Optimize() function is called.
func (j *JVM) NewMethod(class *Class, index int) (*Method, error) {
	classFile := class.File
	if (index < 0) || (index >= len(classFile.Methods)) {
		return nil, fmt.Errorf("Invalid method index: %d", index)
	}
	method := classFile.Methods[index]
	codeAttribute, e := method.GetCodeAttribute(classFile)
	if e != nil {
		return nil, fmt.Errorf("Failed getting method code attribute: %s", e)
	}
	codeBytes := codeAttribute.Code
	codeMemory := MemoryFromSlice(codeBytes)
	var instruction Instruction
	address := uint(0)
	instructionCount := 0
	// This initial pass only counts the number of instructions in the method.
	for address < uint(len(codeBytes)) {
		instruction, e = GetNextInstruction(codeMemory, address)
		if e != nil {
			return nil, fmt.Errorf("Error reading instruction: %s", e)
		}
		instructionCount++
		address += instruction.Length()
	}
	toReturn := Method{
		ContainingClass: class,
		Name:            string(method.Name),
		Types:           method.Descriptor,
		MaxLocals:       int(codeAttribute.MaxLocals),
		Instructions:    make([]Instruction, instructionCount),
		CodeBytes:       codeBytes,
		OptimizeDone:    false,
	}
	return &toReturn, nil
}

// This does the "optimization" pass on the method if it hasn't already been
// done. Returns an error if one occurs. Immediately returns nil if
// m.OptimizeDone is already true.
func (m *Method) Optimize() error {
	if m.OptimizeDone {
		return nil
	}
	address := uint(0)
	var e error
	var instruction Instruction
	codeMemory := MemoryFromSlice(m.CodeBytes)
	instructionCount := len(m.Instructions)

	// Create the instruction objects, and make a map of instruction offsets ->
	// indices in the Instructions slice. This map is used in the next pass,
	// when calling the "optimize" function.
	offsetMap := make(map[uint]int)
	for i := 0; i < instructionCount; i++ {
		instruction, e = GetNextInstruction(codeMemory, address)
		if e != nil {
			return fmt.Errorf("Error reading instruction: %s", e)
		}
		m.Instructions[i] = instruction
		offsetMap[address] = i
		address += instruction.Length()
	}

	// Finally, call the "optimize" function on every instruction.
	address = 0
	for i := 0; i < len(m.Instructions); i++ {
		instruction = m.Instructions[i]
		e = instruction.Optimize(m, address, offsetMap)
		if e != nil {
			return fmt.Errorf("Error in optimization pass over %s: %s",
				instruction, e)
		}
		address += instruction.Length()
	}
	m.OptimizeDone = true
	return nil
}

// Adds the given class file to the JVM so that its code
func (j *JVM) LoadClass(class *class_file.Class) error {
	loadedClass, e := NewClass(j, class)
	if e != nil {
		return fmt.Errorf("Error loading class: %s", e)
	}
	j.Classes[string(loadedClass.Name)] = loadedClass
	return nil
}

// Returns a reference to the named class. Returns a ClassNotFoundError if the
// class hasn't been loaded.
func (j *JVM) GetClass(name string) (*Class, error) {
	// TODO: Make a GetOrLoadClass function, that can potentially load classes
	// during the "optimize" pass if they're needed.
	toReturn := j.Classes[name]
	if toReturn == nil {
		return nil, ClassNotFoundError(name)
	}
	return toReturn, nil
}

// Shorthand for acquiring the lock on the list of active threads.
func (j *JVM) lockThreadList() {
	(&(j.threadsLock)).Lock()
}

// Shorthand for releasing the lock on the list of active threads.
func (j *JVM) unlockThreadList() {
	(&(j.threadsLock)).Unlock()
}

// Shorthand for calling GetMethod on the named class.
func (j *JVM) GetMethod(className, methodKey string) (*Method, error) {
	c := j.Classes[className]
	if c == nil {
		return nil, ClassNotFoundError(className)
	}
	return c.GetMethod(methodKey)
}

// Spawns a new thread in the JVM, with the given method. The methodKey must
// follow the format returned by the GetMethodKey function.
func (j *JVM) StartThread(className, methodKey string) error {
	method, e := j.GetMethod(className, methodKey)
	if e != nil {
		return e
	}
	// We may need to optimize this method in case this is the first time it's
	// being invoked.
	e = method.Optimize()
	if e != nil {
		return fmt.Errorf("Failed preparing thread's start method for "+
			"execution: %s", e)
	}
	j.lockThreadList()
	threadIndex := len(j.threads)
	newThread := Thread{
		CurrentMethod:    method,
		ParentJVM:        j,
		InstructionIndex: 0,
		LocalVariables:   make([]Object, method.MaxLocals),
		Stack:            NewStack(),
		threadComplete:   make(chan error),
		threadIndex:      threadIndex,
	}
	j.threads = append(j.threads, &newThread)
	j.unlockThreadList()
	e = (&newThread).Run()
	if e != nil {
		return e
	}
	return nil
}

// Waits for all threads. May return any error from any thread if the thread
// has any error other than ThreadExitedError. Will return nil if all threads
// exited successfully.
func (j *JVM) WaitForAllThreads() error {
	var currentThread *Thread
	var toReturn error
	var currentError error
	for {
		j.lockThreadList()
		if len(j.threads) <= 0 {
			j.unlockThreadList()
			break
		}
		currentThread = j.threads[len(j.threads)-1]
		j.unlockThreadList()
		currentError = currentThread.WaitForCompletion()
		// Only returns errors that aren't ThreadExitedErrors
		if currentError != ThreadExitedError {
			if currentError == nil {
				currentError = fmt.Errorf("Invalid nil thread exit value")
			}
			toReturn = currentError
		}
	}
	return toReturn
}

// A simple wrapper around LoadClass that takes a class filename instead of a
// parsed file. Returns the name of the loaded class on success.
func (j *JVM) LoadClassFromFile(classFileName string) (string, error) {
	file, e := os.Open(classFileName)
	if e != nil {
		return "", fmt.Errorf("Failed opening class file: %s", e)
	}
	defer file.Close()
	classFile, e := class_file.ParseClass(file)
	if e != nil {
		return "", e
	}
	className, e := classFile.GetName()
	if e != nil {
		return "", fmt.Errorf("Failed getting class name: %s", e)
	}
	e = j.LoadClass(classFile)
	if e != nil {
		return "", e
	}
	return string(className), nil
}

// Gets the correctly formatted key for looking up the "main" method in our
// internal Methods map.
func getMainMethodKey() string {
	stringArrayType := &class_file.ArrayType{
		Dimensions:  1,
		ContentType: class_file.ClassInstanceType("java/lang/String"),
	}
	tmp := &class_file.Method{
		// public static
		Access: 1 | 8,
		// main
		Name: []byte("main"),
		// void
		Descriptor: &class_file.MethodDescriptor{
			ArgumentTypes: []class_file.FieldType{stringArrayType},
			ReturnType:    class_file.PrimitiveFieldType('V'),
		},
	}
	return GetMethodKey(tmp)
}

// Takes a path to a class file, parses and loads the class, then looks for the
// main function in the class and starts executing it.
func (j *JVM) StartMainClass(classFileName string) error {
	className, e := j.LoadClassFromFile(classFileName)
	if e != nil {
		return e
	}
	e = j.StartThread(className, getMainMethodKey())
	return e
}
