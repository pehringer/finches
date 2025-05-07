# Mapper

Mapper is a library that uses linear genetic programming to automatically synthesize functions that fit input-output data. It’s designed for discovering deterministic mappings from structured data.

## Use Cases

- **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable. For example, Mapper can help uncover the rules behind proprietary black-box systems by modeling their behavior.

- **Data Compression** - Evolve compact functions that approximate large datasets. By replacing raw data with concise models, Mapper enables significant reductions in storage for structured, deterministic data.

## Key Concepts

- **Linear Genetic Programming** – A form of genetic programming where programs are represented as linear sequences of instructions. These programs are evolved through mutation and selection to minimize a fitness function.
  + Linear, instruction-based programs map closely to hardware operations, enabling fast evaluation.
  + Well suited for exploring large, rugged search spaces where traditional methods may struggle.

- **Conditional Execution** – Individual instructions can be conditionally executed based on the results of prior instructions. This enables branching logic without requiring explicit control flow statements like jumps or branches.
  + Traditional branch-based control flow is often brittle and prone to breaking under mutation or crossover. Conditional execution offers a more robust, more position-independent alternative.
  + Reduces control flow complexity (shrinking the search space) while still supporting basic branching behaviors like piecewise logic through simple if-else structures.

- **Accumulator Architecture** – Programs use a single accumulator register for all intermediate computations. This simplifies the instruction set while still allowing for complex operations through sequences of simpler ones.
  - Encourages incremental evolution of functionality. Simple accumulator operations can be combined over generations into richer behaviors.
  - Aligns naturally with single‑input, single‑output programs. Write the input into the accumulator, run the instruction sequence, then read the output from the accumulator.

## Novel Design

The combination of accumulator-based computation and conditional instruction execution as the sole control flow mechanism offers a novel approach within the Linear Genetic Programming domain. It improves robustness under mutation and crossover, enables smoother evolution of complex behaviors, and aligns naturally with single-input/single-output problem domains.

# Instruction Set Architecture

```A``` - Accumulator register.

```F``` - Flag register.
- ```Z``` - Zero flag, set if the accumulator is zero.
- ```N``` - Negative flag, set if the accumulator is negative.

```M[0-255]``` - Memory.
- Immediate values are preloaded into memory before execution.

16-Bit Instructions format.
|Bits   |Field     |Description                       |
|-------|----------|----------------------------------|
|[15-13]|```COND```|Condition for execution.          |
|[12-9] |```OPER```|Operation to execute.             |
|[8]    |```SETF```|Set flag register after execution.|
|[7-0]  |```ADDR```|Address of operand in memory.     |

|```COND```|Mnemonic|Pseudocode                    |
|----------|--------|------------------------------|
|000       |        |if True                       |
|001       |LT      |if ```N```                    |
|010       |LE      |if ```N``` or ```Z```         |
|011       |EQ      |if ```Z```                    |
|100       |NE      |if not ```Z```                |
|101       |GE      |if not ```N```                |
|110       |GT      |if not ```N``` and not ```Z```|
|111       |NV      |if False                      |

|```OPER```|Mnemonic|Pseudocode                           |Protection|
|----------|--------|-------------------------------------|----------|
|0000      |LD      |```A``` = ```M[ADDR]```              |          |
|0001      |ST      |```M[ADDR]``` = ```A```              |          |
|0010      |AD      |```A``` += ```M[ADDR]```             |          |
|0011      |SB      |```A``` -= ```M[ADDR]```             |          |
|0100      |ML      |```A``` *= ```M[ADDR]```             |          |
|0101      |DV      |```A``` /= ```M[ADDR]```             |Zero      |
|0110      |MX      |```A``` = max(```A```, ```M[ADDR]```)|          |
|0111      |MN      |```A``` = min(```A```, ```M[ADDR]```)|          |
|1000      |AB      |```A``` = abs(```M[ADDR]```)         |          |
|1001      |PW      |```A``` = pow(```A```, ```M[ADDR]```)|NaN, Inf  |
|1010      |SQ      |```A``` = sqrt(```M[ADDR]```)        |NaN, Inf  |
|1011      |EX      |```A``` = exp(```M[ADDR]```)         |NaN, Inf  |
|1100      |LG      |```A``` = log(```M[ADDR]```)         |NaN, Inf  |
|1101      |SN      |```A``` = sin(```M[ADDR]```)         |          |
|1110      |CS      |```A``` = cos(```M[ADDR]```)         |          |
|1111      |TN      |```A``` = tan(```M[ADDR]```)         |NaN, Inf  |
