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
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day1.Task1(), "/day1/input", 902451)

            @Test
            fun example1() = testTask(days.day1.Task1(), "/day1/task1/example1", 514579)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day1.Task2(), "/day1/input", 85555470)

            @Test
            fun example1() = testTask(days.day1.Task2(), "/day1/task2/example1", 241861950)
        }
    }

    @Nested
    inner class Day2 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day2.Task1(), "/day2/input", 625)

            @Test
            fun example1() = testTask(days.day2.Task1(), "/day2/task1/example1", 2)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day2.Task2(), "/day2/input", 391)

            @Test
            fun example1() = testTask(days.day2.Task2(), "/day2/task2/example1", 1)
        }
    }

    @Nested
    inner class Day3 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day3.Task1(), "/day3/input", 247)

            @Test
            fun example1() = testTask(days.day3.Task1(), "/day3/task1/example1", 7)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day3.Task2(), "/day3/input", 2983070376)

            @Test
            fun example1() = testTask(days.day3.Task2(), "/day3/task2/example1", 336)
        }
    }

    @Nested
    inner class Day4 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day4.Task1(), "/day4/input", 250)

            @Test
            fun example1() = testTask(days.day4.Task1(), "/day4/task1/example1", 2)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day4.Task2(), "/day4/input", 158)

            @Test
            fun example1() = testTask(days.day4.Task2(), "/day4/task2/example1", 0)

            @Test
            fun example2() = testTask(days.day4.Task2(), "/day4/task2/example2", 4)
        }
    }

    @Nested
    inner class Day5 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day5.Task1(), "/day5/input", 991)

            @Test
            fun example1() = testTask(days.day5.Task1(), "/day5/task1/example1", 357)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day5.Task2(), "/day5/input", 534)
        }
    }

    @Nested
    inner class Day6 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day6.Task1(), "/day6/input", 6633)

            @Test
            fun example1() = testTask(days.day6.Task1(), "/day6/task1/example1", 6)

            @Test
            fun example2() = testTask(days.day6.Task1(), "/day6/task1/example2", 11)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day6.Task2(), "/day6/input", 3202)

            @Test
            fun example1() = testTask(days.day6.Task2(), "/day6/task2/example1", 6)
        }
    }

    @Nested
    inner class Day7 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day7.Task1(), "/day7/input", 213)

            @Test
            fun example1() = testTask(days.day7.Task1(), "/day7/task1/example1", 4)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day7.Task2(), "/day7/input", 38426)

            @Test
            fun example1() = testTask(days.day7.Task2(), "/day7/task2/example1", 32)

            @Test
            fun example2() = testTask(days.day7.Task2(), "/day7/task2/example2", 126)
        }
    }

    @Nested
    inner class Day8 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day8.Task1(), "/day8/input", 2014)

            @Test
            fun example1() = testTask(days.day8.Task1(), "/day8/task1/example1", 5)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day8.Task2(), "/day8/input", 2251)

            @Test
            fun example1() = testTask(days.day8.Task2(), "/day8/task2/example1", 8)
        }

        @Test
        fun loadInstructions() {
            val originalInput = getResource("/day8/input")
                .openStream()
                .bufferedReader()
                .readText()

            val instructions: List<Instruction> = getResource("/day8/input")
                .openStream()
                .bufferedReader()
                .useLines { days.day8.loadInstructions(it) }

            val output = instructions.joinToString("\n") { it.toString() }

            assertThat(output.trim()).isEqualTo(originalInput.trim())
        }
    }

    @Nested
    inner class Day9 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day9.Task1(), "/day9/input", 1504371145)

            @Test
            fun example1() = testTask(days.day9.Task1(5), "/day9/task1/example1", 127)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day9.Task2(), "/day9/input", 183278487)

            @Test
            fun example1() = testTask(days.day9.Task2(127), "/day9/task2/example1", 62)
        }
    }

    @Nested
    inner class Day10 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day10.Task1(), "/day10/input", 2170)

            @Test
            fun example1() = testTask(days.day10.Task1(), "/day10/task1/example1", 35)

            @Test
            fun example2() = testTask(days.day10.Task1(), "/day10/task1/example2", 220)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day10.Task2(), "/day10/input", 24803586664192)

            @Test
            fun example1() = testTask(days.day10.Task2(), "/day10/task2/example1", 8)

            @Test
            fun example2() = testTask(days.day10.Task2(), "/day10/task2/example2", 19208)
        }
    }

    @Nested
    inner class Day11 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day11.Task1(), "/day11/input", 2263)

            @Test
            fun example1() = testTask(days.day11.Task1(), "/day11/task1/example1", 37)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day11.Task2(), "/day11/input", 2002)

            @Test
            fun example1() = testTask(days.day11.Task2(), "/day11/task1/example1", 26)
        }
    }

    @Nested
    inner class Day12 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day12.Task1(), "/day12/input", 1294)

            @Test
            fun example1() = testTask(days.day12.Task1(), "/day12/task1/example1", 25)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day12.Task2(), "/day12/input", 20592)

            @Test
            fun example1() = testTask(days.day12.Task2(), "/day12/task2/example1", 286)
        }
    }

    @Nested
    inner class Day13 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day13.Task1(), "/day13/input", 4722)

            @Test
            fun example1() = testTask(days.day13.Task1(), "/day13/task1/example1", 295)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day13.Task2(), "/day13/input", 825305207525452)

            @Test
            fun example1() = testTask(days.day13.Task2(), "/day13/task2/example1", 3417)

            @Test
            fun example2() = testTask(days.day13.Task2(), "/day13/task2/example2", 754018)

            @Test
            fun example3() = testTask(days.day13.Task2(), "/day13/task2/example3", 779210)

            @Test
            fun example4() = testTask(days.day13.Task2(), "/day13/task2/example4", 1261476)

            @Test
            fun example5() = testTask(days.day13.Task2(), "/day13/task2/example5", 1202161486)
        }
    }

    @Nested
    inner class Day14 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day14.Task1(), "/day14/input", 9879607673316)

            @Test
            fun example1() = testTask(days.day14.Task1(), "/day14/task1/example1", 165)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day14.Task2(), "/day14/input", 3435342392262)

            @Test
            fun example1() = testTask(days.day14.Task2(), "/day14/task2/example1", 208)
        }
    }

    @Nested
    inner class Day15 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day15.Task1(), "/day15/input", 694)

            @Test
            fun example1() = testTask(days.day15.Task1(), "/day15/task1/example1", 436)

            @Test
            fun example2() = testTask(days.day15.Task1(), "/day15/task1/example2", 1)

            @Test
            fun example3() = testTask(days.day15.Task1(), "/day15/task1/example3", 10)

            @Test
            fun example4() = testTask(days.day15.Task1(), "/day15/task1/example4", 27)

            @Test
            fun example5() = testTask(days.day15.Task1(), "/day15/task1/example5", 78)

            @Test
            fun example6() = testTask(days.day15.Task1(), "/day15/task1/example6", 438)

            @Test
            fun example7() = testTask(days.day15.Task1(), "/day15/task1/example7", 1836)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day15.Task2(), "/day15/input", 21768614)

            @Test
            fun example1() = testTask(days.day15.Task2(), "/day15/task2/example1", 175594)

            @Test
            fun example2() = testTask(days.day15.Task2(), "/day15/task2/example2", 2578)

            @Test
            fun example3() = testTask(days.day15.Task2(), "/day15/task2/example3", 3544142)

            @Test
            fun example4() = testTask(days.day15.Task2(), "/day15/task2/example4", 261214)

            @Test
            fun example5() = testTask(days.day15.Task2(), "/day15/task2/example5", 6895259)

            @Test
            fun example6() = testTask(days.day15.Task2(), "/day15/task2/example6", 18)

            @Test
            fun example7() = testTask(days.day15.Task2(), "/day15/task2/example7", 362)
        }
    }

    @Nested
    inner class Day16 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day16.Task1(), "/day16/input", 32842)

            @Test
            fun example1() = testTask(days.day16.Task1(), "/day16/task1/example1", 71)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day16.Task2(), "/day16/input", 2628667251989)
        }
    }

    @Nested
    inner class Day17 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day17.Task1(), "/day17/input", 359)

            @Test
            fun example1() = testTask(days.day17.Task1(), "/day17/task1/example1", 112)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day17.Task2(), "/day17/input", 2228)

            @Test
            fun example1() = testTask(days.day17.Task2(), "/day17/task1/example1", 848)
        }
    }

    @Nested
    inner class Day18 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day18.Task1(), "/day18/input", 7293529867931)

            @Test
            fun example1() = testTask(days.day18.Task1(), "/day18/task1/example1", 71)

            @Test
            fun example2() = testTask(days.day18.Task1(), "/day18/task1/example2", 51)

            @Test
            fun example3() = testTask(days.day18.Task1(), "/day18/task1/example3", 26)

            @Test
            fun example4() = testTask(days.day18.Task1(), "/day18/task1/example4", 437)

            @Test
            fun example5() = testTask(days.day18.Task1(), "/day18/task1/example5", 12240)

            @Test
            fun example6() = testTask(days.day18.Task1(), "/day18/task1/example6", 13632)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day18.Task2(), "/day18/input", -1)

            @Test
            fun example1() = testTask(days.day18.Task2(), "/day18/task2/example1", 231)

            @Test
            fun example2() = testTask(days.day18.Task2(), "/day18/task2/example2", 51)

            @Test
            fun example3() = testTask(days.day18.Task2(), "/day18/task2/example3", 46)

            @Test
            fun example4() = testTask(days.day18.Task2(), "/day18/task2/example4", 1445)

            @Test
            fun example5() = testTask(days.day18.Task2(), "/day18/task2/example5", 669060)

            @Test
            fun example6() = testTask(days.day18.Task2(), "/day18/task2/example6", 23340)
        }
    }

    @Nested
    inner class Day19 {
        @Nested
        inner class Task1 {
            @Test
            fun solution() = testTask(days.day19.Task1(), "/day19/input", 122)

            @Test
            fun example1() = testTask(days.day19.Task1(), "/day19/task1/example1", 2)
        }

        @Nested
        inner class Task2 {
            @Test
            fun solution() = testTask(days.day19.Task2(), "/day19/input", 287)

            @Test
            fun example1() = testTask(days.day19.Task2(), "/day19/task2/example1", 12)
        }
    }
}
