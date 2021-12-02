package days.day14

import framework.Task
import shared.cartesianProduct
import java.net.URL
import java.util.*

class Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val targetMask = BitSet(36)
        val register: MutableMap<Long, Long> = mutableMapOf()
        var wildcardIndices: List<Int> = listOf()
        var wildCards: List<List<Boolean>> = listOf(listOf())

        input.openStream()
            .bufferedReader()
            .readText()
            .let { linePattern.findAll(it) }
            .forEach { match ->
                val (target, index, value) = match.destructured
                when (target) {
                    "mask" -> {
                        val wildcardMask = BitSet(36)

                        value.reversed().forEachIndexed { i, c ->
                            targetMask.set(i, c == '1')
                            wildcardMask.set(i, c == 'X')

                            wildcardIndices = wildcardMask.stream().toArray().toList()
                            wildCards =
                                wildcardIndices.indices.fold(listOf()) { acc, _ ->
                                    if (acc.isEmpty())
                                        listOf(listOf(true), listOf(false))
                                    else acc.cartesianProduct(
                                        listOf(true, false)
                                    ).map { it.first + listOf(it.second) }
                                }
                        }
                    }
                    "mem" -> {
                        val targetBitSet = BitSet.valueOf(LongArray(1) { index.toLong() });
                        targetBitSet.or(targetMask)

                        wildCards.forEach { wildcard ->
                            wildcardIndices.forEachIndexed { index, i -> targetBitSet.set(i, wildcard[index]) }
                            register[targetBitSet.toLongArray()[0]] = value.toLong()
                        }

                    }
                }
            }

        val result = register.values.map { it }.sum()

        return Result.success(result)
    }
}
