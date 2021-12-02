package days.day20

import shared.flip
import shared.mirrorVertical
import shared.rotateCCW
import shared.rotateCW

data class Placement(val tile: Tile, val rotation: Rotation, val inverted: Boolean) {
    fun sideFacing(direction: Direction): Direction =
        when (val side = Direction.fromDegrees(direction.degrees + rotation.degrees)) {
            Direction.EAST -> if (inverted) Direction.WEST else Direction.EAST
            Direction.WEST -> if (inverted) Direction.EAST else Direction.WEST
            else -> side
        }

    fun pattern(): List<List<Char>> = tile.pattern
        .let { if (inverted) it.mirrorVertical() else it }
        .let {
            when (rotation) {
                Rotation.R0 -> it
                Rotation.R90 -> it.rotateCCW()
                Rotation.R180 -> it.flip()
                Rotation.R270 -> it.rotateCW()
            }
        }
}
