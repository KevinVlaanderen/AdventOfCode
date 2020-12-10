package days.day5

import framework.Task
import java.net.URL

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { line -> calculateSeatId(line) }
            .toList().maxOrNull()!!

        return Result.success(result)
    }
}


