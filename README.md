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

|Name|Description                                      |
|----|-------------------------------------------------|
|N   |Set if the result of CMP is negtive (sign bit).  |
|Z   |Set if the result of CMP is zero (all bits zero).|

### Processor Registers

|Name|Description              |
|----|-------------------------|
|R0  |General purpose register.|
|R1  |General purpose register.|
|R2  |General purpose register.|
|R3  |General purpose register.|
|R4  |General purpose register.|
|R5  |General purpose register.|
|R6  |General purpose register.|
|R7  |General purpose register.|
|R8  |General purpose register.|
|R9  |General purpose register.|
|R10 |General purpose register.|
|R11 |General purpose register.|
|R12 |General purpose register.|
|R13 |General purpose register.|
|R14 |General purpose register.|
|R15 |General purpose register.|

### Instruction Layout

|```31          28```|```      27      ```|```26          24```|```23                20```|```19           16```|```15          0```|
|--------------------|--------------------|--------------------|--------------------------|---------------------|-------------------|
|```Operation Code```|```Immediate Flag```|```Condition Code```|```Destination Register```|```Source Register```|```Operand Value```|

```Operation Code``` specifies the operation to be performed:

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

```Immediate Flag``` indicates whether the ```Operand Value``` is a register or immediate value:

|Flag   |Type           |
|-------|---------------|
|```0```|Register       |
|```1```|Immediate Value|

```Condition Code``` specifies conditions under which the instruction executes:

|Code     |Name          |Description                               |
|---------|--------------|------------------------------------------|
|```000```|UN            |Uncondition                               |
|```001```|LT            |Less than (N flag set).                   |
|```010```|LE            |Less than or equal to (N or Z flags set). |
|```011```|EQ            |Equal to (Z flag set).                    |
|```100```|NE            |Not equal to (Z flag not set).            |
|```101```|GE            |Greater than or equal to (N flag not set).|
|```110```|GT            |Greater than (N and Z flags not set).     |
|```111```|***RESERVED***|                                          |

```Destination Register``` defines the register where the result of the operation is stored.

```Source Register``` defines the register containing the first operand value.

```Operand Value``` defines either the register containing the second operand value or an immediate value:

|Type           |Bits Range   |
|---------------|-------------|
|Register       |```15 - 12```|
|Immediate Value|```15 -  0```|
