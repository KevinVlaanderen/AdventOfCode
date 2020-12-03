package days.day3

import days.day3.tasks.Task1
import framework.Day
import framework.TaskDescription

class Day3: Day {
    override val n = 3
    override val title = "Toboggan Trajectory"
    override val tasks = listOf(
        TaskDescription(2, Task1(),
            "With the toboggan login problems resolved, you set off toward the airport. While travel by toboggan might be easy, it's certainly not safe: there's very minimal steering and the area is covered in trees. You'll need to see which angles will take you near the fewest trees.\n" +
            "\n" +
            "Due to the local geology, trees in this area only grow on exact integer coordinates in a grid. You make a map (your puzzle input) of the open squares (.) and trees (#) you can see. For example:\n" +
            "\n" +
            "..##.......\n" +
            "#...#...#..\n" +
            ".#....#..#.\n" +
            "..#.#...#.#\n" +
            ".#...##..#.\n" +
            "..#.##.....\n" +
            ".#.#.#....#\n" +
            ".#........#\n" +
            "#.##...#...\n" +
            "#...##....#\n" +
            ".#..#...#.#\n" +
            "\n" +
            "These aren't the only trees, though; due to something you read about once involving arboreal genetics and biome stability, the same pattern repeats to the right many times:\n" +
            "\n" +
            "..##.........##.........##.........##.........##.........##.......  --->\n" +
            "#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..\n" +
            ".#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.\n" +
            "..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#\n" +
            ".#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.\n" +
            "..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->\n" +
            ".#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#\n" +
            ".#........#.#........#.#........#.#........#.#........#.#........#\n" +
            "#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...\n" +
            "#...##....##...##....##...##....##...##....##...##....##...##....#\n" +
            ".#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->\n" +
            "\n" +
            "You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).\n" +
            "\n" +
            "The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers); start by counting all the trees you would encounter for the slope right 3, down 1:\n" +
            "\n" +
            "From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.\n" +
            "\n" +
            "The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:\n" +
            "\n" +
            "..##.........##.........##.........##.........##.........##.......  --->\n" +
            "#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..\n" +
            ".#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.\n" +
            "..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#\n" +
            ".#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.\n" +
            "..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->\n" +
            ".#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#\n" +
            ".#........#.#........X.#........#.#........#.#........#.#........#\n" +
            "#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...\n" +
            "#...##....##...##....##...#X....##...##....##...##....##...##....#\n" +
            ".#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->\n" +
            "\n" +
            "In this example, traversing the map using this slope would cause you to encounter 7 trees.\n" +
            "\n" +
            "Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?")
    )
}