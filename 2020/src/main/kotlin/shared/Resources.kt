package shared

import java.net.URL

fun getResource(path: String): URL =
    when (val url = {}::class.java.getResource(path)) {
        is URL -> url
        else -> throw Exception("Failed to read resource $path")
    }
