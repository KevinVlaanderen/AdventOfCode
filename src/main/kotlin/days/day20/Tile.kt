package days.day20

import shared.toLines

data class Tile(val id: Int, val sides: Map<Direction, Side>, val pattern: List<List<Char>>) {
    override fun toString(): String = id.toString()

    class Side(val pattern: String)

    companion object {
        private val titlePattern = """Tile (\d+):""".toRegex()

        fun parse(input: String): Tile {
            val lines = input.toLines()
            val (id) = titlePattern.matchEntire(lines.first())!!.destructured

            val data = lines.drop(1).map { line -> line.map { it } }.toList()

            val sides = mapOf(
                Direction.NORTH to Side(data.first().joinToString("")),
                Direction.EAST to Side(data.map { it.last() }.joinToString("")),
                Direction.SOUTH to Side(data.last().joinToString("").reversed()),
                Direction.WEST to Side(data.map { it.first() }.joinToString("").reversed())
            )

            val pattern = data.drop(1).dropLast(1).map { it.drop(1).dropLast(1) }

            return Tile(id.toInt(), sides, pattern)
        }
    }
}