package days.day9

import framework.Task
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

                generatePermutations(preamble).all { it.first + it.second != target }
            }
            .last()

        return Result.success(result)
    }
}

fun <T> generatePermutations(data: Collection<T>): Sequence<Pair<T, T>> = sequence {
    for (first in data)
        for (second in data.minus(first)) yield(Pair(first, second))
}
