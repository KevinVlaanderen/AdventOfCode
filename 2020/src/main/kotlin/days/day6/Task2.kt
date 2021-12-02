package days.day6

import framework.Task
import shared.toBlocks
import shared.toLines
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input
            .openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .map {
                it
                    .toLines()
                    .map { line -> line.toSet() }
                    .reduce { acc, set -> acc.intersect(set) }
                    .size
            }
            .sum()

        return Result.success(result)
    }
}


