package days.day23

import framework.Task
import java.net.URL

class Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val CUPS = 1000000
        val REPETITIONS = 10000000

        val cupsFromInput = input.openStream().bufferedReader().readLine()!!.toList()
            .map { Character.getNumericValue(it) }
            .asSequence()
        val additionalCups = generateSequence(cupsFromInput.maxOrNull()!! + 1) { it + 1 }
            .take(CUPS - cupsFromInput.count())
        val allCups = cupsFromInput + additionalCups

        val cups = allCups
            .windowed(2, 1)
            .plus(listOf(listOf(allCups.last(), allCups.first())))
            .associate { it.first() to it.last() }
            .toMutableMap()

        var currentCup = cupsFromInput.first()

        val lowestCup = allCups.minOrNull()!!
        val highestCup = allCups.maxOrNull()!!

        repeat(REPETITIONS) {
            val removedCups = generateSequence(cups[currentCup]!!) { cups[it]!! }.take(3)
            cups[currentCup] = cups[removedCups.last()]!!

            val destinationCup = generateSequence(if (currentCup == lowestCup) highestCup else currentCup - 1) {
                if (it == lowestCup) highestCup else it - 1
            }.first { it !in removedCups }

            cups[removedCups.last()] = cups[destinationCup]!!
            cups[destinationCup] = removedCups.first()

            currentCup = cups[currentCup]!!
        }

        val firstResultCup = cups[1]!!
        val secondResultCup = cups[firstResultCup]!!
        val result = firstResultCup.toLong() * secondResultCup.toLong()

        return Result.success(result)
    }
}