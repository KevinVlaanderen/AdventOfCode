package days.day18

import framework.Task
import java.net.URL

class Task2 : Task<Long>() {
    private val expressionPattern = """\(([^()]+)\)""".toRegex()

    override fun run(input: URL): Result<Long> {
        val result = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { line ->
                var current = line
                while (true) {
                    val (subExpression) = expressionPattern.find(current)?.destructured ?: break
                    current = expressionPattern.replaceFirst(current, solveExpression(subExpression))
                }

                solveExpression(current).toLong()
            }
            .sum()

        return Result.success(result)
    }

    private fun solveExpression(text: String): String {
        var current = text
        do {
            val previous = current
            current = listOf(
                Pair(addPattern.find(current), replaceAdd),
                Pair(multiplyPattern.find(current), replaceMultiply),
            )
                .firstOrNull { it.first != null }
                ?.let { it.second(current, it.first!!) } ?: current
        } while (current != previous)

        return current
    }

    private val replaceAdd = fun(text: String, matchResult: MatchResult): String {
        val (left, right) = matchResult.destructured
        return addPattern.replaceFirst(text, (left.toLong() + right.toLong()).toString())
    }

    private val replaceMultiply = fun(text: String, matchResult: MatchResult): String {
        val (left, right) = matchResult.destructured
        return multiplyPattern.replaceFirst(text, (left.toLong() * right.toLong()).toString())
    }
}