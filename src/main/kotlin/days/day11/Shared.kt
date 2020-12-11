package days.day11

fun seatInSight(floor: List<List<Char>>, x: Int, y: Int, vector: Pair<Int, Int>, maxSteps: Int? = null): Boolean {
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