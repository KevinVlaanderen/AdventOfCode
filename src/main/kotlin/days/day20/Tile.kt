package days.day20

import shared.toLines

data class Tile(val id: Int, val data: List<List<Char>>) {
    val sides = mapOf(
        Side.NORTH to data.first().joinToString(""),
        Side.EAST to data.map { it.last() }.joinToString(""),
        Side.SOUTH to data.last().joinToString(""),
        Side.WEST to data.map { it.first() }.joinToString("")
    )

    override fun toString(): String = id.toString()

    companion object {
        private val titlePattern = """Tile (\d+):""".toRegex()

        fun parse(input: String): Tile {
            val lines = input.toLines()
            val (id) = titlePattern.matchEntire(lines.first())!!.destructured

            val data = lines.drop(1).map { line -> line.map { it } }.toList()

            return Tile(id.toInt(), data)
        }
    }
}