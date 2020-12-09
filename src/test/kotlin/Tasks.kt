import days.day8.machine.Instruction
import framework.Task
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Nested
import org.junit.jupiter.api.Test
import shared.getResource

class Tasks {
    companion object {
        fun <T> testTask(task: Task<T>, resourcePath: String, expectedResult: T) {
            val result = task.run(getResource(resourcePath))
            assertThat(result.getOrNull()).isEqualTo(expectedResult)
        }
    }

    @Nested
    inner class Day1 {
        @Test
        fun task1() = testTask(days.day1.Task1, "/day1", 902451)

        @Test
        fun task2() = testTask(days.day1.Task2, "/day1", 85555470)
    }

    @Nested
    inner class Day2 {
        @Test
        fun task1() = testTask(days.day2.Task1, "/day2", 625)

        @Test
        fun task2() = testTask(days.day2.Task2, "/day2", 391)
    }

    @Nested
    inner class Day3 {
        @Test
        fun task1() = testTask(days.day3.Task1, "/day3", 247)

        @Test
        fun task2() = testTask(days.day3.Task2, "/day3", 2983070376)
    }

    @Nested
    inner class Day4 {
        @Test
        fun task1() = testTask(days.day4.Task1, "/day4", 250)

        @Test
        fun task2() = testTask(days.day4.Task2, "/day4", 158)
    }

    @Nested
    inner class Day5 {
        @Test
        fun task1() = testTask(days.day5.Task1, "/day5", 991)

        @Test
        fun task2() = testTask(days.day5.Task2, "/day5", 534)
    }

    @Nested
    inner class Day6 {
        @Test
        fun task1() = testTask(days.day6.Task1, "/day6", 6633)

        @Test
        fun task2() = testTask(days.day6.Task2, "/day6", 3202)
    }

    @Nested
    inner class Day7 {
        @Test
        fun task1() = testTask(days.day7.Task1, "/day7", 213)

        @Test
        fun task2() = testTask(days.day7.Task2, "/day7", 38426)
    }

    @Nested
    inner class Day8 {
        @Test
        fun task1() = testTask(days.day8.Task1, "/day8", 2014)

        @Test
        fun task2() = testTask(days.day8.Task2, "/day8", 2251)

        @Test
        fun loadInstructions() {
            val originalInput = getResource("/day8")
                .openStream()
                .bufferedReader()
                .readText()

            val instructions: List<Instruction> = getResource("/day8")
                .openStream()
                .bufferedReader()
                .useLines { days.day8.loadInstructions(it) }

            val output = instructions.joinToString("\n") { it.toString() }

            assertThat(output.trim()).isEqualTo(originalInput.trim())
        }
    }

    @Nested
    inner class Day9 {
        @Test
        fun task1() = testTask(days.day9.Task1, "/day9", 1504371145)
    }
}
