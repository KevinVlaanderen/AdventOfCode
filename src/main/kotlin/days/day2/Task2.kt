package days.day2

import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task2: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day2")

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

        return TaskResult(count)
    }

    override val number: Int = 2
    override val description: String =
        """
            While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

            The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.
            
            Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.
            
            Given the same example list from above:
            
                1-3 a: abcde is valid: position 1 contains a and position 3 does not.
                1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
                2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
            
            How many passwords are valid according to the new interpretation of the policies?
        """.trimIndent()
}


