package days.day3.tasks

import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task2: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day3").lines().filter { it != "" }

        val count: Long =
            countTrees(input, 1, 1).toLong() *
            countTrees(input, 3, 1).toLong() *
            countTrees(input, 5, 1).toLong() *
            countTrees(input, 7, 1).toLong() *
            countTrees(input, 1, 2).toLong()

        return TaskResult(count)
    }
}


