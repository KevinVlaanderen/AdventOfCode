package days.day1

import framework.Task
import framework.exceptions.AnswerNotFoundException

class Task2 : Task<Int>(2) {
    override fun run(input: String): Result<Int> {
        val data = input.lines().filter { it != "" }.map { it.toInt() }

        for ((firstIndex, first) in data.withIndex())
            for ((secondIndex, second) in data.drop(firstIndex+1).withIndex())
                for (third in data.drop(secondIndex+1)) {
                    if (first + second + third == 2020) {
                        return Result.success(first * second * third)
                    }
        }

        return Result.failure(AnswerNotFoundException())
    }
}


