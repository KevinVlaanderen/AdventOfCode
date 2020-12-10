package days.day5

import framework.Task
import java.net.URL

object Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { line -> calculateSeatId(line) }
            .filter { seat -> seat in 7..1016 }
            .toList()
            .sorted()
            .windowed(2)
            .find { list -> list[1] == list[0] + 2 }
            .let { list -> list!![0] + 1 }


        return Result.success(result)
    }
}


