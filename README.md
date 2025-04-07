# Linear Genetic Programming

### Key Concepts

1) **Linear Genetic Programming (LGP)**: A variant of Genetic Programming where programs are represented as linear sequences of instructions, resembling machine code. LGP operates on a register-based architecture, with each instruction modifying register values. Programs execute sequentially, and evolve through genetic operators like mutation and crossover, which alter the instruction sequence to optimize performance.

2) **Conditionally Executed Instructions**: Instructions execute only if specific processor flags are active. These flags are set by previous operations. This technique, common in CPU architectures like x86 and ARM, avoids traditional branching by embedding control flow within instruction-level logic.

### LGP Control Flow

Some LGP systems include control flow through conditional branching or instruction skipping. While these approaches are functional, they can be fragile in the context of genetic evolution. Small mutations, such as changing or inserting a single instruction, can significantly disrupt the control flow. Because these methods often depend on precise instruction ordering, even a minor change can make the program non-functional. This tight dependency on structure makes traditional control flow difficult to evolve reliably.

A more stable alternative uses conditionally executed instructions. These instructions depend on processor flags, which are set by earlier operations. This approach is more resilient to small changes because the flow of execution depends on local conditions rather than explicit jumps. It allows conditional logic to emerge and evolve gradually without breaking the program. Since LGP already uses a linear, sequential structure, flag-based execution fits naturally and helps preserve the program’s behavior during mutation. This makes it possible to develop complex behaviors while maintaining robustness during evolution. It also aligns more closely with the goals of LGP by allowing useful control flow to form incrementally.

# Instruction Set Architecture

#### Processor Flags:

|Flag|Description        |
|----|-------------------|
|N   |Result was negative|
|O   |Result overflowed  |
|Z   |Result was zero    |

#### Processor Registers:

|Register|Description    |
|--------|---------------|
|00000   |General purpose|
|...     |               |
|11111   |General purpose|

#### Instruction Format:

|[31:27]  |[26]     |[25]    |[24:22]  |[21:17]    |[16:12]|[11:0] |
|---------|---------|--------|---------|-----------|-------|-------|
|Operation|Immediate|SetFlags|Condition|Destination|Source1|Source2|

#### Operation - Operation to Execute:

|Operation|Description                  |Details                 |
|---------|-----------------------------|------------------------|
|00000    |Signed integer addition      |Signed immediate value  |
|00001    |Signed integer subtraction   |Signed immediate value  |
|00010    |Signed integer multiplication|Signed immediate value  |
|00011    |Signed integer division      |Signed immediate value  |
|00100    |Signed integer modulo        |Signed immediate value  |
|00101    |Signed integer maximum       |Signed immediate value  |
|00110    |Signed integer minimum       |Signed immediate value  |
|00111    |Bitwise AND                  |Unsigned immediate value|
|01000    |Bitwise NAND                 |Unsigned immediate value|
|01001    |Bitwise exclusive OR (XOR)   |Unsigned immediate value|
|01010    |Bitwise inclusive OR         |Unsigned immediate value|
|01011    |Bitwise NOR                  |Unsigned immediate value|
|01100    |Bitwise logical left shift   |Unsigned immediate value|
|01101    |Bitwise logical right shift  |Unsigned immediate value|
|01110    |No operation                 |                        |
|...      |                             |                        |
|11111    |No operation                 |                        |

#### Immediate - Source2 Immediate Value:

|Immediate|Description                  |
|---------|-----------------------------|
|0        |Source2 is a register        |
|1        |Source2 is an immediate value|

#### SetFlags - Set Processor Flags:

|SetFlags|Description              |
|--------|-------------------------|
|0       |Do not set flags         |
|1       |Set flags after operation|

#### Condition - Condition for Execution:

|Condition|Description             |Details              |
|---------|------------------------|---------------------|
|000      |Always                  |Unconditional        |
|001      |Less than               |N flag set           |
|010      |Less than or equal to   |N or Z flags set     |
|011      |Equal to                |Z flag set           |
|100      |Not equal to            |Z flag not set       |
|101      |Greater than or equal to|N flag not set       |
|110      |Greater than            |N and Z flags not set|
|111      |Overflow                |O flag set           |

#### Destination - Destination register.

#### Source1 - First operand register.

#### Source2 - Second operand register ([4:0]) or immediate value ([11:0]) based on Immediate field.
