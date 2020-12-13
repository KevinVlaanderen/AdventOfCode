package days.day1

import framework.Task
import framework.exceptions.AnswerNotFoundException
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val data = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { it.toInt() }
            .toList()

        for ((index, lower) in data.withIndex())
            for (upper in data.drop(index + 1)) {
                if (lower + upper == 2020) return Result.success(lower * upper)
            }

        return Result.failure(AnswerNotFoundException())
    }
}


