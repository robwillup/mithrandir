# Computer Architecture

John Von Neuman, mathematician, physicist and computer scientist.

He said his ideas was based on the work of others, particularly John Mauchly and J. Presper Eckert.
These two were the designers of the EDVAC.

John Neumann's architecture said that the Processing unit included processor registers and the ALU.
Control Unit includes both an instruction register and a program counter.
These two form the CPU.

Memory unit stores data and instructions. It is meant for short term storage of information.
So the architecture also includes a Mass Storage Unit for long term persistent storage.
Also, there are input and output devices.

This architecture has persisted for so long because it does not dictates the details of how
its components should be implemented.

The benefits of this architecture are two-fold:

- General purpose computing
  - Load any program
  - Well defined instruction set
  - Stored program

- Data and program stored together

- Arithmetic and Logical Unit: gets two input and returns and output
  - It takes inputs along with Operation
  - It produces output with flags to indicates conditions
- Control Unit
  - Responsible for interacting with memory and registers
  - Keeps track of current instructions
  - Has a program counter
  - Sends decoded operations to the ALU
  - Does something with the flags that come from the ALU
- Registers
  - Short term memory of the CPU
  - Stash data for use by the next operation
  - data and operations are stored in memory
- Memory
  - Up to the control unit to fetch the instructions and execute them

- Cycle by which the CPU processes instructions is composed of 3 discrete operations
  - Fetch: Grabs the next instruction from an address in memory
  - Decode: Figure out what the instruction type is that's going to be executed.
  - Execute: Performs the actual work and increments the program counter when it's done.

- The CPU clock drives the cycle
