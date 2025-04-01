# ***clgp*** - ***C***onditional ***L***inear ***G***enetic ***P***rogramming

### Key Concepts

1) ***Linear Genetic Programming (LGP)***: A type of Genetic Programming (GP) where programs are represented as linear sequences of instructions, similar to machine code. LGP operates on a register-based architecture where instructions modify values stored in registers. Programs execute sequentially, and their evolution is driven by genetic operators such as mutation and crossover, which modify the instruction sequence to optimize performance.

2) ***Conditionally Executed Instructions***: Operations that only execute if a specified condition is met, without requiring explicit branching. Instead of altering the flow of execution with jumps or branches, these instructions check a condition and either perform their operation or do nothing. These type of instructions are found in many popular CPU architectures, such as x86 and ARM, where some instructions can be conditionally executed based on processor flags set by compare instructions.

### Current LGP Conditional Execution Model

Some LGP implementations already incorporate conditional execution, typically through conditional branching or skipping. These mechanisms allow the program to alter its flow based on certain conditions. While functional, these approaches introduce challenges in the context of genetic algorithms.

The primary issue is that small mutations, such as changing an existing instruction or inserting a new one, can drastically alter the control flow. This can lead to the program becoming non-functional, as conditional branching or skipping often relies on a tight instruction ordering to function correctly. This makes such conditional execution mechanisms poorly suited for the evolution of programs through genetic algorithms, where incremental modifications are common.

### An Alternative LGP Conditional Execution Model

An alternative approach to conditional execution in LGP involves using instructions that are conditionally executed based on flags set by comparison operations. This type of execution has a number of advantages over traditional branching or skipping.

Primarily, conditionally executed instructions are more adaptable to evolutionary modifications because these operations rely on a looser, linear, flag-dependent instruction ordering. They are more resilient to small mutations, such as changing an existing instruction or inserting a new one. This results in a more stable evolutionary process, where conditional logic can evolve incrementally without disturbing the integrity of the program’s execution.

Additionally, the flag-setting and flag-dependent execution introduces a linear, order-based dependency that fits naturally with LGP’s linear, order-dependent nature.
