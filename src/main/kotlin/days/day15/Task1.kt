package days.day15

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val data = input.openStream().bufferedReader().readLine()
            .split(',').map { it.toInt() }

        val result = playMemoryGame(data, 2020)

        return Result.success(result)
    }
}
