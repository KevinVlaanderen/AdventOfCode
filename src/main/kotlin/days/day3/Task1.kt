package days.day3

import framework.Task
import shared.toLines

object Task1 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val data = input.toLines().toList()

        val count = countTrees(data, 3, 1)

        return Result.success(count)
    }
}


