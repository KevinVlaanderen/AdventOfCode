package framework.interfaces

import java.net.URL

interface Runnable<out T> {
    fun run(input: URL): Result<T>
}