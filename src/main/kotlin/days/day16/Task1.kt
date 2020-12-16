package days.day16

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    private val rulePattern = """(.+): (.+)""".toRegex()

    private fun parseRule(input: String): TicketRule? {
        val (name, ranges) = rulePattern.find(input)?.destructured ?: return null
        return TicketRule(name, ranges.split("or").map {
            val values = it.trim().split("-")
            Pair(values[0].toInt(), values[1].toInt())
        })
    }

    private fun parseTicket(input: String): List<Int> = input.split(",").map { it.toInt() }

    override fun run(input: URL): Result<Int> {
        val rules = mutableListOf<TicketRule>()
        val nearbyTickets = mutableListOf<List<Int>>()

        input.openStream()
            .bufferedReader()
            .useLines { lines ->
                var currentSection = InputSection.RULES
                for (line in lines) {
                    when (line) {
                        "" -> continue
                        "your ticket:" -> currentSection = InputSection.YOUR_TICKET
                        "nearby tickets:" -> currentSection = InputSection.NEARBY_TICKETS
                        else -> {
                            when (currentSection) {
                                InputSection.RULES -> parseRule(line)?.let { rules.add(it) }
                                InputSection.YOUR_TICKET -> continue
                                InputSection.NEARBY_TICKETS -> parseTicket(line).let { nearbyTickets.add(it) }
                            }
                        }
                    }
                }
            }

        val result = nearbyTickets
            .flatMap { ticket ->
                ticket.filter { value ->
                    rules.none { rule -> rule.validInput(value) }
                }
            }
            .sum()

        return Result.success(result)
    }
}
