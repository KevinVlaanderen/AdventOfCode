package days.day1

import framework.exceptions.AnswerNotFoundException
import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task2: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day1")

        val data = input.lines().filter { it != "" }.map { it.toInt() }

        for ((firstIndex, first) in data.withIndex())
            for ((secondIndex, second) in data.drop(firstIndex+1).withIndex())
                for (third in data.drop(secondIndex+1)) {
                    if (first + second + third == 2020) {
                        return TaskResult(first * second * third)
                    }
        }

        throw AnswerNotFoundException()
    }

    override val number: Int = 2
    override val description: String =
        """
            The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

            Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.
            
            In your expense report, what is the product of the three entries that sum to 2020?
        """.trimIndent()
}


