import days.day1.Day1
import days.day2.Day2
import days.day3.Day3
import framework.TaskRunner
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.DynamicTest
import org.junit.jupiter.api.TestFactory

class TaskTest {
    @TestFactory
    fun testTasks() = listOf(
        Triple(Day1(), 1, "902451"),
        Triple(Day1(), 2, "85555470"),
        Triple(Day2(), 1, "625"),
        Triple(Day2(), 2, "391"),
        Triple(Day3(), 1, "247"),
        Triple(Day3(), 2, "2983070376"))
        .map { (day, task, answer) ->
            DynamicTest.dynamicTest("The answer for ${day.n} task $task is $answer") {
                val result = TaskRunner.run(day.tasks[task-1])
                assertThat(result.answer).isEqualTo(answer)
            }
        }
}
