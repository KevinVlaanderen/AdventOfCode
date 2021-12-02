package days.day16

class TicketRule(private val acceptableRanges: List<Pair<Int, Int>>) {

    fun validInput(input: Int) = acceptableRanges.any { input in (it.first..it.second) }

    companion object {
        private val rulePattern = """(.+): (.+)""".toRegex()

        fun parse(input: String): Pair<String, TicketRule>? {
            val (name, ranges) = rulePattern.find(input)?.destructured ?: return null
            return Pair(name, TicketRule(ranges.split("or").map {
                val values = it.trim().split("-")
                Pair(values[0].toInt(), values[1].toInt())
            }))
        }
    }
}