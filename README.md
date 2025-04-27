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
  + **R[0-63]**
  + All immediate values are preloaded into registers before execution begins.

#### Instruction format:

|[31-24] |[23-22]  |[21-19]  |[18]    |[17-12]    |[11-6] |[5-0]  |
|--------|---------|---------|--------|-----------|-------|-------|
|RESERVED|CONDITION|OPERATION|SETFLAGS|DESTINATION|SOURCE1|SOURCE2|

#### CONDITION - Condition for execution:

|CONDITION|Mnemonic|Pseudocode        |Description |
|---------|--------|------------------|------------|
|00	  |        |                  |Always	   |
|01	  |LT	   |if N              |Less Than   |
|10	  |GT	   |if not N and not Z|Greater Than|
|11	  |EQ	   |if Z              |Equal To    |

#### OPERATION - Operation to execute:

|OPERATION|Mnemonic|Pseudocode                              |Description                       |
|---------|--------|----------------------------------------|----------------------------------|
|00000    |ADD     |R[DESTINATION] = R[SOURCE1] + R[SOURCE2]|Floating Point Addition           |
|00001    |SUB     |R[DESTINATION] = R[SOURCE1] - R[SOURCE2]|Floating Point Subtraction        |
|00010    |MUL     |R[DESTINATION] = R[SOURCE1] * R[SOURCE2]|Floating Point Multiplication     |
|00011    |DIV     |R[DESTINATION] = R[SOURCE1] / R[SOURCE2]|Floating Point Protected Division |
|00100    |NOP     |                                        |No Operation                      |
|00101    |NOP     |                                        |No Operation                      |
|00110    |NOP     |                                        |No Operation                      |
|00111    |NOP     |                                        |No Operation                      |

#### SETFLAGS - Set flags based on the operations result:

**NOTE**: NOP (no operation) cannot set flags since it does not produce a result.

|PEEKFLAGS|Mnemonic|Pseudocode                                     |Description|
|---------|--------|-----------------------------------------------|-----------|
|0        |        |                                               |No         |
|1        |S	   |N = R[DESTINATION] < 0; Z = R[DESTINATION] == 0|Set Flags  |

### DESTINATION - Destination register for the result.

### SOURCE1 - Source register of the first operand.

### SOURCE2 - Source register of the second operand.


