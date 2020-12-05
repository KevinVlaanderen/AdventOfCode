package days.day2

import framework.Task

object Task2 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val regex = """(\d+)-(\d+) ([a-z]): (.*)""".toRegex()

        val count = regex
            .findAll(input)
            .count { matchResult ->
                val groupValues = matchResult.groupValues
                val firstIndex = groupValues[1].toInt() - 1
                val secondIndex = groupValues[2].toInt() - 1
                val char = groupValues[3].toCharArray()[0]
                val password = groupValues[4]

                (password[firstIndex] == char).xor(password[secondIndex] == char)
            }

        return Result.success(count)
    }
}


