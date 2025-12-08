package days.day24

data class Tile(val location: Pair<Int, Int>) {
    private var _side = Side.WHITE

    val side: Side
        get() = _side

    fun flip() {
        _side = if (_side == Side.WHITE) Side.BLACK else Side.WHITE
    }
}
