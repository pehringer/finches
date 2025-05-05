# Mapper

Mapper is a library that uses linear genetic programming (LGP) to automatically synthesize functions that fit input–output data. It’s designed for discovering deterministic mappings from structured data.

## Use Cases

1. **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable. For example, Mapper can help uncover the rules behind proprietary black-box systems by modeling their behavior.

2. **Data Compression** - Evolve compact functions that approximate large datasets. By replacing raw data with concise models, Mapper enables significant reductions in storage for structured, deterministic data.

## Key Concepts

- **Linear Genetic Programming (LGP)** – A form of genetic programming where programs are represented as linear sequences of instructions. These programs are evolved through mutation and selection to minimize a fitness function.

- **Conditional Execution** – Individual instructions can be conditionally executed based on the results of prior instructions. This enables branching logic without requiring explicit control flow statements like jumps or branches.

- **Accumulator Architecture** – Programs use a single accumulator register for all intermediate computations. This simplifies the instruction set while still allowing for complex operations through sequences of simpler operations.




# Instruction Set Architecture

- Flag Register (**F**):
  + Zero flag (**Z**) bit 0.
  + Negative flag (**N**) bit 1.
- Accumulator Register (**A**).
- Memory (**M[0-255]**):
  + Immediate values are preloaded into memory before execution.

#### Instruction format:

|[15-13]  |[12-9]   |[8]    |[7-0]  |
|---------|---------|-------|-------|
|CONDITION|OPERATION|SETFLAG|ADDRESS|

#### CONDITION: Condition for execution:

|CONDITION|Mnemonic|Pseudocode            |Description        |
|---------|--------|----------------------|-------------------|
|000      |        |if True               |Always             |
|001      |LT      |if N                  |Less than.         |
|010      |LE      |if N or Z             |Less than equal.   |
|011      |EQ      |if Z                  |Equal.             |
|100      |NE      |if not Z              |Not equal.         |
|101      |GE      |if not N              |Greater than equal.|
|110      |GT      |if not N and not Z    |Greater than.      |
|111      |NV      |if False              |Never              |

#### OPERATION: Operation to execute:

|OPCODE|Mnemonic|Pseudocode            |Description              |
|------|--------|----------------------|-------------------------|
|0000  |LD      |A = M[OPERAND]        |Load accumulator.        |
|0001  |ST      |M[OPERAND] = A        |Store accumulator        |
|0010  |AD      |A += M[OPERAND]       |Addition.                |
|0011  |SB      |A -= M[OPERAND]       |Subtraction.             |
|0100  |ML      |A *= M[OPERAND]       |Multiplication.          |
|0101  |DV      |A /= M[OPERAND]       |Protected division.      |
|0110  |MX      |A = max(A, M[OPERAND])|Maximum.                 |
|0111  |MN      |A = min(A, M[OPERAND])|Minimum.                 |
|1000  |AB      |A = abs(M[OPERAND])   |Absolute value.          |
|1001  |PW      |A = pow(M[OPERAND])   |Protected exponentiation.|
|1010  |SQ      |A = sqrt(M[OPERAND])  |Protected square root.   |
|1011  |EX      |A = exp(M[OPERAND])   |Protected exponential.   |
|1100  |LG      |A = log(M[OPERAND])   |Protected logarithm.     |
|1101  |SN      |A = sin(M[OPERAND])   |Sine.                    |
|1110  |CS      |A = cos(M[OPERAND])   |Consine.                 |
|1111  |TN      |A = tan(M[OPERAND])   |Protected tangent.       |

#### SETFLAG: Set flag register after execution:

|SETFLAG|Mnemonic|Pseudocode            |Description|
|-------|--------|----------------------|-----------|
|0      |        |F = F                 |Do not set.|
|1      |S       |Z = A == 0; N = A < 0 |Set flags. |

#### ADDRESS: Address of operand.
