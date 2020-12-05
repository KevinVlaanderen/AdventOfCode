package days.day1

import framework.Task
import framework.exceptions.AnswerNotFoundException
import shared.toLines

object Task1 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val data = input
            .toLines()
            .filter { it.isNotBlank() }
            .map { it.toInt() }
            .toList()

        for ((index, lower) in data.withIndex())
            for (upper in data.drop(index + 1)) {
                if (lower + upper == 2020) return Result.success(lower * upper)
            }

        return Result.failure(AnswerNotFoundException())
    }
}


