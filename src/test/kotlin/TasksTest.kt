import days.day1.Day1
import days.day2.Day2
import framework.TaskRunner
import org.junit.Assert.assertEquals
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Nested

class TasksTest {
    @Nested
    inner class Day1Test {
        private val day = Day1()

        @Test
        fun task1() {
            val result = TaskRunner.run(day.tasks[0].implementation)

            assertEquals("902451", result.answer)
        }

        @Test
        fun task2() {
            val result = TaskRunner.run(day.tasks[1].implementation)

            assertEquals("85555470", result.answer)
        }
    }

    @Nested
    inner class Day2Test {
        private val day = Day2()

        @Test
        fun task1() {
            val result = TaskRunner.run(day.tasks[0].implementation)

            assertEquals("625", result.answer)
        }

        @Test
        fun task2() {
            val result = TaskRunner.run(day.tasks[1].implementation)

            assertEquals("391", result.answer)
        }
    }
}
