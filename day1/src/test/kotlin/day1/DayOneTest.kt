package day1

import org.junit.Assert.assertEquals
import org.junit.Test
import shared.FileUtil

class DayOneTest {
    @Test
    fun testPartOne() {
        val classUnderTest = DayOne()

        val data = FileUtil().readResource("/input") ?: return
        val input = data.lines().filter { it != "" }.map { it.toInt() }

        assertEquals(902451, classUnderTest.partOne(input))
    }

    @Test fun testPartTwo() {
        val classUnderTest = DayOne()

        val data = FileUtil().readResource("/input") ?: return
        val input = data.lines().filter { it != "" }.map { it.toInt() }

        assertEquals(85555470, classUnderTest.partTwo(input))
    }
}
