import days.day1.Day1
import days.day2.Day2
import days.day3.Day3
import framework.TaskRunner

val days = listOf(
    Day1(),
    Day2(),
    Day3()
)

fun main() {
    for (day in days) {
        println("Day ${day.n}: ${day.title}")

        for (task in day.tasks) {
            val result = TaskRunner.run(task.implementation)

            println("\tResult for task ${task.n}: ${result.answer}")
        }
    }
}
