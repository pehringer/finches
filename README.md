# Function Generator (fungen)

### Key Concepts

1) **Linear Genetic Programming (LGP)**: A variant of Genetic Programming where programs are represented as linear sequences of instructions, resembling machine code. These instructions are run sequentially, with each instruction modifying the program’s state. Programs evolve through genetic operators like mutation and crossover, which alter the instruction sequence to search for solutions and optimize performance.

2) **Conditionally Executed Instructions**: Instructions execute only if specific processor flags are active. These flags are set by previous operations. This technique, common in CPU architectures like x86 and ARM, avoids traditional branching by embedding control flow within instruction-level logic.

### LGP Control Flow

Some LGP systems include control flow through conditional branching or instruction skipping. While these approaches are functional, they can be fragile in the context of genetic evolution. Small mutations, such as changing or inserting a single instruction, can significantly disrupt the control flow. Because these methods often depend on precise instruction ordering, even a minor change can make the program non-functional. This tight dependency on structure makes traditional control flow difficult to evolve reliably.

A more stable alternative uses conditionally executed instructions. These instructions depend on processor flags, which are set by earlier operations. This approach is more resilient to small changes because the flow of execution depends on local conditions rather than explicit jumps. It allows conditional logic to emerge and evolve gradually without breaking the program. Since LGP already uses a linear, sequential structure, flag-based execution fits naturally and helps preserve the program’s behavior during mutation. This makes it possible to develop complex behaviors while maintaining robustness during evolution. It also aligns more closely with the goals of LGP by allowing useful control flow to form incrementally.

# Instruction Set Architecture

- Flags:
  + **N** (Negative)
  + **Z** (Zero)
- Registers:
  + **R[0-255]**
  + All immediate values are preloaded into registers before execution begins.

#### Instruction format:

|[31-30]  |[29-25]  |[24]    |[23-16]    |[15-8] |[7-0]  |
|---------|---------|--------|-----------|-------|-------|
|CONDITION|OPERATION|SETFLAGS|DESTINATION|SOURCE1|SOURCE2|

#### CONDITION - Condition for execution:

|CONDITION|Mnemonic|Pseudocode        |Description |
|---------|--------|------------------|------------|
|00	  |        |                  |Always	   |
|01	  |LT	   |if N              |Less Than   |
|10	  |GT	   |if not N and not Z|Greater Than|
|11	  |EQ	   |if Z              |Equal To    |

#### OPERATION - Operation to execute:

|OPERATION|Mnemonic|Pseudocode                                      |Description                       |
|---------|--------|------------------------------------------------|----------------------------------|
|00000    |ADD     |R[DESTINATION] = R[SOURCE1] + R[SOURCE2]        |Floating Point Addition           |
|00001    |SUB     |R[DESTINATION] = R[SOURCE1] - R[SOURCE2]        |Floating Point Subtraction        |
|00010    |MUL     |R[DESTINATION] = R[SOURCE1] * R[SOURCE2]        |Floating Point Multiplication     |
|00011    |DIV     |R[DESTINATION] = R[SOURCE1] / R[SOURCE2]        |Floating Point Protected Division |
|00100    |MAX     |R[DESTINATION] = maximum(R[SOURCE1], R[SOURCE2])|Floating Point Maximum            |
|00101    |MIN     |R[DESTINATION] = minimum(R[SOURCE1], R[SOURCE2])|Floating Point Minimum            |
|00110    |ABD     |R[DESTINATION] = |R[SOURCE1] - R[SOURCE2]|      |Floating Point Absolute Difference|
|00111    |AVG     |R[DESTINATION] = (R[SOURCE1] + R[SOURCE2]) / 2  |Floating Point Average            |
|01000    |NOP     |                                                |No Operation                      |
|01001    |NOP     |                                                |No Operation                      |
|01010    |NOP     |                                                |No Operation                      |
|01011    |NOP     |                                                |No Operation                      |
|01100    |NOP     |                                                |No Operation                      |
|01101    |NOP     |                                                |No Operation                      |
|01110    |NOP     |                                                |No Operation                      |
|01111    |NOP     |                                                |No Operation                      |
|10000    |NOP     |                                                |No Operation                      |
|10001    |NOP     |                                                |No Operation                      |
|10010    |NOP     |                                                |No Operation                      |
|10011    |NOP     |                                                |No Operation                      |
|10100    |NOP     |                                                |No Operation                      |
|10101    |NOP     |                                                |No Operation                      |
|10110    |NOP     |                                                |No Operation                      |
|10111    |NOP     |                                                |No Operation                      |
|11000    |NOP     |                                                |No Operation                      |
|11001    |NOP     |                                                |No Operation                      |
|11010    |NOP     |                                                |No Operation                      |
|11011    |NOP     |                                                |No Operation                      |
|11100    |NOP     |                                                |No Operation                      |
|11101    |NOP     |                                                |No Operation                      |
|11110    |NOP     |                                                |No Operation                      |
|11111    |NOP     |                                                |No Operation                      |

#### SETFLAGS - Set flags based on the operations result:

**NOTE**: NOP (no operation) cannot set flags since it does not produce a result.

|PEEKFLAGS|Mnemonic|Pseudocode                                     |Description           |
|---------|--------|-----------------------------------------------|----------------------|
|0        |        |                                               |Do not set flags	  |
|1        |S	   |N = R[DESTINATION] < 0; Z = R[DESTINATION] == 0|Peek top and set flags|

### DESTINATION - Destination register for the result.

### SOURCE1 - Source register of the first operand.

### SOURCE2 - Source register of the second operand.

