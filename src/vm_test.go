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

func TestStackPop(t *testing.T) {
	v := NewVM()
	v.push(19)
	fmt.Println(v.cpu)
	val := v.pop()
	fmt.Println(val)
	if (v.cpu.sp != MEMSIZE-1) && (val != 19) {
		t.Errorf("vals: %d %d", val, v.cpu.sp)
	}
}

func TestProgram(t *testing.T) {
	program := []int32{load, 5, load, 4, isub, iconst, 0, ilt, brf, 4, load, 4, iret, load, 5, iret, iconst, 19, iconst, 91, call, 0, 2, print, halt}
	v := NewVM()
	v.loadProgram(program, 16)
	v.run()
}
