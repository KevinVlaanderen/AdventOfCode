package days.day9

import framework.Task
import framework.exceptions.AnswerNotFoundException
import java.net.URL

object Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val target = 1504371145L

        val data = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { it.toLong() }
            .toList()

        generateSequence(2) { if (it < data.size) it + 1 else null }
            .forEach { windowSize ->
                data
                    .windowed(windowSize, 1)
                    .find { it.sum() == target }
                    ?.run {
                        val sorted = sorted()
                        return Result.success(sorted.first() + sorted.last())
                    }
            }

        return Result.failure(AnswerNotFoundException())
    }
}
