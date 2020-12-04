package days.day2

import framework.Task
import shared.readResource

class Task2 : Task<Int>(2) {
    override fun run(): Result<Int> {
        val input = readResource("/day2")

        val regex = """(\d+)-(\d+) ([a-z]): (.*)""".toRegex()

        val matches = regex.findAll(input)

        val count = matches.count { matchResult ->
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


