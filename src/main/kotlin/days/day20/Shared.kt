package days.day20

fun calculateAvailableConnections(tiles: List<Tile>) =
    tiles.windowed(tiles.size, 1, true).flatMap { window ->
        val currentTile = window.first()

        currentTile.sides.flatMap { currentSide ->
            window.drop(1).flatMap { otherTile ->
                otherTile.sides.mapNotNull { otherTileSide ->
                    when (currentSide.value.pattern) {
                        otherTileSide.value.pattern -> Connection(
                            Pair(currentTile, currentSide.key),
                            Pair(otherTile, otherTileSide.key),
                            true
                        )
                        otherTileSide.value.pattern.reversed() -> Connection(
                            Pair(currentTile, currentSide.key),
                            Pair(otherTile, otherTileSide.key),
                            false
                        )
                        else -> null
                    }
                }
            }
        }
    }