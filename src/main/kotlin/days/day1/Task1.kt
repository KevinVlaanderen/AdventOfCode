package days.day1

import framework.exceptions.AnswerNotFoundException
import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task1: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day1")

        val data = input.lines().filter { it != "" }.map { it.toInt() }

        for ((index, lower) in data.withIndex()) for (upper in data.drop(index+1)) {
            if (lower + upper == 2020) {
                return TaskResult(lower * upper)
            }
        }

        throw AnswerNotFoundException()
    }
}


