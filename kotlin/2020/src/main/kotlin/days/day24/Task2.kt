package days.day24

import framework.Task
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val floor = Floor()

        input.openStream().bufferedReader().lineSequence()
            .map { followSteps(generateSteps(it)) }
            .forEach { floor.flip(it) }

        repeat(100) {
            floor.tiles
                .filter { tile ->
                    val blackNeighbors = Direction.values().count { direction ->
                        val neighbor = calculateNewLocation(tile.location, direction)
                        floor.tile(neighbor).side == Side.BLACK
                    }

                    when (tile.side) {
                        Side.BLACK -> blackNeighbors == 0 || blackNeighbors > 2
                        Side.WHITE -> blackNeighbors == 2
                    }
                }.forEach { it.flip() }
        }

        val result = floor.tiles.count { it.side == Side.BLACK }

        return Result.success(result)
    }
}