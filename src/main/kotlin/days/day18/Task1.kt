package days.day18

import framework.Task
import java.net.URL

class Task1 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val result = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { lex(it).filter { token -> token.first != Token.WHITESPACE }.iterator() }
            .map { parseExpression(it) }
            .map { execute(it) }
            .sum()


        return Result.success(result)
    }

    private fun lex(line: String): Sequence<Pair<Token, String>> = sequence {
        var index = 0
        while (index < line.length) {
            val (token, match) = Token.matchNextToken(line, index)
            index = match.range.last + 1
            yield(Pair(token, match.value))
        }
    }

    private fun parseExpression(tokens: Iterator<Pair<Token, String>>, subExpression: Boolean = false): Expression {
        var result: Expression? = null

        while (tokens.hasNext()) {
            val token = tokens.next()

            result = when (token.first) {
                Token.LEFT_PARENTHESIS -> parseExpression(tokens, true)
                Token.RIGHT_PARENTHESIS -> if (result != null) return result else throw Exception("Unexpected ')'")
                Token.NUMBER -> NumberExpression(token.second.toLong())
                Token.ADD -> parseAddExpression(result, tokens)
                Token.MULTIPLY -> {
                    parseMultiplyExpression(result, tokens)
                }
                else -> throw Exception("Unexpected token ${token.first.name}")
            }
        }

        if (subExpression) throw Exception("Expected ')'")

        if (result == null) throw Exception("Expected expression") else return result
    }

    private fun parseAddExpression(leftHand: Expression?, tokens: Iterator<Pair<Token, String>>): Expression {
        if (leftHand == null) throw Exception("Unexpected '+'; expected expression")

        val nextToken = if (tokens.hasNext()) tokens.next() else throw Exception("Expected expression")

        return when (nextToken.first) {
            Token.NUMBER -> AddExpression(leftHand, NumberExpression(nextToken.second.toLong()))
            Token.LEFT_PARENTHESIS -> AddExpression(leftHand, parseExpression(tokens, true))
            else -> throw Exception("Expected expression")
        }
    }

    private fun parseMultiplyExpression(leftHand: Expression?, tokens: Iterator<Pair<Token, String>>): Expression {
        if (leftHand == null) throw Exception("Unexpected '+'; expected expression")

        val nextToken = if (tokens.hasNext()) tokens.next() else throw Exception("Expected expression")

        return when (nextToken.first) {
            Token.NUMBER -> MultiplyExpression(leftHand, NumberExpression(nextToken.second.toLong()))
            Token.LEFT_PARENTHESIS -> MultiplyExpression(leftHand, parseExpression(tokens, true))
            else -> throw Exception("Expected expression")
        }
    }

    private fun execute(expression: Expression) = expression.evaluate()
}
