package days.day3

import days.day3.tasks.Task1
import days.day3.tasks.Task2
import framework.Day

class Day3: Day {
    override val n = 3
    override val title = "Toboggan Trajectory"
    override val tasks = listOf(Task1(), Task2())
}