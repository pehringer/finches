# Mapper

Mapper is a library that uses linear genetic programming to automatically synthesize functions that fit input-output data. It’s designed for discovering deterministic mappings from structured data.

## Use Cases

- **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable. For example, Mapper can help uncover the rules behind proprietary black-box systems by modeling their behavior.

- **Data Compression** - Evolve compact functions that approximate large datasets. By replacing raw data with concise models, Mapper enables significant reductions in storage for structured, deterministic data.

# Instruction Set Architecture

```D[0-7]``` - Data Registers.
- Holds floating point results of operations. Used for instruction operands.
- Constant values are preloaded into the registers before program execution.

```F[0-7]``` - Flag Register.
- Holds boolean results of comparisons. Used for instruction execution.
- True values are preloaded into the register before program execution.

- **16 Bit Instruction Format:**
|OPERATION|RESULT|FIRST|SECOND|PREDICATE|
|---------|------|-----|------|---------|
|[15-12]  |[11-9]|[8-6]|[5-3] |[2-0]    |

|OPERATION|Mnemonic|Pseudocode                                                 |Protection|
|---------|--------|-----------------------------------------------------------|----------|
|0000     |AD      |```if F[PREDICATE]: R[RESULT] = R[FIRST] + R[SECOND]```    |          |
|0001     |SB      |```if F[PREDICATE]: R[RESULT] = R[FIRST] - R[SECOND]```    |          |
|0010     |ML      |```if F[PREDICATE]: R[RESULT] = R[FIRST] * R[SECOND]```    |          |
|0011     |DV      |```if F[PREDICATE]: R[RESULT] = R[FIRST] / R[SECOND]```    |Zero      |
|0100     |PW      |```if F[PREDICATE]: R[RESULT] = pow(R[FIRST], R[SECOND])```|NaN, Inf  |
|0101     |SQ      |```if F[PREDICATE]: R[RESULT] = sqrt(R[FIRST])```          |NaN, Inf  |
|0110     |EX      |```if F[PREDICATE]: R[RESULT] = exp(R[FIRST])```           |NaN, Inf  |
|0111     |LG      |```if F[PREDICATE]: R[RESULT] = log(R[FIRST])```           |NaN, Inf  |
|1000     |SN      |```if F[PREDICATE]: R[RESULT] = sin(R[FIRST])```           |          |
|1001     |CS      |```if F[PREDICATE]: R[RESULT] = cos(R[FIRST])```           |          |
|1010     |TN      |```if F[PREDICATE]: R[RESULT] = tan(R[FIRST])```           |NaN, Inf  |
|1011     |AB      |```if F[PREDICATE]: R[RESULT] = abs(R[FIRST])```           |          |
|1100     |LT      |```if F[PREDICATE]: F[RESULT] = R[FIRST] < R[SECOND]       |          |
|1101     |LE      |```if F[PREDICATE]: F[RESULT] = R[FIRST] <= R[SECOND]      |          |
|1110     |EQ      |```if F[PREDICATE]: F[RESULT] = R[FIRST] == R[SECOND]      |          |
|1111     |NE      |```if F[PREDICATE]: F[RESULT] = R[FIRST] != R[SECOND]      |          |
