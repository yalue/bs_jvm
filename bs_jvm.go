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
	// The containing JVM.
	ParentJVM *JVM
	// The index of the current instruction in the method.
	InstructionIndex uint
	// The stack for this thread, split into separate groups.
	Stack           DataStack
	References      ReferenceStack
	ReturnLocations CallStack
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
				t.ThreadExitReason = e
				t.threadComplete <- t.ThreadExitReason
				close(t.threadComplete)
				return
			}
			if t.InstructionIndex >= uint(len(t.CurrentMethod.Instructions)) {
				e = fmt.Errorf("Invalid instruction index: %d",
					t.InstructionIndex)
				break
			}
			n = t.CurrentMethod.Instructions[t.InstructionIndex]
			e = n.Execute(t)
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
func (t *Thread) GetReturnFrame() MethodFrame {
	return MethodFrame{
		Method:         t.CurrentMethod,
		ReturnIndex:    t.InstructionIndex + 1,
		ReferenceFrame: t.References.GetFrame(),
		DataFrame:      t.Stack.GetFrame(),
	}
}

// Restores the given method frame. Does not modify f.
func (t *Thread) RestoreFrame(f *MethodFrame) {
	t.CurrentMethod = f.Method
	t.InstructionIndex = f.ReturnIndex
	t.References.SetFrame(f.ReferenceFrame)
	t.Stack.SetFrame(f.DataFrame)
}

// Carries out a method call, including pushing the return location. Returns an
// error if one occurs. Expects the instruction index to point at the
// instruction causing the call.
func (t *Thread) Call(method *Method) error {
	// TODO: Test call/ret when ready--how are arguments passed?
	if (t.InstructionIndex + 1) >= uint(len(t.CurrentMethod.Instructions)) {
		return fmt.Errorf("Invalid return address (inst. index %d)",
			t.InstructionIndex)
	}
	e := t.ReturnLocations.Push(t.GetReturnFrame())
	if e != nil {
		return e
	}
	t.CurrentMethod = method
	t.InstructionIndex = 0
	return nil
}

// Carries out a method return, popping a return location. If the thread's
// initial method returns in the thread, this ends the thread and returns nil.
func (t *Thread) Return() error {
	frame, e := t.ReturnLocations.Pop()
	if e == StackEmptyError {
		t.EndThread(ThreadExitedError)
		return nil
	}
	if e != nil {
		return e
	}
	t.RestoreFrame(&frame)
	return nil
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

// Holds a parsed JVM method.
type Method struct {
	// A reference to the parent JVM.
	ParentJVM *JVM
	// The class in which the method was defined.
	ContainingClass *Class
	// Contains all parsed functions in the method.
	Instructions []Instruction
}

// Parses the given method from the class file into the structure needed by the
// JVM for actual execution. Also carries out pre-optimization. Does *not*
// modify the state of the JVM.
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
		ParentJVM:       j,
		ContainingClass: class,
		Instructions:    make([]Instruction, instructionCount),
	}
	address = 0
	offsetMap := make(map[uint]int)
	// The second pass reads the instructions into the internal array, and
	// builds a map between instruction offsets -> indices for optimization.
	for i := 0; i < instructionCount; i++ {
		instruction, e = GetNextInstruction(codeMemory, address)
		if e != nil {
			return nil, fmt.Errorf("Error reading instruction: %s", e)
		}
		toReturn.Instructions[i] = instruction
		offsetMap[address] = i
		address += instruction.Length()
	}
	// The final pass performs the per-instruction optimization.
	address = 0
	for i := 0; i < instructionCount; i++ {
		instruction = toReturn.Instructions[i]
		e = instruction.Optimize(&toReturn, address, offsetMap)
		if e != nil {
			return nil, fmt.Errorf("Error in optimization pass over %s: %s",
				instruction, e)
		}
	}
	return &toReturn, nil
}

// Holds a loaded JVM class.
type Class struct {
	Methods map[string]*Method
	File    *class_file.Class
}

// Returns the named method from the class. Returns a MethodNotFoundError if
// the method isn't found.
func (c *Class) GetMethod(name string) (*Method, error) {
	toReturn := c.Methods[name]
	if toReturn == nil {
		return nil, MethodNotFoundError(name)
	}
	return toReturn, nil
}

// Takes a class loaded by the class_file package and converts it to the Class
// type needed by the JVM. Does *not* modify the state of the JVM.
func (j *JVM) NewClass(class *class_file.Class) (*Class, error) {
	toReturn := Class{
		Methods: make(map[string]*Method),
		File:    class,
	}
	var methodName []byte
	var method *Method
	var e error
	for i := range class.Methods {
		methodName = class.Methods[i].Name
		method, e = j.NewMethod(&toReturn, i)
		if e != nil {
			return nil, fmt.Errorf("Failed loading method %s: %s", methodName,
				e)
		}
		toReturn.Methods[string(methodName)] = method
	}
	return &toReturn, nil
}

// Adds the given class file to the JVM so that its code
func (j *JVM) LoadClass(class *class_file.Class) error {
	name, e := class.GetName()
	if e != nil {
		return fmt.Errorf("Failed getting class name: %s", e)
	}
	loadedClass, e := j.NewClass(class)
	if e != nil {
		return fmt.Errorf("Error loading class %s: %s", name, e)
	}
	j.Classes[string(name)] = loadedClass
	return nil
}

// Returns a reference to the named class. Returns a ClassNotFoundError if the
// class hasn't been loaded.
func (j *JVM) GetClass(name string) (*Class, error) {
	toReturn := j.Classes[name]
	if toReturn == nil {
		return nil, ClassNotFoundError(name)
	}
	return toReturn, nil
}

// Wraps jvm.GetClass and class.GetMethod into a single function.
func (j *JVM) GetMethod(className, methodName string) (*Method, error) {
	class, e := j.GetClass(className)
	if e != nil {
		return nil, e
	}
	return class.GetMethod(methodName)
}

// Shorthand for acquiring the lock on the list of active threads.
func (j *JVM) lockThreadList() {
	(&(j.threadsLock)).Lock()
}

// Shorthand for releasing the lock on the list of active threads.
func (j *JVM) unlockThreadList() {
	(&(j.threadsLock)).Unlock()
}

// Spawns a new thread in the JVM, with the given method.
func (j *JVM) StartThread(className, methodName string) error {
	method, e := j.GetMethod(className, methodName)
	if e != nil {
		return e
	}
	j.lockThreadList()
	threadIndex := len(j.threads)
	newThread := Thread{
		CurrentMethod:    method,
		ParentJVM:        j,
		InstructionIndex: 0,
		Stack:            NewDataStack(4096 * 4),
		ReturnLocations:  NewCallStack(1024),
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

// Takes a path to a class file, parses and loads the class, then looks for the
// main function in the class and starts executing it.
func (j *JVM) StartMainClass(classFileName string) error {
	file, e := os.Open(classFileName)
	if e != nil {
		return fmt.Errorf("Failed opening class file: %s\n", e)
	}
	classFile, e := class_file.ParseClass(file)
	if e != nil {
		return e
	}
	className, e := classFile.GetName()
	if e != nil {
		return fmt.Errorf("Failed getting class name: %s\n", e)
	}
	e = j.LoadClass(classFile)
	if e != nil {
		return e
	}
	e = j.StartThread(string(className), "main")
	return e
}
