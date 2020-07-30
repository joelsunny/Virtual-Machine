package main

/*
** Lets start by defining the basic data structures
 */

// Register 32 bit
type Register = int

const (
	R0 = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

// VM represents the structure of the virtual machine
type VM struct {
	R  [16]Register
	SP *Register
	PC int
}

// Instruction , single instruction
type Instruction struct {
	Opcode Register
	A0     Register
	A1     Register
	A2     Register
}

// Program , a program to be executed on the VM is represented as a slice of Instructions
type Program []Instruction

// Run program
func (v *VM) Run(p Program) {
	v.PC = 0
	for i := 0; i < len(p); i++ {
		inst := p[i]
		v.step(inst)
	}
}
