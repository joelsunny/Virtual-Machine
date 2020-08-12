package main

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
	pop
	halt
)

func (v *VM) iadd() {

}
