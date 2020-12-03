package days.day3.tasks

import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task1: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day3").lines().filter { it != "" }

        val count = countTrees(input, 3, 1)

        return TaskResult(count)
    }
}


