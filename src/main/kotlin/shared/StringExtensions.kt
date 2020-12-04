package shared

fun String.toBlocks() = this
    .split("""\n{2,}""".toRegex())
    .filter { it.trim() != "" }

fun String.extractDictionary(separator: Char = ':') = this
    .split("""\s+""".toRegex())
    .filter { it.trim() != "" }
    .associate { field ->
        val keyValue = field.split(separator)
        keyValue[0] to keyValue[1]
    }