import framework.Day
import framework.TaskRunner

val daysList = mapOf(
    1 to Day(
        "Report Repair", mapOf(
            1 to days.day1.Task1,
            2 to days.day1.Task2
        ),
        "/day1"
    ),
    2 to Day(
        "Password Philosophy", mapOf(
            1 to days.day2.Task1,
            2 to days.day2.Task2
        ),
        "/day2"
    ),
    3 to Day(
        "Toboggan Trajectory", mapOf(
            1 to days.day3.Task1,
            2 to days.day3.Task2
        ),
        "/day3"
    ),
    4 to Day(
        "Passport Processing", mapOf(
            1 to days.day4.Task1,
            2 to days.day4.Task2
        ),
        "/day4"
    )
)

fun main() {
    for (day in daysList) {
        println("Day ${day.key}: ${day.value.title}")

        for (task in day.value.tasks) {
            val result = TaskRunner.run(task.value, day.value.defaultResourcePath)

            println("\tResult for task ${task.key}: ${result.getOrNull()}")
        }
    }
}
