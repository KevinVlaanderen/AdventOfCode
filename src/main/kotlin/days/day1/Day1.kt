package days.day1

import days.day1.tasks.Task1
import days.day1.tasks.Task2
import framework.Day

class Day1: Day {
    override val n = 1
    override val title = "Report Repair"
    override val tasks = listOf(Task1(), Task2())
}