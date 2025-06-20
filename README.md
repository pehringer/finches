# Mapper

Mapper is a library that uses linear genetic programming (LGP) to automatically synthesize functions that fit input-output data. It’s designed for discovering deterministic mappings from structured data.

# Use Cases

- **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable. For example, Mapper can help uncover the rules behind proprietary black-box systems by modeling their behavior.

- **Data Compression** - Evolve compact functions that approximate large datasets. By replacing raw data with concise models, Mapper enables significant reductions in storage for structured, deterministic data.

# Instruction Set Architecture (ISA)

**Mapper's ISA is specifically designed to optimize LGP evolvability**, creating a smoother solution landscape for programs.

Traditional LGP often rely on branching control flow operations, which easily break under genetic crossover and mutation, leading to a rougher fitness landscape. Mapper, however, **employs branchless programming techniques for its control flow**. This approach integrates control flow directly into computations. By avoiding explicit branching, LGP can incrementally evolve conditional flow logic through symbolic regression, resulting in a significantly smoother solution landscape.

For example, consider how the absolute value function can be implemented:

**Branched Version:**
```
abs(x):
  if x < 0:
    return x * -1
  else:
    return x
```

**Branchless Version:**
```
abs(x):
  return ((x < 0) * (x * -1)) + ((x > 0) * (x)) 
```

Furthermore, this ISA **only contains arithmetic comparison operations** (such as min, max, <, and >). This design choice provides a more incremental and continuous path for forming control flow logic when comparied to all-or-nothing boolean operations (like AND, OR, NOT), further contributing to a smoother solution landscape.

Finally, the **inclusion of inverse operations** (e.g., addition and subtraction, multiplication and division, sine and arcsine) creates a more uniform and balanced search space. This allows instructions to be effectively "undone" or counteracted**, which further contributes to a smoother solution landscape and enhances Mapper's ability to explore.

These design choices, including the **compact uniform 16-bit instruction format**, collectively contribute to a significantly reduced search space for LGP to explore. By constraining these architectural dimensions, Mapper can more efficiently navigate the solution landscape, making the evolutionary process more efficient.

```R[0-15]``` - Registers.
- **Input registers**: The 1st input is loaded into R0, the 2nd into R1, and so on.
- **Output register**: The final program output is stored in R15.
- The remaining registers are pre-loaded with constant values.

- **16 Bit Instruction Format:**
|OPCODE |RESULT|FIRST|SECOND|
|-------|------|-----|------|
|[15-12]|[11-8]|[7-4]|[3-0] |

**NOTE:** Operations return NaN (Not a Number) or Inf (Infinity) for invalid operands.

|OPCODE|Mnemonic|Pseudocode                                        |
|------|--------|--------------------------------------------------|
|0000  |AD      |```R[RESULT] = R[FIRST] + R[SECOND]```            |
|0001  |SB      |```R[RESULT] = R[FIRST] - R[SECOND]```            |
|0010  |ML      |```R[RESULT] = R[FIRST] * R[SECOND]```            |
|0011  |DV      |```R[RESULT] = R[FIRST] / R[SECOND]```            |
|0100  |PW      |```R[RESULT] = pow(R[FIRST], R[SECOND])```        |
|0101  |SQ      |```R[RESULT] = sqrt(R[FIRST])```                  |
|0110  |EX      |```R[RESULT] = exp(R[FIRST])```                   |
|0111  |LG      |```R[RESULT] = log(R[FIRST])```                   |
|1000  |SN      |```R[RESULT] = sin(R[FIRST])```                   |
|1001  |AS      |```R[RESULT] = asin(R[FIRST])```                  |
|1010  |CS      |```R[RESULT] = cos(R[FIRST])```                   |
|1011  |AC      |```R[RESULT] = acos(R[FIRST])```                  |
|1100  |MN      |```R[RESULT] = min(R[FIRST], R[SECOND])```        |
|1101  |MX      |```R[RESULT] = max(R[FIRST], R[SECOND])```        |
|1110  |LT      |```R[RESULT] = 1 if R[FIRST] < R[SECOND] else 0```|
|1111  |GT      |```R[RESULT] = 1 if R[FIRST] > R[SECOND] else 0```|
