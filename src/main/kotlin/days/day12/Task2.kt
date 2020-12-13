package days.day12

import framework.Task
import java.net.URL
import kotlin.math.*

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var shipX = 0
        var shipY = 0
        var waypointX = 10
        var waypointY = 1

        input
            .openStream()
            .bufferedReader()
            .readLines()
            .forEach {
                val action = it[0]
                val value = it.substring(1).toInt()

                when (action) {
                    'N' -> waypointY += value
                    'E' -> waypointX += value
                    'S' -> waypointY -= value
                    'W' -> waypointX -= value
                    'L', 'R' -> {
                        val radius = hypot(waypointX.toDouble(), waypointY.toDouble())
                        val theta = atan2(waypointY.toDouble(), waypointX.toDouble())
                        val angle = Math.toDegrees(theta)
                        val newAngle = when (action) {
                            'L' -> (angle + value) % 360
                            'R' -> (angle - value + 360) % 360
                            else -> angle
                        }

                        waypointX = (cos(Math.toRadians(newAngle)) * radius).roundToInt()
                        waypointY = (sin(Math.toRadians(newAngle)) * radius).roundToInt()
                    }
                    'F' -> {
                        shipX += waypointX * value
                        shipY += waypointY * value
                    }
                }
            }

        return Result.success(shipX.absoluteValue + shipY.absoluteValue)
    }
}
