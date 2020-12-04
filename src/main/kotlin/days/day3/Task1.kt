package days.day3

import framework.Task
import shared.toLines

class Task1 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val data = input.toLines()

        val count = countTrees(data, 3, 1)

        return Result.success(count)
    }
}


