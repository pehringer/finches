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

|[31]     |[30-28]  |[27-24]  |[23-16]    |[15-8] |[7-0]  |
|---------|---------|---------|-----------|-------|-------|
|SETFLAGS||CONDITION|OPERATION|DESTINATION|SOURCE1|SOURCE2|

#### SETFLAGS - Set flags based on the operations result:

|PEEKFLAGS|Mnemonic|Pseudocode                                     |Description|
|---------|--------|-----------------------------------------------|-----------|
|0        |        |N = N; Z = Z                                   |No         |
|1        |S	   |N = R[DESTINATION] < 0; Z = R[DESTINATION] == 0|Set Flags  |

#### CONDITION - Condition for execution:

|CONDITION|Mnemonic|Pseudocode        |Description       |
|---------|--------|------------------|------------------|
|000      |        |if True           |Always            |
|001      |LT      |if N              |Less Than         |
|010      |LE      |if N or Z         |Less Than Equal   |
|011      |EQ      |if Z              |Equal             |
|100      |NE      |if not Z          |Not Equal         |
|101      |GE      |if not N          |Greater Than Equal|
|110      |GT      |if not N and not z|Greater Than      |
|111      |NOP     |if false          |Never             |

#### OPERATION - Operation to execute:

|OPERATION|Mnemonic|Pseudocode                              |Description                       |
|------- -|--------|----------------------------------------|----------------------------------|
|0000     |ADD     |R[DESTINATION] = R[SOURCE1] + R[SOURCE2]|Floating Point Addition           |
|0001     |SUB     |R[DESTINATION] = R[SOURCE1] - R[SOURCE2]|Floating Point Subtraction        |
|0010     |MUL     |R[DESTINATION] = R[SOURCE1] * R[SOURCE2]|Floating Point Multiplication     |
|0011     |DIV     |R[DESTINATION] = R[SOURCE1] / R[SOURCE2]|Floating Point Protected Division |
|0100     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|0101     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|0110     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|0111     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1000     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1001     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1010     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1011     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1100     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1101     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1110     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |
|1111     |MOV     |R[DESTINATION] = R[SOURCE1]             |Move                              |

### DESTINATION - Destination register for the result.

### SOURCE1 - Source register of the first operand.

### SOURCE2 - Source register of the second operand.


