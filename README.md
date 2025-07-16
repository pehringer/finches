# Finches

Finches is a library that uses linear genetic programming (LGP) to automatically synthesize functions that fit input-output data.
It’s designed for discovering deterministic mappings from structured data.

# Use Cases

- **Reverse Engineering** - Infer hidden logic from input–output pairs, even when source code or hardware is unavailable.
For example, Finches can help uncover the rules behind proprietary black-box systems by modeling their behavior.

- **Data Compression** - Evolve compact functions that approximate large datasets.
By replacing raw data with concise models, Finches enables significant reductions in storage for structured, deterministic data.

# Finches Instruction Set Architecture (ISA)

**Finches ISA is specifically designed to optimize LGP evolvability**, creating a smoother solution landscape for programs.
Traditional LGP often rely on branching control flow operations, which easily break under genetic crossover and mutation, leading to a rougher fitness landscape.
Finches, however, **employs branchless programming techniques for its control flow**.
This approach integrates control flow directly into computations.
By avoiding explicit branching, LGP can incrementally evolve conditional flow logic through symbolic regression, resulting in a significantly smoother solution landscape.

Furthermore, this ISA **only contains arithmetic comparison operations** (such as min, max, <, and >).
This design choice provides a more incremental and continuous path for forming control flow logic when compared to all-or-nothing boolean operations (like AND, OR, NOT), further contributing to a smoother solution landscape.

Finally, the **inclusion of inverse operations** (e.g., addition and subtraction, multiplication and division, sine and arcsine) creates a more uniform and balanced search space.
This allows instructions to be effectively "undone" or counteracted, which further contributes to a smoother solution landscape and enhances Finches ability to explore.

These design choices, including the **compact uniform 16-bit instruction format**, collectively contribute to a significantly reduced search space for LGP to explore.
By constraining these architectural dimensions, Finches can more efficiently navigate the solution landscape, making the evolutionary process more efficient.

**Instruction Format:**
|OPCODE |RESULT|FIRST|SECOND|
|-------|------|-----|------|
|[15-12]|[11-8]|[7-4]|[3-0] |

**Instruction Set:**
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

**NOTE:** Operations return NaN (Not a Number) or Inf (Infinity) for invalid operands.

# Finches Genetic Algorithm (GA)

**Finches GA is designed to be as simple as possible, yet still capable of eliciting complex and desirable evolutionary behaviors**.
It promotes slow population convergence, encourages depth-first and breadth-first exploration of the solution space, and **manages genome size and bloat**.
Furthermore, it **minimize program breakage**, ensuring genetic operations are less disruptive to evolving programs.
All these design choices not only contribute to a smoother fitness landscape but also directly enhance performance by simplifying computational overhead.

The major difference in the GA is the **replacement of the traditional crossover operator with the Fission and Transfer operators**.
These two operators simulate horizontal gene transfer, rather then the vertical gene transfer of traditional crossover.

**Algorithm:**
- **Initialization:** Create a population of individuals. Each individual contains a very low fitness score, sixteen random constants, and a single random instruction.
- **Loop:**
  + **Selection:** Two neighboring parents are randomly selected from population. A separate donor individual is also randomly selected from the population.
  + **Replacement:** The parent with the lower fitness score (or if equal, the one with a higher instruction count) is selected to be the new offspring.
  + **Fission:** The offspring's instructions and constants are replaced with a copy of the remaining parent's instructions and constants.
  + **Mutation:** The offspring then undergoes one of the following mutations:
    * A constant is randomly adjusted slightly.
    * An instruction is randomly replaced (with a new random one).
    * An instruction is randomly deleted.
    * An instruction is randomly inserted (with a new random one).
  + **Transfer:** With a small probability, a small random block of instructions is copied from the donor and is randomly inserted into the offspring.
  + **Evaluation:** The offspring's fitness score is calculated by running a series of test cases. NaN or Infinity results are heavily penalized.
- **Termination:** The individual with the highest fitness score is selected. Its constants and instructions are returned.
