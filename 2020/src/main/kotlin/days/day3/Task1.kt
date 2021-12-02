package days.day3

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val data = input
            .openStream()
            .bufferedReader()
            .readLines()

        val count = countTrees(data, 3, 1)

        return Result.success(count)
    }
}


