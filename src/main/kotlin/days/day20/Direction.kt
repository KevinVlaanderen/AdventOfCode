package days.day20

enum class Direction(val degrees: Int) {
    NORTH(0), EAST(90), SOUTH(180), WEST(270);

    fun opposite(): Direction = fromDegrees(this.degrees + 180)

    companion object {
        fun fromDegrees(degrees: Int): Direction =
            values().find { it.degrees == if (degrees < 0) degrees % 360 + 360 else degrees % 360 }
                ?: throw IllegalArgumentException()
    }
}