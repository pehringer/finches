# Function Generator (fungen)

### Key Concepts

1) **Linear Genetic Programming (LGP)**: A variant of Genetic Programming where programs are represented as linear sequences of instructions, resembling machine code. These instructions are run sequentially, with each instruction modifying the program’s state. Programs evolve through genetic operators like mutation and crossover, which alter the instruction sequence to search for solutions and optimize performance.

2) **Conditionally Executed Instructions**: Instructions execute only if specific processor flags are active. These flags are set by previous operations. This technique, common in CPU architectures like x86 and ARM, avoids traditional branching by embedding control flow within instruction-level logic.

### LGP Control Flow

Some LGP systems include control flow through conditional branching or instruction skipping. While these approaches are functional, they can be fragile in the context of genetic evolution. Small mutations, such as changing or inserting a single instruction, can significantly disrupt the control flow. Because these methods often depend on precise instruction ordering, even a minor change can make the program non-functional. This tight dependency on structure makes traditional control flow difficult to evolve reliably.

A more stable alternative uses conditionally executed instructions. These instructions depend on processor flags, which are set by earlier operations. This approach is more resilient to small changes because the flow of execution depends on local conditions rather than explicit jumps. It allows conditional logic to emerge and evolve gradually without breaking the program. Since LGP already uses a linear, sequential structure, flag-based execution fits naturally and helps preserve the program’s behavior during mutation. This makes it possible to develop complex behaviors while maintaining robustness during evolution. It also aligns more closely with the goals of LGP by allowing useful control flow to form incrementally.

# Instruction Set Architecture

- [Harvard Architecture](https://en.wikipedia.org/wiki/Harvard_architecture)
- [Stack Machine](https://en.wikipedia.org/wiki/Stack_machine)
- Has [Processor Flags](https://en.wikipedia.org/wiki/Status_register):
  + **N** (Negative)
  + **Z** (Zero)

#### Instruction Format:

|[39-37]  |[36-33]  |[32]     |[31-0]   |
|---------|---------|---------|---------|
|CONDITION|OPERATION|PEEKFLAGS|IMMEDIATE|

#### CONDITION - Condition for Execution:

|CONDITION|Mnemonic|Pseudocode         |Description             |
|---------|--------|-------------------|------------------------|
|000      |        |IF(TRUE)           |Always                  |
|001      |LT      |IF(N)              |Less than               |
|010      |LE      |IF(N OR Z)         |Less than or equal to   |
|011      |EQ      |IF(Z)              |Equal to                |
|100      |NE      |IF(NOT Z)          |Not equal to            |
|101      |GE      |IF(NOT N)          |Greater than or equal to|
|110      |GT      |IF(NOT N AND NOT Z)|Greater than            |
|111      |NV      |IF(FALSE)          |Never                   |

#### OPERATION - Operation to Execute:

**NOTE**: Instructions are paired together so a least significant bit flip will translate to a similar or immediate version of the instruction.

**NOTE**: Instructions with the significant bit set (or a 4 letter mnemonic) will use the immediate field of the instruction.

|OPERATION|Mnemonic|Pseudocode                            |Description                      |
|---------|--------|--------------------------------------|---------------------------------|
|00000    |ADD     |PUSH(POP() + POP())                   |Floating point add               |
|00001    |ADDI    |PUSH(POP() + IMMEDIATE)               |Floating point add immediate     |
|00010    |SUB     |PUSH(POP() - (POP()))                 |Floating point subtract          |
|00011    |SUBI    |PUSH(POP() - IMMEDIATE)               |Floating point subtract immediate|
|00100    |MUL     |PUSH(POP() * POP())                   |Floating point multiply          |
|00101    |MULI    |PUSH(POP() * IMMEDIATE)               |Floating point multiply immediate|
|00110    |DIV     |PUSH(POP() / (POP()))                 |Floating point divide            |
|00111    |DIVI    |PUSH(POP() / IMMEDIATE)               |Floating point divide immediate  |
|01000    |MAX     |PUSH(MAX(POP(), POP()))               |Floating point maximum           | 
|01001    |MAXI    |PUSH(MAX(POP(), IMMEDIATE))           |Floating point maximum immediate |
|01010    |MIN     |PUSH(MIN(POP(), POP()))               |Floating point minimum           |
|01011    |MINI    |PUSH(MIN(POP(), IMMEDIATE))           |Floating point minimum immediate |
|01100    |POP     |POP()                                 |Pop                              |
|01101    |PUSH    |PUSH(IMMEDIATE)                       |Push immediate                   |
|01110    |SWP     |X = POP(); Y = POP(); PUSH(X); PUSH(Y)|Swaps top two items              |
|01111    |PICK    |PUSH(STACK[TOP - IMMEDIATE])          |Push n-th item from the top      |

#### PEEKFLAGS - Peek And Set Flags:

|PEEKFLAGS|Mnemonic|Pseudocode                     |Description           |
|---------|--------|-------------------------------|----------------------|
|0        |        |N = N; Z = Z                   |Do not set flags      |
|1        |PF      |N = PEEK() < 0; Z = PEEK() == 0|Peek top and set flags|

#### IMMEDIATE -  Immediate Value:

Used by some instructions, unused by others.
