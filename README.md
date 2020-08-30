# About
Toy implementation of a stack based virtual machine.
The ISA and VM architecture is inspired from [this](https://youtu.be/OjaAToVkoTw) talk by Terrence Parr
# Machine Architecture
```
                                +-----------------------------------------------+
                                |                    CPU                        |
                                |                                               |
   memory                       |                                               |
                                |     +------------+                            |
+------------+                  |     |  fetch     |             registers      |
|   code     |                  |     +-----+------+                            |
+------------+ <--------------> |           |                  +-----------+    |
|            |                  |           v                  |    sp     |    |
|            |  ^               |     +-----+------+           +-----------+    |
|            |  |               |     |  decode    |           |    fp     |    |
|            |  |               |     +-----+------+           +-----------+    |
+------------+  | stack         |           |                  |    ip     |    |
|--|stack|---|  |               |           v                  +-----------+    |
+------------+  +               |     +-----+------+                            |
                                |     |  execute   |                            |
                                |     +------------+                            |
                                |                                               |
                                |                                               |
                                +-----------------------------------------------+


```

# Instruction Set Architecture

opcode | instr  |  op1 |  op2 | description  |
---|---|---|---|---|
|0 |iadd   |   |   |   integer add|
| 1|isub  |   |   |   |
| 2|imul   |   |   |   |
| 3|ilt   |   |   | integer less than  |
| 4|ieq  |   |   |   |
| 5|br   | addr  |   | branch to addr  |
| 6|brt   | addr  |   | branch if true, addr relative to instruction pointer  |
| 7|brf   | addr  |   |   |
| 8|iconst   |value   |   |push integer const to stack   |
| 9|load   | addr  |   |  load local variable @`addr` relative to the frame pointer |
| 10|gload   |addr   |   |   |
| 11|store   |addr   |   | store stack top to `addr` in memory  |
| 12|gstore   | addr  |   |   |
| 13|print   |   |   |   |
| 14| call  | addr  | numArgs  |   |
| 15|iret   |   |   |   |
| 16|halt   |   |   |   |


### A test program

```
max: // program to find the max of two numbers
load 5
load 4
isub
iconst 0
ilt
brf 4
load 4 
iret
load 5 // assembler will translate occurrence of label with the address of this instruction
iret

iconst 19
iconst 13
call max 2
print
halt
```

since we don't have an assembler we will load the program into the memory of the vm directly as follows:

```
program := []int32{load, 5, load, 4, isub, iconst, 0, ilt, brf, 4, load, 4, iret, load, 5, iret, iconst, 19, iconst, 91, call, 0, 2, print, halt}
v := NewVM()
v.loadProgram(program, 16) // 16 is the address of the first instruction to be executed
```

# Callstack - A guide to function calls

## Call stack structure
|   |  stack |description   |
|---|------|---|
|   |   :   | previous contents/frame  |
|   |  args  |  arguments to the function |
|   |  nargs  | number of arguments  |
|   |  ip    |  return address |
|   |  fp  |  previous frame pointer |
|   |  local|  local variables to the function |

## What happens on a call instruction?

### Understanding the `frame`

A call instruction is usually of the form

`call addr numArgs`

The sequence of actions necessary to make a function call is handled by the `call` instruction. The call instruction instructs the cpu to push the number of arguments to the function and the return address. The arguments itself are expected to be pushed by the caller of the function. This is followed by setting the instruction pointer to the function body in code memory. Thus a call instruction is better understood as an advanced branch instruction.

### Why do we need a frame pointer?
Ref: https://softwareengineering.stackexchange.com/a/194341

### Keeping track of previous frame pointer
Push it with the new call frame, so that `ret` instruction can restore it

