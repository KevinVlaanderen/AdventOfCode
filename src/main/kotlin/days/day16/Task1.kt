package days.day16

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val (rules, _, nearbyTickets) = input.openStream()
            .bufferedReader()
            .useLines { lines -> readInput(lines) }

        val invalidTickets = nearbyTickets
            .flatMap { ticket ->
                ticket.values.filter { value ->
                    rules.none { rule -> rule.value.validInput(value) }
                }
            }

        val result = invalidTickets.sum()

        return Result.success(result)
    }
}
