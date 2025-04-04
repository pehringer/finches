# Linear Genetic Programming

### Key Concepts

1) ***Linear Genetic Programming (LGP)***: A type of Genetic Programming (GP) where programs are represented as linear sequences of instructions, similar to machine code. LGP operates on a register-based architecture where instructions modify values stored in registers. Programs execute sequentially, and their evolution is driven by genetic operators such as mutation and crossover, which modify the instruction sequence to optimize performance.

2) ***Conditionally Executed Instructions***: Instructions that perform their operation only if a specified condition, based on processor flags, is met. These flags may be optionally set by the current or prior instructions. This approach is common in many CPU architectures, such as x86 and ARM.

### Current LGP Control Flow Model

Some LGP implementations already incorporate control flow, typically through conditional branching or skipping. These mechanisms allow the program to alter its flow based on certain conditions. While functional, these approaches introduce challenges in the context of genetic algorithms.

The primary issue is that small mutations, such as changing an existing instruction or inserting a new one, can drastically alter the control flow. This can lead to the program becoming non-functional, as conditional branching or skipping often relies on a tight instruction ordering to function correctly. This makes such control flow mechanisms poorly suited for the evolution of programs through genetic algorithms, where incremental modifications are common.

### An Alternative LGP Control Flow Model

An alternative approach to control flow in LGP involves using conditionally executed instructions. This type of control flow has a number of advantages over traditional branching or skipping.

Primarily, conditionally executed instructions are more adaptable to evolutionary modifications because these operations rely on a looser, linear, flag-dependent instruction ordering. They are more resilient to small mutations, such as changing an existing instruction or inserting a new one. This results in a more stable evolutionary process, where conditional logic can evolve incrementally without disturbing the integrity of the program’s execution.

Additionally, the flag-setting and flag-dependent execution introduces a linear, order-based dependency that fits naturally with LGP’s linear, order-dependent nature.

# Instruction Set Architecture

|[31 - 27]  |[26]|[25]|[24 - 22]|[21 - 17]|[16 - 12]    |[11 - 0]      |
|-----------|----|----|---------|---------|-------------|--------------|
|Operation  |Type|Set |Condition|Result   |First Operand|Second Operand|

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
