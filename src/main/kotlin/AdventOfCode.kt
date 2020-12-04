import framework.Day
import framework.TaskRunner

val daysList = listOf(
    Day(1, "Report Repair", listOf(days.day1.Task1(), days.day1.Task2())),
    Day(2, "Password Philosophy", listOf(days.day2.Task1(), days.day2.Task2())),
    Day(3, "Toboggan Trajectory", listOf(days.day3.Task1(), days.day3.Task2())),
    Day(4, "Passport Processing", listOf(days.day4.Task1(), days.day4.Task2())),
)

fun main() {
    for (day in daysList) {
        println("Day ${day.number}: ${day.title}")

        for (task in day.tasks) {
            val result = TaskRunner.run(task)

            println("\tResult for task ${task.number}: ${result.getOrNull()}")
        }
    }
}
