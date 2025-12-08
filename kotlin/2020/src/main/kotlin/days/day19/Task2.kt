package days.day19

import framework.Task
import shared.toBlocks
import shared.toLines
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val blocks = input.openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .take(2)
            .map { it.toLines() }
            .toList()

        val rules = blocks[0]
            .associate { line ->
                val (index, pattern) = line.split(':').let { Pair(it[0].toInt(), it[1]) }

                when (index) {
                    8 -> index to "42 | 42 8"
                    11 -> index to "42 31 | 42 11 31"
                    else -> index to pattern//.replace("\"", "")
                }
            }
            .mapValues { (_, rule) ->
                val alternatives = rule.split("|").map { alternative ->
                    AndCondition(alternative.trim().split(" ").map { part ->
                        when {
                            part.contains("\"") -> LiteralCondition(part.replace("\"", ""))
                            else -> RuleReference(part.toInt())
                        }
                    })
                }

                when (alternatives.size) {
                    1 -> alternatives[0]
                    else -> OrCondition(alternatives)
                }
            }

        val messages = blocks[1]

        val result = messages.count { message ->
            val results = rules[0]!!.matches(rules, listOf(message))
            results != null && results.any { it == "" }
        }

        return Result.success(result)
    }
}