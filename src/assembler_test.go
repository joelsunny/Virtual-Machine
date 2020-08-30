package main

import (
	"fmt"
	"testing"
)

// TestAssembler implementation
func TestAssembler(t *testing.T) {
	path := "./tests/code.asm"
	p := assembler(path)

	fmt.Println(p)
	v := NewVM()
	v.loadProgram(p.prog, p.entry)
	v.run()
}
