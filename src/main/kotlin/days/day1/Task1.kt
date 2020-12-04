package days.day1

import framework.exceptions.AnswerNotFoundException
import framework.Task
import shared.readResource

class Task1 : Task<Int>(1) {
    override fun run(): Result<Int> {
        val input = readResource("/day1")

        val data = input.lines().filter { it != "" }.map { it.toInt() }

        for ((index, lower) in data.withIndex()) for (upper in data.drop(index+1)) {
            if (lower + upper == 2020) {
                return Result.success(lower * upper)
            }
        }

        return Result.failure(AnswerNotFoundException())
    }
}


