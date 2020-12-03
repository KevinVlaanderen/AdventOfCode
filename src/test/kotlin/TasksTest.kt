import framework.TaskRunner
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.DynamicTest
import org.junit.jupiter.api.TestFactory

class TaskTest {
    @TestFactory
    fun testTasks() = listOf(
        Pair(days.day1.Task1(), 902451),
        Pair(days.day1.Task2(), 85555470),
        Pair(days.day2.Task1(), 625),
        Pair(days.day2.Task2(), 391),
        Pair(days.day3.Task1(), 247),
        Pair(days.day3.Task2(), 2983070376))
        .map { (task, answer) ->
            DynamicTest.dynamicTest("The answer for task $task is $answer") {
                val result = TaskRunner.run(task)
                assertThat(result.getOrNull()).isEqualTo(answer)
            }
        }
}
