package days.day14

import framework.Task
import shared.to36bitString
import shared.toLong
import java.net.URL
import java.util.*

class Task1 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val mask = BitSet(36)
        val maskValue = BitSet(36)
        val register: MutableMap<Int, BitSet> = mutableMapOf()

        input.openStream()
            .bufferedReader()
            .readText()
            .let { linePattern.findAll(it) }
            .forEach { match ->
                val (target, index, value) = match.destructured
                when (target) {
                    "mask" -> value.reversed().forEachIndexed { i, c ->
                        run {
                            mask.set(i, c != 'X')
                            maskValue.set(i, c == '1')
                        }
                    }
                    "mem" -> {
                        val valueBitSet = BitSet(36)
                        value.toInt().to36bitString().reversed().forEachIndexed { i, c ->
                            run {
                                valueBitSet.set(i, if (mask.get(i)) maskValue[i] else c == '1')
                            }
                        }
                        register[index.toInt()] = valueBitSet
                    }
                }
            }

        val result = register.values.map { it.toLong() }.sum()

        return Result.success(result)
    }
}
