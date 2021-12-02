package days.day3

fun countTrees(input: List<String>, deltaX: Int, deltaY: Int): Int =
    input
        .filterIndexed { index, _ -> index % deltaY == 0 }
        .mapIndexed { index, row -> row[(index * deltaX) % input[0].length] }
        .count { it == '#' }
