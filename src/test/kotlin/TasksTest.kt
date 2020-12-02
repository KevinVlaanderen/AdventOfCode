import days.day1.Task1
import days.day1.Task2
import framework.TaskRunner
import org.junit.Assert.assertEquals
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Nested

class TasksTest {
    @Nested
    inner class Day1 {
        @Test
        fun task1() {
            val result = TaskRunner.run(Task1())

            assertEquals("902451", result.answer)
        }

        @Test
        fun task2() {
            val result = TaskRunner.run(Task2())

            assertEquals("85555470", result.answer)
        }
    }
}
