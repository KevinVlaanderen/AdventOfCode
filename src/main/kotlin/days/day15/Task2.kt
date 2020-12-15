package days.day15

import framework.Task
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val data = input.openStream()
            .bufferedReader()
            .readLine()
            .split(',')
            .map { it.toInt() }

        var last = data.last()
        val lastIndices = data.dropLast(1).withIndex().associate { it.value to it.index }.toMutableMap()

        (lastIndices.size until 30000000 - 1).forEach {
            val lastIndex = lastIndices[last]

            lastIndices[last] = it
            last = if (lastIndex == null) 0 else it - lastIndex
        }

        return Result.success(last)
    }
}
