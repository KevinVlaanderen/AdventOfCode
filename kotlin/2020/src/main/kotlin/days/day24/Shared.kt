package days.day24

fun generateSteps(line: String): Sequence<Direction> {
    val lineIterator = line.iterator()
    return sequence {
        while (lineIterator.hasNext()) {
            when (val first = lineIterator.next()) {
                'e' -> yield(Direction.E)
                'w' -> yield(Direction.W)
                else -> yield(
                    when (lineIterator.next()) {
                        'e' -> if (first == 'n') Direction.NE else Direction.SE
                        'w' -> if (first == 'n') Direction.NW else Direction.SW
                        else -> continue
                    }
                )
            }
        }
    }
}

fun followSteps(steps: Sequence<Direction>): Pair<Int, Int> = steps.fold(Pair(0, 0)) { location, direction ->
    calculateNewLocation(location, direction)
}

fun calculateNewLocation(location: Pair<Int, Int>, direction: Direction) = when (direction) {
    Direction.NW -> Pair(if (location.second % 2 == 0) location.first else location.first - 1, location.second - 1)
    Direction.NE -> Pair(if (location.second % 2 == 0) location.first + 1 else location.first, location.second - 1)
    Direction.E -> Pair(location.first + 1, location.second)
    Direction.SE -> Pair(if (location.second % 2 == 0) location.first + 1 else location.first, location.second + 1)
    Direction.SW -> Pair(if (location.second % 2 == 0) location.first else location.first - 1, location.second + 1)
    Direction.W -> Pair(location.first - 1, location.second)
}