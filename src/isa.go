package main

import "fmt"

// OpCode s are represented as 32 bit in memory
type OpCode = int32

const (
	iadd OpCode = iota
	isub
	imul
	ilt
	ieq
	br
	brt
	brf
	iconst
	load
	gload
	store
	gstore
	print
	call
	iret
	halt
)

func (v *VM) iadd() {
	a := v.pop()
	b := v.pop()
	v.push(a + b) // overflow?
}

func (v *VM) isub() {
	a := v.pop()
	b := v.pop()
	v.push(b - a)
}

func (v *VM) imul() {
	a := v.pop()
	b := v.pop()
	v.push(a * b) // overflow?
}

func (v *VM) ilt() {
	a := v.pop()
	b := v.pop()
	if b < a {
		v.push(1)
		return
	}
	v.push(0) // overflow?
}

func (v *VM) ieq() {
	a := v.pop()
	b := v.pop()
	if a == b {
		v.push(1)
		return
	}
	v.push(0) // overflow?
}

func (v *VM) br() {
	addr := v.mem[v.cpu.ip]
	v.cpu.ip = addr
}

func (v *VM) brt() {
	// pop the stack, check if true, then branch
	c := v.pop()
	if c == 1 {
		addr := v.mem[v.cpu.ip]
		v.cpu.ip = addr
		return
	}
	v.cpu.ip++
}

func (v *VM) brf() {
	// make it relative to the instruction pointer for ease of writing functions
	c := v.pop()
	if c == 0 {
		addr := v.mem[v.cpu.ip]
		v.cpu.ip += addr // relative jump
		return
	}
	v.cpu.ip++
}

func (v *VM) iconst() {
	val := v.mem[v.cpu.ip]
	v.cpu.ip++
	v.push(val)
}

func (v *VM) load() {
	// calculate the absolute address by adding the offset to the frame pointer.
	addr := v.cpu.fp + v.mem[v.cpu.ip]
	v.cpu.ip++
	v.push(v.mem[addr])
}

func (v *VM) gload() {
	addr := v.mem[v.cpu.ip]
	v.cpu.ip++
	v.push(v.mem[addr])
}

func (v *VM) store() {
	val := v.pop()
	addr := v.mem[v.cpu.ip]
	v.cpu.ip++
	v.mem[addr] = val
}

func (v *VM) gstore() {

}

func (v *VM) print() {
	fmt.Println(v.pop())
}

func (v *VM) call() {
	// view call as an advanced branch instruction
	// push the args to stack and return control to the
	// first instruction in the function body
	addr := v.mem[v.cpu.ip]
	v.cpu.ip++
	numArgs := v.mem[v.cpu.ip]
	v.cpu.ip++

	// store the number of arguments, this is used on return to pop the arguments off the stack
	v.push(numArgs)
	// store the ip before function call in the stack, i.e the return address
	v.push(v.cpu.ip)
	// store the previous frame pointer
	v.push(v.cpu.fp)
	// set the current frame pointer value to the top of the stack
	v.cpu.fp = v.cpu.sp
	v.cpu.ip = addr
}

func (v *VM) iret() {
	// return instruction should:
	// 1. restore the frame pointer
	// 2. jump back to the line after the function call
	// 3. iret -> return integer

	retVal := v.pop()
	v.cpu.sp = v.cpu.fp
	v.cpu.fp = v.pop() //  restore previous frame
	v.cpu.ip = v.pop() // return address
	nargs := v.pop()
	v.cpu.sp += nargs // pop the args off the stack
	v.push(retVal)
}

func (v *VM) halt() {

}
