import framework.Day
import framework.TaskDescription
import framework.TaskRunner

val daysList = mapOf(
    1 to Day(
        "Report Repair", mapOf(
            1 to TaskDescription(days.day1.Task1),
            2 to TaskDescription(days.day1.Task2)
        ),
        "/day1"
    ),
    2 to Day(
        "Password Philosophy", mapOf(
            1 to TaskDescription(days.day2.Task1),
            2 to TaskDescription(days.day2.Task2)
        ),
        "/day2"
    ),
    3 to Day(
        "Toboggan Trajectory", mapOf(
            1 to TaskDescription(days.day3.Task1),
            2 to TaskDescription(days.day3.Task2)
        ),
        "/day3"
    ),
    4 to Day(
        "Passport Processing", mapOf(
            1 to TaskDescription(days.day4.Task1),
            2 to TaskDescription(days.day4.Task2)
        ),
        "/day4"
    ),
    5 to Day(
        "Passport Processing", mapOf(
            1 to TaskDescription(days.day5.Task1),
            2 to TaskDescription(days.day5.Task2)
        ),
        "/day4"
    ),
    6 to Day(
        "Passport Processing", mapOf(
            1 to TaskDescription(days.day5.Task1),
            2 to TaskDescription(days.day5.Task2)
        ),
        "/day4"
    )
)

fun main() {
    for (day in daysList) {
        println("Day ${day.key}: ${day.value.title}")

        for (taskDescription in day.value.taskDescriptions) {
            val result = TaskRunner.run(
                taskDescription.value.task,
                taskDescription.value.resourcePathOverride ?: day.value.defaultResourcePath!!
            )

            println("\tResult for task ${taskDescription.key}: ${result.getOrNull()}")
        }
    }
}
