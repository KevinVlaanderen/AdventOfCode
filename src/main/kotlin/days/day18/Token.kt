package days.day18

enum class Token(val pattern: Regex) {
    WHITESPACE("""\s+""".toRegex()),
    LEFT_PARENTHESIS("""\(""".toRegex()),
    RIGHT_PARENTHESIS("""\)""".toRegex()),
    MULTIPLY("""\*""".toRegex()),
    ADD("""\+""".toRegex()),
    NUMBER("""\d+""".toRegex());

    companion object {
        fun matchNextToken(text: String, startIndex: Int = 0): Pair<Token, MatchResult> {
            for (token in values()) {
                val match = token.pattern.find(text, startIndex)
                if (match == null || match.range.first > startIndex)
                    continue

                return Pair(token, match)
            }

            throw Exception("Failed to match token")
        }
    }
}