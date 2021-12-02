package days.day18

import framework.Task
import java.net.URL

class Task1 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val result = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { line ->
                var current = line
                do {
                    val previous = current
                    current = solve(current)
                } while (current != previous)

                current.toLong()
            }
            .sum()

        return Result.success(result)
    }

    private fun solve(text: String): String = listOf(
        Pair(addPattern.find(text), replaceAdd),
        Pair(multiplyPattern.find(text), replaceMultiply),
        Pair(expressionPattern.find(text), replaceExpression)
    )
        .filter { it.first != null }
        .minByOrNull { it.first!!.range.first }
        ?.let { it.second(text, it.first!!) } ?: text

    private val replaceAdd = fun(text: String, matchResult: MatchResult): String {
        val (left, right) = matchResult.destructured
        return addPattern.replaceFirst(text, (left.toLong() + right.toLong()).toString())
    }

    private val replaceMultiply = fun(text: String, matchResult: MatchResult): String {
        val (left, right) = matchResult.destructured
        return multiplyPattern.replaceFirst(text, (left.toLong() * right.toLong()).toString())
    }

    private val replaceExpression = fun(text: String, matchResult: MatchResult): String {
        val (value) = matchResult.destructured
        return expressionPattern.replaceFirst(text, value)
    }
}