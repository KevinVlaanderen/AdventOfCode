package framework.interfaces

interface Runnable<out T> {
    fun run(input: String): Result<T>
}