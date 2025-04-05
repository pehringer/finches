# Linear Genetic Programming

### Key Concepts

1) ***Linear Genetic Programming (LGP)***: A variant of Genetic Programming where programs are represented as linear sequences of instructions, resembling machine code. LGP operates on a register-based architecture, with each instruction modifying register values. Programs execute sequentially, and evolve through genetic operators like mutation and crossover, which alter the instruction sequence to optimize performance.

2) ***Conditionally Executed Instructions***: Instructions execute only if specific processor flags are active. These flags are set by previous operations. This technique, common in CPU architectures like x86 and ARM, avoids traditional branching by embedding control flow within instruction-level logic.

### LGP Control Flow

Some LGP systems include control flow through conditional branching or instruction skipping. While these approaches are functional, they can be fragile in the context of genetic evolution. Small mutations, such as changing or inserting a single instruction, can significantly disrupt the control flow. Because these methods often depend on precise instruction ordering, even a minor change can make the program non-functional. This tight dependency on structure makes traditional control flow difficult to evolve reliably.

A more stable alternative uses conditionally executed instructions. These instructions depend on processor flags, which are set by earlier operations. This approach is more resilient to small changes because the flow of execution depends on local conditions rather than explicit jumps. It allows conditional logic to emerge and evolve gradually without breaking the program. Since LGP already uses a linear, sequential structure, flag-based execution fits naturally and helps preserve the program’s behavior during mutation. This makes it possible to develop complex behaviors while maintaining robustness during evolution. It also aligns more closely with the goals of LGP by allowing useful control flow to form incrementally.

# Instruction Set Architecture

|[31-27]  |[26]|[25]|[24-22]  |[21-17]|[16-12]      |[11-0]        |
|---------|----|----|---------|-------|-------------|--------------|
|Operation|Type|Set |Condition|Result |First Operand|Second Operand|

***Operation***: The operation to be performed.

|     |Name          |Description                      |Details                             |
|-----|--------------|---------------------------------|------------------------------------|
|00000|ADD           |Signed integer addition.         |Signed immediate value.             |
|00001|SUB           |Signed integer subtraction.      |Signed immediate value.             |
|00010|MUL           |Signed integer multiplication.   |Signed immediate value.             |
|00011|DIV           |Signed integer division.         |Signed immediate value.             |
|00100|MOD           |Signed integer modulo.           |Signed immediate value.             |
|00101|MAX           |Signed integer maximum.          |Signed immediate value.             |
|00110|MIN           |Signed integer minimum.          |Signed immediate value.             |
|00111|AND           |Bitwise AND.                     |Unsigned immediate value.           |
|01000|NAN           |Bitwise NAND.                    |Unsigned immediate value.           |
|01001|EOR           |Bitwise exclusive OR (XOR).      |Unsigned immediate value.           |
|01010|IOR           |Bitwise inclusive OR.            |Unsigned immediate value.           |
|01011|NOR           |Bitwise NOR.                     |Unsigned immediate value.           |
|01100|LSH           |Bitwise logical left shift.      |Unsigned immediate value.           |
|01101|RSH           |Bitwise logical right shift.     |Unsigned immediate value.           |
|01110|NOP           |No operation.                    |                                    |
|01111|NOP           |No operation.                    |                                    |
|10000|NOP           |No operation.                    |                                    |
|...  |              |                                 |                                    |
|11101|NOP           |No operation.                    |                                    |
|11110|NOP           |No operation.                    |                                    |
|11111|NOP           |No operation.                    |                                    |

***Type***: The ***Second Operand*** type: ```0``` register number, ```1``` immediate value.

***Set***: Set processor flags if the instruction was executed: ```0``` no, ```1``` yes.

|Name|Description                              |
|----|-----------------------------------------|
|N   |Set if the operation result was negative.|
|O   |Set if the operation result overflowed.  |
|Z   |Set if the operation result was zero.    |

***Condition Code***: The conditions under which the instruction executes.

|   |Name          |Description                               |Details               |
|---|--------------|------------------------------------------|----------------------|
|000|AL            |Always.                                   |Unconditional.        |
|001|LT            |Less than.                                |N flag set.           |
|010|LE            |Less than or equal to.                    |N or Z flags set.     |
|011|EQ            |Equal to.                                 |Z flag set.           |
|100|NE            |Not equal to.                             |Z flag not set.       |
|101|GE            |Greater than or equal to.                 |N flag not set.       |
|110|GT            |Greater than.                             |N and Z flags not set.|
|111|OF            |Overflow.                                 |O flag set.           |

***Result***: The register number for the result.

***First Operand***: The register number for the first operand.

***Second Operand***: Either the register number [5 - 0] or the signed immediate value [11 - 0] for the second operand.

|     |Name|Description              |
|-----|----|-------------------------|
|00000|R0  |General purpose register.|
|00001|R1  |General purpose register.|
|00010|R2  |General purpose register.|
|...  |    |                         |
|11101|R29 |General purpose register.|
|11110|R30 |General purpose register.|
|11111|R31 |General purpose register.|
