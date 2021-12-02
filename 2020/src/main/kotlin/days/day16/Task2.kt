package days.day16

import framework.Task
import java.net.URL

class Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val (rules, yourTicket, nearbyTickets) = input.openStream()
            .bufferedReader()
            .useLines { lines -> readInput(lines) }

        val validTickets = nearbyTickets
            .filter { ticket ->
                ticket.values.all { value ->
                    rules.any { rule -> rule.value.validInput(value) }
                }
            }

        val mapping: Map<String, MutableList<Int>> = rules.keys.associateWith { mutableListOf() }

        yourTicket.values.indices.forEach { valueIndex ->
            var validRuleNames: List<String> = mapping.keys.toList()

            val ticketIterator = validTickets.iterator()
            while (validRuleNames.size > 1 && ticketIterator.hasNext()) {
                val ticket = ticketIterator.next()
                validRuleNames =
                    validRuleNames.filter { ruleName -> rules[ruleName]!!.validInput(ticket.values[valueIndex]) }
            }

            validRuleNames.forEach { mapping[it]!!.add(valueIndex) }
        }

        while (true) {
            val validMappings = mapping.filter { it.value.count() == 1 }
            if (validMappings.count() == rules.count()) break
            mapping.minus(validMappings.keys).forEach { (key, _) ->
                mapping[key]!!.removeAll(validMappings.flatMap { it.value })
            }
        }

        val result = mapping
            .filter { it.key.startsWith("departure") }
            .map { yourTicket.values[it.value.first()] }
            .fold(1) { acc: Long, i: Int -> acc * i }

        return Result.success(result)
    }
}
