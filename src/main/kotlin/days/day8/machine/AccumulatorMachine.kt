package days.day8.machine

object AccumulatorMachine {
    private var instructionPointer = 0
    private var stack = mutableListOf<Instruction>()
    private var registers = mutableMapOf("acc" to 0)

    private fun executeInstruction(instruction: Instruction) {
        when (instruction.operation) {
            Operation.ACCUMULATE -> {
                registers["acc"] = registers["acc"]!! + instruction.argument
                instructionPointer += 1
            }
            Operation.JUMP -> instructionPointer += instruction.argument
            Operation.NOOP -> instructionPointer += 1
        }
    }

    fun run(instructions: List<Instruction>): ExecutionState {
        instructionPointer = 0
        stack.clear()
        registers = mutableMapOf("acc" to 0)

        while (instructionPointer < instructions.count()) {
            val nextInstruction = instructions[instructionPointer]

            stack.add(nextInstruction)

            if (stack.dropLast(1).contains(nextInstruction)) throw InfiniteLoopException(captureExecutionState())

            executeInstruction(nextInstruction)
        }

        return captureExecutionState()
    }

    private fun captureExecutionState() = ExecutionState(
        stack,
        registers
    )
}