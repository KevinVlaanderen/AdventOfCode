package days.day16

class Ticket(val values: List<Int>) {
    companion object {
        fun parse(input: String): Ticket = Ticket(input.split(",").map { it.toInt() })
    }
}