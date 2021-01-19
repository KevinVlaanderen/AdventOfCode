package days.day24

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val result = input.openStream().bufferedReader().lineSequence()
            .map { followSteps(generateSteps(it)) }
            .groupBy { Tile(it.x, it.y) }
            .filter { it.value.count() % 2 != 0 }
            .map { it.key }
            .count()

        return Result.success(result)
    }


    private fun generateSteps(line: String): Sequence<Direction> {
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

    private fun followSteps(steps: Sequence<Direction>): Tile = steps.fold(Tile(0, 0)) { location, step ->
        when (step) {
            Direction.NW -> Tile(if (location.y % 2 == 0) location.x else location.x - 1, location.y - 1)
            Direction.NE -> Tile(if (location.y % 2 == 0) location.x + 1 else location.x, location.y - 1)
            Direction.E -> Tile(location.x + 1, location.y)
            Direction.SE -> Tile(if (location.y % 2 == 0) location.x + 1 else location.x, location.y + 1)
            Direction.SW -> Tile(if (location.y % 2 == 0) location.x else location.x - 1, location.y + 1)
            Direction.W -> Tile(location.x - 1, location.y)
        }
    }
}