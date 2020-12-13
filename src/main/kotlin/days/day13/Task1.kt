package days.day13

import framework.Task
import java.net.URL

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input.openStream()
            .bufferedReader()
            .readLines()
            .let { lines ->
                val target = lines[0].toInt()
                val busses = lines[1].split(',').filterNot { it == "x" }.map { it.toInt() }

                busses
                    .map { Pair(it, it - (target % it)) }
                    .minByOrNull { it.second }!!
                    .let { it.first * it.second }
            }


        return Result.success(result)
    }
}
