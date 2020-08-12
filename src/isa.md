# Instruction Set Architecture

opcode | instr  |  op1 |  op2 | description  |
---|---|---|---|---|
| |iadd   |   |   |   integer add|
| |isub  |   |   |   |
| |imul   |   |   |   |
| |ilt   |   |   | integer less than  |
| |ieq  |   |   |   |
| |br   | addr  |   | branch to addr  |
| |brt   | addr  |   | branch if true  |
| |brf   |   |   |   |
| |iconst   |value   |   |push integer const to stack   |
| |load   | addr  |   |  load local |
| |gload   |addr   |   |   |
| |store   |addr   |   |   |
| |gstore   | addr  |   |   |
| |print   |   |   |   |
| | pop  |   |   |pop stack top   |
| | call  | addr  | numArgs  |   |
| |ret   |   |   |   |
| |halt   |   |   |   |


# Reference

The ISA and VM architecture is inspired from [this](https://youtu.be/OjaAToVkoTw) talk by Terrence Parr