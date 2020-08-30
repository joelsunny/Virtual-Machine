# Assembler design

Assembler is a program that takes in a file containing the assembly code and translates it to machine code. In addition to simply translating the instructions to their corresponding opcodes, an assembler also makes possible some higher level syntax such as:
1. `label` usage to mark instructions
2. the use of functions marked with `proc` and `end` keywords

Assembler makes use of a symbol table (a map between label definitions and the corresponding code positions) to keep track of labels and function names. It is the function of the assembler to find the correct entry point to the program. Assembler returns a struct of type `Program` defined as below.

```
type Program struct {
	prog  []Word
	entry Word
}
```