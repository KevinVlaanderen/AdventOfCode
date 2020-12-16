package days.day16

class TicketRule(val name: String, val acceptableRanges: List<Pair<Int, Int>>) {
    fun validInput(input: Int) = acceptableRanges.any { input in (it.first..it.second) }
}