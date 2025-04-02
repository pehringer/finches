# ***clgp*** - ***C***onditional ***L***inear ***G***enetic ***P***rogramming

### Key Concepts

1) ***Linear Genetic Programming (LGP)***: A type of Genetic Programming (GP) where programs are represented as linear sequences of instructions, similar to machine code. LGP operates on a register-based architecture where instructions modify values stored in registers. Programs execute sequentially, and their evolution is driven by genetic operators such as mutation and crossover, which modify the instruction sequence to optimize performance.

2) ***Conditionally Executed Instructions***: Operations that execute only if a specified condition is met. Rather than altering the flow of execution with branches, these instructions check a condition and either perform their operation or do nothing. This approach is common in many CPU architectures, such as x86 and ARM, where some instructions can be conditionally executed based on processor flags set by comparison operations.

### Current LGP Conditional Execution Model

Some LGP implementations already incorporate conditional execution, typically through conditional branching or skipping. These mechanisms allow the program to alter its flow based on certain conditions. While functional, these approaches introduce challenges in the context of genetic algorithms.

The primary issue is that small mutations, such as changing an existing instruction or inserting a new one, can drastically alter the control flow. This can lead to the program becoming non-functional, as conditional branching or skipping often relies on a tight instruction ordering to function correctly. This makes such conditional execution mechanisms poorly suited for the evolution of programs through genetic algorithms, where incremental modifications are common.

### An Alternative LGP Conditional Execution Model

An alternative approach to conditional execution in LGP involves using instructions that are conditionally executed based on flags set by comparison operations. This type of execution has a number of advantages over traditional branching or skipping.

Primarily, conditionally executed instructions are more adaptable to evolutionary modifications because these operations rely on a looser, linear, flag-dependent instruction ordering. They are more resilient to small mutations, such as changing an existing instruction or inserting a new one. This results in a more stable evolutionary process, where conditional logic can evolve incrementally without disturbing the integrity of the program’s execution.

Additionally, the flag-setting and flag-dependent execution introduces a linear, order-based dependency that fits naturally with LGP’s linear, order-dependent nature.

# CLGP Instruction Set Architecture

### Processor Flags

### Processor Registers

|Code       |Name|Description|
|-----------|----|-----------|
|```00000```|R0  |           |
|```00001```|R1  |           |
|```00010```|R2  |           |
|```00011```|R3  |           |
|```00100```|R4  |           |
|```00101```|R5  |           |
|```00110```|R6  |           |
|```00111```|R7  |           |
|```01000```|R8  |           |
|```01001```|R9  |           |
|```01010```|R10 |           |
|```01011```|R11 |           |
|```01100```|R12 |           |
|```01101```|R13 |           |
|```01110```|R14 |           |
|```01111```|R15 |           |
|```10000```|R16 |           |
|```10001```|R17 |           |
|```10010```|R18 |           |
|```10011```|R19 |           |
|```10100```|R20 |           |
|```10101```|R21 |           |
|```10110```|R22 |           |
|```10111```|R23 |           |
|```11000```|R24 |           |
|```11001```|R25 |           |
|```11010```|R26 |           |
|```11011```|R27 |           |
|```11100```|R28 |           |
|```11101```|R29 |           |
|```11110```|R30 |           |
|```11111```|R31 |           |

### Instruction Layout

|```31          28```|```      27      ```|```26          24```|```23                19```|```18           14```|```13          0```|
|--------------------|--------------------|--------------------|--------------------------|---------------------|-------------------|
|```Operation Code```|```Immediate Flag```|```Condition Code```|```Destination Register```|```Source Register```|```Operand Value```|

***Operation Code***: Specifies the operation to be performed.

|Code      |Name          |Description                      |
|----------|--------------|---------------------------------|
|```0000```|CMP           |Compare (subtract) and set flags.|
|```0001```|ADD           |Signed integer addition.         |
|```0010```|SUB           |Signed integer subtraction.      |
|```0011```|MUL           |Signed integer multiplication.   |
|```0100```|DIV           |Signed integer division.         |
|```0101```|MOD           |Signed integer modulo.           |
|```0110```|AND           |Bitwise AND.                     |
|```0111```|NAN           |Bitwise NAND.                    |
|```1000```|EOR           |Bitwise exclusive OR (XOR).      |
|```1001```|IOR           |Bitwise inclusive OR.            |
|```1010```|NOR           |Bitwise NOR.                     |
|```1011```|LSH           |Bitwise logical left shift.      |
|```1100```|RSH           |Bitwise logical right shift.     |
|```1101```|***RESERVED***|                                 |
|```1110```|***RESERVED***|                                 |
|```1111```|***RESERVED***|                                 |

***Immediate Flag***: Indicates whether the ***Operand Value*** is a register or immediate value.

|Flag   |Type           |
|-------|---------------|
|```0```|Register       |
|```1```|Immediate Value|

***Condition Code***: Defines conditions under which the instruction executes.

|Code     |Name          |Description              |Flag State                                   |
|---------|--------------|-------------------------|---------------------------------------------|
|```000```|UN            |Uncondition.             |Any or none set.                             |
|```001```|LT            |Less than.               |***N*** set (CMP result < 0).                |
|```010```|LE            |Less than or equal to.   |***N*** or ***Z*** set (CMP result <= 0).    |
|```011```|EQ            |Equal to                 |***Z*** set (CMP result == 0).               |
|```100```|NE            |Not equal to.            |***Z*** not set (CMP result != 0).           |
|```101```|GE            |Greater than or equal to.|***N*** not set (CMP result >= 0).           |
|```110```|GT            |Greater than.            |***N*** and ***Z*** not set (CMP result > 0).|
|```111```|***RESERVED***|                         |                                                 |

***Destination Register***: The register where the result of the operation is stored.

***Source Register***: The register containing the first operand value.

***Operand Value***: Either the register containing the second operand value or an immediate value, depending on the ***Immediate Flag***.

|Type           |Bits Range  |
|---------------|------------|
|Register       |```13 - 9```|
|Immediate Value|```13 - 0```|

