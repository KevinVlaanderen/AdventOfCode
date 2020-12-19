package days.day17

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var space = mapOf(0 to input.openStream() // zyx
            .bufferedReader()
            .useLines { lines ->
                lines.mapIndexed { lineIndex, line ->
                    lineIndex to line.mapIndexed { colIndex, col ->
                        colIndex to when (col) {
                            '#' -> true
                            else -> false
                        }
                    }.associate { it.first to it.second }
                }.associate { it.first to it.second }
            })

        var bounds = Triple(Pair(0, 0), Pair(0, space[0]!!.size - 1), Pair(0, space[0]!![0]!!.size - 1))

        val vectors = (-1..1).flatMap { x ->
            (-1..1).flatMap { y ->
                (-1..1).map { z ->
                    Triple(x, y, z)
                }
            }
        }.minus(Triple(0, 0, 0))

        repeat(6) {
            bounds = Triple(
                Pair(bounds.first.first - 1, bounds.first.second + 1),
                Pair(bounds.second.first - 1, bounds.second.second + 1),
                Pair(bounds.third.first - 1, bounds.third.second + 1),
            )

            space = (bounds.first.first..bounds.first.second).map { z ->
                z to (bounds.second.first..bounds.second.second).map { y ->
                    y to (bounds.third.first..bounds.third.second).map { x ->
                        val current = space[z]?.get(y)?.get(x) ?: false

                        val activeNeighbourCount = vectors.count {
                            space[z + it.first]?.get(y + it.second)?.get(x + it.third) ?: false
                        }

                        x to when (current) {
                            true -> activeNeighbourCount == 2 || activeNeighbourCount == 3
                            false -> activeNeighbourCount == 3
                        }
                    }.associate { it.first to it.second }
                }.associate { it.first to it.second }
            }.associate { it.first to it.second }
        }

        val result = space.keys.sumOf { z ->
            space[z]!!.keys.sumOf { y ->
                space[z]!![y]!!.keys.count { x ->
                    space[z]!![y]!![x]!!
                }
            }
        }

        return Result.success(result)
    }
}
