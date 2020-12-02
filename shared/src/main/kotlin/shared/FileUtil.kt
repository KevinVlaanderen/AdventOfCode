package shared

class FileUtil {
    fun readResource(path: String): String? =
        when (val data = this::class.java.getResource(path)?.readText(Charsets.UTF_8)) {
            is String -> data
            else -> {
                println("Failed to read resource $path")
                null
            }
        }
}