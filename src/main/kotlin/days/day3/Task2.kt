package days.day3

import framework.Task

class Task2 : Task<Long>(2) {
    override fun run(input: String): Result<Long> {
        val data = input.lines().filter { it != "" }

        val count: Long =
            countTrees(data, 1, 1).toLong() *
            countTrees(data, 3, 1).toLong() *
            countTrees(data, 5, 1).toLong() *
            countTrees(data, 7, 1).toLong() *
            countTrees(data, 1, 2).toLong()

        return Result.success(count)
    }
}


