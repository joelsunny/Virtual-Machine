# Callstack - A guide to function calls

## Call stack structure
|   |  stack |description   |
|---|------|---|
|   |   :   | previous contents/frame  |
|   |  args  |  arguments to the function |
|   |  nargs  | number of argument  |
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

