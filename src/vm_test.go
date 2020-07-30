package main

import (
	"fmt"
	"testing"
)

// TestAdd add instruction
func TestRun(t *testing.T) {
	v := VM{}
	p := []Instruction{Instruction{Opcode: ld, A0: R0, A1: 103}}
	v.Run(p)
	fmt.Println(v)
	if v.R[R0] != 103 {
		t.Errorf("Expected result to be %d but instead got %d!", 103, v.R[R0])
	}
}
