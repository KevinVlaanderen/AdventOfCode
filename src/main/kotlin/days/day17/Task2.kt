package days.day17

import framework.Task
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        var space = mapOf( // wzyx
            0 to mapOf(0 to input.openStream()
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
        )

        var bounds =
            listOf(Pair(0, 0), Pair(0, 0), Pair(0, space[0]!![0]!!.size - 1), Pair(0, space[0]!![0]!![0]!!.size - 1))

        val vectors: List<List<Int>> = (-1..1).flatMap { x ->
            (-1..1).flatMap { y ->
                (-1..1).flatMap { z ->
                    (-1..1).mapNotNull { w ->
                        if (x != 0 || y != 0 || z != 0 || w != 0) listOf(x, y, z, w) else null
                    }
                }
            }
        }

        repeat(6) {
            bounds = listOf(
                Pair(bounds[0].first - 1, bounds[0].second + 1),
                Pair(bounds[1].first - 1, bounds[1].second + 1),
                Pair(bounds[2].first - 1, bounds[2].second + 1),
                Pair(bounds[3].first - 1, bounds[3].second + 1),
            )

            space = (bounds[0].first..bounds[0].second).map { w ->
                w to (bounds[1].first..bounds[1].second).map { z ->
                    z to (bounds[2].first..bounds[2].second).map { y ->
                        y to (bounds[3].first..bounds[3].second).map { x ->
                            val current = space[w]?.get(z)?.get(y)?.get(x) ?: false

                            val activeNeighbourCount = vectors.count {
                                space[w + it[0]]?.get(z + it[1])?.get(y + it[2])?.get(x + it[3]) ?: false
                            }

                            x to when (current) {
                                true -> activeNeighbourCount == 2 || activeNeighbourCount == 3
                                false -> activeNeighbourCount == 3
                            }
                        }.associate { it.first to it.second }
                    }.associate { it.first to it.second }
                }.associate { it.first to it.second }
            }.associate { it.first to it.second }
        }

        val result = space.keys.sumOf { w ->
            space[w]!!.keys.sumOf { z ->
                space[w]!![z]!!.keys.sumOf { y ->
                    space[w]!![z]!![y]!!.keys.count { x ->
                        space[w]!![z]!![y]!![x]!!
                    }
                }
            }
        }

        return Result.success(result)
    }
}
