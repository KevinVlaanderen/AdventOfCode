package days.day8

import days.day8.machine.Instruction
import days.day8.machine.Operation

val instructionRegex = """(.*) ([+|-]\d+)""".toRegex()

fun loadInstructions(lines: Sequence<String>): List<Instruction> = lines.mapIndexedNotNull { index, line ->
    instructionRegex.find(line)?.let {
        Instruction(
            index,
            Operation.getOperationByKeyword(it.groupValues[1]),
            it.groupValues[2].toInt()
        )
    }
}.toList()