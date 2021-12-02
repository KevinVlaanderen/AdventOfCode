package days.day12

import framework.Task
import java.net.URL
import kotlin.math.absoluteValue

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var x = 0
        var y = 0
        var heading = 90

        input
            .openStream()
            .bufferedReader()
            .readLines()
            .forEach {
                val action = it[0]
                val value = it.substring(1).toInt()

                when (action) {
                    'N' -> y += value
                    'E' -> x += value
                    'S' -> y -= value
                    'W' -> x -= value
                    'L' -> heading = (heading - value + 360) % 360
                    'R' -> heading = (heading + value) % 360
                    'F' -> when (heading) {
                        0 -> y += value
                        90 -> x += value
                        180 -> y -= value
                        270 -> x -= value
                    }
                }
            }

        return Result.success(x.absoluteValue + y.absoluteValue)
    }
}
