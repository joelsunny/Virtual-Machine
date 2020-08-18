package main

import "fmt"

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

func (v *VM) loadProgram(program []Word, entry int32) {
	for i := 0; i < len(program); i++ {
		v.mem[i] = program[i]
	}
	v.cpu.ip = entry
}

func (v *VM) run() {
	v.cpu.sp = MEMSIZE - 1
	v.cpu.fp = v.cpu.sp

	for {
		if v.step() {
			break
		}
	}
}

// stack functions
func (v *VM) pop() Word {
	if v.cpu.sp < MEMSIZE-1 {
		v.cpu.sp++
	}
	return v.mem[v.cpu.sp]
}

func (v *VM) push(w Word) {
	v.mem[v.cpu.sp] = w
	v.cpu.sp--
	// check for stack overflow
	// todo
}

func (v *VM) printStack() string {
	return fmt.Sprintf("%v", v.mem[v.cpu.sp+1:MEMSIZE])
}

// Implementation of the fetch -> decode -> execution cycle
func (v *VM) step() bool {
	// fetch the next instruction
	instr := v.mem[v.cpu.ip]
	v.cpu.ip++

	// decode
	switch instr {
	case iadd:
		v.iadd()
	case isub:
		v.isub()
	case imul:
		v.imul()
	case ilt:
		v.ilt()
	case ieq:
		v.ieq()
	case br:
		v.br()
	case brt:
		v.brt()
	case brf:
		v.brf()
	case iconst:
		v.iconst()
	case load:
		v.load()
	case gload:
		v.gload()
	case store:
		v.store()
	case gstore:
		v.gstore()
	case print:
		v.print()
	case call:
		v.call()
	case iret:
		v.iret()
	case halt:
		return true
	}
	// println(instr, v.printStack())
	return false
}
