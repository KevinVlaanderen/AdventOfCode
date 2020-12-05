package days.day5

import framework.Task
import java.net.URL

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        input
            .openStream()
            .bufferedReader()
            .useLines { lines ->
                val result = lines
                    .filter { it.isNotBlank() }
                    .map { line -> calculateSeatId(line) }
                    .sortedDescending()
                    .first()

                return Result.success(result)
            }
    }
}


