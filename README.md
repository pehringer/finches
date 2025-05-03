# Function Generator (fungen)

# Instruction Set Architecture

- State Register: **S**
- Accumulator Register: **A**
- General Purpose Registers: **R[0-63]**
  + Immediate values are preloaded into these registers before execution.

#### Instruction format:

|[15-10]|[9-6] |[5-0]  |
|-------|------|-------|
|STATE  |OPCODE|OPERAND|

#### Instructions:

|OPCODE|Mnemonic|Pseudocode                           |Description                  |
|------|--------|-------------------------------------|-----------------------------|
|0000  |LD      |if S == STATE: A = R[OPERAND]        |Load Accumulator             |
|0001  |ST      |if S == STATE: R[OPERAND] = A        |Store Accumulator            |
|0010  |AD      |if S == STATE: A += R[OPERAND]       |Floating Point Addition      |
|0011  |SB      |if S == STATE: A -= R[OPERAND]       |Floating Point Subtraction   |
|0100  |ML      |if S == STATE: A *= R[OPERAND]       |Floating Point Multiplication|
|0101  |DV      |if S == STATE: A /= R[OPERAND]       |Floating Point Division      |
|0110  |LT      |if S == STATE and A < 0: S = OPERAND |Transition Less Than         |
|0111  |GT      |if S == STATE and A > 0: S = OPERAND |Transition Greater Than      |
|1000  |EQ      |if S == STATE and A == 0: S = OPERAND|Transition Equal             |
|1001  |NE      |if S == STATE and A != 0: S = OPERAND|Transition Not Equal         |
|1010  |NOP     |                                     |No Operation                 |
|1011  |NOP     |                                     |No Operation                 |
|1100  |NOP     |                                     |No Operation                 |
|1101  |NOP     |                                     |No Operation                 |
|1110  |NOP     |                                     |No Operation                 |
|1111  |NOP     |                                     |No Operation                 |
