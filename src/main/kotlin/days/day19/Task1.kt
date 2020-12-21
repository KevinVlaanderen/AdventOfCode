package days.day19

import framework.Task
import shared.toBlocks
import shared.toLines
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val (rules, messages) = input.openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .toList()
            .let { blocks ->
                val (rulesBlock, messagesBlock) = blocks.take(2)
                val rules = rulesBlock.toLines().associate { line ->
                    val (index, pattern) = line.split(':').let { Pair(it[0].toInt(), it[1]) }
                    index to "(?:${pattern.replace("\"", "")})"
                }
                Pair(rules, messagesBlock)
            }

        val pattern = "^${fillInRules(rules, rules[0]!!).replace(" ", "")}$"

        val result = pattern.toRegex(RegexOption.MULTILINE).findAll(messages).count()

        return Result.success(result)
    }

    private fun fillInRules(rules: Map<Int, String>, current: String): String = """\d+""".toRegex().replace(current) {
        val childIndex = it.value.toInt()
        fillInRules(rules, rules[childIndex]!!)
    }
}