package days.day2

import framework.Task
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val regex = """(\d+)-(\d+) ([a-z]): (.*)""".toRegex()

        val count = input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .mapNotNull { line -> regex.find(line) }
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


