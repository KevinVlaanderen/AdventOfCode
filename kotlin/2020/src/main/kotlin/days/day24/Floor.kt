package days.day24

class Floor {
    private val rows = mutableMapOf<Int, Row>()

    val tiles
        get() = rows.values.flatMap { it.tiles }

    fun tile(location: Pair<Int, Int>): Tile {
        if (rows[location.second] == null) rows[location.second] = Row(location.second)
        return rows[location.second]!!.tile(location.first)
    }

    fun flip(location: Pair<Int, Int>) {
        Direction.values().forEach { direction -> tile(calculateNewLocation(location, direction)) }

        tile(location).flip()
    }

    inner class Row(val y: Int) {
        private val columns = mutableMapOf<Int, Tile>()

        val tiles
            get() = columns.values

        internal fun tile(x: Int): Tile {
            if (columns[x] == null) columns[x] = Tile(Pair(x, y))
            return columns[x]!!
        }
    }
}
