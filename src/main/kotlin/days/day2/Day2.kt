package days.day2

import days.day2.tasks.Task1
import days.day2.tasks.Task2
import framework.Day

class Day2: Day {
    override val n = 2
    override val title = "Password Philosophy"
    override val tasks = listOf(Task1(), Task2())
}