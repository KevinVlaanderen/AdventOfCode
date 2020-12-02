package day1

import shared.FileUtil

class DayOne {
    fun partOne(input: List<Int>): Int? {
        for ((index, lower) in input.withIndex()) for (upper in input.drop(index+1)) {
            if (lower + upper == 2020) {
                return lower * upper
            }
        }

        return null
    }

    fun partTwo(input: List<Int>): Int? {
        for ((firstIndex, first) in input.withIndex()) for ((secondIndex, second) in input.drop(firstIndex+1).withIndex()) for (third in input.drop(secondIndex+1)) {
            if (first + second + third == 2020) {
                return first * second * third
            }
        }

        return null
    }
}

fun main() {
    val data = FileUtil().readResource("/input") ?: return
    val input = data.lines().filter { it != "" }.map { it.toInt() }

    val answerToPartOne = DayOne().partOne(input)
    println("Part 1: $answerToPartOne")

    val answerToPartTwo = DayOne().partTwo(input)
    println("Part 2: $answerToPartTwo")
}
