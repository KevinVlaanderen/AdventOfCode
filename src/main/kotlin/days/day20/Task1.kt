package days.day20

import framework.Task
import shared.toBlocks
import java.net.URL

class Task1 : Task<Long>() {

    override fun run(input: URL): Result<Long> {
        val tiles = input.openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .map { Tile.parse(it) }
            .toList()

        val availableConnections = calculateAvailableConnections(tiles)

        val corners = tiles
            .filter { tile ->
                availableConnections.count { it.from.first == tile || it.to.first == tile } == 2
            }

        val result = corners.map { it.id.toLong() }.reduce { acc, i -> acc * i }

        return Result.success(result)
    }

    private fun calculateAvailableConnections(tiles: List<Tile>) =
        tiles.windowed(tiles.size, 1, true).flatMap { window ->
            val currentTile = window.first()

            currentTile.sides.flatMap { currentSide ->
                window.drop(1).flatMap { otherTile ->
                    otherTile.sides.mapNotNull { otherTileSide ->
                        when (currentSide.value) {
                            otherTileSide.value -> Connection(
                                Pair(currentTile, currentSide.key),
                                Pair(otherTile, otherTileSide.key),
                                false
                            )
                            otherTileSide.value.reversed() -> Connection(
                                Pair(currentTile, currentSide.key),
                                Pair(otherTile, otherTileSide.key),
                                true
                            )
                            else -> null
                        }
                    }
                }
            }
        }
}