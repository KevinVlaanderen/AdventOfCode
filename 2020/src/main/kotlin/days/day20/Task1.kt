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
}