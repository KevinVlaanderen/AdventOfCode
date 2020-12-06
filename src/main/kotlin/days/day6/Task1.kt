package days.day6

import framework.Task
import shared.toBlocks
import java.net.URL

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input
            .openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .map { block ->
                block
                    .replace(System.lineSeparator(), "")
                    .toSet()
                    .size
            }
            .sum()

        return Result.success(result)
    }
}


