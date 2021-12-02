package days.day16

fun readInput(lines: Sequence<String>): Triple<Map<String, TicketRule>, Ticket, List<Ticket>> {
    val rules = mutableMapOf<String, TicketRule>()
    var yourTicket: Ticket? = null
    val nearbyTickets = mutableListOf<Ticket>()

    var currentSection = InputSection.RULES
    for (line in lines) {
        when (line) {
            "" -> continue
            "your ticket:" -> currentSection = InputSection.YOUR_TICKET
            "nearby tickets:" -> currentSection = InputSection.NEARBY_TICKETS
            else -> {
                when (currentSection) {
                    InputSection.RULES -> TicketRule.parse(line)?.let { rules[it.first] = it.second }
                    InputSection.YOUR_TICKET -> yourTicket = Ticket.parse(line)
                    InputSection.NEARBY_TICKETS -> Ticket.parse(line).let { nearbyTickets.add(it) }
                }
            }
        }
    }

    return Triple(rules, yourTicket!!, nearbyTickets)
}