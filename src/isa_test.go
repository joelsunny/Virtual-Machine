package main

import (
	"fmt"
	"testing"
)

// TestAdd add instruction
func TestAdd(t *testing.T) {
	v := VM{}
	v.ld(R0, 1)
	v.ld(R1, 2)
	v.add(R0, R1, R2)
	fmt.Println(v)
	if v.R[R2] != 3 {
		t.Errorf("Expected result to be %d but instead got %d!", 3, v.R[R2])
	}
}
