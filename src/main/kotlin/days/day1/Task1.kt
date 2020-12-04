package days.day1

import framework.Task
import framework.exceptions.AnswerNotFoundException

class Task1 : Task<Int>(1) {
    override fun run(input: String): Result<Int> {
        val data = input.lines().filter { it != "" }.map { it.toInt() }

        for ((index, lower) in data.withIndex()) for (upper in data.drop(index+1)) {
            if (lower + upper == 2020) {
                return Result.success(lower * upper)
            }
        }

        return Result.failure(AnswerNotFoundException())
    }
}


