package days.day2

import framework.Task

object Task1 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val regex = """(\d+)-(\d+) ([a-z]): (.*)""".toRegex()

        val count = regex
            .findAll(input)
            .count { matchResult ->
                val groupValues = matchResult.groupValues
                val min = groupValues[1].toInt()
                val max = groupValues[2].toInt()
                val char = groupValues[3].toCharArray()[0]
                val password = groupValues[4]

                val num = password.count { it == char }

                num in min..max
            }

        return Result.success(count)
    }
}


