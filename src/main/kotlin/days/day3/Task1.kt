package days.day3

import framework.Task
import shared.readResource

class Task1 : Task<Int>(1) {
    override fun run(): Result<Int> {
        val input = readResource("/day3").lines().filter { it != "" }

        val count = countTrees(input, 3, 1)

        return Result.success(count)
    }
}


