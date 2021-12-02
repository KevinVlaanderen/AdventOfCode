package days.day15

fun playMemoryGame(data: List<Int>, target: Int): Int {
    val startData = data.dropLast(1)
    val lastIndices = MutableList(target) { index -> startData.indexOf(index) }

    return (data.size - 1 until target - 1).fold(data.last()) { value, index ->
        val lastIndex = lastIndices[value]

        lastIndices[value] = index
        if (lastIndex == -1) 0 else index - lastIndex
    }
}