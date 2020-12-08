package days.day8

import framework.Task
import java.net.URL

val instructionRegex = """(.*) ([+|-]\d+)""".toRegex()

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val instructions = mutableListOf<Instruction>()

        input
            .openStream()
            .bufferedReader()
            .useLines { lines ->
                lines.forEach {
                    instructionRegex.find(it)?.run {
                        instructions.add(
                            Instruction(
                                Operation.getOperationByKeyword(groupValues[1]),
                                groupValues[2].toInt()
                            )
                        )
                    }
                }
            }

        val result = AccumulatorMachine().run(instructions)

        return Result.success(result)
    }
}

class AccumulatorMachine {
    private lateinit var instructions: List<Instruction>
    private var currentInstruction = 0
    private var executedInstructions = mutableListOf<Int>()

    var accumulator: Int = 0
        private set

    private fun executeInstruction(instruction: Instruction) {
        println(
            "Executing ${instruction.operation} with argument ${instruction.argument}"
        )

        when (instruction.operation) {
            Operation.ACCUMULATE -> {
                accumulator += instruction.argument
                currentInstruction += 1
            }
            Operation.JUMP -> currentInstruction += instruction.argument
            Operation.NOOP -> currentInstruction += 1
        }
    }

    fun run(instructions: List<Instruction>): Int {
        this.instructions = instructions
        currentInstruction = 0
        executedInstructions.clear()

        while (!executedInstructions.contains(currentInstruction)) {
            executedInstructions.add(currentInstruction)
            executeInstruction(instructions[currentInstruction])
        }

        println("Executed ${executedInstructions.size} instructions before exiting with accumulator value $accumulator")

        return accumulator
    }
}
