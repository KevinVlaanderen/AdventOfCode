package days.day8

import days.day8.machine.*
import framework.Task
import shared.replace
import java.net.URL

object Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val instructions: List<Instruction> = input
            .openStream()
            .bufferedReader()
            .useLines { loadInstructions(it) }

        var executionState: ExecutionState? = null

        try {
            executionState = AccumulatorMachine.run(instructions)
        } catch (e: InfiniteLoopException) {
            e.state.stack
                .filter { it.operation == Operation.JUMP || it.operation == Operation.NOOP }
                .reversed()
                .forEach retry@{ lastJumpOrNoop ->
                    val newInstruction = Instruction(
                        lastJumpOrNoop.pointer,
                        when (lastJumpOrNoop.operation) {
                            Operation.JUMP -> Operation.NOOP
                            Operation.NOOP -> Operation.JUMP
                            else -> lastJumpOrNoop.operation
                        },
                        lastJumpOrNoop.argument
                    )

                    val newInstructions = instructions.replace(newInstruction) { it.pointer == lastJumpOrNoop.pointer }

                    try {
                        executionState = AccumulatorMachine.run(newInstructions)
                        return@retry
                    } catch (e: Throwable) {
                    }
                }
        }

        return Result.success(executionState?.registers?.get("acc") ?: 0)
    }
}
