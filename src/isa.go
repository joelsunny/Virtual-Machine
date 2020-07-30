package main

/*
Instruction Set
*/
const (
	ld = Register(iota)
	add
)

func (v *VM) step(instr Instruction) {
	op := instr.Opcode
	switch op {
	case ld:
		v.ld(instr.A0, instr.A1)
	}
}

// Load : ld Reg Value
func (v *VM) ld(r int, val Register) {
	v.R[r] = val
}

func (v *VM) add(r0 int, r1 int, r2 int) {
	v.R[r2] = v.R[r0] + v.R[r1]
}
