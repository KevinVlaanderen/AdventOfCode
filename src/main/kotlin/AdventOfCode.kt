import days.day1.Day1
import framework.TaskRunner

fun main() {
    val days = listOf(Day1())

    for (day in days) {
        println("Day ${day.n}: ${day.title}")

        for (task in day.tasks) {
            val result = TaskRunner.run(task.implementation)

            println("\tResult for task ${task.n}: ${result.answer}")
        }
    }
}
