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
}


