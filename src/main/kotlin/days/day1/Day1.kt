package days.day1

import days.day1.tasks.Task1
import days.day1.tasks.Task2
import framework.Day
import framework.TaskDescription

class Day1: Day {
    override val n = 1
    override val title = "Report Repair"
    override val tasks = listOf(
        TaskDescription(1, Task1(),
            "After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.\n" +
            "\n" +
            "The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.\n" +
            "\n" +
            "To save your vacation, you need to get all fifty stars by December 25th.\n" +
            "\n" +
            "Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!\n" +
            "\n" +
            "Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.\n" +
            "\n" +
            "Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.\n" +
            "\n" +
            "For example, suppose your expense report contained the following:\n" +
            "\n" +
            "1721\n" +
            "979\n" +
            "366\n" +
            "299\n" +
            "675\n" +
            "1456\n" +
            "\n" +
            "In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.\n" +
            "\n" +
            "Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?"),
        TaskDescription(2, Task2(),
            "The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.\n" +
            "\n" +
            "Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.\n" +
            "\n" +
            "In your expense report, what is the product of the three entries that sum to 2020?")
    )
}