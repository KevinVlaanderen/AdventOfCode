package days.day19

fun fillInRules(rules: Map<Int, String>, current: String): String {
    val parsed = """\d+""".toRegex().replace(current) {
        val childIndex = it.value.toInt()
        fillInRules(rules, rules[childIndex]!!)
    }
    return when (current.contains("|")) {
        true -> "(?:$parsed)"
        false -> parsed
    }
}