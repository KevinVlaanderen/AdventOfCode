import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.DynamicContainer
import org.junit.jupiter.api.DynamicTest
import org.junit.jupiter.api.TestFactory
import shared.readResource

class Tasks {
    @TestFactory
    fun answers() = mapOf(
        1 to mapOf(
            1 to Triple(days.day1.Task1, "/day1", 902451),
            2 to Triple(days.day1.Task2, "/day1", 85555470)),
        2 to mapOf(
            1 to Triple(days.day2.Task1, "/day2", 625),
            2 to Triple(days.day2.Task2, "/day2", 391)),
        3 to mapOf(
            1 to Triple(days.day3.Task1, "/day3", 247),
            2 to Triple(days.day3.Task2, "/day3", 2983070376)),
        4 to mapOf(
            1 to Triple(days.day4.Task1, "/day4", 250),
            2 to Triple(days.day4.Task2, "/day4", 158)))
        .map { (dayNumber, tasks) ->
            DynamicContainer.dynamicContainer("Day $dayNumber", tasks.map { (task, details) ->
                DynamicTest.dynamicTest("The answer for task $task is ${details.third}") {
                    val result = details.first.run(readResource(details.second))
                    assertThat(result.getOrNull()).isEqualTo(details.third)
                }
            })
        }
}
