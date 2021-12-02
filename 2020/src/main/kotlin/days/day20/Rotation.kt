package days.day20

enum class Rotation(val degrees: Int) {
    R0(0), R90(90), R180(180), R270(270);

    companion object {
        fun withSide(side: Direction, facing: Direction, inverted: Boolean): Rotation {
            val originDegrees = when (side) {
                Direction.EAST -> if (inverted) R270 else R90
                Direction.WEST -> if (inverted) R90 else R270
                else -> fromDegrees(side.degrees)
            }

            return fromDegrees(originDegrees.degrees - facing.degrees)
        }

        private fun fromDegrees(degrees: Int): Rotation =
            values().find { it.degrees == if (degrees < 0) degrees % 360 + 360 else degrees % 360 }
                ?: throw IllegalArgumentException()
    }
}