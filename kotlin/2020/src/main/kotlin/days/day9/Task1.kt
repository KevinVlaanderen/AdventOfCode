package days.day9

import framework.Task
import shared.generatePermutations
import java.net.URL

class Task1(private val preambleLength: Int = 25) : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val result = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { it.toLong() }
            .toList()
            .windowed(preambleLength + 1, 1)
            .first { window ->
                val target = window.last()
                val preamble = window.dropLast(1)

                preamble.generatePermutations().all { it.first + it.second != target }
            }
            .last()

        return Result.success(result)
    }
}
