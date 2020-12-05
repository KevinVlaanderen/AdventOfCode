package days.day1

import framework.Task
import framework.exceptions.AnswerNotFoundException
import shared.toLines

object Task2 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val data = input
            .toLines()
            .filter { it.isNotBlank() }
            .map { it.toInt() }
            .toList()

        for ((firstIndex, first) in data.withIndex())
            for ((secondIndex, second) in data.drop(firstIndex + 1).withIndex())
                for (third in data.drop(secondIndex + 1)) {
                    if (first + second + third == 2020) {
                        return Result.success(first * second * third)
                    }
                }

        return Result.failure(AnswerNotFoundException())
    }
}


