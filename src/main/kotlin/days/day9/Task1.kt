package days.day9

import framework.Task
import shared.generatePermutations
import java.net.URL

object Task1 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val result = input
            .openStream()
            .bufferedReader()
            .useLines { lines -> lines.map { it.toLong() }.toList() }
            .windowed(26, 1)
            .first { window ->
                val target = window.last()
                val preamble = window.dropLast(1)

                preamble.generatePermutations().all { it.first + it.second != target }
            }
            .last()

        return Result.success(result)
    }
}
