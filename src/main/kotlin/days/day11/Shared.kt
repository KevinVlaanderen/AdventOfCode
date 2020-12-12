package days.day11

import shared.transformUntilUnchanged

private val vectors = (-1..1).flatMap { x -> (-1..1).map { y -> Pair(x, y) } }.minus(Pair(0, 0))

fun seatsWhenStabilized(floor: List<List<Char>>, minSeatsBeforeVacate: Int, maxDistance: Int? = null): Int =
    floor.transformUntilUnchanged {
        (it.indices).map ymap@{ y ->
            (it[y].indices).map xmap@{ x ->
                val count by lazy { vectors.count { vector -> seatInSight(it, x, y, vector, maxDistance) } }

                when (it[y][x]) {
                    'L' -> if (count == 0) '#' else 'L'
                    '#' -> if (count >= minSeatsBeforeVacate) 'L' else '#'
                    else -> it[y][x]
                }
            }
        }
    }.sumOf { line -> line.count { it == '#' } }

private fun seatInSight(
    floor: List<List<Char>>,
    x: Int,
    y: Int,
    vector: Pair<Int, Int>,
    maxSteps: Int? = null
): Boolean {
    var newX = x
    var newY = y
    var steps = 0

    while (maxSteps == null || maxSteps > steps++) {
        newX += vector.first
        newY += vector.second

        if (newY !in floor.indices || newX !in floor[newY].indices) return false
        when (floor[newY][newX]) {
            'L' -> return false
            '#' -> return true
        }
    }

    return false
}