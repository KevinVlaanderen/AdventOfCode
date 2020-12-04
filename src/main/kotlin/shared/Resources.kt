package shared

fun readResource(path: String): String =
    when (val data = {}::class.java.getResource(path)?.readText(Charsets.UTF_8)) {
        is String -> data
        else -> throw Exception("Failed to read resource $path")
    }