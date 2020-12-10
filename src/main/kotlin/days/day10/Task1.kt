package days.day10

import framework.Task
import java.net.URL

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input
            .openStream()
            .bufferedReader()
            .useLines { lines ->
                lines
                    .map { it.toInt() }
                    .plus(0)
                    .toList()
                    .sorted()
                    .windowed(2, 1)
                    .map { it[1] - it[0] }
                    .groupBy { it }
                    .let { it[1]!!.count() * (it[3]!!.count() + 1) }
            }

        return Result.success(result)
    }
}
