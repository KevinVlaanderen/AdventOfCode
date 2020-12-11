package days.day11

import framework.Task
import java.net.URL

object Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var floor = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .map { it.toList() }
            .toList()

        val vectors = (-1..1).flatMap { x -> (-1..1).map { y -> Pair(x, y) } }.minus(Pair(0, 0))

        var previousFloor: List<List<Char>>
        do {
            previousFloor = floor
            floor = (floor.indices).map ymap@{ y ->
                (floor[y].indices).map xmap@{ x ->
                    val count by lazy { vectors.count { vector -> seatInSight(floor, x, y, vector) } }

                    when (floor[y][x]) {
                        'L' -> if (count == 0) '#' else 'L'
                        '#' -> if (count >= 5) 'L' else '#'
                        else -> floor[y][x]
                    }
                }
            }
        } while (previousFloor != floor)

        val result = floor.sumOf { line -> line.count { it == '#' } }

        return Result.success(result)
    }
}
