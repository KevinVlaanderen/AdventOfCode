package days.day3

import framework.Task

class Task1 : Task<Int>(1) {
    override fun run(input: String): Result<Int> {
        val data = input.lines().filter { it != "" }

        val count = countTrees(data, 3, 1)

        return Result.success(count)
    }
}


