package days.day19

abstract class Condition {
    abstract fun matches(rules: Map<Int, Condition>, input: List<String>): List<String>?
}

class AndCondition(private val conditions: List<Condition>) : Condition() {
    override fun matches(rules: Map<Int, Condition>, input: List<String>): List<String>? = conditions
        .fold(input) { validInput, condition ->
            if (validInput.isEmpty()) validInput else condition.matches(
                rules,
                validInput
            ) ?: listOf()
        }
        .let { if (it.isEmpty()) null else it }
}

class OrCondition(private val conditions: List<Condition>) : Condition() {
    override fun matches(rules: Map<Int, Condition>, input: List<String>): List<String>? = conditions
        .flatMap { it.matches(rules, input) ?: listOf() }
        .let { if (it.isEmpty()) null else it }
}

class RuleReference(private val ruleNumber: Int) : Condition() {
    override fun matches(rules: Map<Int, Condition>, input: List<String>): List<String>? =
        rules[ruleNumber]!!.matches(rules, input)
}

class LiteralCondition(private val value: String) : Condition() {
    override fun matches(rules: Map<Int, Condition>, input: List<String>): List<String>? = input
        .filter { it.startsWith(value) }
        .map { it.substring(value.length) }
        .let { if (it.isEmpty()) null else it }
}