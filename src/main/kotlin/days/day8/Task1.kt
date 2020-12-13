package days.day8

import days.day8.machine.AccumulatorMachine
import days.day8.machine.InfiniteLoopException
import days.day8.machine.Instruction
import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val instructions: List<Instruction> = input
            .openStream()
            .bufferedReader()
            .useLines { loadInstructions(it) }

        try {
            AccumulatorMachine().run(instructions)
        } catch (e: InfiniteLoopException) {
            return Result.success(e.state.registers["acc"]!!)
        }

        return Result.success(-1)
    }
}
