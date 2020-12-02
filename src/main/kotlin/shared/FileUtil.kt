package shared

import java.lang.Exception

object FileUtil {
    fun readResource(path: String): String =
        when (val data = this::class.java.getResource(path)?.readText(Charsets.UTF_8)) {
            is String -> data
            else -> throw Exception("Failed to read resource $path")
        }
}