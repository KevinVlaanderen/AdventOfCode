import framework.Day
import framework.TaskDescription
import framework.TaskRunner

val daysList = mapOf(
    1 to Day(
        "Report Repair", mapOf(
            1 to TaskDescription(days.day1.Task1()),
            2 to TaskDescription(days.day1.Task2())
        ),
        "/day1"
    ),
    2 to Day(
        "Password Philosophy", mapOf(
            1 to TaskDescription(days.day2.Task1()),
            2 to TaskDescription(days.day2.Task2())
        ),
        "/day2"
    ),
    3 to Day(
        "Toboggan Trajectory", mapOf(
            1 to TaskDescription(days.day3.Task1()),
            2 to TaskDescription(days.day3.Task2())
        ),
        "/day3"
    ),
    4 to Day(
        "Passport Processing", mapOf(
            1 to TaskDescription(days.day4.Task1()),
            2 to TaskDescription(days.day4.Task2())
        ),
        "/day4"
    ),
    5 to Day(
        "Binary Boarding", mapOf(
            1 to TaskDescription(days.day5.Task1()),
            2 to TaskDescription(days.day5.Task2())
        ),
        "/day5"
    ),
    6 to Day(
        "Custom Customs", mapOf(
            1 to TaskDescription(days.day6.Task1()),
            2 to TaskDescription(days.day6.Task2())
        ),
        "/day6"
    ),
    7 to Day(
        "Handy Haversacks", mapOf(
            1 to TaskDescription(days.day7.Task1()),
            2 to TaskDescription(days.day7.Task2())
        ),
        "/day7"
    ),
    8 to Day(
        "Handheld Halting", mapOf(
            1 to TaskDescription(days.day8.Task1()),
            2 to TaskDescription(days.day8.Task2())
        ),
        "/day8"
    ),
    9 to Day(
        "Encoding Error", mapOf(
            1 to TaskDescription(days.day9.Task1()),
            2 to TaskDescription(days.day9.Task2())
        ),
        "/day9"
    ),
    10 to Day(
        "Adapter Array", mapOf(
            1 to TaskDescription(days.day10.Task1()),
            2 to TaskDescription(days.day10.Task2())
        ),
        "/day10"
    ),
    11 to Day(
        "Seating System", mapOf(
            1 to TaskDescription(days.day11.Task1()),
            2 to TaskDescription(days.day11.Task2()),
        ),
        "/day11"
    ),
    12 to Day(
        "Rain Risk", mapOf(
            1 to TaskDescription(days.day12.Task1()),
            2 to TaskDescription(days.day12.Task2()),
        ),
        "/day12"
    ),
    13 to Day(
        "Shuttle Search", mapOf(
            1 to TaskDescription(days.day13.Task1()),
            2 to TaskDescription(days.day13.Task2()),
        ),
        "/day13"
    ),
    14 to Day(
        "Docking Data", mapOf(
            1 to TaskDescription(days.day14.Task1()),
            2 to TaskDescription(days.day14.Task2()),
        ),
        "/day14"
    ),
    15 to Day(
        "Rambunctious Recitation", mapOf(
            1 to TaskDescription(days.day15.Task1()),
            2 to TaskDescription(days.day15.Task2()),
        ),
        "/day15"
    ),
    16 to Day(
        "Ticket Translation", mapOf(
            1 to TaskDescription(days.day16.Task1()),
            2 to TaskDescription(days.day16.Task2()),
        ),
        "/day16"
    ),
    17 to Day(
        "Conway Cubes", mapOf(
            1 to TaskDescription(days.day17.Task1()),
            2 to TaskDescription(days.day17.Task2()),
        ),
        "/day17"
    ),
    18 to Day(
        "Operation Order", mapOf(
            1 to TaskDescription(days.day18.Task1()),
            2 to TaskDescription(days.day18.Task2()),
        ),
        "/day18"
    ),
    19 to Day(
        "Monster Messages", mapOf(
            1 to TaskDescription(days.day19.Task1()),
            2 to TaskDescription(days.day19.Task2()),
        ),
        "/day19"
    ),
    20 to Day(
        "Jurassic Jigsaw", mapOf(
            1 to TaskDescription(days.day20.Task1()),
            2 to TaskDescription(days.day20.Task2()),
        ),
        "/day20"
    ),
    21 to Day(
        "Allergen Assessment", mapOf(
            1 to TaskDescription(days.day21.Task1()),
            2 to TaskDescription(days.day21.Task2()),
        ),
        "/day21"
    ),
    22 to Day(
        "Crab Combat", mapOf(
            1 to TaskDescription(days.day22.Task1()),
            2 to TaskDescription(days.day22.Task2()),
        ),
        "/day22"
    ),
    23 to Day(
        "Crab Cups", mapOf(
            1 to TaskDescription(days.day23.Task1()),
            2 to TaskDescription(days.day23.Task2()),
        ),
        "/day23"
    ),
    24 to Day(
        "Lobby Layout", mapOf(
            1 to TaskDescription(days.day24.Task1()),
        ),
        "/day24"
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
