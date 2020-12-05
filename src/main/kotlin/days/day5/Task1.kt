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
                    .map { line ->
                        val row = Integer.parseInt(
                            line.take(7)
                                .replace('F', '0')
                                .replace('B', '1'), 2
                        )
                        val column = Integer.parseInt(
                            line.takeLast(3)
                                .replace('L', '0')
                                .replace('R', '1'), 2
                        )

                        row * 8 + column
                    }
                    .sortedDescending()
                    .first()

                return Result.success(result)
            }
    }
}


