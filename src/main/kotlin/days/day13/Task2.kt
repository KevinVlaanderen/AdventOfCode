package days.day13

import framework.Task
import java.net.URL
import kotlin.math.pow

fun findMinX(num: IntArray, rem: IntArray, k: Int): Int {
    var x = 1
    while (true) {
        var j = 0
        while (j < k) {
            if (x % num[j] != rem[j]) break
            j++
        }

        if (j == k) return x else x++
    }
}

object Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val result = input.openStream()
            .bufferedReader()
            .readLines()
            .let { lines ->
                val busses = lines[1].split(',')

                var timeZero = System.nanoTime()
                var timestamp = 100000000000000L

                generateSequence(timestamp) { it + 1 }
                    .first { nextTimestamp ->
                        val timeDelta = System.nanoTime() - timeZero
                        if (timeDelta / 10.0.pow(9.0) >= 1) {
                            println("Processing ${nextTimestamp - timestamp} timestamps per second (current timestamp: $nextTimestamp)")
                            timeZero = System.nanoTime()
                            timestamp = nextTimestamp
                        }

                        busses.withIndex().all { bus ->
                            when (bus.value) {
                                "x" -> true
                                else -> {
                                    val busId = bus.value.toInt()
                                    val waitingTime = (nextTimestamp + bus.index) % busId
//                                    if (bus.index.toLong() > 0) waitingTime %= bus.index.toLong()
                                    waitingTime == 0L
                                }
                            }
                        }
                    }
            }


        return Result.success(result)
    }
}
