package days.day13

import framework.Task
import shared.chineseRemainder
import java.net.URL

class Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val data = input.openStream()
            .bufferedReader()
            .readLines()
            .let { lines ->
                lines[1].split(',')
                    .withIndex()
                    .filter { it.value != "x" }
                    .map { IndexedValue(it.index, it.value.toLong()) }
            }

        val result =
            chineseRemainder(
                data.map { it.value }.toLongArray(),
                data.map { it.value - it.index.toLong() }.toLongArray()
            )

        return Result.success(result)
    }
}
