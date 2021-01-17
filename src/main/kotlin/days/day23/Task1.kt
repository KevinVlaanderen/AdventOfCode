package days.day23

import framework.Task
import java.net.URL
import java.util.*

class Task1 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val cups: LinkedList<Int> =
            LinkedList(input.openStream().bufferedReader().readLine()!!.toList().map { Character.getNumericValue(it) })

        val numberOfCups = cups.count()
        val lowestCupValue = cups.minOrNull()!!
        val highestCupValue = cups.maxOrNull()!!

        var currentCupValue = cups[0]

        repeat(100) {
            val removedCups = (0..2).map { cups.removeAt((cups.indexOf(currentCupValue) + 1) % (numberOfCups - it)) }

            var destinationCupValue = if (currentCupValue == lowestCupValue) highestCupValue else currentCupValue - 1
            while (destinationCupValue in removedCups || destinationCupValue !in cups) {
                destinationCupValue =
                    if (destinationCupValue == lowestCupValue) highestCupValue else destinationCupValue - 1
            }

            val destinationCupIndex = cups.indexOf(destinationCupValue)

            cups.addAll(destinationCupIndex + 1, removedCups)

            currentCupValue = cups[(cups.indexOf(currentCupValue) + 1) % numberOfCups]
        }

        val result = generateSequence(cups.indexOf(1)) { (it + 1) % numberOfCups }
            .drop(1)
            .take(numberOfCups - 1)
            .joinToString("") { cups[it].toString() }.toLong()

        return Result.success(result)
    }
}