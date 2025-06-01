# Mapper

Mapper is a library that uses linear genetic programming to automatically synthesize functions that fit input-output data. It’s designed for discovering deterministic mappings from structured data.

## Use Cases

- **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable. For example, Mapper can help uncover the rules behind proprietary black-box systems by modeling their behavior.

- **Data Compression** - Evolve compact functions that approximate large datasets. By replacing raw data with concise models, Mapper enables significant reductions in storage for structured, deterministic data.

# Instruction Set Architecture 

```R[0-15]``` - Registers.
- **Input registers**: The 1st input is loaded into R0, the 2nd into R1, and so on.
- **Output register**: The final program output is stored in R15.
- **Constant registers**: Remaining registers are pre-loaded with constant values.

- **16 Bit Instruction Format:**
|OPCODE |RESULT|FIRST|SECOND|
|-------|------|-----|------|
|[15-12]|[11-8]|[7-4]|[3-0] |

|OPCODE|Mnemonic|Pseudocode                                    |Protection|
|------|--------|----------------------------------------------|----------|
|0000  |AD      |```R[RESULT] = R[FIRST] + R[SECOND]```        |          |
|0001  |SB      |```R[RESULT] = R[FIRST] - R[SECOND]```        |          |
|0010  |ML      |```R[RESULT] = R[FIRST] * R[SECOND]```        |          |
|0011  |DV      |```R[RESULT] = R[FIRST] / R[SECOND]```        |Zero      |
|0100  |MN      |```R[RESULT] = min(R[FIRST], R[SECOND])```    |          |
|0101  |MX      |```R[RESULT] = max(R[FIRST], R[SECOND])```    |          |
|0110  |LT      |```R[RESULT] = R[FIRST] < R[SECOND] ? 1 : 0```|          |
|0111  |GT      |```R[RESULT] = R[FIRST] > R[SECOND] ? 1 : 0```|          |
|1000  |PW      |```R[RESULT] = pow(R[FIRST], R[SECOND])```    |NaN, Inf  |
|1001  |SQ      |```R[RESULT] = sqrt(R[FIRST])```              |NaN, Inf  |
|1010  |EX      |```R[RESULT] = exp(R[FIRST])```               |NaN, Inf  |
|1011  |LG      |```R[RESULT] = log(R[FIRST])```               |NaN, Inf  |
|1100  |SN      |```R[RESULT] = sin(R[FIRST])```               |          |
|1101  |CS      |```R[RESULT] = cos(R[FIRST])```               |          |
|1110  |NOP     |                                              |          |
|1111  |NOP     |                                              |          |
