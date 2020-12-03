package days.day3.tasks

import framework.Task
import framework.TaskResult
import shared.FileUtil

class Task1: Task {
    override fun run(): TaskResult {
        val input = FileUtil.readResource("/day3")

        val lines = input.lines().filter { it != "" }
        val height = lines.size
        val width = lines[0].length

        val steps = (0 until height).map { (it*3)%width }
        val count = steps.mapIndexed { row, step -> lines[row][step] }.count { it == '#' }

        return TaskResult(count)
    }
}


