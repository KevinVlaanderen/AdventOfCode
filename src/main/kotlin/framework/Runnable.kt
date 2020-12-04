package framework

interface Runnable<out T> {
    fun run(input: String): Result<T>
}