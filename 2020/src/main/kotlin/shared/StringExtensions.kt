package shared

fun String.toLines() = this
    .lineSequence()
    .map { it.trim() }
    .filter { it.isNotBlank() }

fun String.toBlocks() = this
    .splitToSequence(System.lineSeparator() + System.lineSeparator())
    .map { it.trim() }
    .filter { it.isNotBlank() }

fun String.extractMap(separator: Char = ':') = this
    .splitToSequence(" ", System.lineSeparator())
    .map { it.trim() }
    .filter { it.isNotBlank() }
    .associate { field ->
        val keyValue = field.split(separator)
        keyValue[0] to keyValue[1]
    }