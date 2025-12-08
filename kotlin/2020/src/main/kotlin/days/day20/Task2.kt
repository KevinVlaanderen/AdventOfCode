package days.day20

import framework.Task
import framework.exceptions.AnswerNotFoundException
import shared.*
import java.net.URL

class Task2 : Task<Int>() {

    override fun run(input: URL): Result<Int> {
        val tiles = input.openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .map { Tile.parse(it) }
            .toList()

        val connections = calculateAvailableConnections(tiles).let { connections ->
            connections + connections.map { Connection(it.to, it.from, it.inverted) }
        }

        val startTilePlacement = tiles
            .first { tile ->
                connections.count { it.from.first == tile } == 2 &&
                        connections.find { it.from.first == tile && it.from.second == Direction.EAST } != null &&
                        connections.find { it.from.first == tile && it.from.second == Direction.SOUTH } != null
            }.let { Placement(it, Rotation.R0, false) }

        val board = collectTilePlacements(connections, startTilePlacement, Direction.SOUTH)
            .map { collectTilePlacements(connections, it, Direction.EAST) }
            .flatMap {
                it.map { placement -> placement.pattern() }.reduce { acc, col ->
                    acc.mapIndexed { index, line -> line + col[index] }
                }
            }.toList()

        val monster = """
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
""".lines().filter { it.isNotEmpty() }.map { it.toList() }

        for (rotation in listOf(Rotation.R0, Rotation.R90, Rotation.R180, Rotation.R270))
            for (inverted in listOf(false, true)) {
                val rotatedMonster = monster
                    .let { if (inverted) it.mirrorVertical() else it }
                    .let {
                        when (rotation) {
                            Rotation.R0 -> it
                            Rotation.R90 -> it.rotateCCW()
                            Rotation.R180 -> it.flip()
                            Rotation.R270 -> it.rotateCW()
                        }
                    }

                val numberOfSeaMonsters = countPattern(board, rotatedMonster)

                if (numberOfSeaMonsters > 0) {
                    val hashesOnBoard = board.count('#')
                    val hashesPerMonster = monster.count('#')
                    val result = hashesOnBoard - (numberOfSeaMonsters * hashesPerMonster)

                    return Result.success(result)
                }
            }

        throw AnswerNotFoundException()
    }

    private fun countPattern(searchArea: List<List<Char>>, pattern: List<List<Char>>): Int {
        val searchAreaWidth = searchArea[0].count()

        val patternWidth = pattern[0].count()
        val patternHeight = pattern.count()

        return searchArea.windowed(patternHeight).sumOf { rows ->
            (0 until searchAreaWidth - patternWidth).count { startIndex ->
                val region = rows.map { it.subList(startIndex, startIndex + patternWidth) }

                for (y in pattern.indices) for (x in pattern[y].indices) {
                    if (pattern[y][x] == '#' && region[y][x] != '#') return@count false
                }

                return@count true
            }
        }
    }

    private fun collectTilePlacements(connections: List<Connection>, start: Placement, direction: Direction) =
        sequence<Placement> {
            var current: Placement? = start

            while (current != null) {
                yield(current)

                val (fromTile, _, fromInverted) = current
                val fromSide = current.sideFacing(direction)

                current = connections.find { it.from.first == fromTile && it.from.second == fromSide }?.let {
                    val (_, to, connectionInverted) = it
                    val (toTile, toSide) = to
                    val toInverted = if (connectionInverted) !fromInverted else fromInverted
                    val toRotation = Rotation.withSide(toSide, direction.opposite(), toInverted)

                    Placement(toTile, toRotation, toInverted)
                }
            }
        }
}