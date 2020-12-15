package days.day15

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var last: Int
        var first: MutableList<Int>

        input.openStream()
            .bufferedReader()
            .readLine()
            .split(',')
            .map { it.toInt() }
            .run { last = this.last(); first = this.dropLast(1).toMutableList() }

        (first.size until 2020).forEach {
            val lastIndex = first.lastIndexOf(last)

            first.add(last)
            last = if (lastIndex >= 0) it - lastIndex else 0
        }

        return Result.success(first.last())
    }
}
