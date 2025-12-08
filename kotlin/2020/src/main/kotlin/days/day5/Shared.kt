package days.day5

fun calculateSeatId(line: String): Int {
    val row = mapStringToInteger(line.take(7), mapOf('F' to '0', 'B' to '1'))
    val col = mapStringToInteger(line.takeLast(3), mapOf('L' to '0', 'R' to '1'))

    return row * 8 + col
}

private fun mapStringToInteger(input: String, mapping: Map<Char, Char>): Int =
    Integer.parseInt(
        mapping.entries.fold(input, { result, entry -> result.replace(entry.key, entry.value) }), 2
    )
