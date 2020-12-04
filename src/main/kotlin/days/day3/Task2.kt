package days.day3

import framework.Task
import shared.readResource

class Task2 : Task<Long>(2) {
    override fun run(): Result<Long> {
        val input = readResource("/day3").lines().filter { it != "" }

        val count: Long =
            countTrees(input, 1, 1).toLong() *
            countTrees(input, 3, 1).toLong() *
            countTrees(input, 5, 1).toLong() *
            countTrees(input, 7, 1).toLong() *
            countTrees(input, 1, 2).toLong()

        return Result.success(count)
    }
}


