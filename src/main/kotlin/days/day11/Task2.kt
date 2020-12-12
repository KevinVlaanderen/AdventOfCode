package days.day11

import framework.Task
import java.net.URL

object Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val floor = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { it.toList() }
            .toList()

        val result = seatsWhenStabilized(floor, 5)

        return Result.success(result)
    }
}
