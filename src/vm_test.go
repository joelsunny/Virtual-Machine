package main

import (
	"fmt"
	"testing"
)

// TestStackPush checks the stack push functionality
func TestStackPush(t *testing.T) {
	v := NewVM()
	v.push(1)
	fmt.Println(v.cpu)
	if v.cpu.sp != MEMSIZE-2 {
		t.Errorf("Expected result to be %d but instead got %d!", MEMSIZE-2, v.cpu.sp)
	}
}
