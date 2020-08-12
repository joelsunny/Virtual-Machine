package main

// MEMSIZE defines the size of the vm memory in words
const MEMSIZE = 1024

// Register 32 bit
type Register = int32

// Word is 32 bits
type Word = int32

// CPU definition
type CPU struct {
	fp Register // frame pointer
	ip Register // instruction pointer
	sp Register // stack pointer, stack grows upwards
	// todo: heap
}

// Memory 1024 words for now
type Memory [MEMSIZE]Word

// VM represents the structure of the virtual machine
type VM struct {
	cpu CPU
	mem Memory
}

// NewVM  instantiate the virtual machine
func NewVM() VM {
	cpu := CPU{0, 0, MEMSIZE - 1}
	return VM{cpu: cpu}
}

// stack functions
func (v *VM) pop() {
	if v.cpu.sp < MEMSIZE-1 {
		v.cpu.sp++
	}
}

func (v *VM) push(w Word) {
	v.mem[v.cpu.sp] = w
	v.cpu.sp--
	// check for stack overflow
	// todo
}
